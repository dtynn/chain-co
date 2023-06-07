package main

import (
	"fmt"
	"os"

	"github.com/ipfs-force-community/sophon-co/api"
)

func main() {
	pkgName := "proxy"

	var proxy api.Proxy
	var local api.Local
	var unsupport api.UnSupport

	targets := []struct {
		def        interface{}
		structName string
		outPath    string
	}{
		{
			def:        &proxy,
			structName: "Proxy",
			outPath:    "./proxy/proxy.go",
		},
		{
			def:        &local,
			structName: "Local",
			outPath:    "./proxy/local.go",
		},
		{
			def:        &unsupport,
			structName: "UnSupport",
			outPath:    "./proxy/unsupport.go",
		},
	}

	for _, t := range targets {
		code, err := Gen(pkgName, t.structName, t.def)
		if err != nil {
			fmt.Println("ERR:", err)
			os.Exit(1)
		}

		if err := os.WriteFile(t.outPath, code, 0o644); err != nil {
			fmt.Println("ERR:", err)
			os.Exit(1)
		}
	}
}
