package main

import (
	"fmt"
	"time"

	c "github.com/dsinecos/go-generator/chanutil"
)

func main() {
	shutDownChan1 := make(chan int)
	shutDownChan2 := make(chan int)
	shutDownChan3 := make(chan int)

	onComplete := c.AndShutdown(shutDownChan1, shutDownChan2, shutDownChan3)

	go func(onComplete <-chan int) {
		<-onComplete
		fmt.Println("All channels closed")
	}(onComplete)

	time.Sleep(1 * time.Second)
	close(shutDownChan1)
	time.Sleep(2 * time.Second)
	close(shutDownChan2)
	time.Sleep(1 * time.Second)
	close(shutDownChan3)

	time.Sleep(2 * time.Second)
}
