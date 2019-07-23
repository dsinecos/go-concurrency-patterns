package chanutil

import (
	"reflect"
	"sync"
)

// Merge combines input from multiple channels into a single channel
func Merge(input ...chan int) chan int {

	out := make(chan int)
	// To synchronize closing the output channel when all input channels
	// are closed
	var shutdownSignal sync.WaitGroup

	// For each input channel a goroutine is launched which reads from the
	// input channel and publishes to the output channel
	for _, inputChan := range input {

		// Waitgroup is incremented for each goroutine launched (to read from assigned input channel)
		shutdownSignal.Add(1)

		// Goroutine reads from assigned input channel and writes to output channel.
		// Goroutine takes channel as an argument instead of accessing via closures to avoid
		// data race conditions, since goroutine is launched from within a for loop
		// [Refer for race conditions inside for loops in Go](https://dsinecos.github.io/blog/Asynchronous-execution-inside-for-loops)
		go func(inputChan chan int) {
			// Decrement the Waitgroup once goroutine exits
			defer shutdownSignal.Done()

			// Block execution of the goroutine.
			// Read from the assigned input channel.
			// Write to the output channel
			for num := range inputChan {
				out <- num
			}
		}(inputChan)
	}

	// A goroutine monitors when all upstream channels close.
	// When all the upstream channels are closed, it'll close
	// the downstream channel
	go func() {
		// Block until all the goroutines launched to read from input channels have
		// been closed
		shutdownSignal.Wait()
		close(out)
	}()

	return out
}

// Split TODO
func Split(input chan int, outputs ...chan int) {

	go func() {

		/*
			Create n WaitGroup, each corresponding to goroutines spawned for
			a single output channel
		*/
		syncGoroutines := make([]sync.WaitGroup, len(outputs))

		for index, outputChan := range outputs {
			syncGoroutines[index] = sync.WaitGroup{}
			syncGoroutines[index].Add(1)

			/*
				The following goroutine monitors if all the goroutines spawned to publish
				to the respective output channel have closed using 'wg' and the 'input'
				channel has been closed, to close the 'output' channel
			*/
			go func(outputChan chan int, wg *sync.WaitGroup) {
				wg.Wait()
				close(outputChan)
			}(outputChan, &syncGoroutines[index])

		}

		for inputVal := range input {
			for index, outputChan := range outputs {

				syncGoroutines[index].Add(1)

				go func(inputVal int, outputChan chan int, wg *sync.WaitGroup) {
					outputChan <- inputVal
					wg.Done()
				}(inputVal, outputChan, &syncGoroutines[index])

			}
		}

		/*
			Since the earlier for-range loop is blocking and will only release
			once the 'input' channel is closed. We can define code here to be
			executed once the 'input' channel is closed
		*/
		for index := range syncGoroutines {
			syncGoroutines[index].Done()
		}

	}()
}

// SplitRnd TODO
func SplitRnd(input chan int, outputs ...chan int) {

	for _, outputChan := range outputs {

		go func(outputChan chan int) {

			defer close(outputChan)

			for num := range input {
				outputChan <- num
			}

		}(outputChan)
	}

}

// Pipeline TODO
func Pipeline(input chan int, filters ...func(task int) bool) chan int {

	out := make(chan int)

	filter := filters[0]

	go func(out chan int, filter func(task int) bool) {
		defer close(out)
		for value := range input {
			if filter(value) {
				out <- value
			}
		}
	}(out, filter)

	if len(filters) != 1 {
		filters = filters[1:]
		return Pipeline(out, filters...)
	}

	return out
}

/*
OrShutdown combines multiple signalling channels and returns a
single signalling channel (The output channel is closed if any
of the input signalling channels is closed)
*/
func OrShutdown(inputs ...<-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		cases := make([]reflect.SelectCase, len(inputs))
		for idx, inputChan := range inputs {
			cases[idx] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(inputChan)}
		}

		reflect.Select(cases)

	}()

	return out
}

/*
AndShutdown combines multiple signalling channels and returns a
single signalling channel (The output channel is closed when all
of the input signalling channels are closed)
*/
func AndShutdown(inputs ...<-chan int) <-chan int {
	out := make(chan int)

	var syncGoroutines sync.WaitGroup

	for _, inputChan := range inputs {

		syncGoroutines.Add(1)

		go func(inputChan <-chan int, wg *sync.WaitGroup) {
			defer wg.Done()
			<-inputChan
		}(inputChan, &syncGoroutines)

	}

	go func(wg *sync.WaitGroup) {
		wg.Wait()
		close(out)
	}(&syncGoroutines)

	return out
}

/*
Pool TODO
*/
func Pool(shutdown <-chan int, input <-chan int, poolSize int, process func(int) int) <-chan int {
	out := make(chan int)

	var syncGoroutine sync.WaitGroup

	for i := 0; i < poolSize; i++ {

		syncGoroutine.Add(1)

		go func(shutdown <-chan int, inputChan <-chan int, wg *sync.WaitGroup) {
			defer wg.Done()

			for {
				select {
				case <-shutdown:
					return
				case num, ok := <-inputChan:
					if !ok {
						return
					}
					select {
					case <-shutdown:
						return
					case out <- process(num):
					}
				}
			}

		}(shutdown, input, &syncGoroutine)

	}

	go func(wg *sync.WaitGroup) {
		syncGoroutine.Wait()
		close(out)
	}(&syncGoroutine)

	return out
}
