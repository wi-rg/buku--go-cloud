# Serialisasi

## Pengertian Serialisasi

Serialisasi adalah proses pengubahan suatu objek menjadi urutan `bit` agar dapat disimpan pada media penyimpanan (seperti berkas komputer, atau pada memori) atau ditransimisikan melalui saluran koneksi jaringan. Sewaktu rangkaian `bit` ini dibaca ulang sesuai dengan format serialisasinya, ia dapat digunakan untuk menciptakan `clone` identik secara `semantis` dari objek aslinya. Bagi banyak objek kompleks, misalnya objek yang banyak menggunakan rujukan, proses ini tidak dapat dilakukan begitu saja. Proses serialisasi objek ini biasa disebut penyusunan (`marshalling`) objek. Sedangkan operasi kebalikannya, yaitu penyusunan kembali bit menjadi struktur data utuh disebut pembongkaran (`unmarshalling`) objek. 

## Penggunaan Serialisasi

Serialisasi digunakan sebagai :
*Cara untuk melakukan penyimpanan objek yang lebih mudah dibandingkan menuliskan properti atas objek-objek tersebut ke dalam berkas teks, dan mengembalikannya sebagai objek.
*Salah satu proses yang dilakukan dalam pemanggilan prosedur jarak jauh (`remote procedure call`), contoh `SOAP`
*Salah satu cara untuk mendistribusikan objek, khususnya dalam arsitektur berbasis komponen, contoh, `COM`, `CORBA`, dll.
*Salah satu cara untuk mendeteksi perubahan data dalam satu periode waktu tertentu.

Agar fungsi serialisasi dapat memberikan manfaat seperti tujuan awalnya, arsitektur perangkat lunak yang independen harus dikelola secara konsisten. Misalnya, agar dapat berfungsi maksimal dalam hal pendistribusian, komputer yang memiliki arsitektur hardware yang berbeda harus pula dapat merekonstruksi aliran data yang telah diserialisasi secara reliabel. Tujuan dari serialisasi struktur data dalam arsitektur independen adalah agar data tersebut secara reliabel dapat dibaca, direkonstruksi secara mudah pada platform-platform lain. Hal ini berarti data yang berasal dari prosedur konvensional yang serderhana, berunjuk kerja tinggi, yang secara langsung menyalin blok-blok memori komputer biasanya tidak akan dapat digunakan pada arsitektur yang lain.
 
## Format Serialisasi

Pada akhir 1990-an, atas desakan kebutuhan untuk menyediakan layanan protokol standar serialisasi alternatif yang lebih dapat dimengerti atau mudah dipahami oleh orang maka terciptalah `text-based encoding XML`. Pengkodean ini berguna untuk objek tetap yang dapat dibaca dan dipahami oleh manusia, atau komunikasi antar sistem tanpa memperhatikan bahasa pemrograman. Kekurangan dari pengkodean ini adalah kemampuan `byte-stream-based encoding`-nya berkurang banyak, tetapi seiring dengan perkembangan media penyimpanan data dan kapasitas pengiriman yang semakin pesat membuat masalah ini tidak terlalu penting seperti diawal era dunia komputer. `Binary XML` diperkenalkan, dimana `Binary XML` ini tidak dapat dibaca oleh text-editor biasa namun memiliki kemampuan penyimpanan data yang lebih baik dri `XML` biasa. Pada tahun 2000-an, XML sering digunakan untuk asinkronus transfer data antara `client` dan `server` dalam `web applications AJAX`.

`JSON` yang notabene lebih ringan daripada `XML` yang mana lebih umum digunakan untuk komunikasi antar `client-server` dalam aplikasi web. `JSON` berbasis `Javascript syntax` namun tetap support dengan bahasa pemrograman lainnya.

Alternatif lainnya adalah `YAML` yang tidak lain adalah pasangan terbaik untuk `JSON` membuatnya lebih powerful dalam serialisasi, lebih `human friendly`, dan jauh lebih ringkas. Fitur-fitur ini termasuk konsep penandaan jenis data, dukungan untuk struktur data non-hirarkis, pilihan untuk membuat struktur data dengan indentasi, dan berbagai bentuk `scalar data quoting`.

Format serialiasi lain yang dapat dipahami manusia yaitu `property list` yang sering digunakan pada `NeXTSTEP`, `GNUstep`, dan `OS X Cocoa`.

Untuk volume data ilmiah yang lebih besar, sepeti data satelit dan output numerik dari iklim, cuaca, atau model laut, telah dikembangkan standar serialisasi khusus diantaranya `HDF`, `netCDF` dan `GRIB` yang lama.  

## Go dan XML



## Go dan JSON



## Go dan YAML
Paket `yaml` memungkinkan program `go`
 untuk encode dan decode nilai pada `yaml`. Paket ini dikembangkan dalam `Canonical` yang mana sebagian dari proyek `juju`  dan didasarkan pada port Go murni dari libyaml C library terkenal untuk mengurai dan menghasilkan data YAML dengan baik.

#### Instalasi
Sebelum melakukan instalasi buat direktori untuk golang terlebih dahulu, disini saya contohkan diletakkan pada direktori `$HOME` dandengan nama `gocode` kemudian lakukan perintah berikut untuk menentukan `PATH` yang dibutuhkan :
```
export GOPATH=$HOME/gocode
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```
Karen saya menggunakan ubuntu dan yang akan di install adalah yaml versi 2 maka menggunakan perintah :
```
go get gopkg.in/yaml.v2
111

#### Contoh Program Go-Yaml
```
package main

import (
    "fmt"
    "log"
    "gopkg.in/yaml.v2"
)

var configText = `
  Message: Welcome
  UpdateInterval: 5
  EmailAddresses:
    - rohimmt@yahoo.com
    - raigeki.rmt@yahoo.com`

type Config struct {
    Message        string   `yaml:"Message"`
    UpdateInterval int      `yaml:"UpdateInterval"`
    EmailAddresses []string `yaml:"EmailAddresses"`
}

func main() {
    var c Config
    if err := yaml.Unmarshal([]byte(configText), &c); err != nil {
        log.Fatalf("Failed to parse yaml: %v", err)
    }

    fmt.Printf("Config: %+v\n", c)
}
```
Maka output yang dihasilkan adalah


## Go dan MessagePack



## Go dan Protocol Buffer



