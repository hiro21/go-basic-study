package main

func main() {
	ns := []int{19, 86, 1, 12}
	var sum int
	for i := 0; i < len(ns); i++ {
		sum += ns[i]
	}
	println(sum)
}
