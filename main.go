package main

import (
	"fmt"

	c "github.com/dsinecos/go-generator/chanutil"
	g "github.com/dsinecos/go-generator/generator"
)

func main() {
	const MAX = 100

	generateInteger := g.Integer(MAX)

	for value := range c.Pipeline(generateInteger, isDivisibleBy3, isDivisibleBy5, isEven) {
		fmt.Println(value)
	}

}

func isDivisibleBy3(task int) bool {
	if task%3 == 0 {
		return true
	}

	return false
}

func isDivisibleBy5(task int) bool {
	if task%5 == 0 {
		return true
	}

	return false
}

func isOdd(task int) bool {
	if task%2 == 0 {
		return false
	}

	return true
}

func isEven(task int) bool {
	return !isOdd(task)
}
