package bubbles

import "bytes"

type World struct {
	x, y string
}

func (w World) Frame() []byte {
	return []byte(w.x + " " + w.y)
}

func (w *World) Input(b []byte) {
	coordinates := bytes.Split(b, []byte(" "))
	w.x = string(coordinates[0])
	w.y = string(coordinates[1])
}
