// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

abstract contract TokenHolder {

    address immutable public erdstall;

    constructor(address _erdstall) {
        erdstall = _erdstall;
    }

    modifier onlyErdstall() {
        require(msg.sender == erdstall, "TokenHolder: not Erdstall");
        _;
    }

    // To be used by all but the ETHHolder.
    modifier noEther() {
        require(msg.value == 0, "TokenHolder: ether present");
        _;
    }

    /*
     * Token Holders must implement these functions.
     * All implementations must use modifier `onlyErdstall`.
     */

    function mint(address token, address owner, bytes calldata value) virtual external;

    // deposit should transfer the token to the TokenHolder contract, making it
    // the owner.
    //
    // Derived contracts can use the following default implementation if they
    // implement `transferFrom`.
    function deposit(address token, address depositor, bytes calldata value)
    virtual external payable onlyErdstall noEther
    {
        // depositor must have called `approve` before.
        transferFrom(token, depositor, address(this), value);
    }

    // transfer should transfer the tokens in `values` from the TokenHolder
    // contract to the recipient.
    //
    // Note that values will be several token values encoded. Merge these first
    // before doing the actual token transaction to potentially save gas.
    //
    // transfer will be called by Erdstall's withdrawal or frozen contract
    // recovery operations.
    //
    // Derived contracts can use the following default implementation if they
    // implement `transferFrom`.
    function transfer(address token, address recipient, bytes calldata value)
    virtual external onlyErdstall
    {
        transferFrom(token, address(this), recipient, value);
    }

    // Derived contracts can implement this function and then don't need to
    // implement `deposit` or `transfer`.
    function transferFrom(address token, address from, address to, bytes calldata value) internal virtual;
}
