package main

import (
	"fmt"

	"github.com/dsinecos/go-generator/generator"
)

func main() {
	fmt.Println("Hello World!")

	intChan := make(chan int)
	const MAX = 10

	generator.IsDivisibleBy(MAX, intChan, 3)

	for integer := range intChan {
		fmt.Println(integer)
	}
}
