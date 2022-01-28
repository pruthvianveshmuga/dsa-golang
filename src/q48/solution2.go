package q48

import "github.com/pruthvianveshmuga/dsa/utils"

func Solution2(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2 + n%2; i++ {
		for j := 0; j < n/2; j++ {
			utils.Swap(&matrix[i][j], &matrix[n-j-1][i])
			utils.Swap(&matrix[n-j-1][i], &matrix[n-i-1][n-j-1])
			utils.Swap(&matrix[n-i-1][n-j-1], &matrix[j][n-i-1])
		}
	}
}