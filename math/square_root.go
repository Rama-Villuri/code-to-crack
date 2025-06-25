// Package math provides basic mathematical functions.
package math

import (
	"fmt"
	"math"
)

// SquareRoot returns the square root of a given number x.
func SquareRoot(x float64) float64 {
	fmt.Println("Calculating square root of", x)
	if x < 0 {
		panic("cannot calculate square root of a negative number")
	}
	if x == 0 {
		return 0
	}
	return math.Sqrt(x)
}
