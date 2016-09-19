package main

import (
	"fmt"
	"reflect"
	s "strings"
)

// Definisi string
var str1 string = "STMIK AKAKOM"
var str2 = "Yogyakarta"
var str3 = "stmik akakom"

func main() {

	// Lihat https://golang.org/pkg/strings/
	fmt.Println(str1)
	fmt.Println(len(str1))
	fmt.Println(s.Contains(str1, "AKAKOM"))
	fmt.Println(s.Title(str3))
	fmt.Println(str1[2])
	fmt.Println(s.Join([]string{str1, str2}, " "))
	fmt.Println(reflect.TypeOf(str1))
	fmt.Println(reflect.TypeOf(str2))
	fmt.Println()

}
