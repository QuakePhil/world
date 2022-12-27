package bouncyball

import "math"

func (obj *bouncyball) bounce() {
	obj.a = obj.collision - obj.a
	// dissipate energy
	obj.v = obj.v * 0.9 // TODO: better formula?
}

func (obj *bouncyball) think() {
	for obj.collide() {
		// TODO: detect when a previously calculated dx,dy pair is considered again
		// and break out of this (possibly) infinite loop
		obj.bounce()
	}
	obj.x += obj.dx
	obj.y += obj.dy
}

// calculate
func (obj *bouncyball) collide() bool {
	obj.dx = math.Cos(obj.a) * obj.v
	obj.dy = math.Sin(obj.a) * obj.v
	x := obj.x + obj.dx
	y := obj.y + obj.dy
	if x < obj.mass || x >= width-obj.mass {
		obj.collision = math.Pi
		return true
	}
	if y < obj.mass || y >= height-obj.mass {
		obj.collision = 0
		return true
	}
	return false
}
