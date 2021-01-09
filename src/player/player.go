package player

type Player struct {
	Name    string
	Symbol  string
	GetMove func() int
}

func (p Player) Move() int {
	return p.GetMove()
}
