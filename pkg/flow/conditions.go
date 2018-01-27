package flow

import "math/rand"

func ifStatement(count int) (float64, int) {
	if r := rand.Float64(); r > 0.9 {
		return r, count
	}
	return ifStatement(count + 1)
}
