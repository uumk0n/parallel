package main

import (
	"fmt"
	"time"

	"github.com/uumk0n/parallel/lab_2/internal/methods"
)

func main() {
	sizes := []int{3, 10, 20}

	// for _, size := range sizes {
	// 	fmt.Println("n =", size)
	// 	binaries.CreateBin("matrixA", size)
	// }
	// for _, size := range sizes {
	// 	fmt.Println("n =", size)
	// 	binaries.CreateBin("matrixB", size)
	// }
	// return
	for _, size := range sizes {
		fmt.Printf("n = %d, matrixC\n", size)

		binAName := fmt.Sprintf("matrixA%d.bin", size)
		binBName := fmt.Sprintf("matrixB%d.bin", size)

		start := time.Now()
		_ = methods.RowByCol(binAName, binBName)
		elapsed := time.Since(start)
		fmt.Printf("1)%s\n", elapsed)

		start = time.Now()
		_ = methods.RowByColTransposed(binAName, binBName)
		elapsed = time.Since(start)
		fmt.Printf("2)%s\n", elapsed)

		start = time.Now()
		_ = methods.RowByRow(binAName, binBName)
		elapsed = time.Since(start)
		fmt.Printf("3)%s\n", elapsed)

		start = time.Now()
		_ = methods.RowByRowTransposed(binAName, binBName)
		elapsed = time.Since(start)
		fmt.Printf("4)%s\n", elapsed)

		start = time.Now()
		_ = methods.ColumnByColumn(binAName, binBName)
		elapsed = time.Since(start)
		fmt.Printf("5)%s\n", elapsed)

		start = time.Now()
		_ = methods.ColumnByColumnTransposed(binAName, binBName)
		elapsed = time.Since(start)
		fmt.Printf("6)%s\n", elapsed)

		start = time.Now()
		_ = methods.ColumnByRow(binAName, binBName)
		elapsed = time.Since(start)
		fmt.Printf("7)%s\n", elapsed)

		start = time.Now()
		_ = methods.ColumnByRowTransposed(binAName, binBName)
		elapsed = time.Since(start)
		fmt.Printf("8)%s\n", elapsed)
	}

	for _, size := range sizes {
		fmt.Printf("n = %d, matrixD\n", size)

		binAName := fmt.Sprintf("matrixA%d.bin", size)
		binBName := fmt.Sprintf("matrixB%d.bin", size)

		start := time.Now()
		_ = methods.RowByCol(binBName, binAName)
		elapsed := time.Since(start)
		fmt.Printf("1)%s\n", elapsed)

		start = time.Now()
		_ = methods.RowByColTransposed(binBName, binAName)
		elapsed = time.Since(start)
		fmt.Printf("2)%s\n", elapsed)

		start = time.Now()
		_ = methods.RowByRow(binBName, binAName)
		elapsed = time.Since(start)
		fmt.Printf("3)%s\n", elapsed)

		start = time.Now()
		_ = methods.RowByRowTransposed(binBName, binAName)
		elapsed = time.Since(start)
		fmt.Printf("4)%s\n", elapsed)

		start = time.Now()
		_ = methods.ColumnByColumn(binBName, binAName)
		elapsed = time.Since(start)
		fmt.Printf("5)%s\n", elapsed)

		start = time.Now()
		_ = methods.ColumnByColumnTransposed(binBName, binAName)
		elapsed = time.Since(start)
		fmt.Printf("6)%s\n", elapsed)

		start = time.Now()
		_ = methods.ColumnByRow(binBName, binAName)
		elapsed = time.Since(start)
		fmt.Printf("7)%s\n", elapsed)

		start = time.Now()
		_ = methods.ColumnByRowTransposed(binBName, binAName)
		elapsed = time.Since(start)
		fmt.Printf("8)%s\n", elapsed)
	}
}
