package q344

func Solution1(s []byte) {
	n := len(s)
	for i := 0; i < n/2; i++ {
		Swap(&s[i], &s[n-i-1])
	}
}