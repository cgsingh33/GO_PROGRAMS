package main

import "fmt"

type square struct {
	side float64
}

type triangle struct {
	height float64
	base   float64
}

type shape interface {
	getArea() float64
}

func (s square) getArea() float64 {
	return s.side * s.side
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func main() {
	t := triangle{base: 21.34, height: 42.56}
	s := square{side: 24.56}
	printArea(t)
	printArea(s)
}
