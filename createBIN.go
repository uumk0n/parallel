package main

import (
	"encoding/binary"
	"math/rand"
	"os"
	"time"
)

func CreateBin(filename string) {
	// Задаем размерность матрицы и диапазон случайных чисел
	const n = 5
	const min = 0
	const max = 100

	// Создаем файл для записи
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Создаем генератор случайных чисел
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// Записываем размерность матрицы в файл
	err = binary.Write(file, binary.LittleEndian, int32(n))
	if err != nil {
		panic(err)
	}

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
