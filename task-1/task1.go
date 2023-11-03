package main

import (
	"fmt"
	"math"
)

func main() {
	a := 0.0
	b := 0.0
	fmt.Scan(&a)
	fmt.Scan(&b)
	a = math.Pow(a, 2)
	b = math.Pow(b, 2)
	fmt.Println(math.Pow(a+b, 0.5))

}
