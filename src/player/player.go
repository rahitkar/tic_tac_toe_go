package player

type Player struct {
	Name    string
	Symbol  string
	GetMove func(format string, a ...interface{}) (n int, err error)
}

func (p Player) Move() int {
	var position int
	p.GetMove("%d", &position)
	return position
}
