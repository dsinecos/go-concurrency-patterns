package main

import (
	"fmt"
	"time"

	c "github.com/dsinecos/go-generator/chanutil"
	g "github.com/dsinecos/go-generator/generator"
)

func main() {
	generateInteger := g.Integer(10)
	shutdown := make(chan int)

	outputChan := c.Pool(shutdown, generateInteger, 4, doubleNum)

	// To simulate the scenario where all the values are received
	// until the input channel is closed
	for value := range outputChan {
		fmt.Println("Value received ", value)
	}

	// To simulate the scenario where fewer values are received
	// than sent by the input channel. Here the shutdown channel
	// is used to signal to close all the outstanding goroutines
	// and prevent goroutine leaks
	for i := 0; i < 1; i++ {
		fmt.Println("Value received ", <-outputChan)
	}
	time.Sleep(2 * time.Second)
	close(shutdown)
	time.Sleep(2 * time.Second)

}

func doubleNum(num int) int {
	return num * 2
}
