// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/utils/Strings.sol";

contract PerunArt is ERC721 {
    string private baseURI;
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
    constructor(string memory _name, string memory _symbol,
                string memory _uriBase, address[] memory _minters)
    ERC721(_name, _symbol)
    {
        string memory thisAddr = Strings.toHexString(uint(uint160(address(this))), 20);
        baseURI = string(abi.encodePacked(_uriBase, thisAddr, "/"));
        owner = msg.sender;
        for (uint i=0; i < _minters.length; i++) {
            minters[_minters[i]] = true;
        }
    }

    function _baseURI() internal view override returns (string memory) {
        return baseURI;
    }

    function addMinter(address _minter) external onlyOwner {
            minters[_minter] = true;
    }

    function mint(address to, uint256 id) external virtual onlyMinter {
        _mint(to, id);
    }
}
