package q350

import "sort"

func Solution2(nums1 []int, nums2 []int) []int {
    res := []int{}
    sort.Slice(nums1, func(i, j int) bool {
        return nums1[i] < nums1[j]
    })
    sort.Slice(nums2, func(i, j int) bool {
        return nums2[i] < nums2[j]
    })
    i, j := 0, 0
    for i < len(nums1) && j < len(nums2) {
        if nums1[i] == nums2[j] {
            res = append(res, nums1[i])
            i++
            j++
        } else if nums1[i] < nums2[j] {
            i++
        } else {
            j++
        }
    }
    return res
}