package main

import "fmt"

var z int

func add(x int, y int) int {
	return x + y
}

func sub(x int, y int) int {
	return x - y
}

func main() {
	sevenPlusFive := add(7, 5)
	sevenMinusFive := sub(7, 5)
	fmt.Printf("7 + 5 = %d \n", sevenPlusFive)
	fmt.Printf("7 - 5 = %d \n", sevenMinusFive)
}
