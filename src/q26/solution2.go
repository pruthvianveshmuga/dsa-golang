package q26

func Solution2(nums []int) int {
    k := 0
    for i := 0; i < len(nums); i++ {
        if i == 0 || nums[i] != nums[i-1] {
            nums[k] = nums[i]
            k++
        }
    }
    return k
}