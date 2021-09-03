// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";

contract PerunArt is ERC721 {
    address public immutable owner;
    mapping(address => bool) minters;

    modifier onlyOwner() {
        require(msg.sender == owner, "PerunArt: not owner");
        _;
    }

    modifier onlyMinter() {
        require(minters[msg.sender], "PerunArt: not minter");
        _;
    }

    /**
      * @dev Creates a new PerunArt ERC721 contract.
      *
      * The _minters are able to mint tokens using `mint`.
      * Additional minters can later be added by the owner=deployer using
      * `addMinter`.
      */
    constructor(string memory _name, string memory _symbol, address[] memory _minters)
    ERC721(_name, _symbol)
    {
        owner = msg.sender;
        for (uint i=0; i < _minters.length; i++) {
            minters[_minters[i]] = true;
        }
    }

    function addMinter(address _minter) external onlyOwner {
            minters[_minter] = true;
    }

    function mint(address to, uint256 id) public onlyMinter {
        _mint(to, id);
    }
}
