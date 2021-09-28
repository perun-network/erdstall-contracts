// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

import "./TokenHolder.sol";
import "./lib/Bytes.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";

contract ERC721Holder is TokenHolder {
    using Bytes for bytes;

    constructor(address erdstall) TokenHolder(erdstall) {}

    function deposit(address token, uint256[] memory ids) external
    {
        transferFrom(token, msg.sender, address(this), ids);
        bytes memory value = abi.encodePacked(ids);
        erdstall.deposit(msg.sender, token, value);
    }

    function transfer(address token, address recipient, bytes calldata value)
    virtual override external onlyErdstall
    {
        uint256[] memory ids = value.asUint256sInplace();
        transferFrom(token, address(this), recipient, ids);
    }

    function transferFrom(address _token, address from, address to, uint256[] memory ids)
    internal
    {
        IERC721 token = IERC721(_token);
        for (uint i=0; i < ids.length; i++) {
            token.transferFrom(from, to, ids[i]);
        }
    }
}

interface IERC721Minter is IERC721 {
    function mint(address to, uint256 id) external;
}

contract ERC721MintableHolder is ERC721Holder {
    using Bytes for bytes;

    constructor(address erdstall) ERC721Holder(erdstall) {}

    function transfer(address _token, address recipient, bytes calldata value)
    override external onlyErdstall
    {
        IERC721Minter token = IERC721Minter(_token);
        uint256[] memory ids = value.asUint256sInplace();
        for (uint i=0; i < ids.length; i++) {
            transferOrMint(token, recipient, ids[i]);
        }
    }

    function transferOrMint(IERC721Minter token, address to, uint256 id) internal {
        try token.ownerOf(id) returns (address owner) {
            if (owner == address(this)) {
                token.transferFrom(address(this), to, id);
            }
            // else token exists, but owned by someone else, so we skip
            // minting it. In a production setting, there has to be some
            // mechanism in place to guarantee that off-chain minted tokens
            // are not double-minted.
        } catch Error(string memory) {
            // If the token doesn't exist yet, it means it was minted
            // off-chain, so it is now minted on-chain.
            token.mint(to, id);
        }
    }

    function mint(address token, address owner, bytes calldata value) internal {
        uint256[] memory ids = value.asUint256sInplace();

        IERC721Minter minter = IERC721Minter(token);
        for (uint i=0; i < ids.length; i++) {
            minter.mint(owner, ids[i]);
        }
    }
}

