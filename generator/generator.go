package generator

import "fmt"

// Integer TODO
func Integer(max int) chan int {

	out := make(chan int)

	go func() {
		defer close(out)
		for i := 1; i <= max; i++ {
			out <- i
		}
	}()

	return out
}

// OddInteger TODO
func OddInteger(input chan int) chan int {

	out := make(chan int)

	go func() {
		defer close(out)

		for integer := range input {
			if integer%2 != 0 {
				out <- integer
			}
		}
	}()

	return out
}

// IsDivisibleBy TODO
func IsDivisibleBy(divisor int) func(chan int) chan int {
	return func(input chan int) chan int {
		out := make(chan int)

		go func() {
			defer close(out)

			for integer := range input {
				if integer%divisor == 0 {
					out <- integer
				}
			}
		}()

		return out
	}
}

// Repeat TODO
func Repeat(shutdown <-chan int, input int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for {
			select {
			case <-shutdown:
				return
			case out <- input:
			}
		}

	}()

	return out
}

// Take TODO
func Take(shutdown <-chan int, input <-chan int, size int) <-chan int {

	out := make(chan int)

	go func() {
		defer close(out)
		defer fmt.Println("Exiting Take goroutine")

		for i := 0; i < size; i++ {
			select {
			case <-shutdown:
				fmt.Println("Shutdown channel activated")
				return
			case num := <-input:
				fmt.Println("Blocked at num ", num)
				out <- num
			}
		}

	}()

	return out

}

// BatchToStream TODO
func BatchToStream(shutdown <-chan int, inputs []int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for _, num := range inputs {
			select {
			case <-shutdown:
				return
			case out <- num:
			}
		}

	}()

	return out
}
