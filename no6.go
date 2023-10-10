package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64

	fmt.Printf("Enter value of x ")
	fmt.Scan(&x)

	var y = (math.Pow(x, 3) + (3 * x)) / (math.Pow(x, 4) + 3*math.Pow(x, 2) + 4)
	fmt.Print(y)
}
