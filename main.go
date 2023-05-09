

/*
This program is a basic implementation of a hash function that takes a name as input in the cli,
sums up each ASCII characters corresponding to each letter of the name and divides that sum by 100.

Side Note: 
After reading the user input, it checks if the input consists of only alphabetic letters by iterating over each character and using the unicode.IsLetter() function.
If a non-alphabetic character is found, an error is thrown with log.Fatalf() and the program exits.
Otherwise, it prints the value.
*/

package main

import (
	"bufio" // This package provides buffered I/O.
	"fmt"   // This package provides formatted I/O.
	"log"   // This package provides logging capabilities;
	"os"    // This package provides platform-independent interface to OS functionality.
	"unicode"
)

func main() {
	// Creating a new buffered reader that reads from the standard input.
	reader := bufio.NewReader(os.Stdin)

	// Prompting the user to enter a name.
	fmt.Print("Please enter a name: ")

	// Reading a line of input from the user up to and including the newline character ('\n').
	// The input is stored in a variable called `name` and any error occurring is stored in the `error` variable.
	name, error := reader.ReadString('\n')

	// Checking if an error occurred during reading the input.
	// If an error occurred, then it is logged and the program exits with a non-zero status code.
	if error != nil {
		log.Fatalf("Error handling your input: %v", error)
	}

	// Checking if the input consists of only alphabetic letters.
	for _, c := range name[:len(name)-1] {
		if !unicode.IsLetter(c) {
			log.Fatalf("Input contains non-alphabetic characters: %v", name)
		}
	}

	// Initializing a variable called `sum` to store the sum of the ASCII values of the characters in the name.
	sum := 0

	// Iterating over each character `c` in the `name` string, using a for loop.
	for _, c := range name[:len(name)-1] {
		// The ASCII value of each character is added to the `sum` variable.
		sum += int(c)
	}

	// Calculating the result as the remainder of the division of the `sum` variable by 100.
	result := sum % 100

	// Printing out the result of the standard output.
	fmt.Printf("The result is: %d\n", result)
}
