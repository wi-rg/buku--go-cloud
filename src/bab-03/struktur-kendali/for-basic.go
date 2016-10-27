package main

import "fmt"

func main() {
	fmt.Println("for init; condition; post { }")
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println(sum)
	}
	fmt.Println()
	fmt.Println("for condition { }")
	sum1 := true
	for sum1 {
		sum += 2
		fmt.Println(sum)
		if sum >= 70 {
			sum1 = false
		}
	}
	fmt.Println()
	fmt.Println("for { }")
	for {
		fmt.Println("never ending loop ... press Ctrl-C to stop")
	}
}
