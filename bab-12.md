# Akses Basis Data
Basis data merupakan aspek yang sangat penting dalam sistem informasi dimana basis data merupakan gudang penyimpanan data yang akan diolah lebih lanjut. Basis data menjadi penting karena dapat mengorganisasi data, menghidari duplikasi data, hubungan antar data yang tidak jelas dan juga update yang rumit.
## Sekilas Tentang Basis Data
Apa itu basis data ?
Basis data atau yang juga biasa di kenal database adalah himpunan kelompok data, yang diorganisasi sedemikian rupa, sehingga kelak dapat dimanfaatkan kembali dengan cepat.

Komponen sistem dalam basis data meliputi :
Perangkat keras (Hardware)
Sistem Operasi (OS)
Basis Data
DBMS (Database Management System)
Pemakai (User)
Perangkat lunak lain

Bahasa Basis Data :

? DBMS merupakan perantara antara user dengan database.
? Cara komunikasi diatur dalam suatu bahasa khusus yang telah ditetapkan oleh DBMS.
    Contoh: SQL, dBase, QUEL, dsb.
? Bahasa database, dibagi dalam 2 bentuk:
    - Data Definition Language (DDL)
       Digunakan dalam membuat tabel baru, indeks, mengubah tabel, menetukan struktur tabel, dsb.
    - Data Manipulation Language (DML)
Digunakan dalam memanipulasi dan pengambilan data pada database.
Manipulasi data, dapat mencakup:
    - Pemanggilan data yang tersimpan dalam database (query)
    - Penyisipan/penambahan data baru ke database
    - Penghapusan data dari database
    - Pengubahan data pada database


### Basis Data SQL
Apa itu SQL? / Pengertian SQL
SQL (Structured Query Language) adalah sebuah bahasa yang dipergunakan untuk mengakses data dalam basis data relasional. Bahasa ini secara de facto merupakan bahasa standar yang digunakan dalam manajemen basis data relasional. Saat ini hampir semua server basis data yang ada mendukung bahasa ini untuk melakukan manajemen datanya.
Standarisasi SQL
Standarisasi SQL dimulai pada tahun 1986, ditandai dengan dikeluarkannya standar SQL oleh ANSI. Standar ini sering disebut dengan SQL86.Standar tersebut kemudian diperbaiki pada tahun 1989 kemudian diperbaiki lagi pada tahun 1992. Versi terakhir dikenal dengan SQL92. Pada tahun 1999 dikeluarkan standar baru yaitu SQL99 atau disebut juga SQL99, akan tetapi kebanyakan implementasi mereferensi pada SQL92.

Saat ini sebenarnya tidak ada server basis data yang 100% mendukung SQL92. Hal ini disebabkan masing-masing server memiliki dialek masing-masing.

Pemakaian Dasar SQL
Secara umum, SQL terdiri dari dua bahasa, yaitu Data Definition Language (DDL) dan Data Manipulation Language (DML). Implementasi DDL dan DML berbeda untuk tiap sistem manajemen basis data (SMBD)[1], namun secara umum implementasi tiap bahasa ini memiliki bentuk standar yang ditetapkan ANSI. Artikel ini akan menggunakan bentuk paling umum yang dapat digunakan pada kebanyakan SMBD.

Data Definition Language (DDL)
DDL digunakan untuk mendefinisikan, mengubah, serta menghapus basis data dan objek-objek yang diperlukan dalam basis data, misalnya tabel, view, user, dan sebagainya. Secara umum, DDL yang digunakan adalah CREATE untuk membuat objek baru, USE untuk menggunakan objek, ALTER untuk mengubah objek yang sudah ada, dan DROP untuk menghapus objek. DDL biasanya digunakan oleh administrator basis data dalam pembuatan sebuah aplikasi basis data.

CREATE

CREATE digunakan untuk membuat basis data maupun objek-objek basis data. SQL yang umum digunakan adalah:

