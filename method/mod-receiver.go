package main

type MyInt int

// ポインタにしないとインクリメントが意図通りでない動きになる
func (n *MyInt) Inc() { *n++ }

func main() {
	var n MyInt
	println(n)
	n.Inc()
	println(n)
}
