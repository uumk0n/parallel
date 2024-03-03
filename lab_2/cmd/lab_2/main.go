package main

import "github.com/uumk0n/parallel/lab_2/internal/binaries"

func main() {
	// matrixName := "matrix"

	// binAName := "matrixA.bin"
	// binBName := "matrixB.bin"

	// start := time.Now()
	// _ = methods.RowByCol(binAName, binBName)
	// elapsed := time.Since(start)
	// fmt.Printf("Row by Column\t%s\t%s\n", matrixName, elapsed)

	// start = time.Now()
	// _ = methods.RowByColTransposed(binAName, binBName)
	// elapsed = time.Since(start)
	// fmt.Printf("Row by Column (Transposed)\t%s\t%s\n", matrixName, elapsed)

	// start = time.Now()
	// _ = methods.RowByRow(binAName, binBName)
	// elapsed = time.Since(start)
	// fmt.Printf("Row by Row\t%s\t%s\n", matrixName, elapsed)

	// start = time.Now()
	// _ = methods.RowByRowTransposed(binAName, binBName)
	// elapsed = time.Since(start)
	// fmt.Printf("Row by Row (Transposed)\t%s\t%s\n", matrixName, elapsed)

	binaries.CreateBin("matrixA", 3)
	binaries.CreateBin("matrixB", 3)

	binaries.CreateBin("matrixA", 10)
	binaries.CreateBin("matrixB", 10)

	binaries.CreateBin("matrixA", 20)
	binaries.CreateBin("matrixB", 20)
}
