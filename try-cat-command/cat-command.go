package main

import (
	"flag"
	"fmt"
	"os"
)

/**
catコマンドの実装
引数のファイルパスを受け取り、そのファイル中身を表示する
(複数指定可能)
-nを指定された場合に行番号をつけて表示する。
*/
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

	// フラグを覗いた引数を取得するレシピ
	excludeFlag := flag.Args()
	fmt.Println("flag.args", excludeFlag)
	fmt.Println("flag.args.len", len(excludeFlag))
	for i := 0; i < len(excludeFlag); i++ {
		fmt.Println("excludeArgs", excludeFlag[i])
	}
}
