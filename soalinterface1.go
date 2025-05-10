package main

import "fmt"

func main() {

	circle := Circle{Radius: 10}
	PrintArea(circle)

	rectangle := Rectangle{Width: 5, Height: 8}
	PrintArea(rectangle)
}

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

type Rectangle struct {
	Width float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func PrintArea(s Shape) {
	fmt.Println("Area :", s.Area())
}