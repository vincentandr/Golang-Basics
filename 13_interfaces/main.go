package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func getArea(s Shape) float64 {
	return s.area()
}

// ----------------

type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct {
	temp int
}

type TCPWriter struct {}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

func main() {
	circle := Circle{0,0,5}
	rectangle := Rectangle{10, 5}

	fmt.Printf("Circle Area: %f\n", getArea(circle))
	fmt.Printf("circle Area: %f\n", getArea(rectangle))

	// Test writer
	var w Writer = ConsoleWriter{} // Polymorphism. Can also be &ConsoleWriter{}, this enables line 44 func parameter to be cw *ConsoleWriter, thus enables to directly change field values, e.g. temp. Otherwise, they must be value parameter, not pointer.
	w.Write([]byte("Hello World!"))
	
	// To access temp variable in ConsoleWriter, cast w to ConsoleWriter
	cw := w.(ConsoleWriter)
	cw.temp = 5
	fmt.Println(cw.temp)

}

