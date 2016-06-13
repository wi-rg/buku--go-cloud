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
Apa itu NOSQL

NOSQL menurut Wikipedia adalah sistem menejemen database yang berbeda dari sistem menejemen database relasional yang klasik dalam beberapa hal. NOSQL mungkin tidak membutuhkan skema tabel dan umumnya menghindari operasi join dan berkembang secara horisontal. Akademisi menyebut database seperti ini sebagai structured storage, istilah yang didalamnya mencakup sistem menejemen database relasional. NOSQL adalah database generasi terbaru yang mengarahkan kepada database yang tidak berelasi (non-relational), dapat disebarkan kepada siapapun (open-source) dan berskala horisontal (horizontal scale).

Johan Oskarsson dari Last.fm memperkenalkan kembali istilah NOSQL pada awal 2009 ketika ia menyelenggarakan sebuah acara untuk membahas “Distributed Open Source dan Non-relational Database”. Nama berusaha untuk label munculnya peningkatan jumlah non-relasional, didistribusikan menyimpan data, termasuk kloning open source dari Google Bigtable/MapReduce dan Amazon Dynamo.

Berbeda dengan basis data SQL dimana meskipun berbeda-beda pembuat namun cara kerja NOSQL maupun cara penggunaannya relatif sama. Contohnya sama-sama menggunakan tabel yang dihubungkan oleh relasi-relasi, manipulasi data dengan bahasa SQL dan sb. Basis data NOSQL bisa sangat berbeda satu sama lain.

Sebagian besar sistem NOSQL pada masa awal tidak berusaha untuk memberikan atomicity, konsistensi, isolasi dan daya tahan jaminan, bertentangan dengan praktik yang berlaku di antara sistem database relasional. Namun di kemudian hari, beberapa database NOSQL dengan pembaruan terkini sudah mampu mengintegrasikan database yang non-relasional ke dalam bentuk database relasional sehingga dapat mempermudah pengguna yang masih belum akrab dengan bahasa standar yang diterapkan NOSQL.

Dilihat dari cara penyimpanan data saja basis data NOSQL tersebar dari cara penyimpanan :

    Key-value based (disimpan dalam bentuk kunci-isi berpasangan)

Kunci-nilai/Key-value (KV) toko menggunakan array asosiatif (juga dikenal sebagai peta atau kamus) sebagai model data fundamental mereka. Dalam model ini, data direpresentasikan sebagai kumpulan pasangan kunci-nilai, sehingga setiap tombol mungkin muncul paling banyak sekali dalam koleksi. Model kunci-nilai adalah salah satu model data non-sepele sederhana, dan model data yang lebih kaya sering diimplementasikan di atas itu. Model kunci-nilai dapat diperluas untuk model memerintahkan yang mempertahankan kunci agar leksikografis. Ekstensi ini sangat kuat, dalam hal ini secara efisien dapat memproses rentang kunci. Toko kunci-nilai dapat menggunakan model konsistensi mulai dari konsistensi akhirnya ke serializability. Beberapa dukungan memesan kunci. Beberapa mempertahankan data dalam memori (RAM), sementara yang lain menggunakan solid-state drive atau disk

    Document based

Dokumen merangkum dan melakukan data encode (atau informasi) dalam beberapa format standar atau encoding. Pengkodean digunakan termasuk XML, YAML, dan JSON serta bentuk biner seperti BSON. Dokumen dibahas dalam database melalui kunci unik yang mewakili dokumen itu. Salah satu karakteristik mendefinisikan lain dari database berorientasi dokumen adalah bahwa di samping kunci pencarian yang dilakukan oleh sebuah toko kunci-nilai, database menawarkan API atau query bahasa yang mengambil dokumen berdasarkan isinya     Implementasi yang berbeda menawarkan cara yang berbeda mengatur dan / atau pengelompokan dokumen. Dibandingkan dengan database relasional, misalnya, koleksi dapat dianggap analog dengan tabel dan dokumen analog dengan catatan. Tetapi mereka berbeda: setiap record dalam sebuah tabel memiliki urutan yang sama bidang, sementara dokumen dalam koleksi mungkin memiliki bidang yang sama sekali berbeda.

    Column based (disimpan dalam kolom-kolom)
    Graph based

