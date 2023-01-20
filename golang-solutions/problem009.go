package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	var a, b, c, temp int

	for i := 1; i < 500; i++ {
		for j := 1; j < 500; j++ {
			temp = 1000 - (i + j)
			if i*i+j*j == temp*temp {
				a, b, c = i, j, 1000-(a+b)
				break
			}
		}
	}

	fmt.Printf("\nAnswer to Problem 9: %d\n", a*b*c)
	fmt.Printf("Time Taken: %v\n\n", time.Since(start))
}
