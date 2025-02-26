package main

import (
	"math/rand"
	"sync"
	"time"
)

// Constants for benchmarking
const (
	numThreads    = 10
	numOperations = 1000
)

// Worker function for mixed operations
func worker(l interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < numOperations; i++ {
		key := rand.Intn(1000)
		op := rand.Intn(3)

		switch list := l.(type) {
		case *ConcurrentLinkedList:
			if op == 0 {
				list.Insert(key)
			} else if op == 1 {
				list.Search(key)
			} else {
				list.Delete(key)
			}
		case *HandOverHandLinkedList:
			if op == 0 {
				list.Insert(key)
			} else if op == 1 {
				list.Search(key)
			} else {
				list.Delete(key)
			}
		}
	}
}

// Benchmark function
func benchmark(l interface{}) time.Duration {
	var wg sync.WaitGroup
	start := time.Now()

	for i := 0; i < numThreads; i++ {
		wg.Add(1)
		go worker(l, &wg)
	}
	wg.Wait()

	return time.Since(start)
}
