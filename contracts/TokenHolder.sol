// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

import "./Erdstall.sol";

abstract contract TokenHolder {

    Erdstall immutable public erdstall;

    constructor(address _erdstall) {
        erdstall = Erdstall(_erdstall);
    }

    modifier onlyErdstall() {
        require(msg.sender == address(erdstall), "TokenHolder: not Erdstall");
        _;
    }

    /*
     * Token Holders must implement transfer and it must use modifier
     * `onlyErdstall`.
     *
     * A token holder also needs means to receive deposits. Each deposit
     * implementation must call `deposit` on the Erdstall contract to register
     * the deposit in the system.
     */

    // transfer should transfer the tokens in `values` from the TokenHolder
    // contract to the recipient.
    //
    // transfer will be called by Erdstall's withdrawal or frozen contract
    // recovery operations.
    function transfer(address token, address recipient, bytes calldata value) virtual external;
}
