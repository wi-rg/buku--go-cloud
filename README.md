# Membangun Aplikasi Cloud Menggunakan Go

Buku ini merupakan buku pemrograman Go yang ditujukan untuk pengembangan aplikasi SaaS di Cloud. Meskipun titik berat dari buku ini ada pada pemrograman aplikasi di Cloud, buku ini juga meliputi dasar-dasar Go serta penggunaan Go untuk backend dan frontend. 

# Penulis

Buku ini ditulis bersama-sama oleh dosen dan mahasiswa [STMIK AKAKOM Yogyakarta](http://www.akakom.ac.id). Meskipun demikian, kontribusi dari siapapun itu akan diterima dengan senang hati. Untuk saat ini, beberapa nama ini berperan dalam penulisan:

* [Bambang Purnomosidi D. P.](http://bpdp.xyz): penulis utama, editor, maintainer, all hands person.
* ... tambahkan ...

# Menggunakan Buku

Meskipun bisa langsung membaca dari repo Github ini melalui file README.md, ada beberapa cara lainnya yang bisa digunakan:

* `make`: membuat 3 format buku (EPUB, HTML, dan PDF), hasil akan diletakkan di direktori sesuai dengan isi variable `BUILD` pada `Makefile`
* `make pdf` untuk membuat format PDF, `make html` untuk membuat format HTML, dan `make epub` untuk membuat format EPUB.

# Isi Buku

Buku ini akan dibagi menjadi beberapa bagian:
* Pengenalan
* Dasar-dasar Pemrograman Go: Sintaksis, Semantik, dan Pustaka Standar
* Go untuk Backend
* Frontend Programming dan Go
* Go di Cloud
* Penutup

## Pengenalan

Pada bagian ini, akan dibahas instalasi peranti pengembangan bahasa pemrograman Go serta peranti yang biasanya digunakan pada saat `coding` menggunakan Go (IDE - `Integrated Development Environment`). Setalah membaca, memahmi, dan mengikuti instruksi pada bagian ini, pembaca akan mempunyai peranti pengembangan Go serta IDE untuk `coding` terinstall pada komputer. Pembaca juga akan memahami struktur kode sumber Go dan siap untuk mempelajari komponen Go lebih lanjut.

* [Bab 1: Pengenalan Go](bab-01.md)
* [Bab 2: IDE untuk Go](bab-02.md)

## Dasar Pemrograman `Go`: Sintaksis, Semantik, dan Pustaka Standar

Bagian ini membahas tentang komponen dasar dan inti dari bahasa pemrograman Go. Semua pembahasan pada bagian ini terdapat pada instalasi standar peranti pengembangan Go.

* [Bab 3: Dasar-dasar Pemrograman Go](bab-03.md)
* [Bab 4: Fungsi / `Function](bab-04.md)
* [Bab 5: Penanganan Kesalahan](bab-05.md)
* [Bab 6: Struktur Data Lanjut](bab-06.md)
* [Bab 7: Method](bab-07.md)
* [Bab 8: Testing](bab-08.md)
* [Bab 9: Konkurensi](bab-09.md)

## Go untuk Backend 

Bagian ini membahas tentang penggunaan Go sebagai peranti pengembang untuk backend. Beberapa bagian dari pembahasan ini menggunakan pustaka standar dari Go dan juga beberapa software atau pustaka pihak ketiga. Setiap pembahasan akan dimulai dengan peranti pustaka standar serta peranti pustaka pihak ketiga yang akan digunakan.

## Frontend Programming dan Go

Bagian ini membahas tentang pemrograman pada sisi frontend serta bagaimana menggunakan Go sebagai bagian dari frontend programming tersebut.

## Penutup

* [Daftar Pustaka](daftar-pustaka.md)

# Lisensi

![CC-BY-SA 4.0](images/cc-by-sa-4.png)
[CC-BY-SA 4.0](http://creativecommons.org/licenses/by-sa/4.0/)

# Lain-lain

Buku ini dibuat menggunakan markdown dan dikonversi menjadi EPUB - HTML - PDF menggunakan [pandoc](http://pandoc.org). Template untuk buku diambil dari [evangoer/pandoc-ebook-template](https://github.com/evangoer/pandoc-ebook-template).
