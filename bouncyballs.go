package main

import (
	"bytes"
	"fmt"
	"strconv"
)

type bouncyball struct {
	x, y float64
	a, v float64
}

func (obj bouncyball) bytes() (s []byte) {
	return []byte(fmt.Sprintf("%f %f %f %f", obj.x, obj.y, obj.a, obj.v))
}

var world []bouncyball

func bouncyballs() {
	for _, obj := range world {
		ws.Broadcast(obj.bytes())
	}
	// todo: Broadcast here instead
	// and clear canvas on draw
}

func bouncyball_spawn(b []byte) {
	var obj bouncyball
	var err error
	coordinates := bytes.Split(b, []byte(" "))
	// probably can parsefloat from bytes directly, skipping string() ?
	obj.x, err = strconv.ParseFloat(string(coordinates[0]), 64)
	check(err)
	obj.y, err = strconv.ParseFloat(string(coordinates[1]), 64)
	check(err)
	obj.a, err = strconv.ParseFloat(string(coordinates[2]), 64)
	check(err)
	obj.v, err = strconv.ParseFloat(string(coordinates[3]), 64)
	check(err)
	world = append(world, obj)
}
