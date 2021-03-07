package main

import (
	"fmt"
	operations "manuel-lang/Go-GitHub-Actions-Dependency"
)

func main() {
	sevenPlusFive := operations.Add(7, 5)
	sevenMinusFive := operations.Sub(7, 5)
	fmt.Printf("7 + 5 = %d \n", sevenPlusFive)
	fmt.Printf("7 - 5 = %d \n", sevenMinusFive)
}
