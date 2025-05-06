package _map

import "fmt"

func TestMap() {
	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
	}

	age, ok := ages["bob"]
	if !ok {
		fmt.Println("bob is not a key in the map, ", ok, age)
		age = 26
	}

	age, ok = ages["alice"]
	fmt.Println("alice is , ", ok, age)

}
