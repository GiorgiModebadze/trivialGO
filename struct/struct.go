package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	// Commenting because I must have to
	Height float64
	Weight float64
	Male   bool
}

func main() {

	var People [2]Person

	People[0] = Person{174, 70, true}
	People[1] = Person{170, 56, false}

	a := [5]int{1, 2, 3, 4, 5}

	fmt.Println(reflect.TypeOf(People))
	fmt.Println(People)
	fmt.Println(a)
}
