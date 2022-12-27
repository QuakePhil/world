package bouncyball

import (
	"bytes"
	"log"
)

type World struct {
	objects []bouncyball
}

func (w World) Frame() []byte {
	var b bytes.Buffer
	for i := range w.objects {
		if i > 0 {
			b.Write([]byte(" "))
		}
		b.Write([]byte(w.objects[i].string()))
		w.objects[i].think()
	}
	return b.Bytes()
}

func (w *World) Input(b []byte) {
	count := bytes.Count(b, []byte(" "))
	if count == 1 {
		resize(b)
	} else {
		w.objects = append(w.objects, bouncyballFromBytes(b))
		log.Println("spawned:", w.objects[len(w.objects)-1])
	}
}
