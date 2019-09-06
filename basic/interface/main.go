package main

import (
	"fmt"
)

// FlyBehavior interface
type FlyBehavior interface {
	Fly()
}

// RocketFlyBehavior is fly behavior implementation with rocket
type RocketFlyBehavior struct{}

// Fly with rocket
func (r *RocketFlyBehavior) Fly() {
	fmt.Println("fly with rocket")
}

// WingsFlyBehavior is fly behavior implementation with wings
type WingsFlyBehavior struct{}

// Fly with wings
func (r *WingsFlyBehavior) Fly() {
	fmt.Println("fly with wings")
}

func main() {
	var flyBehavior FlyBehavior
	flyBehavior = &RocketFlyBehavior{}
	flyBehavior.Fly()

	flyBehavior = &WingsFlyBehavior{}
	flyBehavior.Fly()

}
