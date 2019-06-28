## Motivation
To understand Golang's concurrent constructs (Goroutines, Channels etc) and different Concurrency patterns.

## Learnt

1. Simulate a generator using Goroutines and channels

2. Iterating over a channel using `range`

3. Composing multiple goroutines as a pipeline
   1. Closing multiple downstream channels when an upstream input channel is closed
   2. Closing a single downstream channel when all the upstream input channels are closed

4. Using `go run -race` to detect Race conditions

## API and Usage

_There is a potential for Goroutine leaks if the output channels returned from package functions are not read from_
