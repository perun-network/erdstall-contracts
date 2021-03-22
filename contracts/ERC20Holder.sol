// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

import "./TokenHolder.sol";
import "./lib/Bytes.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

interface IERC20Minter {
    function mint(address to, uint256 amount) external;
}

contract ERC20Holder is TokenHolder {
    using Bytes for bytes;

    constructor(address erdstall) TokenHolder(erdstall) {}

    function mint(address token, address owner, bytes calldata value)
    override external onlyErdstall {
        uint256 amount = value.asUint256();

        IERC20Minter(token).mint(owner, amount);
    }

    function transferFrom(address _token, address from, address to, bytes calldata value)
    override internal
    {
        uint256 amount = value.asUint256();

        IERC20 token = IERC20(_token);
        require(token.transferFrom(from, to, amount),
                "ERC20Holder: transfer failed");
    }
}
