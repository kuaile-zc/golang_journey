package main

import (
	"fmt"
	"golang_journey/utils/array"
	"golang_journey/utils/function"
	"golang_journey/utils/inf"
	_map "golang_journey/utils/map"
	"golang_journey/utils/object"
	"golang_journey/utils/slice"
	"golang_journey/utils/sort"
	"reflect"
)

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)

	//fmt.Printf("The 100 C is %f F", convert.CToF(100))

	type Currency int

	const (
		USD Currency = iota // 美元
		EUR                 // 欧元
		GBP                 // 英镑
		RMB                 // 人民币
	)

	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}

	fmt.Println(RMB, symbol[RMB], reflect.TypeOf(symbol))
	fmt.Printf("%T\n", symbol)

	array.TestArray()
	slice.TestSlice()
	_map.TestMap()
	object.TestObject()
	function.BigSlowOperation()
	function.Double(2)
	function.FuncTest()
	// function.PanicTest(3)  //测试panic 和 defer

	fmt.Println("-----")
	inf.InterfaceTest()
	inf.CalPriceTest()

	x := []int{1, 2, 3}
	y := []int{1, 2, 3}
	//fmt.Println(x == y)  // 切片不可直接比较
	fmt.Println(reflect.DeepEqual(x, y)) // 使用反射包中的 DeepEqual 比较切片)

	sort.DemoTest()

}
