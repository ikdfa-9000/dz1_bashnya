package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.Open("input.txt") // Открытие файла
	if err != nil {
		println(err)
		log.Fatal("Не удалось открыть файл")
	}
	fileScanner := bufio.NewScanner(input)
	fileScanner.Scan()
	input.Close()
	matrixSize, _ := strconv.Atoi(fileScanner.Text())
	matrix := make([][]int, matrixSize) // Создание двумерного слайса
	for i := range matrix {
		matrix[i] = make([]int, matrixSize) // Заполнение двумерного слайса
	}
	for i := 0; i < matrixSize; i++ {
		fileScanner.Scan() // Вызываем .Scan() для каждой новой строчки
		for j := 0; j < matrixSize; j++ {
			add, _ := strconv.Atoi(strings.Split(fileScanner.Text(), " ")[j]) // Конвертация строки в число
			matrix[i][j] = add
		}
	}
	for i := 0; i < matrixSize; i++ {
		for j := i + 1; j < matrixSize-i; j++ {
			if matrix[i][j] != matrix[j][i] {
				fmt.Println("NO")
				return // Досрочно завершаем программу return-ом
			}
		}
	}
	fmt.Println("YES")
}
