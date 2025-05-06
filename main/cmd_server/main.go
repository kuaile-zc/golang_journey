package main

import (
	"flag"
	"fmt"
	"golang_journey/utils/inf"
)

func main() {
	name := flag.String("name", "Guest", "Your name")
	age := flag.Int("age", 0, "Your age")
	temp := inf.CelsiusFlag("temp", 20, "the temperature")

	flag.Parse()

	fmt.Printf("Hello, %s! You are %d years old. \n", *name, *age)
	fmt.Println(temp)
}
