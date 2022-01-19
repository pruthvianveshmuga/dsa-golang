package q136

import "sort"

func Solution3(nums []int) int {
    sort.Slice(nums, func(i, j int) bool {
        return nums[i] < nums[j]
    })
    result := 0
    for i := 0; i < len(nums); i += 2 {
        if i == len(nums)-1 || nums[i] != nums[i+1] {
            result = nums[i]
            break
        }
    }
    return result
}