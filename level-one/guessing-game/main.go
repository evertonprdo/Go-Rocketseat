package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Guessing game")
	fmt.Println("A random number will be drawn. Try to guess. The number is an integer between 0 and 100.")

	x := rand.Int64N(101)
	scanner := bufio.NewScanner(os.Stdin)
	attempts := [10]int64{}

	for i := range attempts {
		fmt.Println("What's your guess?")
		scanner.Scan()

		attempt := scanner.Text()
		attempt = strings.TrimSpace(attempt)

		attemptInt, err := strconv.ParseInt(attempt, 10, 64)
		if err != nil {
			fmt.Println("Your guess must be a whole number")
			return
		}

		switch {
			case attemptInt < x:
				fmt.Println("You Wrong. The number drawn is higher than the number you guessed.", attemptInt)

			case attemptInt > x:
				fmt.Println("You Wrong. The number drawn is lower than the number you guessed.", attemptInt)

			case attemptInt == x:
				fmt.Printf(
					"Congratulation, you get the number right, which was: %d\n"+
					"You hit in %d attempts."+
					"\nThese are your attempts: %v\n",
					x, i + 1, attempts[:i],
				)
				return
		}

		attempts[i] = attemptInt
	}

	fmt.Printf(
		"Unfortunately, you didn't get the number right, which was: %d\n"+
		"You had 10 attempts."+
		"\nThese are your attempts: %v\n",
		x, attempts,
	)
}