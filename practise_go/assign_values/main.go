package main

import "fmt"

func main() {

	a, b := 8, 4
	fmt.Println("Assigned values:\ta =", a, "\tb =", b)

	a += b
	fmt.Println("Adding and assign:\ta =", a)

	a -= b
	fmt.Println("Subtract and assign:\ta =", a)

	a *= b
	fmt.Println("Multiply and assign:\ta =", a)

	a /= b
	fmt.Println("Divide and assign:\ta =", a)

	a %= b
	fmt.Println("Remainder and assign:\ta =", a)

}
