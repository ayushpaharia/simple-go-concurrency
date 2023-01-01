package main

import (
	"fmt"
	"time"
)

func main() {
	started := time.Now()
	needs := []string{"food", "water", "shelter", "company"}
	for _, f := range needs {
		go func(f string) {
			get(f)
		}(f)
	}
	fmt.Printf("Got in %v\n", time.Since(started))
}

func get(need string) {
	fmt.Printf("Getting %s...\n", need)
	time.Sleep(2 * time.Second)
	fmt.Printf("Got %s\n", need)
	fmt.Println("")
}
