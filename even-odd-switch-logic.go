package main

func main() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%2 == 0:
			println("%d-偶数", i)
		default:
			println("%d-奇数", i)
		}
	}
}
