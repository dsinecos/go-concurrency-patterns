package main

import (
	"fmt"

	g "github.com/dsinecos/go-generator/generator"
)

func main() {
	fmt.Println("Hello World!")

	const MAX = 30
	isDivisibleByThree := g.IsDivisibleBy(3)
	isDivisibleByFive := g.IsDivisibleBy(5)

	fmt.Println("Printing integers divisible by 5 and 3")
	for num := range isDivisibleByFive(isDivisibleByThree(g.Integer(MAX))) {
		fmt.Println(num)
	}
}
