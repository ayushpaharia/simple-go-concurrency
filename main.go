package main

import "simple-golang-concurrency/examples"

func main() {
	// examples.WithChannel([]string{"a", "b", "c", "d", "e"})
	// examples.WithWaitGroup([]string{"a", "b", "c", "d", "e"})
	examples.RunFanInFanOut()
}
