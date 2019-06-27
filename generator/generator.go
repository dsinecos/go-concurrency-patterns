package generator

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

// OddInteger To fetch odd numbers equal to or less than 'max'
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
