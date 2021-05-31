package main

import (
	"fmt"
	"go/build"
	"os"
)

func main() {
	pkg, err := build.Import("github.com/filecoin-project/lotus/build", ".", build.FindOnly)
	if err != nil {
		fmt.Println("resolve import path: ", err)
		os.Exit(1)
	}

	fmt.Println(pkg.Dir)
}
