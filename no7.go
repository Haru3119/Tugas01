package main

import (
	"fmt"
)

func main() {
	var lilbro, bigbro int
	fmt.Printf("Enter lilbro number (0-9) : ")
	fmt.Scan(&lilbro)

	fmt.Printf("Enter bigbro number (0-9) : ")
	fmt.Scan(&lilbro)

	if (lilbro%2 == 1 && bigbro == 0) || (lilbro%2 == 0 && bigbro == 1) {
		fmt.Print("lilbro wins !!!")
	} else {
		fmt.Printf("bigbro wins !!!")
	}

}