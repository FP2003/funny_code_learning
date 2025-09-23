package main

import "fmt"

func main() {
	// for loop - 0 to 10

	// for i := 0; i <= 10; i++ {
	// 	fmt.Println("i:", i)
	// }

	// for loop - 0 to 25 (only odd numbers)

	// for i := 0; i <= 25; i++ {
	// 	if i%2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// for loop - 0 to 20 (only even numbers)

	for i := 0; i <= 20; i++ {
		if i%2 == 1 {
			continue
		}
		fmt.Println(i)
	}
}
