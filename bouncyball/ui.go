package bouncyball

import (
	"bytes"
	"strconv"
)

var width, height float64

func resize(b []byte) {
	parts := bytes.Split(b, []byte(" "))
	width, _ = strconv.ParseFloat(string(parts[0]), 64)
	height, _ = strconv.ParseFloat(string(parts[1]), 64)
}
