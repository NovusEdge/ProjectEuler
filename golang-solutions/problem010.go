package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	start := time.Now()
	var answer int64

	for i := int64(0); i < 2000000; i++ {
		if isPrime(i) {
			answer += i
		}
	}

	fmt.Printf("\nAnswer to Problem 10: %d\n", answer-1)
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
