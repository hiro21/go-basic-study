package main

import "fmt"

type GameScore struct {
	user  string
	point int
}

func main() {
	gameResult := GameScore{user: "userA", point: 15}
	fmt.Printf("%sのゲームポイントは%dです", gameResult.user, gameResult.point)
}
