// Buatlah program dalam bahasa Go untuk menghitung volume bola.
// Masukan berupa jari-jari bola dalam bilangan real.
// Keluaran berupa volume bola dalam bilangan real.
// Petunjuk: Rumus volume bola adalah sebagai berikut
// volume bola = \( \frac{4}{3} \pi r^3 \)
// gunakan \( \pi \) = 3.14
// r = jari-jari

package main

import (
	"fmt"
	"math"
)

func main() {
	var r float64

	fmt.Printf("Enter a radius value : ")
	_, err := fmt.Scan(&r)

	if err != nil {
		fmt.Println("Invaild Values !")
		return
	}

	var v = (4/3) * math.Pi * math.Pow(r, 3)
	fmt.Print(v) 
}