package q136

func Solution2(nums []int) int {
    count := map[int]int{}
    for i := 0; i < len(nums); i++ {
        count[nums[i]]++
    }
    result := 0
    for i := 0; i < len(nums); i++ {
        if count[nums[i]] < 2 {
            result = nums[i]
            break
        }
    }
    return result
}