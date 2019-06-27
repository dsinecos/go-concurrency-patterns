package main

import (
	"fmt"

	"github.com/dsinecos/go-generator/generator"
)

func main() {
	fmt.Println("Hello World!")

	intChan := make(chan int)
	const MAX = 10

	generator.OddInteger(MAX, intChan)

	for integer := range intChan {
		fmt.Println(integer)
	}
}
