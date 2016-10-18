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

	fmt.Println("Dipanggil langsung tanpa goroutine")
	printStr()
	printInt()

	fmt.Println("Di source code ini diletakkan di atas 2 goroutine")
	go printStr()
	fmt.Println("Di source code ini diletakkan di antara 2 goroutine")
	go printInt()
	fmt.Println("Di source code ini diletakkan di bawah 2 goroutine")

	var input string
	fmt.Scanln(&input)
	fmt.Println("selesai")

}
