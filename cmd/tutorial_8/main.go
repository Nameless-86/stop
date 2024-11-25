package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

//var dbData = []string{"id1", "id2", "id3", "id4", "id5"}

func main() {
	t0 := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go count() //concurrency
	}

	wg.Wait()
	fmt.Printf("\nTotal exec time: %v", time.Since(t0))
}

func count() {
	var res int
	for i := 0; i < 10000000; i++ {
		res += 1
	}
	wg.Done()
}
