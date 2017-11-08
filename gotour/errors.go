package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt struct{
	What float64
}

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e.What))
}

func ESqrt(x float64) (float64, error) {
	if x < 0 {
		var v = ErrNegativeSqrt{x}
		return x, v
	} else {
		z := 1.0
		for eps:=0.00001; math.Abs(z * z - x) > eps; {
			z -= (z*z - x) / (2*z)
		}
		return z, nil
	}
}

func main() {
	fmt.Println(ESqrt(2))
	fmt.Println(ESqrt(-2))
}
