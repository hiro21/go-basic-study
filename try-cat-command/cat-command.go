package main

import (
	"flag"
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

	// -n取得のレシピ（後で最適化する）
	nFlag := flag.Bool("n", false, "trueの場合にnumberを出力するフラグ")
	flag.Parse()
	fmt.Println(*nFlag)
}
