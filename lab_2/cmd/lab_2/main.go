package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

func main() {
	fileA, err := os.Open("matrixA.bin")
	if err != nil {
		panic(err)
	}
	defer fileA.Close()

	fileB, err := os.Open("matrixB.bin")
	if err != nil {
		panic(err)
	}
	defer fileB.Close()

	resultFile, err := os.Create("result.bin")
	if err != nil {
		panic(err)
	}
	defer resultFile.Close()

	var nA int32
	err = binary.Read(fileA, binary.LittleEndian, &nA)
	if err != nil {
		panic(err)
	}

	var nB int32
	err = binary.Read(fileB, binary.LittleEndian, &nB)
	if err != nil {
		panic(err)
	}

	if nA != nB {
		panic("Размерности матриц не совпадают")
	}

	for i := int32(0); i < nA; i++ {
		for j := int32(0); j < nA; j++ {
			var numA, numB int32
			err := binary.Read(fileA, binary.LittleEndian, &numA)
			if err != nil {
				panic(err)
			}
			_, err = fileB.Seek(int64((i*nA+j)*4), 0)
			if err != nil {
				panic(err)
			}
			err = binary.Read(fileB, binary.LittleEndian, &numB)
			if err != nil {
				panic(err)
			}
			result := numA * numB
			err = binary.Write(resultFile, binary.LittleEndian, result)
			if err != nil {
				panic(err)
			}
			err = binary.Write(resultFile, binary.LittleEndian, int32(' '))
			if err != nil {
				panic(err)
			}
		}
		err := binary.Write(resultFile, binary.LittleEndian, int32('\n'))
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Умножение матриц успешно завершено. Результат записан в файл result.txt")
}
