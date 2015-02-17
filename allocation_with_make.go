package main

import (
	"fmt"
)

func main() {
	var v  []int = make([]int, 5, 10)

	fmt.Printf("Capacity before is %d\n", cap(v))
	fmt.Printf("Length before is %d\n", len(v))

	fmt.Printf("Initial value of v[0] is %d\n", v[0])

	for i := 0; i < 6; i++ {
		v = append(v, i)
	}

	fmt.Printf("Capacity after is %d\n", cap(v))
	fmt.Printf("Length after is %d\n", len(v))
}
