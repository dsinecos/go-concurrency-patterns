## Motivation

I wrote this pkg to understand Golang's concurrency constructs (Goroutines, Channels, Select-case etc) by implementing different concurrency patterns.

## Documentation

You can read the annotated code [here](https://cranky-franklin-0e76f2.netlify.app/)

## Concurrency Patterns

**Merge**

Merge combines input from multiple channels into a single channel

**Split**

Split broadcasts (duplicates) input from a single channel across multiple output channels

**SplitRnd**

SplitRnd distributes input from a single channel across multiple output channels randomly

**Pipeline**

Pipeline recursively creates a multi-stage asynchronous pipeline to filter values from the input channel in order of the functions provided and publish qualifying values to the output channel

**OrShutdown**

OrShutdown combines multiple signalling channels and returns a single signalling channel (The output channel is closed if any of the input signalling channels is closed)

**AndShutdown**

AndShutdown combines multiple signalling channels and returns a single signalling channel (The output channel is closed when all of the input signalling channels are closed)

**Pool**

Pool invokes multiple goroutines to process values on the input channel concurrently

## Resources

- [Concurrency in Go: Tools and Techniques for Developers](https://www.amazon.com/Concurrency-Go-Tools-Techniques-Developers/dp/1491941197)

- [Go Blog - Go Concurrency Patterns: Pipelines and cancellation](https://blog.golang.org/pipelines)