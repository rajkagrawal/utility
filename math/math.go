package math

import (
	"math"
)

func AbsoluteDiffInt(a, b int ) int {
	return int(math.Abs(float64(a-b)))
}

func DigitLen(a int) int {
	return 	int(math.Log10(float64(a)))+1
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

