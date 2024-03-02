package main

import (
	"github.com/uumk0n/parallel/lab_2/internal/binaries"
)

func main() {
	binaries.CreateBin("matrixA.bin")
	binaries.CreateBin("matrixB.bin")

	// fmt.Println(binaries.ReadBin("matrixA.bin"))
	// fmt.Println(binaries.ReadBin("matrixB.bin"))
}
