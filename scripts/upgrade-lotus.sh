#!/bin/bash

set -e

upgrade() {
	exist=$(grep lotus go.mod | awk '{print $2}')
	echo $exist
	if [[ $exist == $1 ]]; then
		echo "dep is already upgraded"
		exit 0;
	fi

	echo "upgrade go package"
	go get -v github.com/filecoin-project/lotus@$1
	go mod tidy

	echo "init submodule"
	rm -rf ./tmp-clone
	git clone --quiet -b $1 --depth 1  https://github.com/filecoin-project/lotus.git ./tmp-clone

	git submodule init
	submodules=$(git submodule status | awk '{print $2}')
	for submod in $submodules;
		do commit=$(git -C ./tmp-clone submodule status | grep $submod | awk '{print $1}' | sed 's/-//');
			echo "USE $commit FOR $submod";
			git -C $submod fetch origin;
			git -C $submod checkout $commit;
	done
	rm -rf ./tmp-clone
}

upgrade "$@"
