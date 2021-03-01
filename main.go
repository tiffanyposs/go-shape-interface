package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	base   int
	height int
}

type square struct {
	sideLength int
}

func main() {
	t := triangle{
		base:   30,
		height: 40,
	}

	s := square{
		sideLength: 40,
	}

	printArea(t)
	printArea(s)
}

func printArea(s shape) {
	area := s.getArea()
	fmt.Println(area)
}

func (t triangle) getArea() float64 {
	return (0.5 * float64(t.base) * float64(t.height))
}

func (s square) getArea() float64 {
	return float64(s.sideLength) * float64(s.sideLength)
}
