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
	fmt.Println(math.Hypot(a, b))
}
