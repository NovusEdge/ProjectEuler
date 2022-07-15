package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	sum := 0
	for i := 1; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	fmt.Printf("\nAnswer to Problem 1: %d\n", sum)
	fmt.Printf("Time Taken: %v\n\n", time.Since(start))
}
