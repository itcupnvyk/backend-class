package main

import (
	"fmt"
)

func main() {
	// Default value int = 0
	var age1 int
	var age2 int32
	var age3 int64
	// Default value float = 0.0
	var price1 float32
	var price2 float64
	// Default value bool = false
	var isFound bool
	// Default value string = "" -> empty string
	var name string

	fmt.Println("age1:", age1)
	fmt.Println("age2:", age2)
	fmt.Println("age3:", age3)
	fmt.Println("price1:", price1)
	fmt.Println("price2:", price2)
	fmt.Println("isFound:", isFound)
	fmt.Println("name:", name)
}
