package q36

import byteUtils "github.com/pruthvianveshmuga/dsa/utils/byte"

func isValid(nums []int) bool {
	hash := map[int]bool{}
	for i := 0; i < len(nums); i++ {
		if hash[nums[i]] {
			return false
		}
		hash[nums[i]] = true
	}
	return true
}

func Solution1(board [][]byte) bool {
	n := len(board)
	//rows
	for i := 0; i < n; i++ {
		nums := []int{}
		for j := 0; j < n; j++ {
			if val := byteUtils.ToInt(board[i][j]); val != 0 {
				nums = append(nums, val)
			}
		}
		if !isValid(nums) {
			return false
		}
	}
	//cols
	for j := 0; j < n; j++ {
		nums := []int{}
		for i := 0; i < n; i++ {
			if val := byteUtils.ToInt(board[i][j]); val != 0 {
				nums = append(nums, val)
			}
		}
		if !isValid(nums) {
			return false
		}
	}
	//boxes
	w := 3
	for x := 0; x < n/w; x++ {
		for y := 0; y < n/w; y++ {
			nums := []int{}
			for i := 0; i < w; i++ {
				for j := 0; j < w; j++ {
					if val := byteUtils.ToInt(board[(x*w)+i][(y*w)+j]); val != 0 {
						nums = append(nums, val)
					}
				}
			}
			if !isValid(nums) {
				return false
			}
		}
	}
	return true
}