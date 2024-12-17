package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	randomDuration := time.Duration(rand.Intn(1000)) * time.Millisecond
	defer wg.Done()
	fmt.Printf("worker %d starting\n", id)

	time.Sleep(randomDuration) //simulate some work
	fmt.Printf("worker %d worked for %d miliseconds\n", id, randomDuration)
	fmt.Printf("worker %d done\n", id)
}

func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1) //increment counter for each goroutine
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("All workers done")

}
