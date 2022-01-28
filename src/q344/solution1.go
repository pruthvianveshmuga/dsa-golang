package q344

import byteUtils "github.com/pruthvianveshmuga/dsa/utils/byte"

func Solution1(s []byte) {
	n := len(s)
	for i := 0; i < n/2; i++ {
		byteUtils.Swap(&s[i], &s[n-i-1])
	}
}