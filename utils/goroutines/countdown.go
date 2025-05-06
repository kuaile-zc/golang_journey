package goroutines

import (
	"fmt"
	"time"
)

func Countdown() {
	fmt.Println("Commencing countdown.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-tick
	}

}

func MemorySyc() {
	var x, y int

	go func() {
		x = 1
		fmt.Print("y:", y, " ") // y is not initialized
	}()

	go func() {
		y = 1
		fmt.Print("x:", x, " ") // x is not initialized
	}()

	time.Sleep(time.Second)
}
