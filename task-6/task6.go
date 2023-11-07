package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type emptyThing struct{}

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
	existingNumbersMap := map[string]emptyThing{}
	var b emptyThing
	for _, number := range arrNumbersString {
		if _, numberExist := existingNumbersMap[number]; !numberExist {
			existingNumbersMap[number] = b
		}
	}
	fmt.Println(len(existingNumbersMap))
}
