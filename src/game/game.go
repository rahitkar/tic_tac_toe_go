package game

import (
	"tic_tac_toe/src/player"
)

func Play(player1, player2 player.Player, printer func(a ...interface{})) {
	gameBoard := board{cells: [9]string{}, boardStatus: status{won: false, progress: true}}
	p1, p2 := player1, player2
	gameBoard.drawBoard(printer)

	for gameBoard.boardStatus.progress {
		gameBoard = update(p1, gameBoard, printer)
		gameBoard.boardStatus = gameBoard.deriveStatus()
		gameBoard.drawBoard(printer)
		p1, p2 = swap(p1, p2)
	}
	if gameBoard.boardStatus.won {
		printer(p2.Name, " wins")
		return
	}
	printer("game draw")
}

func update(p1 player.Player, b board, printer func(a ...interface{})) board {
	position := p1.Move()
	err, updatedBoard := b.update(position, p1.Symbol)
	if err != nil {
		printer(err.Error())
		return update(p1, b, printer)
	}
	return updatedBoard
}

func swap(player1, player2 player.Player) (player.Player, player.Player) {
	return player2, player1
}
