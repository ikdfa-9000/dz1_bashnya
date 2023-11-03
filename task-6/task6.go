package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	// "strconv"
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
	arrNumbersString := strings.Split(fileScanner.Text(), " ") // Получили слайс из всех чисел в формате string.
	fmt.Println(arrNumbersString)
	input.Close()
	existingNumbersMap := map[string]int{}
	uniqueNumCount := 0
	var b int
	for _, number := range arrNumbersString {
		if _, numberExist := existingNumbersMap[number]; !numberExist {
			existingNumbersMap[number] = b
			uniqueNumCount++
		}
	}
	fmt.Println(uniqueNumCount)
}
