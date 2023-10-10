package main

import (
	"fmt"
)

func main() {
	var x float64

	fmt.Printf("Enter value of x :")
	fmt.Scan(&x)

	var y = (3 * x - 5) * (2 * x + 1)
	fmt.Print(y)
}