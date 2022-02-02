package q387

func Solution1(s string) int {
	freq := map[byte]int{}
	for i := 0; i < len(s); i++ {
		freq[s[i]]++
	}
	for i := 0; i < len(s); i++ {
		if freq[s[i]] < 2 {
			return i
		}
	}
	return -1
}