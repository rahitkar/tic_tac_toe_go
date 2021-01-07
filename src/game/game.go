package game

import (
	"tic_tac_toe/src/player"
)

func Play(player1, player2 player.Player, print func(a ...interface{}) (n int, err error)) {
	var gameBoard board
	p1, p2 := player1, player2
	gameBoard.drawBoard(print)
	gameStatus := gameBoard.driveStatus()
	for gameStatus.progress {
		gameBoard = update(p1, gameBoard, print)
		gameStatus = gameBoard.driveStatus()
		gameBoard.drawBoard(print)
		p1, p2 = swap(p1, p2)
	}
	if gameStatus.won {
		print(p2.Name, " wins")
		return
	}
	print("game draw")
}

func update(p1 player.Player, b board, print func(a ...interface{}) (n int, err error)) board {
	position := p1.Move()
	err, updatedBoard := b.update(position, p1.Symbol)
	if err != nil {
		print(err.Error())
		update(p1, b, print)
	}
	return updatedBoard
}

func swap(player1, player2 player.Player) (player.Player, player.Player) {
	return player2, player1
}
