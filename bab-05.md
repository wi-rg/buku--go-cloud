# Penanganan Kesalahan

## Penggunaan Kode Error


## Panic dan Recover 

Go menyediakan konstruksi *panic* dan *recover* untuk menangani kesalahan yang tidak bisa "ditolerir". Sebagai contoh, jika aplikasi kita mutlak memerlukan suatu file konfigurasi dan file konfigurasi tersebut tidak ada, maka kita bisa menggunakan *panic* untuk menghentikan eksekusi aplikasi dan kemudian memberikan pesan kesalahan. Jika kita hanya menggunakan *panic*, maka program akan berakhir dengan pesan error serta dump dari instruksi biner yang menyebabkan error tersebut. Hal ini seringkali tidak dikehendaki sehingga diperlukan *recover* untuk mengakhiri program dengan baik.



