package main

import (
	"fmt"
	"os"
)

func main() {

	if _, err := os.Stat("not-exist.txt"); err != nil {
		panic("File not exist")
	}

	// Jika file not-exist.txt tidak ada di current dir
	// maka kode di bawah ini tidak akan pernah di eksekusi

	fmt.Println("Setelah memeriksa file ... melanjutkan")

	// kemudian muncul berikut ini:
	//
	// panic: File not exist

	// goroutine 1 [running]:
	// panic(0x467c00, 0xc42000a310)
	// /opt/software/go-dev-tools/go/go1.7.1/src/runtime/panic.go:500 +0x1a1
	// main.main()
	//   /home/bpdp/kerjaan/git-repos/wi-rg/buku--go-cloud/src/bab-05/panic.go:8 +0x9e
	//   exit status 2

}
