package game

import (
	"tic_tac_toe/src/player"
)

type Game struct {
	Player1 player.Player
	Player2 player.Player
	board   [9]string
}

func (game Game) Play(print func(a ...interface{}) (n int, err error)) {
	game.board = [9]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	gameStatus := status{won: false, draw: false, progress: false}
	drawBoard(print, game.board)
	for gameStatus = gameStatus.deriveStatus(game.board); !gameStatus.progress; {
		position := game.Player1.Move()
		game.board[position-1] = game.Player1.Symbol
		game.Player1, game.Player2 = swap(game.Player1, game.Player2)
		drawBoard(print, game.board)
	}
	if gameStatus.won {
		print(game.Player2, "won")
		return
	}
	print("game draw")
}

func drawBoard(print func(a ...interface{}) (n int, err error), board [9]string) {
	for i := 0; i < len(board); i += 3 {
		print(board[i], board[i+1], board[i+2], "\n")
	}
}

func swap(player1, player2 player.Player) (player.Player, player.Player) {
	return player2, player1
}
