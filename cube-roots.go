package main

import (
	"fmt"
	"math/cmplx"
)

func Cbrt(x complex128) complex128 {
	last_z, z := x, complex128(1.0)

	for cmplx.Abs(last_z-z) >= 1.0e-6 {
		last_z, z = z, z-(z*z*z-x)/(3*z*z)
	}

	return z
}

func main() {
	fmt.Println("own Cbrt(2):\t", Cbrt(2))
	fmt.Println("cmplx.Pow(2, -1.0/3.0):\t", cmplx.Pow(2, 1.0/3.0))
}
