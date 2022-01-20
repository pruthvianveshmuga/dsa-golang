package q1

func Solution2(nums []int, target int) []int {
	valToInd := map[int][]int{}
	for i := 0; i < len(nums); i++ {
		valToInd[nums[i]] = append(valToInd[nums[i]], i)
	}
	for i := 0; i < len(nums); i++ {
		val := target - nums[i]
		for j := 0; j < len(valToInd[val]); j++ {
			if valToInd[val][j] != i {
				return []int{i, valToInd[val][j]}
			}
		}
	}
	return []int{-1, -1}
}