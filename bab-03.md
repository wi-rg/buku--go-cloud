# Dasar-dasar Pemrograman Go

## Struktur Program Go

## Program Aplikasi Sederhana - 1 File `binary executable` Utama

Suatu aplikasi `executable` (artinya bisa dijalankan secara langsung oleh sistem operasi) mempunyai struktur seperti yang terlihat pada listing berikut ini:

~~~go
/*
	aplikasi.go

	Contoh program sederhana untuk menjelaskan
	struktur program Go untuk aplikasi executable

	(c) bpdp.xyz

*/

// Program Go diawali dengan nama paket.
// Paket untuk aplikasi executable selalu berada
// pada paket main.
package main

// pustaka standar yang diperlukan
// Jika hanya satu:
// import "fmt"
// Jika lebih dari satu:
import (
	"fmt"
	"os"
)

// "Fungsi" merupakan satuan terintegrasi dari
// program Go, selalu diberi nama "main" untuk
// aplikasi executable.
func main() {

	// ini adalah kode sumber / program Go
	// akan dijelaskan lebih lanjut, abaikan
	// jika belum paham
	var (
		user    string
		homeDir string
		goHome  string
	)

	user = os.Getenv("USER")
	homeDir = os.Getenv("HOME")
	goHome = os.Getenv("GOROOT")

	fmt.Printf("Halo %s", user)
	fmt.Printf("\nHome anda di %s", homeDir)
	fmt.Printf("\nAnda menggunakan Go di %s", goHome)
	fmt.Printf("\n")

}
~~~

Untuk menjalankan kode sumber di atas, ikuti langkah-langkah berikut:

### Tanpa Proses Kompilasi

~~~bash
$ go run aplikasi.go 
Halo bpdp
Home anda di /home/bpdp
Anda menggunakan Go di /home/bpdp/software/go-dev-tools/go/go1.7.3
~~~

### Mengkompilasi Menjadi *Binary Executable*

~~~bash
$ go build aplikasi.go 
$ ls -la
total 1620
drwxr-xr-x 1 bpdp users      38 Sep 12 22:49 .
drwxr-xr-x 1 bpdp users     288 Aug 15 11:24 ..
-rwxr-xr-x 1 bpdp users 1654480 Sep 12 22:49 aplikasi
-rw-r--r-- 1 bpdp users     900 Sep 12 22:49 aplikasi.go
$ ./aplikasi 
Halo bpdp
Home anda di /home/bpdp
Anda menggunakan Go di /opt/software/go-dev-tools/go/go1.7.3
$ file aplikasi
aplikasi: ELF 64-bit LSB executable, x86-64, version 1 (SYSV), statically linked, not stripped
$
~~~

## Pustaka / Library / Package

