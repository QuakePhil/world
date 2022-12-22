package main

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"strconv"
)

type bouncyballs struct {
	objects []bouncyball
}

func (obj bouncyball) string() string {
	return fmt.Sprintf("%.1f %.1f %.1f %.1f %.1f", obj.x, obj.y, obj.a, obj.v, obj.mass)
}

func (w bouncyballs) frame() []byte {
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

func (w *bouncyballs) input(b []byte) {
	w.objects = append(w.objects, bouncyballFromBytes(b))
	log.Println("spawned:", w.objects[len(w.objects)-1])
}

type bouncyball struct {
	x, y, a, v, dx, dy, mass float64
}

func bouncyballFromBytes(b []byte) (obj bouncyball) {
	coordinates := bytes.Split(b, []byte(" "))
	// probably can parsefloat from bytes directly, skipping string() ?
	obj.x, _ = strconv.ParseFloat(string(coordinates[0]), 64)
	obj.y, _ = strconv.ParseFloat(string(coordinates[1]), 64)
	obj.mass, _ = strconv.ParseFloat(string(coordinates[4]), 64)
	x2, _ := strconv.ParseFloat(string(coordinates[2]), 64)
	y2, _ := strconv.ParseFloat(string(coordinates[3]), 64)
	obj.vectorize(x2, y2)
	return
}

func (obj *bouncyball) vectorize(x2, y2 float64) {
	y := y2 - obj.y
	x := x2 - obj.x
	obj.a = math.Atan2(y, x)
	obj.v = math.Sqrt(x*x + y*y)
}

func (obj *bouncyball) think() {
	for obj.checkDeltas() {
		obj.a -= math.Pi / 2.0 // TODO: this is clearly incorrect, was too lazy, fix me plz
	}
	obj.x += obj.dx
	obj.y += obj.dy
}

func (obj *bouncyball) checkDeltas() bool {
	obj.dx = math.Cos(obj.a) * obj.v
	obj.dy = math.Sin(obj.a) * obj.v
	x := obj.x + obj.dx
	y := obj.y + obj.dy
	return x < 0 || y < 0 || x >= config.width || y >= config.height
}
