package main

import (
	"fmt"
	"tic_tac_toe/src/game"
	"tic_tac_toe/src/player"
)

func printer(a...interface{}) {
	fmt.Print(a...)
}

func scaner() int {
	var input int
	fmt.Scanf("%d", &input)
	return input
}

func main() {
	player1 := player.Player{Name: "rahit", Symbol: "X", GetMove: scaner}
	player2 := player.Player{Name: "raja", Symbol: "O", GetMove: scaner}

	game.Play(player1, player2, printer)
}