Jenis database dirancang untuk data yang hubungan baik diwakili sebagai grafik (unsur saling berhubungan dengan jumlah yang belum ditentukan hubungan antara mereka). Jenis data bisa hubungan sosial, jaringan transportasi umum, peta jalan atau topologi jaringan.

Bersumber dari halaman resmi NOSQL (http://nosql-database.org), NOSQL mengelompokkan database NOSQL ke dalam beberapa kategori, yaitu


    Wide Column Store / Column Families

Contoh : Accumulo, Casssandra, Clouddata, ConcourseDB, Hadoop /Hbase, Hypertable, kdb+, dll.

    Document Store

Contoh : CouchDB, Couchbase, Clusterpoint, Elasticserach, Mark-Logic, MongoDB, OrientDB, dll.

    Key Value / Tuple Store

Contoh : Aerospike, BangDB, BerkeleyDB, Chordless, DynamoDB, GenieDB, LevelDB, Riak, Redis, Oracle NOSQL Database, Tarantool, dll.

    Graph Databases

Contoh : ArangoDB, OrientDB, Infinite Graph, Neo4J, OpenLink Virtuoso, Stardog, WhiteDB, dll.

    Multimodel Databases

Contoh : Alchemy DB, ArangoDB, CortexDB, Datomic,FoundationDB, OrientDB, WonderDB, dll.

    Object Databases

Contoh : db4o, GemStone/S, HSS DB, Magma, Objectivity, siaqodb, Star-counter, VelocityDB, Versant, ZODB.

    Grid & Cloud Database

Contoh : Crate Data, GigaSpaces, GemFire, Hazelcast, Coherence, Queplix, dll.

    Solutions XML Databases

Contoh : BaseX, EMC Documentum xDB, eXist, Qizx,Sedna,  dll.

    Multidimensional Databases

Contoh : Globals, GT.M, Inter-systems Cache, MiniM DB, SciDB, dll.

    Multidivalue Databases

Contoh : Model204 DB, OpenInsight, OpenQM, Reality, U2, dll.

    Event Sourcing

Contoh : Eventstore.

    Time Series Databases

Contoh : Axibase.

 

Diantara banyak database NOSQL yang ada berdasarkan riset salah satu website pada Juli 2015 menunjukkan database NOSQL paling populer saat ini secara berurutan adalah MongoDB, Apache Cassandra, Redis, Solr, ElasticSearch, HBase, Splunk, memcached, dan Neo4j.

Terdapat tiga hal besar yang mempengaruhi perkembangan ini yaitu jumlah user yang banyak, jumlah data yang besar dan cloud computing. Dan dengan 3 hal besar diatas juga menjadikan sistem database harus mampu bergerak secara :

Data harus bisa bergerak secara flexible,

Harus mampu bergerak secara cepat dengan data dan user yang besar;

Dan yang terakhir peningkatan performa untuk dapat memuaskan user yang menginginkan pengolahan data yang cepat.

 
Tinjauan Paper Terkait NOSQL

Paper penelitian internasional pertama yang akan ditinjau berjudul “Type of NOSQL Databases and its Comparison with Relational Databases” karya Ameya Nayak, Anil Poriya, dan Dikshay Poojary dari Thakur College of Engineering and Technology University of Mumbai, India. Jurnal ini diterima pada bulan Maret 2013 dan kemudian diterbitkan dalam International Journal of Applied Information Systems (IJAIS) dengan ISSN : 2249-0868 Jurnal dapat diakses melalui link http://citeseerx.ist.psu.edu/viewdoc/download?doi=10.1.1.402.3372&rep=rep1&type=pdf.

Penulis ingin mendiskusikan model data dalam NOSQL, tipe penyimpanan data dalam NOSQL, karateristik dan fitur dari setiap penyimpanan data, bahasa query yang digunakan dalam NOSQL, kelebihan dan kekurangan NOSQL dan kemungkinan NOSQL di masa depan.

Penulis mengategorikan NOSQL dalam 5 tipe, yaitu key-value store databases, column-oriented databases, document store databases, graph databases, dan object oriented databases. Disertai pula dalam paper ini beberapa contoh dari masing-masing tipe dan menjelaskan karateristik umum dan khusus dari produk-produk NoSQL menurut tipe-tipe yang ada.

Keistimewaan NOSQL menurut penulis-penulis adalah tidak menggunakan SQL sama sekali yang sering digunakan oleh relational database lainnya sebagai bahasa query. Selain itu, NOSQL tidak memiliki bahasa standar query. Alasannya adalah kebanyakan database NOSQL memiliki bahasa query sendiri, seperti Cassandra yang memiliki CQL (Cassandra Query Language). Begitu juga MongoDB dengan Mongo Query Language.

Akhirnya sulit bagi user untuk mengganti dari database NOSQL satu ke database NOSQL lainnya. Karenanya untuk menutupi kekurangan NOSQL tersebut, perlu ada bahasa query umum seperti SQL yang dapat digunakan di seluruh database NOSQL. Itulah yang menjadi salah satu keistimewaan sekaligus kekurangan NOSQL.

Penulis menyimpulkan keuntungan dan kerugian menggunakan NOSQL. Keuntungannya adalah menyediakan jangkauan luas dari model data untuk memilih kondisi, skala mudah diubah (scalable), tidak membutuhkan administrator (pengelola) database, lebih cepat, efisien, dan fleksibel, tersusun pada langkah yang tinggi, serta mampu menangani kesalahan pada hardware. Sementara kerugian menggunakan NOSQL adalah struktur belum matang, tidak memiliki bahasa query standar, beberapa database tidak menjamin sifat ACID (Atomicity, Consistensy, Isolation, dan Durability).

Meskipun NOSQL kelambatan.

Although NOSQL has evolved at a very high pace, it still lags behind relational database in terms of number of users. The main reason behind this is that the users are more familiar with SQL while NOSQL databases lack a standard query language. If a standard query language for NOSQL is introduced, it will surely be a game changer. There are a few DBaaS providers over the cloud like Xeround which works on the hybrid database model, that is, they have the familiar SQL in the frontend and NOSQL in the backend. Databases ini mungkin tidak lebih cepat dari database NOSQL  murni tapi mereka masih menyediakan fitur dari relasional dan NOSQL kepada user. Jadi banyak kerugian dari NOSQL dan relational database yang dapat ditutupi. Dengan beberapa kemajuan di arsitektur anakan dari RDBMS ini, prospek masa depan untuk databases NOSQL sebagai database untuk pelayanan (DBaaS) akan lebih baik.

+ + +

Paper penelitian internasional kedua yang akan ditinjau berjudul “Choosing the right NOSQL database for the job: a quality attribute evaluation” karya João Ricardo Lourenço, et al. Jurnal ini diterbitkan dalam Journal of Big Data yang dipubllikasikan 14 Agustus 2015 dengan ISSN 2196-1115. Jurnal ini dapat diakses melalui link http://download.springer.com/static/pdf/999/art%253A10.1186%252Fs40537-015-0025-0.pdf.

Paper yang ditulis bersama empat rekannya, yaitu Bruno Cabral, Paulo Carreiro, Marco Vieira dan Jorge Bernardino memfokuskan tulisannya pada database NOSQL dan melakukan pembandingan dengan berpedoman pada metode-metode dalam software quality assurance.

Pada bagian latar belakang dan ulasan literatur penulis-penulis.

Analisa
Bagan 1 : Bagan yang menunjukkkan teorema CAP, singkatan dari Consistency, Avaibility, dan Partition-Tolerance

Perbandingan setiap produk terapan NOSQL juga ditampilkan. Berikut tabelnya.
  	Aerospike 	Cassandra 	Couchbase 	CouchDB 	Hbase 	MongoDB 	Voldemort
Category 	Key-value 	Column-store 	Document-store 	Document-store 	Column-store 	Document-store 	Key-value
CAP 	AP 	AP/CP 	CP 	AP 	CP 	CP 	AP
Consistency 	Dapat dikonfigurasi (dengan pilihan) 	Dapat dikonfigurasi (dengan pilihan) 	Eventual / Ter-jadi di akhir. 	Eventual / Ter-jadi di akhir. 	Configurable (konsistensi kuat dan eventual) 	Dapat dikonfigurasi (dengan pilihan) 	Read-Repair (client kendalikan bentrok)
Durability 	Diberitahu tertulis ke node-node hasil replikasi 	Dapat dikonfigurasi (dengan beberapa opsi/pilihan) 	Dapat dikonfigurasi (dengan beberapa opsi/pilihan) 	Dapat dikonfigurasi (diberitahu tertulis paling sedikit pada satu disk) 	Dapat dikonfigurasi (dengan beberapa opsi/pilihan) 	Dapat dikonfigurasi (dengan beberapa opsi/pilihan) 	Diberitahu tertulis ke node-node yang diinginkan
Querying 	Internal API 	CQL 	Internal API (MapReduce) 	Internal API (MapReduce) 	Internal API 	Internal API, MapReduce, complex query 	Internal API (get, put, delete)
Concurren-cy Control 	Read-commited isolation level 	MVCC 	MVCC 	MVCC 	Optimistic locking with MVCC 	Master-slave with multi-granularity locking 	Optimistic locking with MVCC
Partitioning Scheme 	Proprietary (basis Paxos) 	Consistent Hashing 	Consistent Hashing 	Consistent Hashing (pihak ketiga) 	Range Based 	Consistent Hashing 	Consistent Hashing
Native Partitioning 	Ya 	Ya 	Ya 	Tidak 	Ya 	Ya 	Ya

 

D

 
Ulasan Produk Terapan NOSQL

Salah satu produk terapan NOSQL yang terkenal adalah MongoDB. MongoDB adalah database open source berbasis dokumen (Document-Oriented Database) yang awalnya dibuat dengan bahasa C++. MongoDB sendiri sudah dikembangkan oleh 10gen sejak Oktober 2007, namun baru dipublikasikan pada Februari 2009. Selain karena performanya  4 kali lebih cepat dibandingkan MySQL serta mudah diaplikasikan, karena telah tergabung juga sebagai modul PHP.

Dalam konsep MongoDB tidak ada yang namanya tabel, kolom ataupun baris yang ada hanyalah collection (ibaratnya tabel), document (ibaratnya record). Data modelnya sendiri disebut BSON dengan struktur mirip dengan JSON. Strukturnya cukup mudah dibaca, contohnya seperti ini.

{

“nama” : “Dede Gunawan”,

“kontak” : {

“alamat” : “Kp. Cioray, Ds. Jatiwaras”,

“kota” : “Tasikmalaya”,

“kodepos” : “46191”,

“telp” : “0812xxxxxx”,

}

}

Dengan konsep key-value yang ada pada MongoDB, setiap document otomatis memiliki index id yang unik. Hal ini membantu mempercepat proses pencarian data secara global.

MongoDB hadir dengan beberapa kelebihan yaitu :

    Performa yang ditawarkan MongoDB lebih cepat dibandingkan MySQL ini disebabkan oleh memcached dan format dokumennya yang berbentuk seperti JSON
    Replikasi, adalah fitur yang sangat bermanfaat untuk backup data secara realtime. MongoDB sangat cocok digunakan untuk portal berita ataupun blog, namun belum cocok untuk digunakan pada sistem informasi yang berkaitan dengan keuangan karena MongoDB tidak mendukung transaction SQL.
    Auto-sharding, merupakan fitur untuk memecah database yang besar menjadi beberapa bagian demi optimalisasi performa database.
    MongoDB juga sudah mendukung C, C++, C#, Erlang, Haskell, Java, JavaScript, .NET(C# F#, PowerShell), Lips, Perl, PHP, Python, Ruby dan Scala
    Cross-platform, sehingga dapat digunakan di Windows, Linux, OS X dan Solaris
    Proses CRUD (Create, Read, Update, Delete) terasa sangat ringan
    Map/Reduce, akan sangat membantu ketika kita melakukan operasi agregasi. Outputnya pun akan menjadi collection. Kalau di MySQL biasanya kita menggunakan query GROUP BY
    GridFS, spesifikasi yang digunakan untuk menyimpan data yang sangat besar.

 

Kelebihan dan kekurangan MongoDB

Fitur-fitur MongoDB

Dikutip dari situs resminya, CouchDB adalah database yang benar-benar merangkul web. Menyimpan data Anda dengan dokumen JSON. Mengakses dokumen Anda dengan browser web Anda, melalui HTTP. Query, menggabungkan, dan mengubah dokumen Anda dengan JavaScript. CouchDB bekerja dengan baik dengan web modern dan aplikasi mobile. Anda bahkan dapat melayani aplikasi web langsung dari CouchDB. Dan Anda dapat mendistribusikan data Anda, atau aplikasi Anda, efisien menggunakan replikasi tambahan CouchDB itu. CouchDB mendukung master-master setup dengan deteksi konflik otomatis. Written in: C++

    Main point: Retains some friendly properties of SQL. (Query, index)
    License: AGPL (Drivers: Apache)
    Protocol: Custom, binary (BSON)
    Master/slave replication (auto failover with replica sets)
    Sharding built-in
    Queries are javascript expressions
    Run arbitrary javascript functions server-side
    Better update-in-place than CouchDB
    Uses memory mapped files for data storage
    Performance over features
    Journaling (with –journal) is best turned on
    On 32bit systems, limited to ~2.5Gb
    Text search integrated
    GridFS to store big data + metadata (not actually an FS)
    Has geospatial indexing
    Data center aware

 

Best used: If you need dynamic queries. If you prefer to define indexes, not map/reduce functions. If you need good performance on a big DB. If you wanted CouchDB, but your data changes too much, filling up disks.

 

For example: For most things that you would do with MySQL or PostgreSQL, but having predefined columns really holds you back.

CouchDB hadir dengan serangkaian fitur, seperti on-the-fly transformasi dokumen dan real-time pemberitahuan perubahan, yang membuat web app pengembangan mudah. Bahkan datang dengan mudah untuk menggunakan konsol web administrasi. Anda menebak itu, disajikan langsung dari CouchDB! Kami peduli banyak tentang skala didistribusikan. CouchDB sangat tersedia dan partisi toleran, tetapi juga akhirnya konsisten. Dan kita peduli banyak tentang data Anda. CouchDB memiliki mesin fault-tolerant penyimpanan yang menempatkan keamanan data Anda terlebih dahulu.

    Written in: Erlang
    feaMain point: DB consistency, ease of use
    License: Apache
    Protocol: HTTP/REST
    Bi-directional (!) replication,
    continuous or ad-hoc,
    with conflict detection,
    thus, master-master replication. (!)
    MVCC – write operations do not block reads
    Previous versions of documents are available
    Crash-only (reliable) design
    Needs compacting from time to time
    Views: embedded map/reduce
    Formatting views: lists & shows
    Server-side document validation possible
    Authentication possible
    Real-time updates via ‘_changes’ (!)
    Attachment handling
    thus, CouchApps (standalone js apps)

 

Best used: For accumulating, occasionally changing data, on which pre-defined queries are to be run. Places where versioning is important.

 

For example: CRM, CMS systems. Master-master replication is an especially interesting feature, allowing easy multi-site deployments.

Untuk mengunduh database CouchDB dapat mengunjungi alamat situs https://couchdb.apache.org.

Cassandra merupakan projek top level dari Apache yang mampu melahirkan Facebook dan dibangun dalam Google Big Table dan Amazon Dynamo. Casssandra adalah database.

Cassandra ditulis dalam bahasa Java dengan lisensi dari Apache. Cassandra menggunakan protokol CQL (versi terbaru adalah CQL3 ‒ red) dan Thrift. CQL memiliki beberapa batasan tertentu seperti tidak ada sintaks JOIN, dan beberapa kumpulan. CQL3 is now the official interface. Don’t look at Thrift, unless you’re working on a legacy app. This way, you can live without understanding ColumnFamilies, SuperColumns, etc.

    Querying by key, or key range (secondary indices are also available)
    Tunable trade-offs for distribution and replication (N, R, W)
    Data can have expiration (set on INSERT).
    Writes can be much faster than reads (when reads are disk-bound)
    Map/reduce possible with Apache Hadoop
    All nodes are similar, as opposed to Hadoop/HBase
    Very good and reliable cross-datacenter replication
    Distributed counter datatype.
    You can write triggers in Java.


### Basis Data NewSQL



## Paket Standar Go untuk Akses Basis Data




## Driver Go untuk Akses Basis Data




## Go dan SQL: PostgreSQL



## Go dan NoSQL: RethinkDB




