/*
This program is a basic implementation of a hash function that takes a name as input in the cli, 
sums up each ASCII characters corresponding to each letter of the name and divides that sum by 100.
*/

package main 

import (
	"bufio" // This package provides buffered I/O . 
	"fmt"  // This package provides formatted I/O . 
	"log"  // This package provides logging capabilities ;
	"os"   // This package provides platform-independent interface to OS functionality . 

)

func main() { 
	// Creating a new buffered reader that reads from the standard input.
	reader := bufio.NewReader(os.Stdin) 

  // Prompting the user to enter a name . 
	fmt.Print("Please enter a name : ") 


  // Reading a line of input from the user up to and including the newline character ('\n') . 
	// The input is stored in a variable called `name ` and any error occuring is stored in the `error`  variable . 
	name, error := reader.ReadString('\n') 


	// Checking if an error occured during reading the input.
	// If error occured, then it is logged and the program exits with a non-zero status code . 
	if error != nil { 
		log.Fatalf("Error handling your input: %v", error) 
	}

	name = name[:len(name)-1] // remove new line character

	// Initializing a variable called `sum` to store the sum of the ASCII values of the characters in the name.
	sum := 0

	 // Iterating over each character `c` in the `name` string , using a For loop . 
	for _, c:= range name { 
		// The ASCII value of each character is added to the `sum` varaible . 
		sum += int(c) 

	}
  // Calculating the result as the remainder of the division of the `sum` variable by 100 . 
	result := sum %100 

   // Printing out the result of the standard output. 
	fmt.Printf("The result is:%d\n", result)


	}