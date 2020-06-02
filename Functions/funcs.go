package main

import (
	"fmt"
)

func main() {
	sayHello(1)
	fmt.Println(letsSumUp(-1, 10))
}

func letsSumUp(a int, b int) (int, error) {

	if a < 0 {
		return 0, fmt.Errorf("This is mistake my Man %d", a)
	}
	return a + b, nil
}

func sayHello(times int) {

	for i := 0; i < times; i++ {
		fmt.Println("Hello")

	}
}
