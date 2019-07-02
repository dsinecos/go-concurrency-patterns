package main

import (
	"fmt"
	"time"

	c "github.com/dsinecos/go-generator/chanutil"
)

func main() {
	input1 := make(chan int)
	input2 := make(chan int)
	input3 := make(chan int)

	shutdownChan := c.OrShutdown(input1, input2, input3)

	go func() {
		time.Sleep(1 * time.Second)
		close(input1)
		fmt.Println("Closing inputChan1")
	}()

	go func() {
		time.Sleep(2 * time.Second)
		close(input2)
		fmt.Println("Closing inputChan2")

	}()

	go func() {
		time.Sleep(3 * time.Second)
		close(input3)
		fmt.Println("Closing inputChan3")

	}()

	start := time.Now()

	select {
	case <-shutdownChan:
		fmt.Println("Shutdown Channel activated")
		fmt.Println(time.Since(start))
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Exiting main goroutine")

}
