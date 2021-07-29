package reverse_integer

import (
	"math"
)

func Reverse(x int) int {
	result := 0
	for x != 0 {
		lastDigit := x % 10
		x /= 10
		if result > (math.MaxInt32 / 10) || result == (math.MaxInt32 / 10) && lastDigit > 7 {
			return 0
		}
		if result < (math.MinInt32 / 10) || result == (math.MinInt32 / 10) && lastDigit < -8 {
			return 0
		}
		result = (result * 10) + lastDigit
	}
	return result
}
