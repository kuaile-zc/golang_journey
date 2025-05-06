package array

import "fmt"

func TestArray() {
	var a [10]int
	fmt.Println("Before a is ", a)
	changeArray(a)
	fmt.Println("After1 a is ", a)
	changeArray2(&a)
	fmt.Println("After2 a is ", a)
}

func changeArray(array [10]int) {
	array[0] = 100
}

func changeArray2(array *[10]int) {
	array[0] = 100
}
