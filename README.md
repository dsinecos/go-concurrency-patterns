## Motivation
To understand Golang's concurrency constructs (Goroutines, Channels etc) and implement different concurrency patterns.

## Learnt

1. Simulate a generator using Goroutines and channels

2. Iterating over a channel using `range`

3. Composing multiple goroutines as a pipeline
   1. Closing multiple downstream channels when an upstream input channel is closed
   2. Closing a single downstream channel when all the upstream input channels are closed
   3. Creating a Pipeline connected by channels where each filter is executed in a separate goroutine

4. Using `go run -race` to detect Race conditions

## API and Usage

_There is a potential for Goroutine leaks if the output channels returned from package functions are not read from_

## ToDo

1. How to extend the `chanutil` pkg to work for channels of different types without repeating code?
   1. [Applied Go - How to make generic data containers type safe](https://www.youtube.com/watch?v=rco7GEg3v0I)
   2. Using `go generate`
      1. [Generate implementations for generic types in Go](https://flaviocopes.com/golang-generic-generate/)
      2. [Using code generation to survive without generics in Go](https://www.calhoun.io/using-code-generation-to-survive-without-generics-in-go/)
      3. [Reducing boilerplate with go generate](https://blog.gopheracademy.com/advent-2015/reducing-boilerplate-with-go-generate/)
   3. Using a stage in the Pipeline to assert on types

2. Review `chanutil` and `generator` pkg to implement lexical confinement 
   1. By restricting operations on channels (to read or write) and 
   2. Providing variables through arguments to goroutines instead of closures (to avoid data races)

3. Review `chanutil` and `generator` pkg to identify points of Goroutine leaks | Add appropriate fix to prevent Goroutine leaks
   
   1. Allow to close a pipeline using a `shutdown` channel passed by the developer and prevent Goroutine leaks (Refer Pipelines article on Go Blog)
      
      1. Use a `shutdown` channel (Add an `orDone` method to reduce boilerplate when adding the `shutdown` channel)
      
      2. Use the Context pkg
   
   2. Is composition of the different methods possible when a `shutdown` channel is taken as an input argument?

4. Review `chanutil` and `generator` pkg to implement adequate error handling

5. Completed - `generator`
   1. Completed - Add `Repeat` and `Take` methods to `generator` pkg
   2. Completed - Add method `BatchToStream` to convert a batch of data to a stream to pkg `generator`

6. `chanutil`

   1. Completed - Add `Pool` method to `chanutil` pkg to allow simultaneous processing of tasks across multiple goroutines (`func (input chan int, poolSize int, processFunc func()) chan int)`)
   
   2. Add benchmarks to `Pool` method to compare performance across single and multiple goroutines for a CPU bound/ intensive task [Refer](https://play.golang.org/p/iAUyHaaUk1H)
      1. Benchmark for different Pool sizes
   
   3. For the `Pool` method use a buffered channel the size of which is equal to the number of resources available that are to be used among the Goroutines
   
   4. Completed - Add a method `OrShutdown` that combines multiple signalling channels and returns a single signalling channel (The output channel is closed if any of the input signalling channels is closed)

   5. Completed - Add a method `AndShutdown` that combines multiple signalling channels and returns a single signalling channel (The output channel is closed when all of the input signalling channels are closed)

   6. Add versions for the `Split` and `SplitRnd` method
       1.  Add functionality to return a slice of channels (The length of the slice would be derived from an argument to the function)
       2.  Add functionality to return a channel of channels (The number of channels written into the meta channel will be derived from an argument to the function)

7. Add documentation
    1.  Completed - To `chanutil` pkg using `docco` [Website](http://ashkenas.com/docco/)
    2.  Add steps to generate documentation in `Readme.md` using `docco`

8. Write tests for the `chanutil` pkg


