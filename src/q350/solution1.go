package q350

func Solution1(nums1 []int, nums2 []int) []int {
    nums2Map := map[int]int{}
    for _, num := range nums2 {
        nums2Map[num]++
    }
    res := []int{}
    for _, num := range nums1 {
        if nums2Map[num] > 0 {
            res = append(res, num)
            nums2Map[num]--
        }
    }
    return res
}