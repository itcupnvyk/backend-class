package main

import (
	"fmt"
)

// User represent user entity
type User struct {
	Name    string
	Age     int
	ammount float64
}

// GetAmmount return ammount
// method with pointer receiver
func (u *User) GetAmmount() float64 {
	return u.ammount
}

// Deposit saldo
func (u *User) Deposit(ammount float64) {
	u.ammount += ammount
}

func main() {
	u := &User{}
	u.Deposit(10000)

	ammount := u.GetAmmount()
	fmt.Println(ammount)
}
