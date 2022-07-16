package main

import (
	"fmt"
	"math"
	"sort"
	"time"
)

func main() {
	start := time.Now()
	num_factors := factors(600851475143)
	var prime_factors []int

	for _, i := range num_factors {
		if isPrime(i) {
			prime_factors = append(prime_factors, i)
		}
	}

	sort.Ints(prime_factors)

	fmt.Printf("\nAnswer to Problem 3: %d\n", prime_factors[len(prime_factors)-1])
	fmt.Printf("Time Taken: %v\n\n", time.Since(start))
}

func isPrime(n int) bool {
	return len(factors(n)) == 2
}

func factors(n int) (factors []int) {
	for i := 1; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			factors = append(factors, i)
			factors = append(factors, n/i)
		}
	}

	return
}
