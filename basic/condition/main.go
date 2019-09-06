package main

import (
	"fmt"
)

func main() {
	number := 1

	if number == 1 {
		fmt.Println("one")
	} else if number == 2 {
		fmt.Println("two")
	} else {
		fmt.Println("not one or two")
	}

	age := 20
	switch age {
	case 20:
		fmt.Println("age: 20")
	case 10:
		fmt.Println("age: 10")
	default:
		fmt.Println("default")
	}

	switch {
	case age < 10:
		fmt.Println("young")
	case age >= 10 && age < 18:
		fmt.Println("teenager")
	default:
		fmt.Println("default")
	}
}
