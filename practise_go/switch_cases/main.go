package main

import "fmt"

func main() {

	num := 2
	char := 'B'

	switch num {
	case 1:
		fmt.Println("\nNumber is one")
	case 2:
		fmt.Println("\nNumber is two")
	case 3:
		fmt.Println("\nNumber is three")
	case 4:
		fmt.Println("\nNumber is four")
	default:
		fmt.Println("\nNumber is not recognised")
	}

	switch char {
	case 'A':
		fmt.Println("\nLetter is A")
	case 'B':
		fmt.Println("\nLetter is B")
	case 'C':
		fmt.Println("\nLetter is C")
	}

}
