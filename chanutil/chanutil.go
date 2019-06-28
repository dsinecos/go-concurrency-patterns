package chanutil

import "sync"

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
