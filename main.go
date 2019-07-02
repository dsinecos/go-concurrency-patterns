package main

import (
	"fmt"

	g "github.com/dsinecos/go-generator/generator"
)

func main() {
	shutdown := make(chan int)

	inputInts := []int{1, 2, 3, 4, 5, 6}

	for num := range g.BatchToStream(shutdown, inputInts) {
		fmt.Println(num)
	}

	inputChan := g.BatchToStream(shutdown, inputInts)

	for i := 0; i < 3; i++ {
		fmt.Println(<-inputChan)

	}
	close(shutdown)
}
