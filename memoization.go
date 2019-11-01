package main

import (
	"fmt"
	"time"
)

const LIMIT = 50

var fibboNumbers [LIMIT]uint64

func main() {
	var result uint64 = 0

	start := time.Now() // start timing

	// looping the fibonacci function i=50 times as you store the result to the 'result' variable
	result = fibonacci(i)
	for i := 0; i < LIMIT; i++ {
		fmt.Printf("fibonnaci(%d) is: %d\n", i, result)
	}
	end := time.Now()       // stop timing
	delta := end.Sub(start) // storing duration in the variable delta
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}

// logic of fibonacci sequence
func fibonacci(n int) (res uint64) {
	//memoization:check if the fibonacci(n) is already known in the array:
	if fibboNumbers[n] != 0 {
		res = fibboNumbers[n]
		return
	}

	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}

	fibboNumbers[n] = res
	return
}
