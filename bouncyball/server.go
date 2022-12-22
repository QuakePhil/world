package bouncyball

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"strconv"
)

type World struct {
	objects []bouncyball
}

var width, height float64

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

func resize(b []byte) {
	parts := bytes.Split(b, []byte(" "))
	width, _ = strconv.ParseFloat(string(parts[0]), 64)
	height, _ = strconv.ParseFloat(string(parts[1]), 64)
}

type bouncyball struct {
	x, y, a, v, dx, dy, mass float64
}

func bouncyballFromBytes(b []byte) (obj bouncyball) {
	coordinates := bytes.Split(b, []byte(" "))
	// probably can parsefloat from bytes directly, skipping string() ?
	obj.x, _ = strconv.ParseFloat(string(coordinates[0]), 64)
	obj.y, _ = strconv.ParseFloat(string(coordinates[1]), 64)
	obj.a, _ = strconv.ParseFloat(string(coordinates[2]), 64)
	obj.v, _ = strconv.ParseFloat(string(coordinates[3]), 64)
	obj.mass, _ = strconv.ParseFloat(string(coordinates[4]), 64)
	return
}

func (obj bouncyball) string() string {
	return fmt.Sprintf("%.1f %.1f %.1f %.1f %.1f", obj.x, obj.y, obj.a, obj.v, obj.mass)
}

func (obj *bouncyball) think() {
	for obj.checkDeltas() {
		obj.a -= math.Pi / 2.0 // TODO: this is clearly incorrect, was too lazy, fix me plz
		// dissipate energy
		obj.v = obj.v * 0.9 // TODO: better formula?
	}
	obj.x += obj.dx
	obj.y += obj.dy
}

func (obj *bouncyball) checkDeltas() bool {
	obj.dx = math.Cos(obj.a) * obj.v
	obj.dy = math.Sin(obj.a) * obj.v
	x := obj.x + obj.dx
	y := obj.y + obj.dy
	return x < 0 || y < 0 || x >= width || y >= height
}
