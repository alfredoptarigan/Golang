package main

import (
	"fmt"
	"pertama/calc"
)

func main() {
	fmt.Println("Halo Belajar Golang")

	add := calc.Add(9, 9)
	multiply := calc.Multiply(6, 6)

	fmt.Println(add)
	fmt.Println(multiply)
}
