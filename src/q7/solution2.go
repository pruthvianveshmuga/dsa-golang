package q7

import "math"

func Solution2(x int) int {
	result := 0
	for x != 0 {
		if (result <= (math.MaxInt32 - x%10)/10) && (result >= (math.MinInt32 + x%10)/10) {
			result *= 10
			result += x%10
			x /= 10
		} else {
			return 0
		}
	}
	return result
}