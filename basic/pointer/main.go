package main

import (
	"fmt"
)

func change(number *int) {
	*number = 99
}

func main() {
	var p *int

	fmt.Println(p)
	p = new(int)
	fmt.Println(p)
	*p = 10
	fmt.Println(*p)
	fmt.Println(&p)

	number := 10
	change(&number)
	fmt.Println(number)
}
