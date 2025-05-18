package main

import (
	"fmt"
)

func GenDisplaceFn(a float64, initialVelocity float64, initialDisplacement float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*t*t + initialVelocity*t + initialDisplacement
	}
}

func maissn() {
	var a, v0, s0, t float64

	fmt.Print("Enter acceleration: ")
	fmt.Scan(&a)

	fmt.Print("Enter initial velocity: ")
	fmt.Scan(&v0)

	fmt.Print("Enter initial displacement: ")
	fmt.Scan(&s0)

	fmt.Print("Enter time: ")
	fmt.Scan(&t)

	fn := GenDisplaceFn(a, v0, s0)

	displacement := fn(t)
	fmt.Printf("Displacement after %.2f seconds: %.4f\n", t, displacement)
}
