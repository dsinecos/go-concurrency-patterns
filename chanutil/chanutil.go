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

// Split broadcasts (duplicates) input from a single channel across multiple output channels
func Split(input chan int, outputs ...chan int) {

	go func() {

		// Create n WaitGroups, each corresponding to goroutines spawned for
		// writing to a single output channel
		syncGoroutines := make([]sync.WaitGroup, len(outputs))

		for index, outputChan := range outputs {
			syncGoroutines[index] = sync.WaitGroup{}

			// Waitgroup incremented to account for when the input channel closes
			syncGoroutines[index].Add(1)

			// The following goroutine monitors if all the goroutines spawned to publish
			// to the respective output channel have closed using 'wg' and the 'input'
			// channel has been closed, to close the 'output' channel
			go func(outputChan chan int, wg *sync.WaitGroup) {
				wg.Wait()
				close(outputChan)
			}(outputChan, &syncGoroutines[index])

		}

		// Range over the input channel. Spawn a goroutine for each output channel each time a value
		// is read from the input channel.
		for inputVal := range input {
			for index, outputChan := range outputs {

				// Increment the waitgroup for the assigned output channel to account
				// for the goroutine that will be launched to write the present value
				// to the output channel
				syncGoroutines[index].Add(1)

				go func(inputVal int, outputChan chan int, wg *sync.WaitGroup) {
					outputChan <- inputVal
					wg.Done()
				}(inputVal, outputChan, &syncGoroutines[index])

			}
		}

		// Since the earlier for-range loop is blocking and will only release
		// once the 'input' channel is closed. We can define code here to be
		// executed once the 'input' channel is closed
		// Range over n Waitgroups and signal their end since the input channel
		// is closed
		for index := range syncGoroutines {
			syncGoroutines[index].Done()
		}

	}()
}

// SplitRnd distributes input from a single channel across multiple output channels randomly
func SplitRnd(input chan int, outputs ...chan int) {

	// Spawn a goroutine for each output channel. The goroutine
	// is responsible for reading from the input channel and writing
	// to the output channel
	for _, outputChan := range outputs {

		// Takes output channel as an argument and not via closures to avoid
		// data race conditions
		go func(outputChan chan int) {
			// Close the output channel once the input channel is closed
			defer close(outputChan)

			// Blocks and runs until the input channel is closed
			for num := range input {
				outputChan <- num
			}

		}(outputChan)
	}

}

// Pipeline recursively creates a multi-stage asynchronous pipeline to filter values from the input channel in order of the functions provided and publish qualifying values to the output channel
func Pipeline(input chan int, filters ...func(task int) bool) chan int {

	out := make(chan int)

	filter := filters[0]

	// Goroutine creates a stage of the pipeline using a filter function and
	// an output channel
	go func(out chan int, filter func(task int) bool) {
		// Closes the output channel once the goroutine exits
		defer close(out)
		// Blocks and reads from the input channel until it is closed
		for value := range input {
			if filter(value) {
				out <- value
			}
		}
	}(out, filter)

	// If there are more than one filter functions, Pipeline is invoked recursively
	if len(filters) != 1 {
		filters = filters[1:]
		return Pipeline(out, filters...)
	}

	// If there is only a single filter function, output channel is returned
	return out
}

// OrShutdown combines multiple signalling channels and returns a
// single signalling channel (The output channel is closed if any
// of the input signalling channels is closed)
func OrShutdown(inputs ...<-chan int) <-chan int {
	out := make(chan int)

	// Goroutine creates a dynamically sized select-case statement using the reflect package.
	go func() {
		defer close(out)

		cases := make([]reflect.SelectCase, len(inputs))
		// For each input channel a case statement is initialized to read from the respective channel
		for idx, inputChan := range inputs {
			cases[idx] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(inputChan)}
		}

		// Goroutine blocks here until one of the input channels closes
		reflect.Select(cases)

	}()

	return out
}

// AndShutdown combines multiple signalling channels and returns a
// single signalling channel (The output channel is closed when all
// of the input signalling channels are closed)
func AndShutdown(inputs ...<-chan int) <-chan int {
	out := make(chan int)

	// WaitGroup is used to synchronize closing the output channel when
	// all the input channels are closed
	var syncGoroutines sync.WaitGroup

	// For each input channel, a goroutine is spawned. When the input channel
	// closes, the goroutine decrements the WaitGroup and exits
	for _, inputChan := range inputs {

		// Increment WaitGroup for each goroutine launched
		syncGoroutines.Add(1)

		go func(inputChan <-chan int, wg *sync.WaitGroup) {
			defer wg.Done()
			// Goroutine blocks here until the input channel closes
			<-inputChan
		}(inputChan, &syncGoroutines)

	}

	// Goroutine monitors when all the input channels are closed (signalled by respective
	// goroutines spawned from the for loop).
	go func(wg *sync.WaitGroup) {
		// Blocks until all goroutines spawned to monitor individual channels exit
		wg.Wait()
		close(out)
	}(&syncGoroutines)

	return out
}

// Pool invokes multiple goroutines to process values on the input channel concurrently
func Pool(shutdown <-chan int, input <-chan int, poolSize int, process func(int) int) <-chan int {
	out := make(chan int)

	// WaitGroup initialized to synchronize closing all the goroutines once the input
	// channel is closed
	var syncGoroutine sync.WaitGroup

	// Spawn 'poolSize' goroutines where each goroutine reads from the input channel
	// runs the 'process' function on the value and publishes result to the output channel
	for i := 0; i < poolSize; i++ {

		syncGoroutine.Add(1)

		go func(shutdown <-chan int, inputChan <-chan int, wg *sync.WaitGroup) {
			defer wg.Done()

			for {
				// Select-Case used to allow to close the goroutine if the 'shutdown' channel is closed
				select {
				case <-shutdown:
					return
				case num, ok := <-inputChan:
					// To allow the goroutine to exit when the input channel is closed
					if !ok {
						return
					}
					// The following select-case block covers the case where value was read from the input channel
					// and while the goroutine was waiting to write on the output channel, 'shutdown' channel was
					// closed. Without the following select-case block, the goroutine would attempt to write to
					// the output channel despite the 'shutdown' channel being signalled to close
					select {
					case <-shutdown:
						return
					case out <- process(num):
					}
				}
			}

		}(shutdown, input, &syncGoroutine)

	}

	// The following goroutine closes the output channel once all the goroutines reading from the input
	// channel have exited
	go func(wg *sync.WaitGroup) {
		syncGoroutine.Wait()
		close(out)
	}(&syncGoroutine)

	return out
}
