package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/dtynn/chain-co/api"
	"github.com/dtynn/chain-co/proxy-gen"
)

func main() {
	var proxy api.Proxy
	output, err := gen.Gen("proxy", "Proxy", &proxy)
	if err != nil {
		panic(err)
	}

	if err := ioutil.WriteFile("./proxy/proxy.go", output, 0644); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
