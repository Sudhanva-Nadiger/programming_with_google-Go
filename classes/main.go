package main

import (
	"math"
	"strconv" // Importing strconv for string conversion
)

type Person struct {
	Name       string
	Age        string
	Occupation string
}

func (p *Person) PrintDetails() {
	println("Name:", p.Name)
	println("Age:", p.Age)
	println("Occupation:", p.Occupation)
}

func (p *Person) isEligible() bool {
	return p.Age >= "18"
}

type Point struct {
	x int
	y int
}

func (p Point) GetX() int {
	return p.x
}

func (p Point) GetY() int {
	return p.y
}

func (p1 Point) getDist(p2 Point) float64 {
	return math.Sqrt(float64((p2.x-p1.x)*(p2.x-p1.x) + (p2.y-p1.y)*(p2.y-p1.y)))
}

func main_() {
	p := &Person{
		Name:       "John Doe",
		Age:        "30",
		Occupation: "Software Engineer",
	}
	p.PrintDetails()

	if p.isEligible() {
		println(p.Name, "is eligible.")
	} else {
		println(p.Name, "is not eligible.")
	}

	p1 := Point{x: 3, y: 4}
	p2 := Point{x: 7, y: 1}

	distance := p1.getDist(p2)
	println("Distance between points:", strconv.FormatFloat(distance, 'f', 8, 64))
}
