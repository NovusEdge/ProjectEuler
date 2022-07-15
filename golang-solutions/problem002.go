package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	var a, b, sum uint64 = 1, 2, 2
	for a+b < 4000000 {
		next := a + b
		if next%2 == 0 {
			sum += next
		}
		a, b = b, next
	}

	fmt.Printf("\nAnswer to Problem 2: %d\n", sum)
	fmt.Printf("Time Taken: %v\n\n", time.Since(start))
}
