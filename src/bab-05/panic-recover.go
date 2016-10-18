package main

import (
	"fmt"
	"os"
)

func main() {

	// Fungsi berikut untuk melakukan recover supaya
	// program berakhir dengan baik, tidak menampilkan
	// run time error seperti di panic.go
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("Masalah:", e)
		}
	}()

	if _, err := os.Stat("not-exist.txt"); err != nil {
		panic("File yang diperlukan tidak ada!")
	}

	// Jika file not-exist.txt tidak ada di current dir
	// maka kode di bawah ini tidak akan pernah di eksekusi

	fmt.Println("Setelah memeriksa file ... melanjutkan")

}
