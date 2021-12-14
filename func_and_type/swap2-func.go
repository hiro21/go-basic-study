package main

func main() {
	n, m := 10, 20
	// &はポインタになる。
	// 設定した値が書き換わるイメージ
	swap2(&n, &m)
	println(n, m)
}

func swap2(arg1, arg2 *int) {
	tmp := *arg1
	*arg1 = *arg2
	*arg2 = tmp
}
