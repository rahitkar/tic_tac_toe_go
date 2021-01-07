package main

import (
	"fmt"
	"tic_tac_toe/src/game"
	"tic_tac_toe/src/player"
)

func main() {
	player1 := player.Player{Name: "rahit", Symbol: "X", GetMove: fmt.Scanf}
	player2 := player.Player{Name: "raja", Symbol: "O", GetMove: fmt.Scanf}

	game.Play(player1, player2, fmt.Print)
}
