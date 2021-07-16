// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

// Minimal bytes library, inspired by
// https://github.com/GNSPS/solidity-bytes-utils
library Bytes {
    function asUint256(bytes memory x) internal pure returns (uint256 r) {
        require(x.length == 32, "Bytes: not length 32");
        assembly {
            r := mload(add(x, 0x20))
        }
    }

    function asUint256s(bytes memory x) internal pure returns (uint256[] memory) {
        require(x.length % 32 == 0, "Bytes: length not multiple of 32");

        uint256[] memory y = new uint256[](x.length/32);

        assembly {
            let a := add(x, 0x20) // initial bytes pointer
            let b := add(y, 0x20) // initial uint[] pointer
            for
                {
                    let i := 0 // relative mem pointer
                    let end := mload(x) // length of x is first word of x
                }
                lt(i, end)
                { i := add(i, 0x20) }
            {
                mstore(add(b, i), mload(add(a, i)))
            }
        }

        return y;
    }

    // Returns a pointer to the same memory array as x and destructively sets
    // the first word, which is the length of the array, to the bytes length
    // divided by 32, effectively turning the bytes array into an uint256[].
    //
    // The original reference bytes x must not be used after this destructive
    // operation.
    function asUint256sInplace(bytes memory x) internal pure returns (uint256[] memory r) {
        require(x.length % 32 == 0, "Bytes: length not multiple of 32");
        assembly {
            mstore(x, div(mload(x), 0x20))
            r := x
        }
    }
}
