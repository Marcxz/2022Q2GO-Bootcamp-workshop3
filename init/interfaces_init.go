package main

import "math"

// 1.- Create interface shape

// 2.- Create struct circle
type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c circle) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
}

// 3.- Create struct rectangle
type rectangle struct {
	width, height float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

// Implements print shape method
func main() {
	// declaring a circle and a rectangle values

	// circle and rectangle both implements the geometry interface  because they implement all methods of the interface
    // an interface is implicitly implemented in Go


	// do assertions


	// empty interfaces
}