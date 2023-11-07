package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	lineToChange := fileScanner.Text()
	input.Close()
	lineToChange = strings.Replace(lineToChange, "1", "one", -1)
	fmt.Println(lineToChange)
}
