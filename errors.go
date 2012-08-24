package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

func Sqrt(f float64) (float64, error) {
	if f <= 0 {
		return 0, ErrNegativeSqrt(f)
	}

	last_z, z := f, 1.0

	for math.Abs(z-last_z) >= 1.0e-6 {
		last_z, z = z, z-(z*z-f)/(2*z)
	}

	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
