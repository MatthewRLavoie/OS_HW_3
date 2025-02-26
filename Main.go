package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Benchmarking Single Lock List")
	singleLockList := &ConcurrentLinkedList{}
	durationSingle := benchmark(singleLockList)
	fmt.Printf("Execution time: %v\n", durationSingle)

	fmt.Println("\nBenchmarking Hand-Over-Hand Lock List")
	handOverHandList := &HandOverHandLinkedList{}
	durationHOH := benchmark(handOverHandList)
	fmt.Printf("Execution time: %v\n", durationHOH)
}
