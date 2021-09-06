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

	f, err := p.Lookup("Reponame")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}

	name := f.(func() string)()
	fmt.Println(name)
}
