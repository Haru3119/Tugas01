package main

import (
	"fmt"
)

func main() { 
	var length, height, width, weight int
	
	fmt.Print("Enter package length (cm) : ")
	_, err1 := fmt.Scan(&length)

	fmt.Print("Enter package width (cm) : ")
	_, err2 := fmt.Scan(&width)

	fmt.Print("Enter package height (cm) : ")
	_, err3 := fmt.Scan(&height)

	fmt.Print("Enter package weight (gr) : ")
	_, err4 := fmt.Scan(&weight)

	if err1 != nil  || err2 != nil || err3 != nil || err4 != nil {
		fmt.Println("Not Valid !!!")
	}

	vol := int(length * width * height) / 1000000
	packageWeight := int(weight) / 1000

	acceptable := vol <= 1 && packageWeight <= 30

	fmt.Print("Package can be delivered",  acceptable)
}