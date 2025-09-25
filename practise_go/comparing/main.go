package main

import "fmt"

func main() {

	zero, num, max := 0, 0, 1
	up, dn := 'A', 'a'

	result := (num == zero)
	fmt.Println("Equality:\tnum == zero\t", result)

	result = (up == dn)
	fmt.Println("Equality:\tup == dn\t", result)

	result = (max != zero)
	fmt.Println("Inequality:\tmax != zero\t", result)

	result = (zero > max)
	fmt.Println("Greater:\tzero = max\t", result)

	result = (max <= zero)
	fmt.Println("Less or Equal:\tmax <= zero\t", result)

}
