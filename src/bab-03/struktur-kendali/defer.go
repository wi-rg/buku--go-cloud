package main

import "fmt"

func printStr() {

	fmt.Println("A")
	fmt.Println("B")
	fmt.Println("C")

}

func printInt() {

	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

}

func main() {

	fmt.Println()
	fmt.Println("Fungsi printStr() di-defer disini")
	defer printStr()

	fmt.Println("Fungsi printInt() di-defer disini")
	defer printInt()

	fmt.Println()
	fmt.Println("Ini aliran eksekusi normal")
	fmt.Println("--------------------------")
	printStr()
	printInt()
	fmt.Println()
	fmt.Println("Akhir dari main()")
	fmt.Println("Setelah akhir dari main()")
	fmt.Println("fungsi yang di-defer akan dikerjakan ")
	fmt.Println("secara LIFO - yang di-defer terakhir akan")
	fmt.Println("dikerjakan lebih dulu")

}
