package main

import (
	"fmt"

	g "github.com/dsinecos/go-generator/generator"
)

func main() {
	const MAX = 30
	isDivisibleByThree := g.IsDivisibleBy(3)
	isDivisibleByFive := g.IsDivisibleBy(5)

	integer3 := g.Integer(MAX)
	integer5 := g.Integer(MAX)

	isDivisibleByThreeOrFive := g.Merge(isDivisibleByThree(integer3), isDivisibleByFive(integer5))

	fmt.Println("Printing integers divisible by 5 or 3")
	for num := range isDivisibleByThreeOrFive {
		fmt.Println(num)
	}

}
