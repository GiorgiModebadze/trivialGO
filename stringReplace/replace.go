package main

import (
	"fmt"
	"strings"
)

func main() {
	broker := "Hom# Sw##t Hom#"
	replacer := strings.NewReplacer("#", "e")

	newVal := replacer.Replace(broker)

	fmt.Println(newVal)
}
