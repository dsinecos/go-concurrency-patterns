package main

import (
	"fmt"

	"github.com/dsinecos/go-generator/generator"
)

func main() {
	fmt.Println("Hello World!")

	const MAX = 10
	intChan := generator.Integer(MAX)
	oddChan := generator.OddInteger(MAX)
	isDivChan := generator.IsDivisibleBy(MAX, 3)

	fmt.Println("Printing integers")
	for integer := range intChan {
		fmt.Println(integer)
	}

	fmt.Println("Printing odd integers")
	for odd := range oddChan {
		fmt.Println(odd)
	}

	fmt.Println("Printing integers divisible by 3")
	for num := range isDivChan {
		fmt.Println(num)
	}
}
