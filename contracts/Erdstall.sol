// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.7.0;
pragma experimental ABIEncoderV2;

import "./lib/Sig.sol";
import "../vendor/openzeppelin-contracts/contracts/token/ERC20/IERC20.sol";

contract Erdstall {
    // The epoch-balance statements signed by the TEE.
    struct Balance {
        uint64 epoch;
        address account;
        bool exit;
        uint256 value;
    }

    uint64 constant notFrozen = uint64(-2); // use 2nd-highest number to indicate not-frozen

    // Parameters set during deployment.
    address public immutable tee; // yummi 🍵
    uint64 public immutable bigBang; // start of first epoch
    uint64 public immutable epochDuration; // number of blocks of one epoch phase
    IERC20 public immutable token; // ERC20 token handled by this Erdstall

    mapping(uint64 => mapping(address => uint256)) public deposits; // epoch => account => deposit value
    mapping(uint64 => mapping(address => bool)) public withdrawn; // epoch => account => withdrawn-flag
    mapping(uint64 => mapping(address => bool)) public challenges; // epoch => account => challenge-flag
    mapping(uint64 => uint256) public numChallenges; // epoch => numChallenges
    uint64 public frozenEpoch = notFrozen; // epoch at which contract was frozen

    event Deposited(uint64 indexed epoch, address indexed account, uint256 value);
    event Withdrawn(uint64 indexed epoch, address indexed account, uint256 value);
    event Challenged(uint64 indexed epoch, address indexed account);
    event ChallengeResponded(uint64 indexed epoch, address indexed account, uint256 value, bytes sig);
    event Frozen(uint64 indexed epoch);

    constructor(address _tee, uint64 _epochDuration, address _token) {
        tee = _tee;
        bigBang = uint64(block.number);
        epochDuration = _epochDuration;
        token = IERC20(_token);
    }

    modifier onlyAlive {
        require(!isFrozen(), "Erdstall frozen");
        // in case freeze wasn't called yet...
        require(!isLastEpochChallenged(), "Erdstall freezing");
        _;
    }

    //
    // Normal Operation
    //

    // deposit deposits `amount` token into the Erdstall system, crediting the
    // sender of the transaction. The sender must have set an allowance by
    // calling `allowance` on the token contract before calling `deposit`.
    function deposit(uint256 amount) external onlyAlive {
        uint64 epoch = depositEpoch();
        require(token.transferFrom(msg.sender, address(this), amount),
                "deposit: token transfer failed");
        deposits[epoch][msg.sender] += amount;

        emit Deposited(epoch, msg.sender, amount);
    }

    // withdraw lets a user withdraw their funds from the system. It is only
    // possible to withdraw with an exit proof, that is, a balance proof with
    // field `exit` set to true.
    //
    // `sig` must be a signature created with
    // `signText(keccak256(abi.encode(balance)))`.
    function withdraw(Balance calldata balance, bytes calldata sig) external onlyAlive {
        require(balance.epoch <= sealedEpoch(), "withdraw: too early");
        require(balance.account == msg.sender, "withdraw: wrong sender");
        require(balance.exit, "withdraw: no exit proof");
        verifyBalance(balance, sig);

        _withdraw(balance.epoch, balance.value);
    }

    function _withdraw(uint64 epoch, uint256 value) internal {
        require(!withdrawn[epoch][msg.sender], "already withdrawn");
        withdrawn[epoch][msg.sender] = true;

        require(token.transfer(msg.sender, value), "withdraw: token transfer failed");
        emit Withdrawn(epoch, msg.sender, value);
    }

    //
    // Challenge Functions
    //

    // Challenges the operator to post the user's last epoch's exit proof.
    // The user needs to pass the latest balance proof, that is, of the just
    // sealed epoch, to proof that they are part of the system.
    //
    // After a challenge is opened, the operator (anyone, actually) can respond
    // to the challenge using function `respondChallenge`.
    function challenge(Balance calldata balance, bytes calldata sig) external onlyAlive {
        require(balance.account == msg.sender, "challenge: wrong sender");
        require(balance.epoch == sealedEpoch(), "challenge: wrong epoch");
        verifyBalance(balance, sig);

        registerChallenge(balance.value);
    }

    // challengeDeposit should be called by a user if they deposited but never
    // received a balance proof from the operator.
    //
    // After a challenge is opened, the operator (anyone, actually) can respond
    // to the challenge using function `respondChallenge`.
    function challengeDeposit() external onlyAlive {
        registerChallenge(0);
    }

    function registerChallenge(uint256 sealedValue) internal {
        uint64 epoch = challengeEpoch();
        require(!challenges[epoch][msg.sender], "already challenged");

        uint256 value = sealedValue + deposits[epoch][msg.sender];
        require(value > 0, "no value in system");

        challenges[epoch][msg.sender] = true;
        numChallenges[epoch]++;

        emit Challenged(epoch, msg.sender);
    }

    function respondChallenge(Balance calldata balance, bytes calldata sig) external onlyAlive {
        require(balance.epoch == responseEpoch(), "respondChallenge: wrong epoch");
        require(balance.exit, "respondChallenge: no exit proof");
        verifyBalance(balance, sig);

        uint64 challEpoch = balance.epoch - 1;
        require(challenges[challEpoch][balance.account], "respondChallenge: no challenge registered");
        challenges[challEpoch][balance.account] = false;
        numChallenges[challEpoch]--;

        // The challenging user can create their exit proof from this data and
        // withdraw with `withdraw`.
        emit ChallengeResponded(balance.epoch, balance.account, balance.value, sig);
    }

    // withdrawFrozen lets any user withdraw all funds locked in the frozen
    // contract. Parameter `balance` needs to be the balance proof of the last
    // unchallenged epoch.
    //
    // Implicitly calls ensureFrozen to ensure that the contract state is set to
    // frozen if the last epoch has an unanswered challenge.
    function withdrawFrozen(Balance calldata balance, bytes calldata sig) external {
        ensureFrozen();

        require(balance.account == msg.sender, "withdrawFrozen: wrong sender");
        require(balance.epoch == frozenEpoch, "withdrawFrozen: wrong epoch");
        verifyBalance(balance, sig);

        // Also recover deposits from broken epoch
        uint256 value = balance.value + deposits[frozenEpoch+1][msg.sender];

        _withdraw(frozenEpoch, value);
    }

    // ensureFrozen ensures that the state of the contract is set to frozen if
    // the last epoch has at least one unanswered challenge.
    //
    // It is implicitly called by withdrawFrozen but can be called seperately if
    // the contract should be frozen before anyone wants to withdraw.
    function ensureFrozen() public {
        if (isFrozen()) { return; }
        require(isLastEpochChallenged(), "no challenge in last epoch");

        // freezing to previous, that is, last unchallenged epoch
        uint64 epoch = challengedEpoch() - 1;
        frozenEpoch = epoch;

        emit Frozen(epoch);
    }

    function isLastEpochChallenged() internal view returns (bool) {
        return numChallenges[challengedEpoch()] > 0;
    }

    function isFrozen() internal view returns (bool) {
        return frozenEpoch != notFrozen;
    }

    //
    // Epoch Counter Abstractions
    //

    function depositEpoch() internal view returns (uint64) {
        return epoch();
    }

    function challengeEpoch() internal view returns (uint64) {
        return epoch()-1;
    }

    function responseEpoch() internal view returns (uint64) {
        return epoch()-1;
    }

    function sealedEpoch() internal view returns (uint64) {
        return epoch()-2;
    }

    function challengedEpoch() internal view returns (uint64) {
        return epoch()-2;
    }

    // epoch returns the current epoch. It should not be used directly in public
    // functions, but the fooEoch functions instead, as they account for the
    // correct shifts.
    function epoch() internal view returns (uint64) {
        return (uint64(block.number) - bigBang) / epochDuration;
    }

    function verifyBalance(Balance memory balance, bytes memory sig) public view {
        require(Sig.verify(encodeBalanceProof(balance), sig, tee), "invalid signature");
    }

    function encodeBalanceProof(Balance memory balance) public view returns (bytes memory) {
        return abi.encode(
            "ErdstallBalance",
            address(this),
            balance.epoch,
            balance.account,
            balance.value,
            balance.exit);
    }
}
