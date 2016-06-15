# Konkurensi
program-program besar yang sering terdiri dari banyak sub-program yang lebih kecil. Misalnya web server menangani permintaan yang dibuat dari web browser dan menyajikan halaman web HTML di respon. Setiap permintaan ditangani seperti program kecil.

Ini akan menjadi ideal untuk program-program seperti ini untuk dapat menjalankan komponen yang lebih kecil pada saat yang sama (dalam kasus server web untuk menangani beberapa permintaan). Membuat kemajuan pada lebih dari satu tugas secara bersamaan dikenal sebagai konkurensi. Go memiliki dukungan yang kaya untuk concurrency menggunakan goroutines dan saluran. 

##Goroutines

Sebuah goroutine adalah fungsi yang mampu berjalan secara bersamaan dengan fungsi lainnya. Untuk membuat goroutine kita menggunakan kata kunci go diikuti dengan doa fungsi: 
~~~go
paket utama

 impor "fmt"

 func f (n int) {
   untuk i: = 0;  i <10;  i ++ {
     fmt.Println (n, ":", i)
   }
 }

 func main () {
   pergi f (0)
   string input var
   fmt.Scanln (& input)
 } 
 ~~~
 Program ini terdiri dari dua goroutines. The goroutine pertama adalah implisit dan fungsi utama itu sendiri. The goroutine kedua tercipta saat kita sebut go f(0) . Biasanya ketika kita memanggil fungsi program kami akan mengeksekusi semua pernyataan dalam fungsi dan kemudian kembali ke garis berikutnya doa tersebut. Dengan goroutine kita segera kembali ke baris berikutnya dan jangan menunggu untuk fungsi untuk menyelesaikan. Inilah sebabnya mengapa panggilan ke Scanln fungsi telah dimasukkan; tanpa itu program akan keluar sebelum diberi kesempatan untuk mencetak semua nomor.

Goroutines yang ringan dan kita dapat dengan mudah membuat ribuan dari mereka. Kita dapat memodifikasi program kami untuk menjalankan 10 goroutines dengan melakukan hal ini:
~~~bash
  func main () {
   untuk i: = 0;  i <10;  i ++ {
     pergi f (i)
   }
   string input var
   fmt.Scanln (& input)
 } 
~~~
Anda mungkin telah memperhatikan bahwa ketika Anda menjalankan program ini tampaknya untuk menjalankan goroutines dalam rangka daripada secara bersamaan. Mari kita tambahkan beberapa penundaan untuk fungsi menggunakan time.Sleep dan rand.Intn : 
~~~bash
paket utama

 impor (
   "Fmt"
   "waktu"
   "Matematika / rand"
 )

 func f (n int) {
   untuk i: = 0;  i <10;  i ++ {
     fmt.Println (n, ":", i)
     amt: = time.Duration (rand.Intn (250))
     time.sleep (time.Millisecond * amt)
   }
 }

 func main () {
   untuk i: = 0;  i <10;  i ++ {
     pergi f (i)
   }
   string input var
   fmt.Scanln (& input)
 } 
~~~
f mencetak angka dari 0 sampai 10, menunggu antara 0 dan 250 ms setelah masing-masing. The goroutines sekarang harus dijalankan secara bersamaan.
saluran

Saluran menyediakan cara bagi dua goroutines untuk berkomunikasi dengan satu sama lain dan melakukan sinkronisasi eksekusi mereka. Berikut adalah contoh program menggunakan saluran:
~~~bash
  paket utama

 impor (
   "Fmt"
   "waktu"
 )

 func pinger (c chan string) {
   untuk i: = 0;  ;  i ++ {
     c <- "ping"
   }
 }

 printer func (c chan string) {
   untuk {
     msg: = <- c
     fmt.Println (msg)
     time.sleep (time.Second * 1)
   }
 }

 func main () {
   var c chan string = make (chan string)

   pergi pinger (c)
   pergi printer (c)

   string input var
   fmt.Scanln (& input)
 } 
~~~
Program ini akan mencetak "ping" selamanya (tekan enter untuk menghentikannya). Sebuah jenis saluran diwakili dengan kata kunci chan diikuti oleh jenis hal-hal yang diteruskan saluran (dalam hal ini kita mengirimkan string). The <- (panah kiri) operator yang digunakan untuk mengirim dan menerima pesan pada saluran. c <- "ping" berarti mengirim "ping" . msg := <- c berarti menerima pesan dan menyimpannya dalam msg . The fmt baris bisa juga ditulis seperti ini: fmt.Println(<-c) dalam hal ini kita bisa menghapus baris sebelumnya.

