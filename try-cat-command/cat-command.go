package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	argsLength := len(args)
	if argsLength > 1 {
		for i := 1; i < argsLength; i++ {
			fmt.Println(args[i])
		}
	} else {
		fmt.Println("no setting args")
	}
	fmt.Println("finished")
}
