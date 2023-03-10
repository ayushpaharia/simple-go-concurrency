package examples

import (
	"fmt"
	"sync"
	"time"
)

func WithChannel(needs []string) {
	started := time.Now()
	ch := make(chan string)
	for _, f := range needs {
		go func(f string) {
			get(f)
			ch <- f
		}(f)
	}
	for range needs {
		<-ch
	}
	fmt.Printf("Got in %v\n", time.Since(started))
	defer close(ch)
}

func WithWaitGroup(needs []string) {
	started := time.Now()
	var wg sync.WaitGroup
	wg.Add(len(needs))
	for _, f := range needs {
		go func(f string) {
			get(f)
			wg.Done()
		}(f)
	}
	wg.Wait()
	fmt.Printf("Got in %v\n", time.Since(started))
}

func get(need string) {
	fmt.Printf("Getting %s...\n", need)
	time.Sleep(1 * time.Second)
	fmt.Printf("Got %s", need)
	fmt.Println("")
}
