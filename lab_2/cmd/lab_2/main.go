package main

import (
	"fmt"

	"github.com/uumk0n/parallel/lab_2/internal/binaries"
)

func main() {
	fmt.Println(binaries.ReadBin("matrixA.bin"))
	fmt.Println(binaries.ReadBin("matrixB.bin"))
}
