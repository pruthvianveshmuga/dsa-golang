package q283

func Solution2(nums []int) {
	zeroIndex := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 && zeroIndex == -1 {
			zeroIndex = i
		} else if nums[i] != 0 && zeroIndex != -1 {
			nums[zeroIndex], nums[i] =  nums[i], nums[zeroIndex]
				zeroIndex++;
		}
	}
}