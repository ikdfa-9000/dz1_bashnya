package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.Open("input.txt") // Открытие файла
	if err != nil {
		println(err)
		os.Exit(1)
	}
	fileScanner := bufio.NewScanner(input)         // Создаём сканнер
	fileScanner.Scan()                             // Сканирование файла, находим первую строку
	arrSize, _ := strconv.Atoi(fileScanner.Text()) // Первая строка - размер слайса, конвертация строки .Text() в int
	fileScanner.Scan()                             // Сканим файл ещё раз, находим вторую строку
	numbersArr := make([]int, arrSize)             // Создание слайса
	for i := 0; i < arrSize; i++ {
		add, _ := strconv.Atoi(strings.Split(fileScanner.Text(), " ")[i])
		numbersArr[i] = add
	}
	for i := 0; i < arrSize-1; i++ {
		numbersArr[i], numbersArr[i+1] = numbersArr[i+1], numbersArr[i]
	}
	for _, numberSlice := range numbersArr {
		fmt.Print(numberSlice, " ")
	}
	input.Close()
}
