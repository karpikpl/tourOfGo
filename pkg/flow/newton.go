package flow

import "math"

// https://tour.golang.org/flowcontrol/8
func sqrt(x float64) float64 {

	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / 2 * z
	}

	return z
}

func sqrt2(x float64) (float64, int) {

	z := 1.0
	for i := 0; ; i++ {
		newZ := z - (z*z-x)/2*z

		if math.Abs(newZ-z) < 0.01 {
			return newZ, i
		}
		z = newZ
	}
}

func sqrt3(x float64) (float64, int) {
	return sqrtRec(1.0, 0, x)
}

func sqrtRec(z float64, iter int, x float64) (float64, int) {
	newZ := z - (z*z-x)/2*z
	iter++

	if math.Abs(newZ-z) < 0.01 {
		return newZ, iter
	}
	return sqrtRec(newZ, iter, x)
}
