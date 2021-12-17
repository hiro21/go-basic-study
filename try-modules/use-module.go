package main

import (
	"time"

	"github.com/tenntenn/greeting"
	greeting2 "github.com/tenntenn/greeting/v2"
)

func main() {
	result := greeting.Do()
	println(result)

	g2 := greeting2.Do(time.Now())
	println(g2)
}
