# Akses Basis Data

## Sekilas Tentang Basis Data



### Basis Data SQL



### Basis Data NoSQL
Pengenalan Basisdata NoSQL
NoSQL adalah sebuah konsep mengenai penyimpanan data non-relasional. Berbeda dengan model basis data relasional yang selama ini digunakan, NoSQL menggunakan beberapa metode yang berbeda-beda. 
Metode ini bergantung dari jenis database yang digunakan. 
Karena NoSQL sendiri merupakan konsep database, dan pada implementasinya, banyak jenis-jenis dari NoSQL ini.

NoSQL sangat berguna pada data-data yang terus-menerus berkembang, dimana  data tersebut sangat kompleks sehingga sebuah database relational tidak lagi bisa mengakomodir. 
Salah satu bentuknya adalah ketika suatu data saling berhubungan satu sama lain, maka akan muncul proses duplikasi data. 
Dimana data saling memanggil ke beberapa permintaan, tambahan data baru, perubahan data, dan lain-lain dengan key yang sama. 
Karena faktor hubungan antar data yang sama terjadi terus-menerus, mendorong faktor redudansi data, data menjadi berlipat-lipat(terjadi kesamaan data pada banyak tempat) dan pada akhirnya akan menyebabkan crash pada database berkonsep RDBMS (Relational database management system).

Contoh nyata penjelasan diatas bisa dilihat pada situs jejaring sosial, FACEBOOK.COM.
Pernahkah terpikirkan 'sebesar apa server penyimpanan facebook dengan anggota ribuan orang tersebut?'
'Bagaimana kalau setengah dari member facebook adalah orang-orang yang narsis sehingga dibutuhkan storage yang sangat besar hanya untuk menyimpan foto seluruh member?'
'Berapa besar ukuran yang dibutuhkan facebook pada servernya?'.

Untuk menanggulangi hal tersebut-lah kemudian dikenal suatu cara penataan Database yang disebut Basisdata NoSQL.


MACAM-MACAM DATABASE NOSQL
Dalam NoSQL, terdapat berbagai macam bentuk database dengan schema dan konsepnya sendiri. 
Sama halnya dengan database SQL yang seperti teman-teman kenal seperti MySQL, SQL Server, Oracle, PostgreSQL, dan lain-lain. 
Dalam NoSQL, terdapat MongoDB, CouchDB, Riak, dan lain-lain. 
Berikut ini adalah beberapa contoh dari database NoSQL yang sudah cukup banyak dikenal dan digunakan oleh kebanyakan orang, namun kami akan bahas beberapa saja, diantaranya yaitu:

1. MongoDB
MongoDB adalah database open source yang sudah dikembangkan sejak Oktober 2007 dan dipublikasikan pada februari 2009. 
MongoDB merupakan salah satu database yang berbasis document (Document-Oriented Database). 
Dari segi performa MongoDB 4 kali lebih cepat dibandingkan MySQL. 
Dikarenakan konsep MongoDB berbasis document maka MongoDB tidak memiliki table, kolom maupun baris namun yang ada adalah collection (tabel) dan document (record). 
Data modelnya disebut dengan BSON yang strukturnya mirip dengan JSON.

Konsep key-value yang ada pada MongoDB, setiap document otomatis memiliki index id yang unik yang membantu mempercepat proses pencarian data secara global.

"Pada MongoDb kita tidak mengenal istilah RELASI karena kita bepikir dalam kerangka dokumen. 
Semua yang kita butuhkan diletakkan didalam sebuah dokumen."

		Kelebihan MongoDB sendiri memiliki beberapa kelebihan yaitu :
		1.  Performa lebih cepat disbanding MySQL
		2.  Replikasi dilakukan secara realtime
		3.  Auto sharding à fitur memecah database yang besar menjadi beberapa bagian
		4.  Mendukung beberapa macam bahasa pemrograman
		5.  Dapat digunakan di berbagai macam platform seperti Windows, Linux dll
		6.  CRUD (Create Read Update Delete) sangat ringan
		7.  GridFS, Spesifikasi yang digunakan untuk menyimpan data yang sangat besar
		8.  Dengan konsep key-value, setiap document otomatis memiliki index id yang unik.

		Adapun Kelemahan dari MongoDB sendiri yaitu :
		1.  Tidak mendukung transaction SQL

