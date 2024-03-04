package binaries

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func CreateBin(filename string, n int) {
	// Задаем размерность матрицы и диапазон случайных чисел
	const min = 0
	const max = 100

	// Создаем файл для записи
	file, err := os.Create(fmt.Sprintf("%s%d.bin", filename, n))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = binary.Write(file, binary.LittleEndian, int32(n))
	if err != nil {
		panic(err)
	}

	// Создаем генератор случайных чисел
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// Записываем саму матрицу в файл
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			num := rand.Intn(max-min+1) + min // Генерируем случайное число в заданном диапазоне
			err := binary.Write(file, binary.LittleEndian, int32(num))
			if err != nil {
				panic(err)
			}
		}
	}
}

func ReadBin(filename string) [][]int {
	// Открываем файл для чтения
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Читаем размерность матрицы из файла
	var n int32
	err = binary.Read(file, binary.LittleEndian, &n)
	if err != nil {
		panic(err)
	}

	// Создаем матрицу для хранения данных из файла
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	// Считываем данные из файла в матрицу
	for i := 0; i < int(n); i++ {
		for j := 0; j < int(n); j++ {
			var num int32
			err := binary.Read(file, binary.LittleEndian, &num)
			if err != nil {
				panic(err)
			}
			matrix[i][j] = int(num)
		}
	}

	return matrix
}
