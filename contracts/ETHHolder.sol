// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

import "./TokenHolder.sol";
import "./lib/Bytes.sol";
import "@openzeppelin/contracts/utils/Address.sol";

contract ETHHolder is TokenHolder {
    using Address for address payable;
    using Bytes for bytes;

    modifier zeroToken(address token) {
        require(token == address(0), "ETHHolder: not zero token address");
        _;
    }

    constructor(address erdstall) TokenHolder(erdstall) {}

    function mint(address, address, bytes calldata) override external pure {
        revert("Wish I could do that.");
    }

    // deposit deposits the message's value into the Erdstall system, crediting
    // the caller.
    function deposit() external payable
    {
        require(msg.value > 0, "ETHHolder: zero value");

        bytes memory value = abi.encodePacked(msg.value);
        erdstall.deposit(msg.sender, address(0), value);
    }

    function transfer(address token, address recipient, bytes calldata value)
    override external onlyErdstall zeroToken(token)
    {
        uint256 amount = value.asUint256();
        payable(recipient).sendValue(amount);
    }
}
