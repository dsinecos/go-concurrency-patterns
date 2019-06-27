package main

import (
	"fmt"

	g "github.com/dsinecos/go-generator/generator"
)

func main() {
	fmt.Println("Hello World!")

	const MAX = 10
	isDivisibleByThree := g.IsDivisibleBy(3)

	fmt.Println("Printing odd integers")
	for num := range isDivisibleByThree(g.Integer(10)) {
		fmt.Println(num)
	}
}
