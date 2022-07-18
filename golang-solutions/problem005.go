package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	answer := 1
	for i := 2; i < 21; i++ {
		answer = LCM(answer, i)
	}

	fmt.Printf("\nAnswer to Problem 5: %d\n", answer)
	fmt.Printf("Time Taken: %v\n\n", time.Since(start))
}

func LCM(a, b int) int {
	return a * b / GCD(a, b)
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
