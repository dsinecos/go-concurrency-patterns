package main

import (
	"fmt"

	g "github.com/dsinecos/go-generator/generator"
)

func main() {
	shutdown := make(chan int)

	repeatInt := 10

	repeatChan := g.Repeat(shutdown, repeatInt)
	takeChan := g.Take(shutdown, repeatChan, 4)

	for value := range takeChan {
		fmt.Println(value)
	}
	close(shutdown)
}