CREATE DATABASE membuat sebuah basis data baru. 
CREATE DATABASE nama_basis_data
CREATE TABLE membuat tabel baru pada basis data yang sedang aktif.
CREATE TABLE nama_tabel
Secara umum, perintah ini memiliki bentuk:

CREATE TABLE [nama_tabel]
(
nama_field1 tipe_data [constraints][,
nama_field2 tipe_data,
...]
)
atau

CREATE TABLE [nama_tabel]
(
nama_field1 tipe_data [,
nama_field2 tipe_data,
...]
[CONSTRAINT nama_field constraints]
)
dengan:

nama_field adalah nama kolom (field) yang akan dibuat. Beberapa sistem manajemen basis data mengizinkan penggunaan spasi dan karakter nonhuruf pada nama kolom.

tipe_data tergantung implementasi sistem manajemen basis data. Misalnya, pada MySQL, tipe data dapat berupa VARCHAR, TEXT, BLOB, ENUM, dan sebagainya.

constraints adalah batasan-batasan yang diberikan untuk tiap kolom. Ini juga tergantung implementasi sistem manajemen basis data, misalnya NOT NULL, UNIQUE, dan sebagainya. Ini dapat digunakan untuk mendefinisikan kunci primer (primary key) dan kunci asing (foreign key).

Satu tabel boleh tidak memiliki kunci primer sama sekali, namun sangat disarankan mendefinisikan paling tidak satu kolom sebagai kunci primer.

Data Manipulation Language (DML)
DML digunakan untuk memanipulasi data yang ada dalam suatu tabel. Perintah yang umum dilakukan adalah:
SELECT untuk menampilkan data
INSERT untuk menambahkan data baru
UPDATE untuk mengubah data yang sudah ada
DELETE untuk menghapus data
SELECT

SELECT adalah perintah yang paling sering digunakan pada SQL, sehingga kadang-kadang istilah query dirujukkan pada perintah SELECT. SELECT digunakan untuk menampilkan data dari satu atau lebih tabel, biasanya dalam sebuah basis data yang sama.

SELECT adalah yang perintah paling sering SQL digunakan dan mempunyai format umum yang berikut:

SELECT    [DISTINCT / ALL] {* I [columnExpression [AS newName]] [,…..]}
FORM    TableName [alias] [,…..]
[WHERE    condition]
[GROUP BY Column list] [HAVING condition]
[ORDER BY Column list]

columnExpression menampllkan suatu nama kolom atau suatu ungkapan, TableName adalah nama dari suatu tabel database ada atau view bahwa mempunyai untuk akses, dan alias adalah suatu singkatan opsional untuk TableName. Urutan memproses suatu statement SELECT adalah :

FROM           Menetapkan tabel atau tabel  itu untuk digunakan
WHERE         Menyaring baris yang tunduk kepada beberapa kondisi 
GROUP BY    Membentuk kelompok baris dengan kolom yang sama nilainya
HAVING        Menyaring kelompok yang tunduk kepada beberapa kondisi 
SELECT         Menetapkan kolom yang adalah untuk nampak keluaran
ORDER BY   Menetapkan order/ pesanan keluaran

Order yang menentukan di dalam statemen SELECT  tidak bisa diubah. Satu-Satunya dua ketentuan wajib dua hal pertama itu: : SELECT and FROM : sisanya adalah opsional 

Perintah select bisa digunakan dengan:
kondisi adalah syarat yang harus dipenuhi suatu data agar ditampilkan.
kondisi_aggregat adalah syarat khusus untuk fungsi aggregat.
Kondisi dapat dihubungkan dengan operator logika, misalnya AND, OR, dan sebagainya.


