package main

import (
	"errors"
	"fmt"
)

func add(a int, b int) int {
	return a + b
}

func sub(a, b int) (result int) {
	result = a - b
	return
}

func div(a, b float32) (float32, error) {
	if b == 0 {
		return 0, errors.New("divided by zero")
	}
	return a / b, nil
}

func main() {
	a := add(1, 2)
	fmt.Println(a)

	b := sub(3, 2)
	fmt.Println(b)

	c, err := div(4, 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(c)

}
