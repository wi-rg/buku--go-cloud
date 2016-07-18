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
