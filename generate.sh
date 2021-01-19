#!/bin/bash

# SPDX-License-Identifier: Apache-2.0

set -e

# Download solc.
wget -nc "https://github.com/ethereum/solidity/releases/download/v0.7.6/solc-static-linux"
chmod +x solc-static-linux
echo -e "Ensure that the newest version of abigen is installed"

solpath="contracts"

# Generates optimized golang bindings and runtime binaries for sol contracts.
# $1  solidity file path, relative to $solpath/.
# $1  golang package name.
# $2… list of contract names.
function generate() {
    file=$1; pkg=$2
    shift; shift   # skip the first two args.
    abigen --pkg $pkg --sol $solpath/$file.sol --out $pkg/$file.go --solc ./solc-static-linux
    ./solc-static-linux --bin-runtime --optimize --allow-paths *, $solpath/$file.sol --overwrite -o $pkg/
    for contract in "$@"; do
        echo -e "package $pkg\n\n // ${contract}BinRuntime is the runtime part of the compiled bytecode used for deploying new contracts.\nvar ${contract}BinRuntime = \`$(<${pkg}/${contract}.bin-runtime)\`" > "$pkg/${contract}BinRuntime.go"
    done
}

# Pragma statements can only be in a solidity file once, so we remove the
# duplicates with awk.
cat ${solpath}/Erdstall.sol \
    <(awk '/^pragma/{p=1;next}{if(p){print}}' ${solpath}/PerunToken.sol) \
    > ${solpath}/Contracts.sol
# Generate bindings
generate "Contracts" "bindings" "Erdstall" "PerunToken"
rm ${solpath}/Contracts.sol

abigen --version --solc ./solc-static-linux
echo -e "Generated bindings"
