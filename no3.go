// Buatlah program dalam bahasa Go untuk menghitung keliling persegi panjang.
// Masukan berupa panjang dan lebar dari persegi panjang dalam bilangan bulat.
// Keluaran berupa keliling persegi panjang dalam bilangan bulat.

package main

import (
	"fmt"
)

func main() {
	var width, height int

	fmt.Print(" Enter value of height and width : " )
	fmt.Scan(&height, &width)

	var p = ( 2 * width ) + (  2 * height )
	fmt.Print(p)

}
