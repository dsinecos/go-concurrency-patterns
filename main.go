package main

import (
	"fmt"
	"time"

	g "github.com/dsinecos/go-generator/generator"
)

func main() {
	generateInteger := g.Integer(20)

	shutdown := make(chan int)

	takeChan := g.Take(shutdown, generateInteger, 10)

	for i := 0; i < 1; i++ {
		fmt.Println(<-takeChan)
	}
	time.Sleep(2 * time.Second)

	close(shutdown)

	time.Sleep(2 * time.Second)
}
