package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	if err := f(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("正常終了")
}

func f() error {
	return errors.New("エラー生成")
}
