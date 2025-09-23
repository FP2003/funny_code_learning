package main

import "fmt"

func main() {
	const pi = 3.14159
	const (
		red = iota + 1
		yellow
		green
		brown
		blue
		pink
		black
	)

	fmt.Printf("Pi approx. %v \n\n", pi)

	fmt.Printf("Red: %v point \n", red)
	fmt.Printf("Blue: %v point \n", blue)
	fmt.Printf("Black: %v point \n", black)

}
