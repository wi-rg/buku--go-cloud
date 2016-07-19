# Akses Basis Data
Database atau basis data adalah kumpulan data yang disimpan secara sistematis di dalam komputer yang dapat diolah atau dimanipulasi menggunakan perangkat lunak (program aplikasi) untuk menghasilkan informasi. Pendefinisian basis data meliputi spesifikasi berupa tipe data, struktur data dan juga batasan-batasan data yang akan disimpan. Basis data merupakan aspek yang sangat penting dalam sistem informasi dimana basis data merupakan gudang penyimpanan data yang akan diolah lebih lanjut. Basis data menjadi penting karena dapat mengorganisasi data, menghidari duplikasi data, hubungan antar data yang tidak jelas dan juga update yang rumit.
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

 DBMS merupakan perantara antara user dengan database.
 Cara komunikasi diatur dalam suatu bahasa khusus yang telah ditetapkan oleh DBMS.
    Contoh: SQL, dBase, QUEL, dsb.
 Bahasa database, dibagi dalam 2 bentuk:
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

Sejarah SQL
Sejarah SQL dimulai dari artikel seorang peneliti dari IBM bernama Jhonny Oracle yang membahas tentang ide pembuatan basis data relasional pada bulan Juni 1970. Artikel ini juga membahas kemungkinan pembuatan bahasa standar untuk mengakses data dalam basis data tersebut. Bahasa tersebut kemudian diberi nama SEQUEL (Structured English Query Language).

Setelah terbitnya artikel tersebut, IBM mengadakan proyek pembuatan basis data relasional berbasis bahasa SEQUEL. Akan tetapi, karena permasalahan hukum mengenai penamaan SEQUEL, IBM pun mengubahnya menjadi SQL. Implementasi basis data relasional dikenal dengan System/R.

Di akhir tahun 1970-an, muncul perusahaan bernama Oracle yang membuat server basis data populer yang bernama sama dengan nama perusahaannya. Dengan naiknya kepopuleran John Oracle, maka SQL juga ikut populer sehingga saat ini menjadi standar de facto bahasa dalam manajemen basis data.

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

Contoh:
Diasumsikan terdapat tabel user yang berisi data sebagai berikut.

username 	passwd 	tanggal_lahir 	jml_transaksi 	total_transaksi
Aris 	6487AD5EF 	09-09-1987 	6 	10.000
Budi 	97AD4erD 	01-01-1994 	0 	0
Charlie 	548794654 	06-12-1965 	24 	312.150
Daniel 	FLKH947HF 	24-04-1980 	3 	0
Erik 	94RER54 	17-08-1945 	34 	50.000

Contoh 1: Tampilkan seluruh data.

    SELECT *
    FROM user

Contoh 2: Tampilkan pengguna yang tidak pernah bertransaksi.

    SELECT *
    FROM user
    WHERE total_transaksi = 0

Contoh 3: Tampilkan username pengguna yang bertransaksi kurang dari 10 dan nilainya lebih dari 1.000.

    SELECT username
    FROM user
    WHERE jml_transakai < 10 AND total_transaksi > 1000

Contoh 4: Tampilkan total nominal transaksi yang sudah terjadi.

    SELECT SUM(total_transaksi) AS total_nominal_transaksi
    FROM user

Contoh 5: Tampilkan seluruh data diurutkan berdasarkan jumlah transaksi terbesar ke terkecil.

    SELECT *
    FROM user
    ORDER BY jml_transaksi DESC

Fungsi aggregat
Beberapa SMBD memiliki fungsi aggregat, yaitu fungsi-fungsi khusus yang melibatkan sekelompok data (aggregat). Secara umum fungsi aggregat adalah:

    SUM untuk menghitung total nominal data
    COUNT untuk menghitung jumlah kemunculan data
    AVG untuk menghitung rata-rata sekelompok data
    MAX dan MIN untuk mendapatkan nilai maksimum/minimum dari sekelompok data.

Fungsi aggregat digunakan pada bagian SELECT. Syarat untuk fungsi aggregat diletakkan pada bagian HAVING, bukan WHERE.
Subquery
Ada kalanya query dapat menjadi kompleks, terutama jika melibatkan lebih dari satu tabel dan/atau fungsi aggregat. Beberapa SMBD mengizinkan penggunaan subquery. Contoh:
Tampilkan username pengguna yang memiliki jumlah transaksi terbesar.

    SELECT username
    FROM user
    WHERE jml_transaksi =
    (
    SELECT MAX(jml_transaksi)
    FROM user
    )

INSERT

Untuk menyimpan data dalam tabel dipergunakan sintaks:

    INSERT INTO [NAMA_TABLE] ([DAFTAR_FIELD]) VALUES ([DAFTAR_NILAI])

Contoh:

    INSERT INTO TEST (NAMA, ALAMAT, PASSWORD) VALUES ('test', 'alamat', 'pass');

UPDATE
Untuk mengubah data menggunakan sintaks:

    UPDATE [NAMA_TABLE] SET [NAMA_KOLOM]=[NILAI] WHERE [KONDISI] 

Contoh:

    UPDATE Msuser set password="123456" where username="abc"

DELETE
Untuk menghapus data dipergunakan sintaks:

    DELETE FROM [nhew andiz] [KONDISI]

Contoh:

    DELETE FROM TEST WHERE NAMA='test';

Overall saya ucapkan banyak terima kasih telah mengizinkan saya untuk menduplicate dan mensunting sedikit isi artikel dari situs belajar terkenal di seluruh dunia. Thanks to : Id.Wikipedia.org


### Basis Data NoSQL

### Basis Data NewSQL

## Paket Standar Go untuk Akses Basis Data

## Driver Go untuk Akses Basis Data

## Go dan SQL: PostgreSQL
PostgreSQL adalah sebuah sistem basis data yang disebarluaskan secara bebas menurut Perjanjian lisensi BSD. Piranti lunak ini merupakan salah satu basis data yang paling banyak digunakan saat ini, selain MySQL dan Oracle. PostgreSQL menyediakan fitur yang berguna untuk replikasi basis data. Fitur-fitur yang disediakan PostgreSQL antara lain DB Mirror, PGPool, Slony, PGCluster, dan lain-lain.

PostgreSQL adalah sistem database yang kuat untuk urusan relasi, open source. Memiliki lebih dari 15 tahun pengembangan aktif dan sudah terbukti segala rancangan arsitekturnya telah mendapat reputasi tentang “kuat”, “handal”, “integritas data”, dan “akurasi data”

Pengguna PostgreSQL

    * Yahoo! untuk analisa prilaku pengguna web, menyimpan 2 petabyte data dan mengklaim sebagai gudang data terbesar. Menggunakan versi PostgreSQL yang dimodifikasi, dengan engine penyimpanan berbasis kolom yang sepenuhnya berbeda.[6][7]
    *MySpace, situs jejaring sosial populer, menggunakan basisdata Aster nCluster untuk gudang data, dibangun diatas PostgreSQL tanpa modifikasi.[8][9]
    *OpenStreetMap, proyek kolaboratif untuk menciptakan peta dunia yang bebas sunting.[10]
    *Afilias, register domain untuk .org, .info, dan sebagainya.[11]
    *Sony Online multiplayer online game.[12]
    *BASF, platform belanja untuk portal agribisnisnya.[13]
    *hi5.com portal jejaring sosial.[14]
    *Skype aplikasi VoIP, basisdata pusat bisnis.[15]
    *Sun xVM, perangkat lunak virtualisasi dan otomasi datacenter milik Sun.

## Go dan NoSQL: RethinkDB


###NoSQL MongoDB
