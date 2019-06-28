package generator

import "sync"

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

// Merge TODO
func Merge(input ...chan int) chan int {

	out := make(chan int)
	var shutdownSignal sync.WaitGroup

	for _, inputChan := range input {

		shutdownSignal.Add(1)
		go func(inputChan chan int) {
			defer shutdownSignal.Done()
			for num := range inputChan {
				out <- num
			}
		}(inputChan)
	}

	/*
		A goroutine will be responsible for monitoring when all upstream
		channels. When all the upstream channels are closed, it'll signal
		the downstream channel
	*/
	go func() {
		shutdownSignal.Wait()
		close(out)
	}()

	return out
}
