package byte

func Swap(a *byte, b *byte) {
	temp := *a
	*a = *b
	*b = temp
}

func DeepEqual(a, b []byte) bool {
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

func ToInt(val byte) int {
	if val == '.' {
		return 0
	} else {
		return int(val-'0')
	}
}