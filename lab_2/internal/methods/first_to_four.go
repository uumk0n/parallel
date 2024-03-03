package methods

import (
	"fmt"
	"time"

	"github.com/uumk0n/parallel/lab_2/internal/binaries"
)

func RowByCol(bin_a_name, bin_b_name string) [][]int {
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

func RowByColTransposed(bin_a_name, bin_b_name string) [][]int {
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

func RowByRow(bin_a_name, bin_b_name string) [][]int {
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

func RowByRowTransposed(bin_a_name, bin_b_name string) [][]int {
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

func TestOneToFourMethods() {

	matrixName := "matrix"

	binAName := "matrixA.bin"
	binBName := "matrixB.bin"

	start := time.Now()
	_ = RowByCol(binAName, binBName)
	elapsed := time.Since(start)
	fmt.Printf("Row by Column\t%s\t%s\n", matrixName, elapsed)

	start = time.Now()
	_ = RowByColTransposed(binAName, binBName)
	elapsed = time.Since(start)
	fmt.Printf("Row by Column (Transposed)\t%s\t%s\n", matrixName, elapsed)

	start = time.Now()
	_ = RowByRow(binAName, binBName)
	elapsed = time.Since(start)
	fmt.Printf("Row by Row\t%s\t%s\n", matrixName, elapsed)

	start = time.Now()
	_ = RowByRowTransposed(binAName, binBName)
	elapsed = time.Since(start)
	fmt.Printf("Row by Row (Transposed)\t%s\t%s\n", matrixName, elapsed)
}
