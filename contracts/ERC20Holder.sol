// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

import "./TokenHolder.sol";
import "./lib/Bytes.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract ERC20Holder is TokenHolder {
    using Bytes for bytes;

    constructor(address erdstall) TokenHolder(erdstall) {}

    function deposit(address token, uint256 amount) external
    {
        require(IERC20(token).transferFrom(msg.sender, address(this), amount),
                "ERC20Holder: transferFrom failed");
        bytes memory value = abi.encodePacked(amount);
        erdstall.deposit(msg.sender, token, value);
    }

    function transfer(address token, address recipient, bytes calldata value)
    override external onlyErdstall
    {
        uint256 amount = value.asUint256();
        require(IERC20(token).transfer(recipient, amount), "ERC20Holder: transfer failed");
    }
}
