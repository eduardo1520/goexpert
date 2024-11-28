package tax

import (
	"math"
	"time"
)

func CalculateTax(amount float64) float64 {
	if amount <= 0 {
		return 0
	}
	if amount > 1000 {
		return 10.0
	}
	return 5.0
}

func isZero(value float64) bool {
	const epsilon = 1e-9
	return math.Abs(value) < epsilon
}

func CalculateTax2(amount float64) float64 {
	time.Sleep(time.Millisecond)
	if amount == 0 {
		return 0
	}
	if amount > 1000 {
		return 10.0
	}
	return 5.0
}
