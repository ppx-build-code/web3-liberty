package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area () float64
	Perimeter () float64
}

type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area () float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter () float64 {
	return 2 * math.Pi * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func printShapeInfo(s Shape) {
    fmt.Printf("面积: %.2f, 周长: %.2f\n", s.Area(), s.Perimeter())
}

func main() {
	
	c := Circle{Radius:5}
	r := Rectangle{Width:5, Height:10}
	printShapeInfo(c)
	printShapeInfo(r)
}