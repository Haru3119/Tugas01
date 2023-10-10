package main

import (
	"fmt"
)

func main() {
	var celcius int

	fmt.Print("Enter a degree in celcius : ")
	fmt.Scan(&celcius)

	var kelvin = celcius + 273
	fmt.Print(kelvin)
}