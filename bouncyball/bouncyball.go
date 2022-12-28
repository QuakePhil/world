package bouncyball

import (
	"bytes"
	"fmt"
	"strconv"
)

type bouncyball struct {
	x, y, a, v, dx, dy, mass, charge, collision float64
}

func bouncyballFromBytes(b []byte) (obj bouncyball) {
	coordinates := bytes.Split(b, []byte(" "))
	// probably can parsefloat from bytes directly, skipping string() ?
	obj.x, _ = strconv.ParseFloat(string(coordinates[0]), 64)
	obj.y, _ = strconv.ParseFloat(string(coordinates[1]), 64)
	obj.a, _ = strconv.ParseFloat(string(coordinates[2]), 64)
	obj.v, _ = strconv.ParseFloat(string(coordinates[3]), 64)
	obj.mass, _ = strconv.ParseFloat(string(coordinates[4]), 64)
	obj.charge, _ = strconv.ParseFloat(string(coordinates[5]), 64)
	return
}

func (obj bouncyball) string() string {
	return fmt.Sprintf("%.1f %.1f %.1f %.1f %.1f %.1f", obj.x, obj.y, obj.a, obj.v, obj.mass, obj.charge)
}
