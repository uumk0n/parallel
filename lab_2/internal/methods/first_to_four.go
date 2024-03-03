package methods

import "github.com/uumk0n/parallel/lab_2/internal/binaries"

func Row_by_col(bin_a_name, bin_b_name string) [][]int {
	matrix_a := binaries.ReadBin(bin_a_name)
	matrix_b := binaries.ReadBin(bin_b_name)
	n := len(matrix_a)
	matrix_c := make([][]int, n)
	for i := range matrix_c {
		matrix_c[i] = make([]int, n)
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

func Row_by_col_transposed(bin_a_name, bin_b_name string) [][]int {
	matrix_a := binaries.ReadBin(bin_a_name)
	matrix_b := binaries.ReadBin(bin_b_name)
	n := len(matrix_a)
	matrix_c := make([][]int, n)
	for i := range matrix_c {
		matrix_c[i] = make([]int, n)
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

func Row_by_row(bin_a_name, bin_b_name string) [][]int {
	matrix_a := binaries.ReadBin(bin_a_name)
	matrix_b := binaries.ReadBin(bin_b_name)
	n := len(matrix_a)
	matrix_c := make([][]int, n)
	for i := range matrix_c {
		matrix_c[i] = make([]int, n)
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

func Row_by_row_transposed(bin_a_name, bin_b_name string) [][]int {
	matrix_a := binaries.ReadBin(bin_a_name)
	matrix_b := binaries.ReadBin(bin_b_name)
	n := len(matrix_a)
	matrix_c := make([][]int, n)
	for i := range matrix_c {
		matrix_c[i] = make([]int, n)
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
