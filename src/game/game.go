package game

import (
	"tic_tac_toe/src/player"
)

func Play(player1, player2 player.Player, print func(a ...interface{}) (n int, err error)) {
	p1, p2 := player1, player2
	board, gameStatus := setupGame()
	drawBoard(print, board)

	for gameStatus.progress {
		position := p1.Move()
		board[position-1] = p1.Symbol
		gameStatus = gameStatus.deriveStatus(board)
		p1, p2 = swap(p1, p2)
		drawBoard(print, board)
	}
	if gameStatus.won {
		print(p2.Name, " won")
		return
	}
	print("game draw")
}

func setupGame() ([9]string, status) {
	board := [9]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	gameStatus := status{won: false, progress: true}
	return board, gameStatus
}

func drawBoard(print func(a ...interface{}) (n int, err error), board [9]string) {
	for i := 0; i < len(board); i += 3 {
		print(board[i], board[i+1], board[i+2], "\n")
	}
}

func swap(player1, player2 player.Player) (player.Player, player.Player) {
	return player2, player1
}
