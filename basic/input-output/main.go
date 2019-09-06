package main

import (
	"fmt"
)

func main() {
	var name string

	fmt.Print("Nama : ")
	fmt.Scanln(&name)
	fmt.Println("Hello,", name)
}
