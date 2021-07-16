// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";

contract PerunArt is ERC721 {
    // The single allowed token minter
    address public immutable minter;

    modifier onlyMinter() {
        require(msg.sender == minter, "PerunArt: not minter");
        _;
    }

    /**
      * @dev Creates a new PerunArt ERC721 contract.
      * 
      * The _minter is able to mint tokens using `mint`.
      */
    constructor(string memory _name, string memory _symbol, address _minter)
    ERC721(_name, _symbol)
    {
        minter = _minter;
    }

    function mint(address to, uint256 id) public onlyMinter {
        _mint(to, id);
    }
}