### Basis Data NoSQL
Mari Mengenal Database NoSQL (Not Only SQL) - Pada kesempatan kali ini, kita akan membahas tentang sebuah teknologi yang sangat menarik dalam dunia pengolahan database yaitu NoSQL. Jika sebelumnya kita mengenal berbagai macam jenis database sebagai RDBMS (Relational Database Management System) dengan teknologi utamanya adalah SQL (Structured Query Language). Database NoSQL merupakan lawan dari SQL ini. NoSQL merupakan singkatan dari Not Only SQL, dan menurut beberapa sumber NoSQL adalah sebuah konsep mengenai penyimpanan data non-relasional yang berbeda dengan model basis data relasional yang selama ini digunakan. NoSQL menggunakan beberapa metode yang berbeda-beda. Metode ini bergantung dari jenis database yang digunakan. Karena NoSQL sendiri merupakan konsep database dan pada implementasinya, banyak jenis-jenis dari NoSQL ini. 

Kelebihan utama dari NoSQL adalah untuk menangani Big Data (data dengan jumlah besar yang melebihi proses kapasitas konversi dari sistem database yang ada) dimana data terus-menerus berkembang dan dimana data tersebut sangat kompleks sehingga sebuah database relational tidak lagi bisa mengakomodir. Salah satu bentuknya adalah ketika suatu data saling berhubungan satu sama lain, maka akan muncul proses duplikasi data. Dimana data saling memanggil ke beberapa permintaan, tambahan data baru, perubahan data dan lain-lain dengan key yang sama. 

Karena faktor hubungan antar data yang sama terjadi terus-menerus, mendorong faktor redudansi/duplikasi data, data menjadi berlipat-lipat, dan pada akhirnya akan menyebabkan crash pada database yang berkonsep RDBMS itu, maka harus ada cara lain untuk menanggulangi masalah ini, dan NoSQL inilah solusinya. NoSQL menyederhanakan proses yang terjadi dalam sistem basis data relasional. Dimana hal-hal yang menyebabkan redudansi/duplikasi dihilangkan sehingga trafik server akan seimbang. Penyederhanaan proses ini memungkinkan data direplikasi di banyak server secara mudah dan menjamin ketersediaan data. 

Kelebihan lain yang dimiliki oleh database NoSQL ini adalah kecepatan dalam hal pencarian data. Sebagai contoh Google, salah satu perusahaan yang telah mengaplikasikan NoSQL dalam database pencariannya. Pencarian yang ada di Google sangatlah cepat. Hanya dalam beberapa detik saja, data yang ingin kita cari sesuai keyword akan muncul, padahal kita tahu sendiri database Google sangatlah besar. 

### Basis Data NewSQL

Lanskap untuk sistem manajemen database berubah lebih cepat hari ini daripada sejak hari-hari awal DBMS relasional. Kita tidak hanya memiliki serangan sistem database NoSQL dari berbagai bentuk yang berbeda (kolom, dokumen, kunci / nilai, dan database grafik), tapi kami juga melihat pasar yang sedang berkembang untuk manajemen database di memori, di mana DBMS bergantung pada memori utama bukannya disk untuk penyimpanan data, manajemen, dan manipulasi. Tapi ada yang lain "kategori" dari DBMS berkembang yang dikenal sebagai "NewSQL."

The NewSQL jangka membingungkan dan tidak sangat baik didefinisikan. Pada dasarnya, NewSQL DBMSs ditandai oleh relasional, sistem database SQL dengan terdistribusi, arsitektur kesalahan-toleran. Di memori kemampuan adalah fitur tambahan khas persembahan NewSQL, seperti yang berkerumun layanan database dengan kemampuan akan dikerahkan di awan. produk NewSQL DBMS mungkin memiliki fitur yang lebih sedikit dan komponen, dengan jejak yang ringan, membuat mereka lebih mudah untuk mendukung dan belajar dari warisan persembahan relasional.

Tentu saja, fitur yang disebutkan di sini tidak eksklusif untuk sistem database NewSQL, juga tidak semua dari mereka diperlukan untuk dianggap sebagai produk NewSQL. Banyak produk RDBMS yang ada mendukung banyak fitur ini.

