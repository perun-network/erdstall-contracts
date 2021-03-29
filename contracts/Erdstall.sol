// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

import "./TokenHolder.sol";
import "./lib/Sig.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract Erdstall is Ownable {
    // The epoch-balance statements signed by the TEE.
    struct Balance {
        uint64 epoch;
        address account;
        bool exit;
        TokenValue[] tokens;
    }

    struct TokenValue {
        address token;
        bytes value;
    }

    struct MintReceipt {
        address token;
        bytes value;
        bytes sig; // TEE sig on {"ErdstallMint", token, value}
    }
    // TODO:
    //   * encoder and sig verifier for MintReceipt
    //   * function mint(MintReceipt[]) public
    //   * withdrawal with minting function(s) (cheaper combined call)

    // use 2nd-highest number to indicate not-frozen
    uint64 constant notFrozen = 0xfffffffffffffffe;

    // Parameters set during deployment.
    address public immutable tee; // yummi 🍵
    uint64 public immutable bigBang; // start of first epoch
    uint64 public immutable epochDuration; // number of blocks of one epoch

    mapping(address => TokenHolder) public tokenHolders; // token => holder contract

    // epoch => account => token values
    // Multiple deposits can be made in a single epoch, so we have to use an
    // array.
    mapping(uint64 => mapping(address => TokenValue[])) public deposits;
    mapping(uint64 => mapping(address => bool)) public withdrawn; // epoch => account => withdrawn-flag
    mapping(uint64 => mapping(address => bool)) public challenges; // epoch => account => challenge-flag
    mapping(uint64 => uint256) public numChallenges; // epoch => numChallenges
    uint64 public frozenEpoch = notFrozen; // epoch at which contract was frozen

    event TokenRegistered(address indexed token, string tokenType, address holder);
    event Deposited(uint64 indexed epoch, address indexed account, address token, bytes value);
    event Withdrawn(uint64 indexed epoch, address indexed account, TokenValue[] tokens);
    event Challenged(uint64 indexed epoch, address indexed account);
    event ChallengeResponded(uint64 indexed epoch, address indexed account, TokenValue[] tokens, bytes sig);
    event Frozen(uint64 indexed epoch);

    constructor(address _tee, uint64 _epochDuration) {
        tee = _tee;
        bigBang = uint64(block.number);
        epochDuration = _epochDuration;
    }

    function registerTokenHolder(address token, address holder, string calldata tokenType)
    external onlyOwner
    {
        require(address(tokenHolders[token]) == address(0),
                "Erdstall: token holder already set");
        tokenHolders[token] = TokenHolder(holder);

        emit TokenRegistered(token, tokenType, holder);
    }

    function getTokenHolder(address token) internal view returns (TokenHolder) {
        TokenHolder holder = tokenHolders[token];
        require(address(holder) != address(0), "token not registered");
        return holder;
    }


    modifier onlyAlive() {
        require(!isFrozen(), "Erdstall frozen");
        // in case ensureFrozen wasn't called yet...
        require(numChallenges[epoch()-3] == 0, "Erdstall freezing");
        _;
    }

    modifier onlyUnchallenged() {
        require(numChallenges[epoch()-2] == 0, "open challenges");
        _;
    }

    //
    // Normal Operation
    //

    // deposit deposits `amount` token into the Erdstall system, crediting the
    // sender of the transaction. The sender must have set an allowance by
    // calling `allowance` on the token contract before calling `deposit`.
    //
    // Can only be called if there are no open challenges.
    function deposit(address token, bytes calldata value)
    external payable onlyAlive onlyUnchallenged
    {
        uint64 ep = epoch();

        TokenHolder holder = getTokenHolder(token);
        holder.deposit{value: msg.value}(token, msg.sender, value);
        // Record deposit in case of freeze of the next two epochs.
        deposits[ep][msg.sender].push(TokenValue({token: token, value: value}));

        emit Deposited(ep, msg.sender, token, value);
    }

    // withdraw lets a user withdraw their funds from the system. It is only
    // possible to withdraw with an exit proof, that is, a balance proof with
    // field `exit` set to true.
    //
    // `sig` must be a signature created with
    // `signText(keccak256(abi.encode(...)))`.
    // See `encodeBalanceProof` for exact encoding specification.
    //
    // Can only be called if there are no open challenges.
    function withdraw(Balance calldata balance, bytes calldata sig)
    external onlyAlive onlyUnchallenged
    {
        require(balance.epoch <= sealedEpoch(), "withdraw: too early");
        require(balance.account == msg.sender, "withdraw: wrong sender");
        require(balance.exit, "withdraw: no exit proof");
        verifyBalance(balance, sig);

        _withdraw(balance.epoch, balance.tokens);
    }

    // Transfers all tokens to msg.sender, using each token's token holder.
    function _withdraw(uint64 _epoch, TokenValue[] memory tokens) internal {
        require(!withdrawn[_epoch][msg.sender], "already withdrawn");
        withdrawn[_epoch][msg.sender] = true;

        for (uint i=0; i < tokens.length; i++) {
            TokenHolder holder = getTokenHolder(tokens[i].token);
            holder.transfer(tokens[i].token, msg.sender, tokens[i].value);
        }

        emit Withdrawn(_epoch, msg.sender, tokens);
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
        require(!balance.exit, "challenge: exit proof");
        verifyBalance(balance, sig);

        registerChallenge(balance.tokens.length);
    }

    // challengeDeposit should be called by a user if they deposited but never
    // received a balance proof from the operator.
    //
    // After a challenge is opened, the operator (anyone, actually) can respond
    // to the challenge using function `respondChallenge`.
    function challengeDeposit() external onlyAlive {
        registerChallenge(0);
    }

    function registerChallenge(uint256 numSealedValues) internal {
        uint64 challEpoch = epoch() - 1;
        require(!challenges[challEpoch][msg.sender], "already challenged");

        uint256 numValues =
            numSealedValues + deposits[challEpoch][msg.sender].length;
        require(numValues > 0, "no value in system");

        challenges[challEpoch][msg.sender] = true;
        numChallenges[challEpoch]++;

        emit Challenged(challEpoch, msg.sender);
    }

    // respondChallenge lets the operator (or anyone) respond to open challenges
    // that were posted during the previous on-chain epoch.
    //
    // The latest or second latest balance proof needs to be submitted, so the
    // contract can verify that all challenges were seen by the enclave and all
    // challenging users were exited from the system.
    //
    // Note that a challenging user has to subscribe to ChallengeResponded
    // events from both, the latest and second latest epoch to receive their
    // exit proof.
    function respondChallenge(Balance calldata balance, bytes calldata sig) external onlyAlive {
        uint64 challEpoch = epoch() - 2;
        require((balance.epoch == challEpoch) || (balance.epoch == challEpoch+1),
                "respondChallenge: wrong epoch");
        require(balance.exit, "respondChallenge: no exit proof");
        verifyBalance(balance, sig);

        require(challenges[challEpoch][balance.account], "respondChallenge: no challenge registered");
        challenges[challEpoch][balance.account] = false;
        numChallenges[challEpoch]--;

        // The challenging user can create their exit proof from this data and
        // withdraw with `withdraw` starting with the next epoch.
        emit ChallengeResponded(balance.epoch, balance.account, balance.tokens, sig);
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

        // Also recover deposits from broken epochs
        TokenValue[] memory tokens = appendFrozenDeposits(balance.tokens);
        _withdraw(frozenEpoch, tokens);
    }

    // withdrawFrozenDeposit lets any user withdraw the deposits that they made
    // in the challenged and broken epoch.
    //
    // This function should only be called by users who never received a balance
    // proof and thus never had any balance in the system, to recover their
    // deposits.
    function withdrawFrozenDeposit() external {
        ensureFrozen();

        TokenValue[] memory tokens = appendFrozenDeposits(new TokenValue[](0));
        require(tokens.length > 0, "no frozen deposit");

        _withdraw(frozenEpoch, tokens);
    }

    // ensureFrozen ensures that the state of the contract is set to frozen if
    // the last epoch has at least one unanswered challenge. It is idempotent.
    //
    // It is implicitly called by withdrawFrozen and withdrawFrozenDeposit but
    // can be called seperately if the contract should be frozen before anyone
    // wants to withdraw.
    function ensureFrozen() public {
        if (isFrozen()) { return; }
        uint64 challEpoch = epoch() - 3;
        require(numChallenges[challEpoch] > 0, "no open challenges");

        // freezing to previous, that is, last unchallenged epoch
        uint64 frEpoch = challEpoch - 1;
        frozenEpoch = frEpoch;

        emit Frozen(frEpoch);
    }

    function appendFrozenDeposits(TokenValue[] memory tokens)
    internal view returns (TokenValue[] memory)
    {
        TokenValue[] storage deps1 = deposits[frozenEpoch+1][msg.sender];
        TokenValue[] storage deps2 = deposits[frozenEpoch+2][msg.sender];
        TokenValue[] memory deps = new TokenValue[](tokens.length + deps1.length + deps2.length);

        // Concatenate arrays
        for (uint i=0; i < tokens.length; i++) {
            deps[i] = tokens[i];
        }
        for (uint i=0; i < deps1.length; i++) {
            deps[i+tokens.length] = deps1[i];
        }
        for (uint i=0; i < deps2.length; i++) {
            deps[i+tokens.length+deps1.length] = deps2[i];
        }

        return deps;
    }

    function isFrozen() internal view returns (bool) {
        return frozenEpoch != notFrozen;
    }

    function sealedEpoch() internal view returns (uint64) {
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
            balance.tokens,
            balance.exit);
    }
}
