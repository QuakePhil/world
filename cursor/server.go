package cursor

type World struct {
	coordinates []byte
}

func (w World) Frame() []byte {
	return w.coordinates
}

func (w *World) Input(b []byte) {
	w.coordinates = b
}
