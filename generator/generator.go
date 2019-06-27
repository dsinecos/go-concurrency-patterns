package generator

// Integer TODO
func Integer(max int, out chan<- int) {

	go func() {
		defer close(out)
		for i := 1; i <= max; i++ {
			out <- i
		}
	}()

}

// OddInteger To fetch odd numbers equal to or less than 'max'
func OddInteger(max int, out chan<- int) {

	intChan := make(chan int)

	go func() {
		defer close(out)
		Integer(max, intChan)

		for integer := range intChan {
			if integer%2 != 0 {
				out <- integer
			}
		}
	}()

}

// IsDivisibleBy TODO
func IsDivisibleBy(max int, out chan<- int, divisor int) {

	intChan := make(chan int)

	go func() {
		defer close(out)

		Integer(max, intChan)
		for integer := range intChan {
			if integer%divisor == 0 {
				out <- integer
			}
		}
	}()

}
