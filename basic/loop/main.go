package main

import (
	"fmt"
)

func main() {
	// for loop style
	count := 10
	for i := 0; i < count; i++ {
		fmt.Println("loopOne:", i)
	}

	// while loop style
	i := 0
	for i < count {
		fmt.Println("loopTwo:", i)
		i++
	}
}
