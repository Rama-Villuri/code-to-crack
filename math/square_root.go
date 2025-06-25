// Package math provides basic mathematical functions.
package math

import (
	"fmt"
	"math"
)

// SquareRoot returns the square root of a given number x.
func SquareRoot(x float64) float64 {
	fmt.Println("Calculating square root of", x)
	return math.Sqrt(x)
}
