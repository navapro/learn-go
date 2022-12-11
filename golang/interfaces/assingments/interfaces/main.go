package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	h float64
	b float64
}

type square struct {
	s float64
}

func main() {
	t := triangle{b: 10, h: 10}
	s := square{s:10}

	printArea(t)
	printArea(s)
}

func (t triangle) getArea() float64 {
	return  (t.h * t.b)/2
}

func (s square) getArea() float64 {
	return s.s * s.s
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}