Ada kalanya, para software developer membangun pustaka yang berisi berbagai fungsionalitas yang bisa digunakan kembali suatu saat nanti. Untuk keperluan ini, Go menyediakan fasilitas untuk membangun library dalam bentuk kumpulan fungsi. Kumpulan fungsi ini nantinya akan diletakkan pada suatu repo tertentu sehingga bisa langsung di `go get <lokasi repo pustaka>`. Pada penjelasan berikut ini, kita akan membangun suatu aplikasi kecil (hello) yang menggunakan suatu pustaka yang sebelumnya sudah kita bangun (stringutil/Reverse - untuk membalik kata). Kode sumber diambil dari [How to write Go code](https://golang.org/doc/code.html).

### Mengatur Workspace

Ada 3 direktori di workspace yang disiapkan: bin, pkg, dan src:
* bin: berisi hasil kompilasi aplikasi (hello)
* pkg: berisi hasil kompilasi pustaka (stringutil)
* src: kode sumber untuk pustaka serta aplikasi
Pada direktori tersebut, juga dibuat `env.sh`.

~~~bash
$ ls
total 4
drwxr-xr-x 1 bpdp users 30 Sep 12 23:16 .
drwxr-xr-x 1 bpdp users 56 Sep 12 23:05 ..
drwxr-xr-x 1 bpdp users 10 Sep 12 23:16 bin
-rw-r--r-- 1 bpdp users 50 Sep 12 23:05 env.sh
drwxr-xr-x 1 bpdp users 22 Sep 12 23:16 pkg
drwxr-xr-x 1 bpdp users 32 Sep 12 23:07 src
$ cat env.sh 
export GOPATH=`pwd`
export PATH=$PATH:$GOPATH/bin
$ source ~/env/go/go1.7.3
$ source env.sh 
$ 
~~~

Semua kode sumber, baik untuk pustaka ataupun aplikasi akan diletakkan pada pola direktori tertentu. Go menggunakan pola repo untuk penamaan / pengelompokan aplikasi atau pustaka meskipun belum dimasukkan ke repo di Internet. Sebaiknya membiasakan diri sejak awal menggunakan pola tersebut meskipun belum akan dimasukkan ke repositori di Internet. 

### Membuat Pustaka

Kode sumber untuk pustaka ini akan diletakkan di `src/github.com/bpdp/stringutil`. Paket yang dibuat dengan penamaan ini, nantinya akan diacu dalam `import` sebagai `github.com/bpdp/stringutil`.

~~~go
/*	
    src/github.com/bpdp/stringutil/reverse.go 
	diambil dari https://golang.org/doc/code.html
*/
package stringutil

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
~~~

Untuk mengkompilasi:

~~~bash
$ go build github.com/bpdp/stringutil
$
~~~

Jika tidak ada kesalahan, maka akan langsung kembali ke prompt shell. 

### Membuat Aplikasi yang Memanfaatkan Pustaka

Sama halnya dengan pustaka, aplikasi juga menggunakan pola penamaan yang sama. 

~~~go
/* 
    src/github.com/bpdp/hello/hello.go 
    hello.go
    diambil dari https://golang.org/doc/code.html
*/
package main

import (
	"fmt"

	"github.com/bpdp/stringutil"
)

func main() {
	fmt.Printf(stringutil.Reverse("!oG ,olleH"))
}
~~~

Untuk mengkompilasi:

~~~bash
$ go install github.com/bpdp/hello
$ hello
Hello, Go!
$ ls bin/
total 1612
drwxr-xr-x 1 bpdp users      10 Sep 12 23:16 .
drwxr-xr-x 1 bpdp users      30 Sep 12 23:16 ..
-rwxr-xr-x 1 bpdp users 1650409 Sep 12 23:16 hello
$
~~~

## Konstruksi Dasar Bahasa Pemrograman Go

### Komentar

Bagian komentar dimaksudkan untuk dokumentasi dari *source code*. Ada dua cara untuk memberikan komentar:
* Menggunakan `/* ... */` untuk komentar yang meliputi lebih dari satu baris
* Menggunakan `//` di awal baris untuk komentar yang meliputi satu baris saja
Komentar ini sejak awal sebaiknya sudah dibiasakan harus ada karena Go menyediakan fasilitas `godoc` untuk menghasilkan dokumentasi dari *source code*. Bagian yang sebaiknya diberikan komentar / dokumentasi adalah bagian diatas `package` dan di atas setiap definisi fungsi (lihat contoh dari `stringutil` di atas.

### Tipe Data Angka / Numerik

Untuk tipe numerik, pada dasarnya kita bisa menggunakan bilangan bulat (*integer*) dan bilangan pecahan (*floating-point*). Bilangan bulat terdiri atas bilangan bertanda (*signed* - int) dan bilangan tak-bertanda (*unsigned* - uint). Berikut ini adalah daftar lengkap dari tipe data numerik tersebut:


|  Tipe     | Arti | Jangkauan |
|-----------|------|-----------|
| uint8     | unsigned  8-bit integer | 0 sampai  255 |
| uint16    | unsigned 16-bit integer | 0 sampai 65535 |
| uint32    | unsigned 32-bit integer | 0 sampai 4294967295 |
| uint64    | unsigned 64-bit integer | 0 sampai 18446744073709551615 |
|  |  |  |
| int8      | signed  8-bit integer | -128 sampai 127 |
| int16     | signed 16-bit integer | -32768 sampai 32767 |
| int32     | signed 32-bit integer | -2147483648 sampai 2147483647 |
| int64     | signed 64-bit integer | -9223372036854775808 sampai 9223372036854775807 |
|  |  |  |
| float32   | IEEE-754 32-bit floating-point |  |
| float64   | IEEE-754 64-bit floating-point |  |
|  |  |  |
| complex64  | bilangan kompleks dengan float32 riil dan imajiner | ~ |
| complex128 | bilangan kompleks dengan float64 riil dan imajiner | ~ |
|  |  |  |
| byte | alias dari uint8 |  |
| rune | alias dari int32 |  |

Selain definisi di atas, Go juga mempunyai alias penyebutan yang implementasinya tergantung pada arsitektur komputer yang digunakan:

| Tipe | Arti |
|------|------|
| uint | arsitektur 32 atau 64 bit |
| int  | mempunyai ukuran yang sama dengan uint |
| uintptr | bilangan bulat tak bertanda untuk menyimpan nilai pointer |

### String

String digunakan untuk men

~~~go
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
~~~

Hasil:

~~~bash
$ go run contoh-string.go 
STMIK AKAKOM
12
true
Stmik Akakom
77
STMIK AKAKOM Yogyakarta
string
string
~~~

### Boolean

Tipe data Boolean berisi nilai benar (`true`) atau salah (`false`).

~~~go
package main

import "fmt"

var hasilPerbandingan bool
var angka1 uint8 = 21
var angka2 uint8 = 17

func main() {
	hasilPerbandingan = angka1 < angka2
	fmt.Println(hasilPerbandingan)
}
~~~

Hasil:

~~~bash
$ go run boolean.go 
false
~~~

### Variabel


~~~go
// nilai-default-variabel.go
package main

import "fmt"

func main() {

	// unsigned-integer
	var defUint8 uint8
	var defUint16 uint16
	var defUint32 uint32
	var defUint64 uint64
	var defUint uint

	// signed-integer
	var defInt8 int8
	var defInt16 int16
	var defInt32 int32
	var defInt64 int64
	var defInt int

	// string
	var defString string

	// floating-point
	var defFloat32 float32
	var defFloat64 float64

	// complex
	var defComplex64 complex64
	var defComplex128 complex128

	// alias
	var defByte byte
	var defRune rune

	fmt.Println("\nNilai default untuk uint8 = ", defUint8)
	fmt.Println("Nilai default untuk uint16 = ", defUint16)
	fmt.Println("Nilai default untuk uint32 = ", defUint32)
	fmt.Println("Nilai default untuk uint64 = ", defUint64)
	fmt.Println("Nilai default untuk uint = ", defUint)

	fmt.Println("\nNilai default untuk int8 = ", defInt8)
	fmt.Println("Nilai default untuk int16 = ", defInt16)
	fmt.Println("Nilai default untuk int32 = ", defInt32)
	fmt.Println("Nilai default untuk int63 = ", defInt64)
	fmt.Println("Nilai default untuk int = ", defInt)

	fmt.Println("\nNilai default untuk string = ", defString)

	fmt.Println("\nNilai default untuk float32 = ", defFloat32)
	fmt.Println("Nilai default untuk float64 = ", defFloat64)

	fmt.Println("\nNilai default untuk complex64 = ", defComplex64)
	fmt.Println("Nilai default untuk complex128 = ", defComplex128)

	fmt.Println("\nNilai default untuk byte = ", defByte)
	fmt.Println("Nilai default untuk rune = ", defRune)

}
~~~

Hasil eksekusi:

~~~bash
$ go run nilai-default-variabel.go

Nilai default untuk uint8 =  0
Nilai default untuk uint16 =  0
Nilai default untuk uint32 =  0
Nilai default untuk uint64 =  0
Nilai default untuk uint =  0

Nilai default untuk int8 =  0
Nilai default untuk int16 =  0
Nilai default untuk int32 =  0
Nilai default untuk int63 =  0
Nilai default untuk int =  0

Nilai default untuk string =  

Nilai default untuk float32 =  0
Nilai default untuk float64 =  0

Nilai default untuk complex64 =  (0+0i)
Nilai default untuk complex128 =  (0+0i)

Nilai default untuk byte =  0
Nilai default untuk rune =  0
$
~~~

### Konstanta

Konstanta dimaksudkan untuk menampung data yang tidak akan berubah-ubah. Konstanta dideklarasikan menggunakan kata kunci *const*. Konstant bisa bertipe *character*, string, boolean, atau numerik.

## Pointer

Konsep *pointer* sebenarnya sudah ada pada bahasa pemrograman lain, khususnya C/C++ (dengan kompleksitas yang lebih tinggi). Suatu *pointer* merupakan suatu nilai yang menunjuk ke suatu nilai lain. 

## Struktur Kendali

### Perulangan dengan `for`

### Seleksi Kondisi

#### Pernyataan `if`


#### Pernyataan `switch`

### Defer

Defer digunakan untuk mengekesekusi suatu perintah sebelum suatu fungsi berakhir. Jika berada pada suatu fungsi, baris kode sumber yang di-defer akan dikerjakan sebelum menemui akhir (*return*). Kegunaan utama dari *defer* ini adalah untuk keperluan pembersihan (*cleanup*). Saat kita membuat kode sumber Go, sering kali dalam setiap operasi terdapat beberapa hal yang harus kita akhiri dengan kondisi tertentu, misalnya jika kita membuka file maka kita harus menutup file jika kita sudah selesai melakukan operasi dengan file tersebut. *Defer* mempermudah kita untuk memastikan bahwa pekerjaan-pekerjaan pembersihan tersebut selalu bisa dilakukan.


