package main

import (
	"fmt"
)

func main() {
	v := new(int)

	fmt.Printf("Value is : %d \n", *v)
	fmt.Printf("Address of value is : %d", v)
}
