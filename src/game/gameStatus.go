package game

type status struct {
	won      bool
	progress bool
}

var winingCombination = [8][3]int{
	{0, 1, 2},
	{3, 4, 5},
	{6, 7, 8},
	{0, 3, 6},
	{1, 4, 7},
	{2, 5, 8},
	{0, 4, 8},
	{2, 4, 6},
}

func (s status) deriveStatus(board [9]string) status {
	for _, combination := range winingCombination {
		if board[combination[0]] != "" && board[combination[0]] == board[combination[1]] && board[combination[0]] == board[combination[2]] {
			s.won = true
			s.progress = false
			return s
		}
	}
	s.progress = !isDraw(board)
	return s
}

func isDraw(board [9]string) bool {
	for _, symbol := range board {
		if symbol == "" {
			return false
		}
	}
	return true
}
