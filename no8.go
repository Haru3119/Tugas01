package main

import (
	"fmt"
)

func main() {
	var b1, b2, b3 int

	fmt.Print("enter first sulphur weights : ")
	fmt.Scan(&b1)

	fmt.Print("enter second sulphur weights : ")
	fmt.Scan(&b2)
	
	fmt.Print("enter third sulphur weights : ")
	fmt.Scan(&b3)

	total := b1 + b2 + b3

	kilogram := total / 1000
	gram := total % 1000

	fmt.Printf("sulphur total weight is %d kilogram %d gram\n", kilogram, gram)
}