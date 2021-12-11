package main

func main() {
	for i := 1; i <= 100; i++ {
		print(i)
		if isEven(i) {
			println("-偶数")
		} else {
			println("-奇数")
		}
	}
}

// 偶数の場合にtrueを返却する
func isEven(num int) bool {
	if num%2 == 0 {
		return true
	} else {
		return false
	}
}
