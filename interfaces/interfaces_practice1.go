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

func (t triangle) getArea() float64 {
	return (t.base * t.height) / 2
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(shp shape) {
	fmt.Println(shp.getArea())
}

func main() {
	sqr := square{
		sideLength: 10,
	}
	trq := triangle{
		height: 3,
		base:   10,
	}

	printArea(sqr)
	printArea(trq)

}