Menggunakan saluran seperti ini mensinkronisasikan dua goroutines. Ketika pinger mencoba untuk mengirim pesan pada saluran itu akan menunggu sampai printer siap untuk menerima pesan. (Ini dikenal sebagai blocking) Mari menambahkan pengirim lain untuk program dan melihat apa yang terjadi. Tambahkan fungsi ini:
~~~bash
  func ponger (c chan string) {
   untuk i: = 0;  ;  i ++ {
     c <- "pong"
   }
 } 
~~~
Dan memodifikasi main :
~~~bash
  func main () {
   var c chan string = make (chan string)

   pergi pinger (c)
   pergi ponger (c)
   pergi printer (c)

   string input var
   fmt.Scanln (& input)
 } 
~~~
Program ini sekarang akan bergantian mencetak "ping" dan "pong".
saluran Arah

Kita dapat menentukan arah pada jenis saluran sehingga membatasi itu baik mengirim atau menerima. Misalnya fungsi tanda tangan pinger ini dapat diubah untuk ini:

  func pinger (c chan <- string) 

Sekarang c hanya dapat dikirim ke. Mencoba untuk menerima dari c akan mengakibatkan kesalahan kompilator. Demikian pula kita dapat mengubah printer ini:

  func printer (c <-chan string) 

Saluran yang tidak memiliki batasan ini dikenal sebagai bi-directional. Sebuah saluran bi-directional dapat dikirimkan ke fungsi yang mengambil kirim saja atau hanya menerima saluran, tetapi sebaliknya tidak benar.
Memilih

Go memiliki pernyataan khusus yang disebut select yang bekerja seperti switch tapi untuk saluran:
~~~bash
  func main () {
   c1: = make (chan string)
   c2: = make (chan string)

   pergi func () {
     untuk {
       c1 <- "dari 1"
       time.sleep (time.Second * 2)
     }
   } ()

   pergi func () {
     untuk {
       c2 <- "dari 2"
       time.sleep (time.Second * 3)
     }
   } ()

   pergi func () {
     untuk {
       pilih {
       kasus msg1: = <- c1:
         fmt.Println (msg1)
       kasus msg2: = <- c2:
         fmt.Println (msg2)
       }
     }
   } ()

   string input var
   fmt.Scanln (& input)
 } 
~~~
Program ini mencetak "dari 1" setiap 2 detik dan "dari 2" setiap 3 detik. select picks saluran pertama yang siap dan menerima dari itu (atau mengirim untuk itu). Jika lebih dari satu saluran siap maka secara acak mengambil mana yang akan menerima dari. Jika tidak ada saluran siap, blok pernyataan sampai satu menjadi tersedia. 
The select pernyataan sering digunakan untuk melaksanakan timeout:
~~~bash
  pilih {
 kasus msg1: = <- c1:
   fmt.Println ( "Pesan 1", msg1)
 kasus msg2: = <- c2:
   fmt.Println ( "Pesan 2", msg2)
 Kasus <- waktu.Setelah (time.Second):
   fmt.Println ( "timeout")
 } 
~~~
time.After menciptakan saluran dan setelah durasi yang diberikan akan mengirimkan waktu saat ini di atasnya. (kami tidak tertarik pada waktu sehingga kita tidak menyimpannya dalam variabel) Kita juga dapat menentukan default kasus:
~~~bash
  pilih {
 kasus msg1: = <- c1:
   fmt.Println ( "Pesan 1", msg1)
 kasus msg2: = <- c2:
   fmt.Println ( "Pesan 2", msg2)
 Kasus <- waktu.Setelah (time.Second):
   fmt.Println ( "timeout")
 default:
   fmt.Println ( "tidak siap")
 } 
~~~
Kasus default terjadi segera jika tidak ada saluran siap.
Saluran buffered

Ini juga mungkin untuk melewati parameter kedua untuk fungsi make saat membuat saluran:

  c: = membuat (chan int, 1) 

Hal ini menciptakan saluran buffered dengan kapasitas 1. saluran Biasanya yang sinkron; kedua sisi saluran akan menunggu sampai sisi lain siap. Saluran buffered adalah asynchronous; mengirim atau menerima pesan tidak akan menunggu kecuali saluran tersebut sudah penuh.
masalah

    * Bagaimana Anda menentukan arah jenis saluran?

    * Menulis sendiri Sleep fungsi menggunakan time.After .

    * Apa saluran buffered? Bagaimana Anda akan membuat satu dengan kapasitas 20? 