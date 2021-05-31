#!/bin/bash

set -e

ver=$(grep 'github.com/filecoin-project/lotus' go.mod | awk '{print $2}')
echo "CHECK FOR lotus@$ver"

rm -rf ./tmp-clone
git clone --quiet -b $ver --depth 1  https://github.com/filecoin-project/lotus.git ./tmp-clone

submodules=$(git submodule status | awk '{print $2}')
for submod in $submodules;
	do currcommit=$(git submodule status | grep $submod | awk '{print $1}' | sed 's/-//');
		targetcommit=$(git -C ./tmp-clone submodule status | grep $submod | awk '{print $1}' | sed 's/-//')
		echo "CHECK SUB MODULE $submod: $targetcommit";
		diff <(echo $targetcommit) <(echo $currcommit);
done

rm -rf ./tmp-clone
