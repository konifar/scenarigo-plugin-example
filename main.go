package main

import (
	"fmt"
	"os"
	"plugin"
)

// to check gen.so works
func main() {
	p, err := plugin.Open("gen.so")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}

	f, err := p.Lookup("reponame")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}

	f.(func())()
}
