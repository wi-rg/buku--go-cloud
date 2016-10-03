package main

import "fmt"

func main() {

	const version string = "1.0.0"

	fmt.Println(version)

	const c1uint8 uint8 = 27
	const c2float32 float32 = 123.456

	const c3float32 float32 = c2float32 / 21

	fmt.Println(c1uint8)
	fmt.Println(c2float32)
	fmt.Println(c3float32)

	// Berikut ini akan menghasilkan error:
	// 		cannot assign to c1uint8
	// c1uint8 = 21

}
