package main

import (
	"fmt"
)

func A() {
	fmt.Println(" this is a function A")
}

func B() {
	fmt.Println(" this is a function B")
}

func main() {
	defer A()
	defer B()

	//program/this function ends here  
	//deferred functions executed in LIFO (last in first out) order 
}
