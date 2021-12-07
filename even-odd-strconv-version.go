package main

import "strconv"

func main() {
	for i := 1; i <= 100; i = i + 1 {
		var index string = strconv.Itoa(i)
		if i%2 == 1 {
			println(index + "-奇数")
		} else {
			println(index + "-偶数")
		}
	}
}
