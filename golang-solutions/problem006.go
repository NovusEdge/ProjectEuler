package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	answer := squareOfSum(100) - sumOfSquares(100)

	fmt.Printf("\nAnswer to Problem 6: %d\n", answer)
	fmt.Printf("Time Taken: %v\n\n", time.Since(start))
}

func squareOfSum(n int) int {
	return (n * (n + 1) / 2) * (n * (n + 1) / 2)
}

func sumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}
