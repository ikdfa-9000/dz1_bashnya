package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	for _, symbol := range lineToChange {
		if symbol == '1' {
			fmt.Print("one")
		} else {
			fmt.Print(string(symbol))
		}
	}
}
