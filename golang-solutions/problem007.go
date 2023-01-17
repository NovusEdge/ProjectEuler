package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	start := time.Now()
	var dummy, flag int64 = 0, 0

	for flag < 10001 {
		if isPrime(dummy) {
			flag++
		}
		dummy++
	}

	dummy-- // Excluding 0 since it's neither prime nor composite

	fmt.Printf("\nAnswer to Problem 7: %d\n", dummy)
	fmt.Printf("Time Taken: %v\n\n", time.Since(start))
}

func isPrime(n int64) bool {
	for i := int64(2); i <= int64(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
