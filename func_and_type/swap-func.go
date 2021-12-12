package main

func main() {
	n, m := swap(10, 20)
	println(n, m)
}

// [param1]と[param2]の順番を入れ替えてresponseする
func swap(param1 int, param2 int) (response1 int, response2 int) {
	return param2, param1
}
