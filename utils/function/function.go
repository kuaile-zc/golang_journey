package function

import (
	"fmt"
	"log"
	"math"
	"time"
)

type Point struct{ X, Y float64 }
type Path []Point

func FuncTest() {
	p, q := Point{1, 2}, Point{4, 6}
	fmt.Println("Distance is: ", Distance(p, q))
	fmt.Println("Distance is: ", p.Distance(q))

	perimeter := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println("Triangle perimeter is: ", perimeter.Distance())

	distanceP := p.Distance
	fmt.Println("Distance is: ", distanceP(q))
	distance := Point.Distance
	fmt.Println("Distance is: ", distance(p, q))
}

func BigSlowOperation() {
	defer trace("bigSlowOperation")()
	time.Sleep(1 * time.Second)

}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}

func Double(x int) (result int) {
	defer func() { fmt.Printf("double(%d) = %d\n", x, result) }()
	return x + x
}

func PanicTest(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
	defer fmt.Printf("defer %d\n", x)
	PanicTest(x - 1)
}

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}
