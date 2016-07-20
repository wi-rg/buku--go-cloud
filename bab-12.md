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
Bidang Manajemen System Database pada hari ini sudah berkembang begitu cepat sejak pertama kali dikenalkan kepada Dunia sebagai DBMS. Tidak hanya dapat menyelesiakan bentuk System Database NoSQL pada bentuk-bentuk berkeda(Kolom, Dokumen, Key/Value dan Database berbentuk Grafik), tapi juga dapat menyelesaikan masalah dan melihat Pasar yang tengah berkembang berdasarkan Perhitungan dan Memori yang terdapat pada Manajemen Database itu sendiri, dimana DBMS bergantung pada Memory yang paling utama yang dia miliki daripada kepada Memory Penyimpanan, Manajemen Memori dan bentuk Manipulasi Memory. 
Setelah sekian lama mengenal bentuk Database DBMS yang sudah menjadi semakin umum, sekarang muncul lagi bentuk Kategori baru dari DBMS yang baru saja ber-evolusi belakangan ini, dikenal dengan sebutan "NewSQL" .

Apa itu NewSQL ?

NewSQL adalah bentuk Bahasa AKses Database baru yang digadang-gadang lebih baik, hebat dan lebih cerdas dibandingkan dengan SQL. Banyak para Profesional di bidang Database yang menyatakan bahwa SQL, yang berupa Bahasa Database berbasis-Objek, sudah terlalu Usang dan terlalu Rumit untuk digunakan .

Database NewSQL ini pada dasarnya memiliki Performa yang sama dengan yang dikenal dengan Sistem NoSQL dan yang menyediakan para Administrator sebuah Fasilitas yang Detil, Konsisten, Terisolasi(aman), Tahan lama atau yang sering dikenal sebagai bergaransi Berkemampuan ACID (Atomicity, Consistency, Isolation & Durability) .

Kata-kata atau kalimat baru NewSQL dicanangkan oleh Matthew Aslett, seorang analis yang memiliki 'Group 451'. Beliau yang mencanangkan bentuk pengertian produk baru yang menggunakan Data Model Relasi sambil tetap memiliki interface/Antarmuka seperti yang dimiliki SQL.

Aslett menjelaskan:
" 'NewSQL' adalah kependekan tangan dari berbagai bentuk dan jenis baru yang memiliki Performa Tinggi untuk para Vendor Database SQL. Kami sebelumnya sering menyebut produk ini sebagai 'SQLterukur' atau ScalableSQL untuk membedakannya dari bentuk Produk Database Relasional yang biasanya. Dikarenakan Database ini mengedepankan Skalabilitas, yang sebenarnya adalah fitur yang tidak begitu diperlukan ada pada semua produk, kami kemudian menggunakan (istilah) 'NewSQL' [...] Dan untuk mengklarifikasi perbedaannya, seperti NoSQL, NewSQL agar tidak dipahami secara harafiah: sebagai sesuatu yang baru dari Vendor SQL adalah Vendor-nya itu sendiri, bukan SQL-nya. "

NewSQL didesain untuk mendukung SQL dalam upaya menghandel masalah yang paling baru dengan cara transaksi online Tradisional, yang lebih Spesifik, menitik beratkan pada Skalabilitas dan Performa mereka. Memang sedikit mirip dengan NoSQL, kata beberapa Advokat menambahkan.

Jadi, kesimpulannya adalah bahwa NewSQL dibuat dan akan sangat baik saat digunakan perusahaan yang berjalan dan tertarik untuk mulai bermigrasi/memindahkan penggunaan aplikasinya kepada Platform pengembangan tentang Big Data, yang kemudian pada aplikasinya akan mengembangkan penggunaan/aplikasi baru pada skalabilitas Transaksi Online pada proses Sistemnya dan berharap untuk tetap menggunakan pengetahuan mereka yang sudah ada sebelumnya dalam bidang Proses Transaksi Online, menurut keterangan Prasanna Venkatesh dan Nirmala S.

di akhir, "NewSQL" berarti untuk meng-kategorikan beberapa bentuk dari  grup-grup yang mirip, kata ASlett. Ini bukan masalah tentang membandingkan NewSQL melawan NoSQL, karena perbedaannya antara NoSQL dan NewSQL sangat tipis hingga pada titik dimana kita mengharapkan bahwa Pengertian tentang NoSQL dan NewSQL akan menjadi lebih jelas dan Relevan saat fokus penggunaannya menjadi lebih Spesifik.


## Paket Standar Go untuk Akses Basis Data




## Driver Go untuk Akses Basis Data




## Go dan SQL: PostgreSQL



## Go dan NoSQL: RethinkDB




