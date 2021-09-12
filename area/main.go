package main

import (
	"fmt"
)

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func main() {
	tr := &triangle{
		height: 2.5,
		base:   10,
	}

	sq := &square{
		sideLength: 6.4,
	}

	for _, sh := range []shape{tr, sq} {
		printArea(sh)
	}
}

func (t *triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s *square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(sh shape) {
	area := sh.getArea()
	fmt.Println("Area of the shape is:", area)
}
