package main

import (
	"fmt"
)

// User represent user entity
type User struct {
	Name string
	Age  int
}

func main() {
	userOne := User{
		Name: "name",
	}
	fmt.Println(userOne, userOne.Name, userOne.Age)

	userTwo := User{}
	userTwo.Age = 10
	userTwo.Name = "name"
	fmt.Println(userTwo, userTwo.Name, userTwo.Age)

	userThree := User{"name", 11}
	fmt.Println(userThree, userThree.Name, userThree.Age)
}
