package goroutines

import (
	"fmt"
	"time"
)

func Spinner() {
	go spinner(100 * time.Millisecond)
	fmt.Printf("\rfib(45) = %d\n", fib(45))
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
