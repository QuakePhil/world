package main

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
)

type bouncyball struct {
	x, y float64
	a, v float64
}

func (obj bouncyball) bytes() (s []byte) {
	return []byte(fmt.Sprintf("%.1f %.1f %.1f %.1f", obj.x, obj.y, obj.a, obj.v))
}

func (obj *bouncyball) vectorize(x2, y2 float64) {
	y := y2 - obj.y
	x := x2 - obj.x
	obj.a = math.Atan2(y, x)
	obj.v = math.Sqrt(x*x + y*y)
}

var world []bouncyball

func bouncyballs() []byte {
	var b bytes.Buffer
	for _, obj := range world {
		b.Write(obj.bytes())
		b.Write([]byte(" "))
	}
	return b.Bytes()
}

func bouncyballFromBytes(b []byte) (obj bouncyball) {
	coordinates := bytes.Split(b, []byte(" "))
	// probably can parsefloat from bytes directly, skipping string() ?
	obj.x, _ = strconv.ParseFloat(string(coordinates[0]), 64)
	obj.y, _ = strconv.ParseFloat(string(coordinates[1]), 64)
	x2, _ := strconv.ParseFloat(string(coordinates[2]), 64)
	y2, _ := strconv.ParseFloat(string(coordinates[3]), 64)
	obj.vectorize(x2, y2)
	return
}

func bouncyballSpawn(b []byte) {
	world = append(world, bouncyballFromBytes(b))
	fmt.Println("spawned:", world[len(world)-1])
}