untuk contoh Pembahasan, Cara Penggunaan dan Tambahan Topik dapat dilihat langsung pada:
https://en.wikipedia.org/wiki/MongoDB

untuk contoh tambahan Penggunaan Dokumen MongoDB dapat dilihat langsung pada:
https://docs.mongodb.com/ dan www.tutorialspoint.com/mongodb/

untuk Mendapatkan 
https://www.mongodb.com/

2. Cassandra Apache
adalah sebuah aplikasi database berbasis Bigtabel’s Data. 
Dikutip dari situs resminya, Database Apache Cassandra adalah pilihan yang tepat ketika Anda membutuhkan skalabilitas dan ketersediaan tinggi tanpa mengorbankan kinerja. 
Skalabilitas linear dan terbukti toleransi kesalahan pada perangkat keras komoditas atau infrastruktur cloud membuatnya menjadi platform yang sempurna untuk misi-data penting. 
Dukungan Cassandra untuk mereplikasi di beberapa pusat data yang terbaik di kelasnya, memberikan latency rendah untuk pengguna Anda dan ketenangan pikiran mengetahui bahwa Anda dapat bertahan hidup padam daerah. 
ColumnFamily Cassandra data model menawarkan kenyamanan indeks kolom dengan kinerja log-terstruktur update, dukungan yang kuat untuk pandangan terwujud, dan kuat built-in caching. 
Cassandra Dikembangkan oleh APACHE. Aplikasi Inilah yang digunakan facebook untuk penyimpanan miliayaran data hingga saat ini. 
(http://cassandra.apache.org/)

Secara garis besar, Basis data NoSQL Cassandra dapat menangani data dalam jumlah yang besar (Big Data). 
Cassandra secara otomatis mereplikasi Data ke simpul (node) yang mendukung replikasi di beberapa pusat data yang terkait. 
Dengan arsitektur desentralisasi seperti ini risiko kegagalan penyimpanan data dapat meminimalkan secara default.

Beberapa Fitur Utama Cassandra:
- Desentralisasi:
Setiap node di cluster memiliki peran yang sama. Tidak ada single point of failure. 
Data didistribusikan lintas cluster (sehingga setiap node berisi data yang berbeda), namun tidak ada master karena setiap node dapat melayani permintaan apapun yang sama.

- Mendukung replikasi dan replikasi di multi data center:
Strategi replikasi dapat dikonfigurasi. Cassandra dirancang sebagai sistem terdistribusi, untuk penyebaran sejumlah besar node di beberapa pusat data. 
Fitur utama dari arsitektur terdistribusi Cassandra adalah secara khusus dirancang untuk penyebaran di multi-pusat data, untuk redundansi, untuk failover dan pemulihan dari bencana.

- Skalabilitas:
Membaca dan menulis throughput yang baik meningkat secara linear ketika mesin baru ditambahkan, tanpa downtime atau dampak gangguan pada aplikasi.

- Fault-tolerant:
Data secara otomatis direplikasi ke beberapa node untuk toleransi kesalahan. Replikasi di beberapa pusat data didukung. Node gagal dapat diganti tanpa downtime.

- Penyetelan secara konsisten:
Menulis dan membaca menawarkan tingkat penyetelan secara konsisten, sepanjang jalan dari “menulis tidak pernah gagal” untuk “blok untuk semua replika agar dapat dibaca”, dengan tingkat kuorum di tengah.

- Dukungan MapReduce:
Cassandra memiliki kemampuan untuk di integrasi dengan Hadoop, berikut dukungan MapReduce. Juga mendukung Apache Pig dan Apache Hive.

- Bahasa query:
Memperkanalkan CQL (Cassandra Query Language) diperkenalkan, merupakan SQL-like alternatif terhadap antarmuka RPC tradisional. 
Bahasa driver tersedia untuk Java (JDBC), Python (DBAPI2) dan Node.js (Helenus).

### Basis Data NewSQL



## Paket Standar Go untuk Akses Basis Data




## Driver Go untuk Akses Basis Data




## Go dan SQL: PostgreSQL



## Go dan NoSQL: RethinkDB




