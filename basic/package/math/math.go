package math

import (
	"fmt"
)

// Sum return sum of all given numbers
// Public fucntion
func Sum(numbers ...float64) (sum float64) {
	for _, v := range numbers {
		sum += v
	}
	return
}

// Private function
func function() {
	fmt.Println("private")
}
