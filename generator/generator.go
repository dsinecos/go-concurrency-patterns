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
