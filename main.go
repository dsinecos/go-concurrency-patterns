package main

import (
	"fmt"

	c "github.com/dsinecos/go-generator/chanutil"
	g "github.com/dsinecos/go-generator/generator"
)

func main() {
	const MAX = 10
	isDivisibleByThree := g.IsDivisibleBy(3)
	isDivisibleByFive := g.IsDivisibleBy(5)

	integer3Input := make(chan int)
	integer5Input := make(chan int)

	c.Split(g.Integer(MAX), integer3Input, integer5Input)

	// integer3 := g.Integer(MAX)
	// integer5 := g.Integer(MAX)

	isDivisibleByThreeOrFive := c.Merge(isDivisibleByThree(integer3Input), isDivisibleByFive(integer5Input))

	fmt.Println("Printing integers divisible by 5 or 3")
	for num := range isDivisibleByThreeOrFive {
		fmt.Println(num)
	}

}
