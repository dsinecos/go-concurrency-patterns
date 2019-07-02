package main

import (
	"fmt"

	g "github.com/dsinecos/go-generator/generator"
)

func main() {
	shutdown := make(chan int)

	repeatInt := 10

	repeatChan := g.Repeat(shutdown, repeatInt)

	for i := 0; i < 4; i++ {
		fmt.Println(<-repeatChan)
	}
	close(shutdown)
}
