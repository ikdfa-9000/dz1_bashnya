package main

import "fmt"

func main() {
	a := 0
	b := 0
	c := 0
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	if a+b > c && a+c > b && c+b > a {
		println("YES")
	} else {
		println("NO")
	}
}
