package q1

func Solution4(nums []int, target int) []int {
	valToInd := map[int]int{}
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		if _, exists := valToInd[complement]; exists {
			return []int{i, valToInd[complement]}
		}
		valToInd[nums[i]] = i
	}
	return []int{-1, -1}
}