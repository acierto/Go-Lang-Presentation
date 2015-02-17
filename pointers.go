package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v*Vertex) Scale(f float64) {
	v.X = v.X*f
	v.Y = v.Y*f
}

func UpdateByValue(v Vertex, f float64) {
	v.X = v.X*f
	v.Y = v.Y*f
}

func UpdateByReference(v *Vertex, f float64) {
	v.X = v.X*f
	v.Y = v.Y*f
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func main() {
	v := &Vertex{3, 4}
	v2 := Vertex{3, 4}

	fmt.Println("\nBefore modification:\n")
	fmt.Println("v  = ", *v,  v.Abs())
	fmt.Println("v2 = ", v2, v2.Abs())

	UpdateByReference(v, 5)
	UpdateByValue(v2, 5)

	fmt.Println("\nAfter modification:\n")
	fmt.Println("v  = ", *v,  v.Abs())
	fmt.Println("v2 = ", v2, v2.Abs())
}
