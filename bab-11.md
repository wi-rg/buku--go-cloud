# Serialisasi

## Pengertian Serialisasi


## Penggunaan Serialisasi


## Format Serialisasi


## Go dan XML



## Go dan JSON
#### Definisi
`JSON` (JavaScript Object Notation) adalah format pertukaran data yang ringan, mudah dibaca dan ditulis oleh manusia, serta mudah diterjemahkan dan dibuat (generate) oleh komputer. Format ini dibuat berdasarkan bagian dari [Bahasa Pemprograman JavaScript](http://javascript.crockford.com/), [Standar ECMA-262 Edisi ke-3 - Desember 1999](http://www.ecma-international.org/publications/files/ecma-st/ECMA-262.pdf). JSON merupakan format teks yang tidak bergantung pada bahasa pemprograman apapun karena menggunakan gaya bahasa yang umum digunakan oleh programmer keluarga C termasuk C, C++, C#, Java, JavaScript, Perl, Python dll termasuk pada Golang yang akan kita bahas kal ini. Oleh karena sifat-sifat tersebut, menjadikan JSON ideal sebagai bahasa pertukaran-data.

Dengan paket JSON ini kita dapat membaca dan menulis data JSON dari program Go.

#### Encoding
Encoding atau mengkonvert variable dalam Go ke string dengan format json. Fungsi yang dapat digunakan untuk encoding yaitu :
```
func Marshal(v interface{}) ([]byte, error)
```
Contoh :
```
package main

import (
	"fmt"
	"encoding/json"
)

type MyStruct struct {
	Name   string
	Height int32
	Score  float32
	Exam   []string
}

func main() {
	s := &MyStruct{"Apin", 160, 85.5, []string{"Math", "History"}}
	if jsonStr, err := json.Marshal(s); err == nil {
		fmt.Println(string(jsonStr))
	}
}
```
Output :

```
{"Name":"Apin","Height":160,"Score":85.5,"Exam":["Math","History"]}
```
Dari output diatas dapat diartikan bahwa nama key dari json sesuai dengan nama variable dalam struct tersebut. Jika kita menginginkan key berbeda dengan nama variable kita dapat menggunakan struct tag, perhatikan code dibawah ini :

```
package main

import (
	"encoding/json"
	"fmt"
)

type MyStruct struct {
	Name   string   `json:"nama"`
	Height int32    `json:"tinggi"`
	Score  float32  `json:"nilai"`
	Exam   []string `json:"ulangan"`
	Other  string   `json:"-"`
}

func main() {
	s := &MyStruct{"Apin", 160, 85.5, []string{"Math", "History"}, "Lain-lain"}
	if jsonStr, err := json.Marshal(s); err == nil {
		fmt.Println(string(jsonStr))
	}
}
```
Maka outputnya :
```
{"nama":"Apin","tinggi":160,"nilai":85.5,"ulangan":["Math","History"]}
```
Jika kita ingin meng-ignore atau tidak ingin encode ke json, kita dapat menggunakan json:”-” pada tag di struct nya.

#### Decoding
Kebalikan dari encoding, decoding yaitu membuat variable Go dengan string json. Kita dapat menggunakan fungsi berikut :
```
func Unmarshal(data []byte, v interface{}) error
```
Penggunaannya mirip dengan encoding sebelumnya, dan berikut contoh penggunakaannya.

```
package main

import (
	"encoding/json"
	"fmt"
)

type MyStruct struct {
	Name   string   `json:"nama"`
	Height int32    `json:"tinggi"`
	Score  float32  `json:"nilai"`
	Exam   []string `json:"ulangan"`
	Other  string   `json:"-"`
}

func main() {
	str := `{"nama":"Apin","tinggi":160,"nilai":85.5,"ulangan":["Math","History"], "lainnya": "tidak ada"}`
	mystruct := &MyStruct{}
	if err := json.Unmarshal([]byte(str), mystruct); err == nil {
		fmt.Printf("%#v\n", mystruct)
	}
}
```
Output :
```
&main.MyStruct{Name:"Apin", Height:160, Score:85.5, Exam:[]string{"Math", "History"}, Other:""}
```
Sama dengan encoding sebelumnya kita dapat menggunakan tag pada struct. Jika tidak ada tag struct maka akan menggunakan nama dari variable pada struct tersebut.

#### Encoding dan Decoding ke map
Selain encoding dan decoding pada struct, Go juga dapat mengdecoding dan encoding ke map dengan tipe `map[string]interface{}`. Perhatikan contoh dibawah untuk lebih jelas
```
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//encoding
		fmt.Println("--ENCODING--")
	mymap := make(map[string]interface{})
	mymap["nama"] = "Rudi"
	mymap["ulangan"] = []string{"Fisika", "Sejarah"}
	mymap["object"] = map[string]interface{}{"key": "value"}
	jsonstr, _ := json.Marshal(mymap)
	fmt.Println(string(jsonstr))

	//decoding
	fmt.Println("\n--DECODING--")
	str := `{"nama":"Apin","tinggi":160,"nilai":85.5,"ulangan":["Math","History"], "lainnya": "tidak ada"}`
	m := make(map[string]interface{})
	if err := json.Unmarshal([]byte(str), &m); err == nil {
		for key, val := range(m) {
			fmt.Println(key, ":", val)
		}
	} else {
		fmt.Println(err)
	}
}
```
Output :
```
--ENCODING--
{"nama":"Rudi","object":{"key":"value"},"ulangan":["Fisika","Sejarah"]}

--DECODING--
nama : Apin
tinggi : 160
nilai : 85.5
ulangan : [Math History]
lainnya : tidak ada
```
## Go dan YAML



## Go dan MessagePack



## Go dan Protocol Buffer
Protokol Buffer adalah metode serialisasi data terstruktur. Hal ini berguna dalam mengembangkan program-program untuk berkomunikasi satu sama lain melalui sebuah jaringan atau untuk menyimpan data. Metode ini melibatkan bahasa deskripsi antarmuka yang menggambarkan struktur beberapa data dan program yang menghasilkan source codde dari deskripsi untuk menghasilkan atau mengurai aliran byte yang mewakili data terstruktur.

#### Mendefinisikan Format protokol
Untuk membuat aplikasi buku alamat, yang dilakukan pertama kali dengan membuat file `.proto`. Definisi dalam file `.proto` : Dengan menambahkan pesan untuk setiap struktur data yang ingin diserialisasikan, kemudian tentukan nama dan jenis untuk masing-masing bidang dalam pesan. Contohnya file `.proto` yang sudah mendefinisikan pesan `addressbook.proto`.

File `.proto` dimulai dengan deklarasi paket, yang membantu untuk mencegah konflik penamaan antara proyek yang berbeda.
```
syntax = "proto3";
package tutorial;
```
Didalam Go, nama paket digunakan sebagai paket Go, kecuali jika menetapkan sebuah `go_package`. Bahkan jika kita memberikan `go_package`, kita masih harus menentukan paket normal dan juga untuk menghindari tabrakan nama pada Protocol Buffer ruang nama serta bukan dalam bahasa Go.

Selanjutnya, kita mendefinisikan sebuah pesan dengan agregat dengan isi rangkaian yang diketik. Banyak standar tipe data sederhana yang tersedia sebagai jenis isian, termasuk `bool`, `int32`, `float`, `double`, dan `string`. Kita juga dapat menambahkan struktur lebih lanjut untuk pesan dengan menggunakan jenis pesan lain sebagai jenis isian.
```
message Person {
  string name = 1;
  int32 id = 2;  // Nomor ID untuk Person.
  string email = 3;

  enum PhoneType {
    MOBILE = 0;
    HOME = 1;
    WORK = 2;
  }

  message PhoneNumber {
    string number = 1;
    PhoneType type = 2;
  }

  repeated PhoneNumber phones = 4;
}

// file buku alamat.
message AddressBook {
  repeated Person people = 1;
}
```
Contoh di atas, pesan `Person` berisi pesan `PhoneNumber`, sedangkan pesan `AddressBook` berisi pesan `Person`. Kita bahkan dapat menentukan jenis pesan yang bersarang didalam pesan lain - seperti yang kita lihat, jenis `PhoneNumber` didefinisikan dalam `Person`. Kita juga dapat menentukan jenis enum jika kita ingin salah satu dari berapa isian yang kita memiliki dari daftar yang tersedia - disini kita ingin menentukan bahwa nomor telepon dapat menjadi salah satu `MOBILE`, `HOME`, atau `WORK`.

Nilai "= 1", "= 2" sebagai penanda pada setiap elemen mengidentifikasi unik "tag" isian yang digunakan dalam pengkodean biner. Tag nomor 1-15 memerlukan sekurangnya satu byte untuk mengkodekan dari angka yang lebih tinggi, sehingga optimasi kita dapat memutuskan untuk menggunakan tag tersebut untuk elemen yang biasa digunakan atau berulang, meninggalkan tag 16 dan lebih tinggi untuk elemen opsional yang kurang digunakan. Setiap elemen dalam isian yang diulang membutuhkan pengkodean ulang dengan nomor tag, sehingga isian yang diulang adalah kandidat yang sangat baik untuk optimasi ini.

Jika nilai isian tidak diatur, `nilai default`: nol untuk tipe numerik, string kosong untuk string, kesalahan untuk bools. Untuk pesan tertanam, nilai default selalu "ditetapkan" atau "prototipe" pesan, yang tidak memiliki isian yang yang diberikan. Memanggil pengaksesan untuk mendapatkan nilai dari isian yang belum eksplisit selalu diatur untuk mengembalikan nilai default.

Jika isian diulang, mungkin isian dapat diulang beberapa kali (termasuk nol). Urutan nilai berulang akan dipertahankan dalam protocol buffer. Isian yang diulang secara dinamis nilainya berukuran array.

#### Protokol Buffer API
Menghasilkan addressbook.pb.go memberi kita jenis penggunaan sebagai berikut:
+ Struktur `AddressBook` dengan kolom `People`.
+ Struktur `Person` dengan kolom `Name`, `Id`, `Email` dan `Phones`.
+ Struktur `Person_PhoneNumber`, dengan kolom `Number` dan `Type`.
+ Jenis `Person_PhoneType` dan nilai yang ditetapkan untuk setiap nilai dalam enum `Person.PhoneType`.

Berikut ini adalah contoh dari unit test perintah `list_people` tentang bagaimana kita dapat membuat sebuah perumpamaan dari Person:
```
p := pb.Person{
        Id:    1234,
        Name:  "John Doe",
        Email: "jdoe@example.com",
        Phones: []*pb.Person_PhoneNumber{
                {Number: "555-4321", Type: pb.Person_HOME},
        },
}
```
#### Penulisan pesan
Tujuan menggunakan protocol buffer adalah untuk serialisasi data sehingga bisa diurai ditempat lain. Didalam Go, kita menggunakan library `proto` fungsi `Marshal` untuk serialisasi data protocol buffer. Sebuah pointer ke `struct` pesan protocol buffer mengimplementasikan `proto.Message` antarmuka. Pemanggilan `proto.Marshal` untuk mengembalikan protocol buffer yang dikodekan dalam format wire. Sebagai contoh, kita menggunakan fungsi ini di perintah `add_person`:
```
book := &pb.AddressBook{}
// ...

// Menulis kembali buku alamat yang baru kedalam penyimpanan.
out, err := proto.Marshal(book)
if err != nil {
        log.Fatalln("Gagal untuk mengkodekan buku alamat:", err)
}
if err := ioutil.WriteFile(fname, out, 0644); err != nil {
        log.Fatalln("Gagal untuk menulis buku alamat:", err)
}
```
#### Pesan Pembaca
Untuk mengurai penkodean pesan, kita menggunakan library `proto` fungsi `Unmarshal`. Pemanggilan ini digunakan untuk mengurai data dalam `buf` sebagai protocol buffer dan menempatkan hasilnya di pb. Jadi untuk mengurai file dalam perintah `list_people`, kami menggunakan:
```
// Membaca buku alamat yang sudah ada.
in, err := ioutil.ReadFile(fname)
if err != nil {
        log.Fatalln("Kesalahan membaca file:", err)
}
book := &pb.AddressBook{}
if err := proto.Unmarshal(in, book); err != nil {
        log.Fatalln("Gagal mengurai buku alamat:", err)
}
```
