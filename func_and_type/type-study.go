package main

func main() {
	var sum int
	sum = 5 + 6 + 3
	avg := sum / 3
	// 以下だと型が違う(intと浮動小数点型の比較)のでコンパイルエラー
	// if avg > 4.5 {
	if avg > 3 {
		println("good")
	}

}
