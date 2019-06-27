package main

import (
	"fmt"

	"github.com/dsinecos/go-generator/generator"
)

func main() {
	fmt.Println("Hello World!")

	const MAX = 10
	oddChan := generator.OddInteger(generator.Integer(MAX))

	fmt.Println("Printing odd integers")
	for odd := range oddChan {
		fmt.Println(odd)
	}
}
