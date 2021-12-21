package main

import "os"

func main() {
	args := os.Args
	// 1つめはプログラム名
	println(args[0])
}
