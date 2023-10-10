package main

import (
	"fmt"
)

func main() {
	var num int

	fmt.Print("Enter 2 digit of number (10 - 99): ")
	_, err := fmt.Scan(&num)

	if err != nil || num < 10 || num > 99 {
		fmt.Println("Not valid Number ! Enter a valid number ! (10 - 99) ")
		return
	}

	firstDigit := num / 10
	secondDigit := num % 10

	result := firstDigit*1000 + firstDigit*100  + secondDigit*10 + secondDigit

	fmt.Printf("Result is : %d\n", result)
}
