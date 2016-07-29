# Memahami Microservices



## Apa itu Microservices? Apa perbedaan Microservices dengan Molonithic?

### Monolithic

Suatu desain aplikasi yang memiliki arsitektur yang terpusat dan teknologi yang seragam, atau lebih dikenal dengan istilah Monolithic application architecture. Arsitektur aplikasi monolitik ini menggunakan kode sumber dan teknologi yang serupa untuk menjalankan semua tugas-tugasnya.

Untuk memudahkan pemahaman kita ambil contoh aplikasi wordpress. WordPress merupakan contoh paling mudah untuk menggambarkan arsitektur aplikasi yang bersifat monolitik, dimana dalam satu aplikasi kita dapat memiliki frontend sekaligus backend. Semua fitur security, performance, manajemen konten, ad, statistik, semuanya dibangun dengan menggunakan PHP dan database MySQL, dalam kode sumber yang sama.

Perkembangan berikutnya adalah dengan memisahkan aplikasi berdasarkan user role-nya. Ada frontend yaitu aplikasi yang akan diakses oleh pengakses blog (user) dan Backend yang digunakan oleh kontributor konten dan admin blog. Bisa juga memisahkan halaman reporting, member dan lainnya, dengan tujuan agar satu fungsi tidak akan mengganggu fungsi yang lainnya.

Pada praktiknya pemisahan halaman ini dilakukan di level domain. Anggap saja kita gunakan subdomain www untuk frontend, subdomain admin untuk backend, subdomain report untuk reporting dan seterusnya. Dan masing-masing subdomain tersebut dapat dilakukan pemisahan server secara fisik. Pemisahan ini dilakukan agar proses scaling dan debuging dapat dilakukan dengan mudah.

### Microservices

Secara sederhana, arsitektur aplikasi Microservice ini menggunakan desain yang memecah aplikasi berdasarkan fungsinya secara spesifik. Tidak sekedar dengan memisahkan berdasarkan user-role atau subdomain saja, tetapi aplikasi akan di breakdown lebih rinci lagi dari sisi fungsionalitasnya. Aplikasi akan dirancang agar setiap fungsi bekerja secara independent. Dan setiap fungsi dapat menggunakan teknologi stack yang sesuai dengan kebutuhan, walaupun itu artinya akan terdapat teknologi yang berbeda-beda dalam satu aplikasi besar.
 
Dengan pemisahan aplikasi berdasarkan fungsi-nya ini, pada akhirnya kita akan menemui keragaman teknologi dalam sebuah satu layanan digital. Misalkan dari layanan blog, kita bisa coba pecahkan aplikasi blog tersebut menjadi fungsi konten, user management, komentar, rangking, search dan lainnya.

Pada bagian fungsi konten mungkin kita menggunakan PHP dan mysql, tetapi pada komentar kita menggunakan Python dan Mongodb, sedangkan di search menggunakan nodejs dan elasticsearch untuk penyimpanan datanya. Disini terlihat bahwa setiap fungsi / permasalahan teknis dapat diselesaikan dengan cara dan teknologi yang berbeda-beda.

Dalam konsep microservices, kita tidak hanya melalukan pemisahan di level aplikasi saja, tetapi dari sisi infrastruktur pula. Kita akan menemukan keragaman arsitektur infra, konfigurasi dan optimalisasi sistem yang berbeda, dan sering pula ditemukan jumlah dan spesifikasi server yang tidak sama antara service yang satu dengan yang lain.



## Implementasi

### Monolithic

Arsitektur aplikasi monolitik ini lebih sederhana dalam implementasinya. Kita bisa mulai dengan membuat cluster server yang identik kemudian menggunakan konfigurasi system dan aplikasi yang sama antara satu server dengan yang lain. Dan pastikan data (database, cache, session dan files) di simpan diluar cluster aplikasi agar aplikasi dapat berjalan sama di antara semua cluster server.

Kemudian kita ingin membagi load server dengan memisahkan aplikasi backend dan frontend. Kita akan pisahkan dengan menggunakan subdomain yang berbeda. Disini kita menggunakan domain ‘example.com’ sebagai contoh, maka kita setup subdomain www.example.com sebagai frontend dan subdomain admin.example.com sebagai backend.

Dengan menggunakan subdomain yang berbeda ini kita dapat mengarahkan IP nya ke server yang sama (VirtualHost yang beda), atau menggunakan IP yang berbeda. Karena tujuan kita untuk membagi load antar server, maka pilihan menggunakan IP yang berbeda adalah solusi yang tepat.

Sekarang kita punya IP yang berbeda untuk subdomain www.example.com dengan admin.example.com, dan kita bisa pisahkan cluster web diantara kedua aplikasi tersebut. Dan proses diatara kedua cluster tersebut tidak akan saling mempengaruhi secara langsung.

