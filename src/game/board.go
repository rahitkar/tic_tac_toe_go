package game

import "fmt"

type board struct {
	cells       [9]string
	boardStatus status
}

func (b board) drawBoard(print func(a ...interface{})) {
	for index, symbol := range b.cells {
		if symbol == "" {
			print(index + 1)
		}
		print(symbol)
		if (index+1)%3 == 0 {
			print("\n")
		}
	}
}

type positionOccupied struct {
	err int
}

func (e positionOccupied) Error() string {
	return "position " + string(e.err) + "is occupied"
}

func (b board) update(position int, symbol string) (error, board) {
	fmt.Println(position)
	i := position - 1
	if b.cells[i] == "" {
		b.cells[i] = symbol
		return nil, b
	}
	return positionOccupied{err: position}, b
}

func (b board) deriveStatus() status {
	return b.boardStatus.deriveStatus(b.cells)
}
