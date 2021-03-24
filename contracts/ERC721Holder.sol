// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

import "./TokenHolder.sol";
import "./lib/Bytes.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";

interface IERC721Minter {
    function mint(address to, uint256 id) external;
}

contract ERC721Holder is TokenHolder {
    using Bytes for bytes;

    constructor(address erdstall) TokenHolder(erdstall) {}

    function mint(address token, address owner, bytes calldata value)
    override external onlyErdstall
    {
        uint256[] memory ids = value.asUint256sInplace();

        IERC721Minter minter = IERC721Minter(token);
        for (uint i=0; i < ids.length; i++) {
            minter.mint(owner, ids[i]);
        }
    }

    function transferFrom(address _token, address from, address to, bytes calldata value)
    override internal
    {
        uint256[] memory ids = value.asUint256sInplace();

        IERC721 token = IERC721(_token);
        for (uint i=0; i < ids.length; i++) {
            token.safeTransferFrom(from, to, ids[i]);
        }
    }
}
