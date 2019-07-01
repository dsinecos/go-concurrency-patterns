package main

import (
	"fmt"

	c "github.com/dsinecos/go-generator/chanutil"
	g "github.com/dsinecos/go-generator/generator"
)

func main() {
	const MAX = 10

	generateInteger := g.Integer(MAX)

	output1 := make(chan int)
	output2 := make(chan int)
	output3 := make(chan int)

	c.SplitRnd(generateInteger, output1, output2, output3)

	output4 := c.Merge(output1, output2, output3)

	for value := range output4 {
		fmt.Println("Value ", value)
	}
}
