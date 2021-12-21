//go:build ignore
// +build ignore

package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// println(flag.Args())
	var msg = flag.String("msg", "default", "arg message")
	flag.Parse()
	fmt.Println(msg)
	fmt.Println(os.Args)
	fmt.Println(flag.Args())
}
