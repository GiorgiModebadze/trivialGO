/*
the Goal of this Program is to read a grade and decide if we failed or passed
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	passingScore := 60.

	fmt.Print("Enter your Score : ")

	reader := bufio.NewReader(os.Stdin)

	providedGrade, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	parsedGrade := strings.TrimSpace(providedGrade)

	convertedGrade, err := strconv.ParseFloat(parsedGrade, 64)

	if err != nil {
		log.Fatal(err)
	}

	if convertedGrade >= passingScore {
		fmt.Println("You Passed")
	} else {
		fmt.Println("You Failed")
	}

}
