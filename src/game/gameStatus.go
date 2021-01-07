package game

type status struct {
	won      bool
	draw     bool
	progress bool
}

func (s status) deriveStatus([9]string) status {
	return status{
		won:      false,
		draw:     false,
		progress: false,
	}
}