Secara umum, sistem database NewSQL dirancang untuk memerangi serangan NoSQL menggunakan teknologi relasional. Jenis aplikasi yang dirancang untuk mendapatkan keuntungan dari NewSQL DBMS biasanya akan memiliki sejumlah besar transaksi pendek yang mengakses sejumlah kecil data diindeks dan melaksanakan berulang-ulang. Jelas, yang merupakan bagian besar dari aplikasi (tetapi tidak berarti semua aplikasi). Tujuan dari NewSQL adalah untuk memberikan ketersediaan tinggi dan kinerja untuk data modern tanpa mengorbankan persyaratan konsistensi yang kuat dan kemampuan transaksi yang NoSQL sering skimps pada.
Dengan kata lain, NewSQL DBMSs memberikan transaksi ACID. ACID adalah akronim untuk atomicity, konsistensi, isolasi, dan daya tahan. Masing-masing dari empat kualitas ini diperlukan untuk transaksi untuk dapat memastikan integritas data:

Atomicity berarti bahwa transaksi harus menunjukkan "semua atau tidak" perilaku. Baik semua petunjuk dalam transaksi terjadi, atau tidak satupun dari mereka terjadi. Atomicity mempertahankan "kelengkapan" dari proses bisnis.
Konsistensi mengacu pada keadaan data sebelum dan setelah transaksi dijalankan. Sebuah transaksi mempertahankan konsistensi keadaan data. Dengan kata lain, setelah transaksi dijalankan, semua data dalam database adalah "benar."
Isolasi berarti bahwa transaksi dapat berjalan pada waktu yang sama. Transaksi berjalan secara paralel memiliki ilusi bahwa tidak ada concurrency. Dengan kata lain, tampak bahwa sistem berjalan hanya satu transaksi pada suatu waktu. Tidak ada transaksi konkuren lainnya memiliki visibilitas ke database modifikasi uncommitted dibuat oleh transaksi lainnya. Untuk mencapai isolasi, mekanisme penguncian diperlukan.
Daya tahan mengacu pada dampak pemadaman atau kegagalan pada transaksi berjalan. Sebuah transaksi tahan lama tidak akan mempengaruhi keadaan data jika transaksi berakhir tidak normal. Data akan bertahan kegagalan apapun.
Jadi, NewSQL DBMS relasional, memberikan ACID, dan menawarkan kemampuan modern dan arsitektur modern. Dan, tentu saja, menggunakan SQL untuk mengakses dan memodifikasi data.

Meskipun tidak ada buku aturan untuk apa yang merupakan NewSQL, ada beberapa jenis yang berbeda dari sistem yang memanfaatkan moniker NewSQL. Paling umum, NewSQL mengacu pada produk RDBMS yang sama sekali baru yang beroperasi di cluster didistribusikan shared-

tidak ada node, dengan setiap node memiliki subset dari data. Dengan merancang NewSQL RDBMS dari awal, kode dapat tegas ditulis untuk arsitektur terdistribusi dan menghindari beberapa overhead dari RDBMS tradisional. Contoh produk NewSQL termasuk VoltDB, Clustrix, MemSQL, dan Translattice.

Istilah NewSQL juga dapat berlaku untuk mesin penyimpanan SQL seperti Infobright dan sistem sharding transparan seperti Scalebase.


## Paket Standar Go untuk Akses Basis Data
Paket standar go untuk akses basis data adalah sql. 
Paket sql menyediakan antarmuka umum di basisdata sql .
Paket sql harus digunakan bersamaan dengan driver database.



## Driver Go untuk Akses Basis Data
Ada banyak sekali driver go untuk akses basis data, seperti :
- Apache Phoenix/Avatica 
- Couchbase N1QL
- MySQL
- DB2
- Firebird SQL



## Go dan SQL: PostgreSQL



## Go dan NoSQL: RethinkDB




