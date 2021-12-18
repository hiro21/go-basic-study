package main

import (
	"flag"
	"strings"
)

var msg = flag.String("msg", "default", "description")
var n int

func init() {
	flag.IntVar(&n, "n", 1, "count")
}

// go run flag.go -msg=あいう -n=3
// →あいうあいうあいう
func main() {
	flag.Parse()
	println(strings.Repeat(*msg, n))
}