Dengan memiliki cluster yang berbeda akan lebih memudahkan kita untuk me-manage, monitoring dan membuat perencanaan scaling infrastruktur. Kita ambil contoh yang real seperti ini, andaikan akan ada promosi besar di website dan di perkirakan traffic akan melonjak drastis, maka scaling akan lebih efektif karena cukup dilakukan pada frontend saja.

### Microservices

Seperti yang sudah kita ketahui bahwa arsitektur microservices ini memecah aplikasi menjadi service-service kecil. Dengan tujuan agar setiap service tersebut bekerja secara optimal, dengan mempertimbangkan teknologi yang sesuai dengan kebutuhannya.

Aplikasi dipecah tidak sekedar hanya memisahkan antara backend dan frontend, melainkan dipecah menjadi service-service yang berfungsi secara spesifik. Apabila kita menjalankan layanan website blog, maka dengan pendekatan microservices aplikasi bisa di breakdown menjadi layanan khusus konten, statistik, ad, reporting, photo/video, messaging, searching, komentar dan lainnya.

Tetap pada saat kita memecah aplikasi menjadi layanan-layanan yang kecil, kita akan dihadapkan pada cluster aplikasi dengan stack teknologi yang beragam dengan jumlah yang banyak. Mungkin di bagian konten kita menggunakan PHP dan Mysql, tetapi di fitur statistik kita menggunakan stack teknologi seperti Python, Memcached dan Mongodb, dan begitu juga dengan layanan lainnya dengan teknologi yang di adopsinya.

Sebelumnya kita dengan mudah menggunakan subdomain untuk memisahkan antar aplikasi, tetapi apabila menggunakan teknik yang sama untuk microservices ini, maka yang akan terjadi adalah aplikasi akan terlihat fragmented karena menggunakan subdomain yang terlalu banyak.

Ada beberapa kekurangan dari penggunaan subdomain ini, yang pertama adalah aplikasi akan sulit untuk dimaintain, terutama untuk interkoneksi internal dari aplikasi. Misalnya untuk frontend kita menggunakan https://www.example.com, kemudian ada https://stat.example.com/view/5 untuk statistik, dan reporting menggunakan https://report.example.com/2016/12/05 dan sebagainya.

Selain masalah administrasi aplikasi yang terlalu rumit, dari sisi digital marketing pun dinilai tidak SEO-friendly karena menggunakan Subdomain yang terlalu banyak. Begitu pula dari user experience, sewaktu user membuka website akan menjadi lebih lama, karena setiap subdomain yang diakses akan membutuhkan waktu untuk resolving ke server DNS.

Yang perlu kita lakukan adalah membuat kumpulan layanan tersebut menjadi satu kesatuan secara logic. Dalam sebuah aplikasi web kita akan membuatnya lebih sederhana dengan hanya menggunakan satu subdomain saja (www.example.com). Sehingga semua entitas yang berinteraksi dengan aplikasi (user, aplikasi internal dan eksternal, crawler, cron) akan melihatnya seperti sebuah arsitektur aplikasi monolitik yang sederhana.

Apabila kita masih menggunakan contoh yang diatas, maka aplikasi frontend tetap dapat diakses melalui https://www.example.com, statistik dengan alamat https://www.example.com/stat/view/5, dan reporting menggunakan alamat https://www.example.com/report/2016/12/05, atau gallery diakses dengan alamat http://www.example.com/gallery.php. Dari sini terlihat semua layanan walaupun menggunakan stack teknologi yang berbeda-beda, tetapi tetap diakses melalui subdomain/domain yang sama.

## Kesimpulan
Migrasi arsitektur aplikasi dari monolithic ke microservices mungkin akan menambah pekerjaan untuk developer dan system architect. Tetapi hasil yang didapat akan mempengaruhi proses kerja dan hasilnya akan terasa untuk jangka panjang. Sekilas mungkin arsitektur ini akan terlihat lebih rumit, tetapi arsitektur yang seperti inilah yang seharusnya dijadikan standar dalam merancang aplikasi yang besar. Pada saat kita sudah terbiasa dengan arsitektur ini, selanjutnya kita akan lebih fokus pada fungsi-fungsi yang spesifik.
Walaupun kita menyebutnya sebagai arsitektur aplikasi, tetapi pada praktiknya akan melibatkan tim infrastruktur dalam mengimplementasikannya. Tools seperti load balancer menjadi kunci untuk menjadikan microservices menjadi terlihat lebih terintegrasi. Tetapi dibalik kesederhanaanya itu, cluster web akan menjadi lebih mudah di maintain dan proses-proses yang terjadi akan berjalan pada cluster yang independent
