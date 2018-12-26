package main

import "fmt"

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

type shape interface {
	getArea() float64
}

func main() {
	sq := square{sideLength: 10}
	tg := triangle{height: 5, base: 10}

	printArea(sq)
	printArea(tg)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (sq square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}

func (tg triangle) getArea() float64 {
	return 0.5 * tg.base * tg.height
}
