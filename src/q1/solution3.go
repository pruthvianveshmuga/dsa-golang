package q1

func Solution3(nums []int, target int) []int {
	valToInd := map[int]int{}
	for i := 0; i < len(nums); i++ {
		valToInd[nums[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		val := target - nums[i]
		if valToInd[val] != i {
			return []int{i, valToInd[val]}
		}
	}
	return []int{-1, -1}
}