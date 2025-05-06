package inf

import (
	"fmt"
	"golang_journey/utils/function"
)

// AreaAble 定义一个接口，展示接口能力
type AreaAble interface {
	Area() float64
}

type Circle struct {
	function.Point
	radius float64
}

type Rect struct {
	function.Point
	width, height float64
}

type Polygon struct {
	Points []function.Point
}

func (r *Rect) Area() float64 {
	return r.width * r.height
}

func (c *Circle) Area() float64 {
	return c.radius * c.radius * 3.14
}

func calPrice(sharp any) float64 {
	a, ok := sharp.(AreaAble)

	if !ok {
		// 不规则类型按照哦固定价格200
		return 200.00
	}
	return a.Area() * 2
}

func CalPriceTest() {
	p := function.Point{X: 1, Y: 1}
	c := Circle{p, 1}
	fmt.Printf("圆 c 的价格是: %.2f\n", calPrice(&c))
	r := Rect{p, 1, 2}
	fmt.Printf("矩形 r 的价格是: %.2f\n", calPrice(&r))
	polygon := Polygon{[]function.Point{p, {X: 9, Y: 1}, {X: 2, Y: 2}, {X: 1, Y: 2}, {X: 1, Y: 0}}}
	fmt.Printf("多边形 polygon 的价格是: %.2f\n", calPrice(&polygon))

}
