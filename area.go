// EXAMPLE OF CALCULATE THE AREA OF A TRIANGLE AND A SQUARE IN GO LENGUAGE BY DANI95RICO

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

func main() {
	tr := triangle{7, 5}
	sq := square{5}

	fmt.Print("Area of the triangle: ")
	printArea(tr)

	fmt.Print("Area of the square: ")
	printArea(sq)

}

func (t triangle) getArea() float64 {
	return (0.5 * t.height * t.base)
}

func (s square) getArea() float64 {
	return (s.sideLength * s.sideLength)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
