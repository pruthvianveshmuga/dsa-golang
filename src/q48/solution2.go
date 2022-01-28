package q48

import intUtils "github.com/pruthvianveshmuga/dsa/utils/int"

func Solution2(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2 + n%2; i++ {
		for j := 0; j < n/2; j++ {
			intUtils.Swap(&matrix[i][j], &matrix[n-j-1][i])
			intUtils.Swap(&matrix[n-j-1][i], &matrix[n-i-1][n-j-1])
			intUtils.Swap(&matrix[n-i-1][n-j-1], &matrix[j][n-i-1])
		}
	}
}