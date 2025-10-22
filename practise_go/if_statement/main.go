package main

import "fmt"

func main() {

	check := 'A'

	if 5 < 1 {
		fmt.Println("\n1st Condition Met")
	} else if 'A' != check {
		fmt.Println("\n2nd Condition Met")
	} else {
		fmt.Println("\nBoth Conditions False")
	}

}
