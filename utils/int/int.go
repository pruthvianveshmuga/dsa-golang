package int

import "sort"

func DeepEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func DeepEqual2(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if !DeepEqual(v, b[i]) {
			return false
		}
	}
	return true
}

func DeepEqualWithoutOrder(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func Swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}