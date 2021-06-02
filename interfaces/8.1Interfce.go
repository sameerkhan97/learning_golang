package main

import "fmt"

type shape interface {
	Area() float64
}
type square struct {
	side float64
}

//here square implementing shape interface by defining its function area in it
func (s square) Area() float64 {
	return s.side * s.side
}

type circle struct {
	radius float64
}

//here circle implementing shape interface by defining its function area in it
func (c circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}
func main() {
	circle := circle{10.0}
	square := square{2.0}
	getAreaOf(circle)
	getAreaOf(square)
}

//the shape can call function using any object ht implement its Area function
func getAreaOf(sh shape) {
	fmt.Println("Area ", sh.Area())
}
