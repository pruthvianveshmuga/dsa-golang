package q283

func Solution1(nums []int) {
	zero, nonZero := 0, 0
	for zero < len(nums) && nonZero < len(nums) {
		for ;nonZero < len(nums) && nums[nonZero] == 0; nonZero++ {}
		for ;zero < len(nums) && nums[zero] != 0; zero++ {}
		if zero < len(nums) && nonZero < len(nums) && zero < nonZero {
			nums[zero], nums[nonZero] = nums[nonZero], nums[zero]
			zero++
		}
		nonZero++
	}
}