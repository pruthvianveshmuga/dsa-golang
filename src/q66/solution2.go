package q66

func Solution2(digits []int) []int {
		i := len(digits)-1
		for ;i >= 0 && digits[i] == 9; i-- {
			digits[i] = 0
		}
		if i >= 0 {
			digits[i]++
		} else {
			digits = append([]int{1}, digits...)
		}
		return digits
}