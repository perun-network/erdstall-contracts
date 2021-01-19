// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.7.0;

import "../vendor/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol";

contract PerunToken is ERC20 {
    using SafeMath for uint256;

    /**
     * @dev Creates a new PerunToken contract instance with `accounts` being
     * funded with `initBalance` tokens.
     */
    constructor (address[] memory accounts, uint256 initBalance) ERC20("PerunToken", "PRN") {
        for (uint256 i = 0; i < accounts.length; i++) {
            _mint(accounts[i], initBalance);
        }
    }
}
