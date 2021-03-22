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

    function deposit(address token, address, bytes calldata value)
    override external payable onlyErdstall zeroToken(token)
    {
        uint256 amount = value.asUint256();
        require(amount == msg.value, "ETHHolder: value not in msg");
        require(amount > 0, "ETHHolder: zero value");

        // We don't need to store any state, the value is stored in the
        // ETHHolder contract.
    }

    function transfer(address token, address recipient, bytes calldata value)
    override external onlyErdstall zeroToken(token)
    {
        uint256 amount = value.asUint256();
        payable(recipient).sendValue(amount);
    }

    // never used because we overwrite `deposit` and `transfer`.
    function transferFrom(address, address, address, bytes calldata) override internal {}
}
