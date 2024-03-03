package methods

func Row_by_col(matrix_a, matrix_b [][]int32) [][]int32 {
	n := len(matrix_a)
	matrix_c := make([][]int32, n)
	for i := range matrix_c {
		matrix_c[i] = make([]int32, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				matrix_c[i][j] += matrix_a[i][k] * matrix_b[k][j]
			}
		}
	}
	return matrix_c
}

func Row_by_col_transposed(matrix_a, matrix_b [][]int32) [][]int32 {
	n := len(matrix_a)
	matrix_c := make([][]int32, n)
	for i := range matrix_c {
		matrix_c[i] = make([]int32, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				matrix_c[i][j] += matrix_a[i][k] * matrix_b[j][k]
			}
		}
	}
	return matrix_c
}

func Row_by_row(matrix_a, matrix_b [][]int32) [][]int32 {
	n := len(matrix_a)
	matrix_c := make([][]int32, n)
	for i := range matrix_c {
		matrix_c[i] = make([]int32, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				matrix_c[i][j] += matrix_a[i][k] * matrix_b[k][j]
			}
		}
	}
	return matrix_c
}

func Row_by_row_transposed(matrix_a, matrix_b [][]int32) [][]int32 {
	n := len(matrix_a)
	matrix_c := make([][]int32, n)
	for i := range matrix_c {
		matrix_c[i] = make([]int32, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				matrix_c[i][j] += matrix_a[i][k] * matrix_b[j][k]
			}
		}
	}
	return matrix_c
}
