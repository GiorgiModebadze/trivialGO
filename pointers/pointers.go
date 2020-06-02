package main

import "fmt"

func main() {

	i := 4

	var a *int = &i

	fmt.Println(*a)

}
