// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "../PerunArt.sol";

contract RevertToken is PerunArt {
    uint immutable revertModulus;

    modifier checkRevert(uint256 id) {
        require(id % revertModulus != 0, "revert id");
        _;
    }

    constructor(string memory _baseURI, address[] memory _minters, uint _revertModulus)
    PerunArt("Revert", "RVT", _baseURI, _minters)
    {
        revertModulus = _revertModulus;
    }

    function mint(address to, uint256 id) external override onlyMinter checkRevert(id) {
        _mint(to, id);
    }
}
