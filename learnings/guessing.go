package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seed the random number generator with the current time
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 1 and 100
	secretNumber := rand.Intn(100) + 1

	fmt.Println("Welcome to the Guessing Game!")
	fmt.Println("Try to guess the number between 1 and 100.")

	// Main game loop
	for {
		var guess int

		// Prompt user for input
		fmt.Print("Enter your guess: ")
		_, err := fmt.Scanf("%d", &guess)

		// Check for input errors
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			// Clear the input buffer
			fmt.Scanf("%s")
			continue
		}

		// Check the guess
		if guess < 1 || guess > 100 {
			fmt.Println("Please enter a number between 1 and 100.")
			continue
		} else if guess < secretNumber {
			fmt.Println("Too low! Try again.")
		} else if guess > secretNumber {
			fmt.Println("Too high! Try again.")
		} else {
			fmt.Println("Congratulations! You guessed the correct number.")
			break
		}
	}
}
