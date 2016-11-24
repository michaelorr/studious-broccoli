package pascal

func Triangle(n int) [][]int {
	res := make([][]int, n)

	for row := 0; row < n; row++ {
		res[row] = make([]int, row+1)
	}

	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			if (j == 0) || (i == j) {
				res[i][j] = 1
			} else {
				res[i][j] = res[i-1][j] + res[i-1][j-1]
			}
		}
	}
	return res
}
