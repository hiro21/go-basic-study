package main

import (
	"math/rand"
	"time"
)

func main() {
	t := time.Now().UnixNano()
	rand.Seed(t)
	result := rand.Intn(6)
	switch result {
	case 0:
		println("凶")
	case 1, 2:
		println("吉")
	case 3, 4:
		println("中吉")
	case 5:
		println("大吉")
	default:
		println("想定外エラー[%d]", result)
		panic("想定外エラー")
	}

}
