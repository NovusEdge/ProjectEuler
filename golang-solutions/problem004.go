package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {
	start := time.Now()
	var pallindromes []int

	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			if k := i * j; isPallindrome(k) {
				pallindromes = append(pallindromes, k)
			}
		}
	}

	sort.Ints(pallindromes)

	fmt.Printf("\nAnswer to Problem 4: %d\n", pallindromes[len(pallindromes)-1])
	fmt.Printf("Time Taken: %v\n\n", time.Since(start))
}

func isPallindrome(n int) bool {
	nStr := fmt.Sprintf("%d", n)
	return nStr == reverse(nStr)
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
