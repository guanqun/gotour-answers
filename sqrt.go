package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	last_z, z := x, 1.0

	for math.Abs(z - last_z) >= 1.0e-6 {
		last_z, z = z, z - (z * z - x) / (2 * z);
	}

	return z;
}

func main() {
	fmt.Println("own Sqrt(2):\t", Sqrt(2))
	fmt.Println("math.Sqrt(2):\t", math.Sqrt(2))
}
