package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	guessNumber := 10
	rand.Seed(time.Now().UnixNano())

	r := int64(rand.Intn(guessNumber + 1))

	for {
		fmt.Print("Guess The Number Between 1 and ", guessNumber, " :")

		reader := bufio.NewReader(os.Stdin)

		inputNumber, _ := reader.ReadString('\n')

		stripedNumber := strings.TrimSpace(inputNumber)

		convertedNumber, err := strconv.ParseInt(stripedNumber, 10, 0)

		if err != nil {
			log.Fatal(err)
		}

		if convertedNumber == r {
			fmt.Println("You Guessed the Number")
			break
		} else if convertedNumber < r {
			fmt.Println("Too Low")
		} else {
			fmt.Println("Too High")
		}

	}

}
