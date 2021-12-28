package main

import (
	"flag"
	"fmt"
)

/**
catコマンドの実装
引数のファイルパスを受け取り、そのファイル中身を表示する
(複数指定可能)
-nを指定された場合に行番号をつけて表示する。
*/
func main() {
	// -n取得のレシピ（後で最適化する）
	nFlag := flag.Bool("n", false, "trueの場合にnumberを出力するフラグ")
	flag.Parse()
	fmt.Println(*nFlag)

	// フラグを覗いた引数を取得するレシピ
	excludeFlag := flag.Args()
	argsLength := len(excludeFlag)
	for i := 0; i < argsLength; i++ {
		fmt.Println("excludeArgs", excludeFlag[i])
	}
}
