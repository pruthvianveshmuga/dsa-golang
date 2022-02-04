package q242

import "sort"

func Solution1(s string, t string) bool {
	  s_r := []rune(s)
		s_t := []rune(t)
		
    sort.Slice(s_r, func(i, j int) bool {
			return s_r[i] < s_r[j]
		})
		sort.Slice(s_t, func(i, j int) bool {
			return s_t[i] < s_t[j]
		})
		return string(s_r) == string(s_t)
}