package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x / 2
	for i := 0; i < 1000; i++ {
		prev := z
		z = z - (z*z-x)/(2*z)
		if abs(z-prev) < 1e-10 {
			break
		}
	}
	return z
}

func abs(a float64) float64 {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	fmt.Println(Sqrt(2) == math.Sqrt(2))
	fmt.Println(Sqrt(4), math.Sqrt(4))
	fmt.Println(Sqrt(9), math.Sqrt(9))
}
