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
	file_scanner := bufio.NewScanner(input)          // Создаём сканнер
	file_scanner.Scan()                              // Сканирование файла, находим первую строку
	arr_size, _ := strconv.Atoi(file_scanner.Text()) // Первая строка - размер слайса, конвертация строки .Text() в int
	file_scanner.Scan()                              // Сканим файл ещё раз, находим вторую строку
	numbers_arr := make([]int, arr_size)             // Создание слайса
	for i := 0; i < arr_size; i++ {
		add, _ := strconv.Atoi(strings.Split(file_scanner.Text(), " ")[i])
		numbers_arr[i] = add
	}
	for i := 0; i < arr_size-1; i++ {
		numbers_arr[i], numbers_arr[i+1] = numbers_arr[i+1], numbers_arr[i]
	}
	for _, numberSlice := range numbers_arr {
		fmt.Print(numberSlice, " ")
	}
	input.Close()
}
