package q7

import "math"

func Solution1(x int) int {
	result := 0
	for x != 0 {
		result *= 10
		result += x%10
		x /= 10
	}
	if result > math.MaxInt32 || result < math.MinInt32 {
		result = 0
	}
	return result
}