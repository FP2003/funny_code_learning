package main

import "fmt"

func main() {

	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {

			if i == 3 && j == 2 {
				fmt.Println("Continue when i=", i, "j=", j)
				continue
			}

			if i == 2 && j == 2 {
				fmt.Println("Breaks when i=", i, "j=", j)
				break
			}

			fmt.Println("Running i=", i, "j=", j)
		}
	}
}
