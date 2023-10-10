// Buatlah program dalam bahasa Go untuk menghitung luas persegi panjang. Masukan berupa
// dua buah bilangan real yang menyatakan panjang dan lebar dari persegi panjang. Keluaran
// berupa luas persegi panjang dalam bilangan real.

// h = height/tinggi
// w = width/lebar
// a = area/luas


package main 

import "fmt"

func main() {
	var h, w float64

	fmt.Print("Enter Value of Width and Height : ")
	fmt.Scan(&h, &w)

	var a = h * w
	fmt.Print(a)
}
