package main

import "fmt"

var hasilPerbandingan bool
var angka1 uint8 = 21
var angka2 uint8 = 17

func main() {
	hasilPerbandingan = angka1 < angka2
	fmt.Println(hasilPerbandingan)
}
