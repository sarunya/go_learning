package main

import "fmt"

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

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.height * t.base
}

func printArea(sh shape) {
	fmt.Println("Area of the shape is ", sh.getArea())
}

func main() {
	s := square{sideLength: 2}
	printArea(s)

	t := triangle{height: 2, base: 3}
	printArea(t)
}
