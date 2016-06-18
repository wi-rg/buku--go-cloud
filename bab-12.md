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

Artikel utama: NoSQL dan NewSQL. Generasi berikutnya dari database pasca-relasional di tahun 2000-an dikenal sebagai database NoSQL, termasuk kunci-nilai yang cepat toko dan database berorientasi dokumen. Database XML adalah jenis database berorientasi dokumen terstruktur yang memungkinkan query berdasarkan atribut dokumen XML. Database XML sebagian besar digunakan dalam manajemen basis data perusahaan, di mana XML digunakan sebagai standar interoperabilitas data yang mesin-ke-mesin. Database XML sebagian besar sistem perangkat lunak komersial, yang mencakup Clusterpoint, MarkLogic dan Oracle XML DB. Database NoSQL sering sangat cepat, tidak memerlukan skema tabel tetap, menghindari bergabung dengan operasi dengan menyimpan data denormalized, dan dirancang untuk skala horizontal. Sistem NoSQL paling populer termasuk MongoDB, Couchbase, Riak, memcached, Redis, CouchDB, Hazelcast, Apache Cassandra dan HBase, yang semua produk perangkat lunak open-source.

Dalam beberapa tahun terakhir ada permintaan yang tinggi untuk database terdistribusi secara besar-besaran dengan toleransi partisi yang tinggi tetapi menurut teorema CAP adalah mustahil untuk sistem terdistribusi untuk secara bersamaan memberikan konsistensi, ketersediaan dan toleransi partisi jaminan. Sebuah sistem terdistribusi dapat memenuhi dua dari jaminan tersebut pada saat yang sama, tetapi tidak semua tiga. Untuk itu banyak database NoSQL menggunakan apa yang disebut konsistensi akhirnya untuk menyediakan ketersediaan dan toleransi partisi jaminan dengan penurunan tingkat konsistensi data.

NewSQL adalah kelas database relasional modern yang bertujuan untuk memberikan performa yang sama sistem NoSQL untuk proses transaksi online (baca-tulis) beban kerja saat masih menggunakan SQL dan mempertahankan jaminan ACID sistem database tradisional. Database tersebut termasuk Scalebase, Clustrix, EnterpriseDB, MemSQL, NuoDB dan VoltDB.

## Paket Standar Go untuk Akses Basis Data
Instalasi Package Control

Package control merupakan aplikasi 3rd party untuk Sublime, digunakan untuk mempermudah instalasi plugin. Default-nya Sublime tidak menyediakan aplikas ini, kita perlu menginstalnya sendiri. Silakan ikuti petunjuk berikut untuk cara instalasinya.
Buka situs https://packagecontrol.io/installation, copy script yang ada di tab Sublime Text 3 (tab bagian kiri).
Copy script instalasi plugin

Selanjutnya, jalankan aplikasi Sublime, klik menu View > Show Console, lalu paste script yang sudah di-copy tadi, ke inputan kecil di bagian bawah editor. Lalu tekan Enter.
Show console, paste script instalasi package control

Tunggu hingga proses instalasi selesai. Perhatikan karakter sama dengan (=) di bagian kiri bawah editor yang bergerak-gerak. Jika karakter tersebut menghilang, menandakan bahwa proses instalasi sudah selesai.
Setelah selesai, tutup aplikasi, lalu buka kembali. Package control sudah berhasil di-install.

	Paket sql menyediakan antarmuka generik sekitar SQL (atau seperti SQL) database.
sopir	Sopir paket mendefinisikan interface yang harus dilaksanakan oleh database driver seperti yang digunakan oleh paket sql.



## Driver Go untuk Akses Basis Data

SQL database drivers
The database/sql and database/sql/driver packages are designed for using databases from Go and implementing database drivers, respectively.

See the design goals doc:

http://golang.org/src/pkg/database/sql/doc.txt
Drivers
Drivers for Go's sql package include:

Apache Phoenix/Avatica: https://github.com/Boostport/avatica
Couchbase N1QL: https://github.com/couchbase/go_n1ql
DB2: https://bitbucket.org/phiggins/db2cli
Firebird SQL: https://github.com/nakagami/firebirdsql
MS ADODB: https://github.com/mattn/go-adodb
MS SQL Server (pure go): https://github.com/denisenkom/go-mssqldb
MS SQL Server (uses cgo): https://github.com/minus5/gofreetds
MySQL: https://github.com/ziutek/mymysql [*]
MySQL: https://github.com/go-sql-driver/mysql/ [*]
ODBC: https://bitbucket.org/miquella/mgodbc
ODBC: https://github.com/alexbrainman/odbc
Oracle: https://github.com/mattn/go-oci8
Oracle: https://github.com/rana/ora
QL: http://godoc.org/github.com/cznic/ql/driver
Postgres (pure Go): https://github.com/lib/pq [*]
Postgres (uses cgo): https://github.com/jbarham/gopgsqldriver
Postgres (pure Go): https://github.com/jackc/pgx
SAP HANA (pure go): https://github.com/SAP/go-hdb
SQLite: https://github.com/mattn/go-sqlite3 [*]
SQLite: https://github.com/mxk/go-sqlite
Sybase SQL Anywhere: https://github.com/a-palchikov/sqlago
Vitess: https://godoc.org/github.com/youtube/vitess/go/vt/vitessdriver
YQL (Yahoo! Query Language): https://github.com/mattn/go-yql
Drivers marked with a [*] are both included in and pass the compatibility test suite at https://github.com/bradfitz/go-sql-test


## Go dan SQL: PostgreSQL

Memasang

go get github.com/lib/pq 
 Docs

Untuk dokumentasi rinci dan contoh penggunaan dasar, silakan lihat dokumentasi paket di http://godoc.org/github.com/lib/pq .

 tes

go test digunakan untuk pengujian. Sebuah server PostgreSQL berjalan diperlukan, dengan kemampuan untuk log in. The database default untuk terhubung untuk menguji dengan adalah "pqgotest," tapi dapat diganti menggunakan variabel lingkungan.

Contoh:

 PGHOST=/run/postgresql go test github.com/lib/pq 
Opsional, suite benchmark dapat dijalankan sebagai bagian dari tes:

 PGHOST=/run/postgresql go test -bench . 
 fitur

SSL
Menangani koneksi buruk bagi database/sql
Scan time.Time benar (yaitu timestamp[tz] , time[tz] , date )
Memindai gumpalan biner dengan benar (yaitu bytea )
Paket untuk hstore dukungan
COPY FROM dukungan
pq.ParseURL untuk mengkonversi url untuk koneksi string untuk sql.Open.
Banyak libpq variabel lingkungan yang kompatibel
dukungan socket Unix
Pemberitahuan: LISTEN / NOTIFY
dukungan pgpass

## Go dan NoSQL: RethinkDB

GoRethink - RethinkDB driver untuk Go
tag GitHub GoDoc build status

Pergi driver RethinkDB

GoRethink Logo

Versi saat ini: v2.0.4 (RethinkDB v2.3)

Harap dicatat bahwa versi ini pengemudi hanya mendukung versi RethinkDB menggunakan protokol v0.4 (setiap versi driver lebih tua dari RethinkDB 2.0 tidak akan bekerja).

Jika Anda memerlukan bantuan Anda dapat menemukan saya di slack RethinkDB dalam saluran #gorethink.

 Instalasi

go get gopkg.in/dancannon/gorethink.v2 
(Atau v1)

  pergi mendapatkan gopkg. di /dancannon/gorethink.v1 
 Koneksi

 dasar Connection

Pengaturan koneksi dasar dengan RethinkDB sederhana:

  impor (
     r "github.com/dancannon/gorethink"
     "Log"
 )

 sesi var * Sesi r.

 sesi, err:.. = r Connect (r ConnectOpts {
     Alamat: "localhost: 28015",
 })
 jika berbuat salah! = nil {
     log. Fatalln (err. Kesalahan ())
 }
Lihat dokumentasi untuk daftar argumen didukung untuk Hubungkan ().

 koneksi Renang

Sopir menggunakan kolam koneksi setiap saat, secara default menciptakan dan membebaskan koneksi otomatis. Ini aman untuk digunakan bersamaan oleh beberapa goroutines.

Untuk mengkonfigurasi kolam koneksi MaxIdle , MaxOpen dan Timeout dapat ditentukan saat sambungan. Jika Anda ingin mengubah nilai MaxIdle atau MaxOpen selama runtime maka fungsi SetMaxIdleConns dan SetMaxOpenConns dapat digunakan.

  sesi var * Sesi r.

 sesi, err:.. = r Connect (r ConnectOpts {
     Alamat: "localhost: 28015",
     Database: "test",
     MaxIdle: 10,
     MaxOpen: 10,
 })
 jika berbuat salah! = nil {
     log. Fatalln (err. Kesalahan ())
 }

 sesi. SetMaxOpenConns (5) 
 Terhubung ke cluster

Untuk menyambung ke cluster RethinkDB yang memiliki beberapa node Anda dapat menggunakan sintaks berikut. Saat menghubungkan ke cluster dengan beberapa node permintaan akan didistribusikan antara node tersebut.

  sesi var * Sesi r.

 sesi, err:.. = r Connect (r ConnectOpts {
     Alamat: [] String { "localhost: 28015", "localhost: 28016"},
     Database: "test",
     AuthKey: "14daak1cad13dj",
     DiscoverHosts: benar,
 })
 jika berbuat salah! = nil {
     log. Fatalln (err. Kesalahan ())
 } 
Ketika DiscoverHosts benar setiap node ditambahkan ke cluster setelah koneksi awal maka node baru akan ditambahkan ke kolam node yang tersedia digunakan oleh GoRethink. Sayangnya alamat kanonik setiap server di cluster HARUS ditetapkan sebagai dinyatakan klien akan mencoba untuk menyambung ke node basis data lokal. Untuk informasi lebih lanjut tentang cara mengatur server RethinkDB alamat kanonik mengatur halaman ini http://www.rethinkdb.com/docs/config-file/ .

 Otentikasi pengguna

Login dengan username dan password pertama Anda harus membuat pengguna, ini bisa dilakukan dengan menulis kepada users tabel sistem dan kemudian memberikan bahwa akses pengguna ke setiap meja atau database mereka memerlukan akses ke. pertanyaan ini juga dapat dijalankan di RethinkDB admin konsol.

  err:. = r DB ( "rethinkdb") Tabel ( "pengguna") Sisipkan (peta [string] String {..
     "Id": "john",
     "Password": "p455w0rd",
 }). Exec (sesi)
 ...
 err = r. DB ( "blog"). Tabel ( "posting"). Hibah ( "john", peta [string] bool {
     "Membaca": true,
     "Menulis": true,
 }). Exec (sesi)
 ... 
Akhirnya username dan password harus diteruskan ke Connect saat membuat sesi Anda, misalnya:

  sesi, err:.. = r Connect (r ConnectOpts {
     Alamat: "localhost: 28015",
     Database: "blog",
     Username: "john",
     Password: "p455w0rd",
 }) 
Harap dicatat bahwa DiscoverHosts tidak akan bekerja dengan otentikasi pengguna saat ini karena fakta bahwa RethinkDB membatasi akses ke tabel sistem yang diperlukan.

 Fungsi Query

Perpustakaan ini didasarkan pada driver resmi sehingga kode pada API halaman harus memerlukan sedikit perubahan untuk bekerja.

Untuk melihat dokumentasi lengkap untuk fungsi query memeriksa referensi API atau GoDoc

Slice Expr Contoh

  r. Expr ([] antarmuka {} {1, 2, 3, 4, 5}) Run (sesi). 
Peta Expr Contoh

  r Expr (peta [string] antarmuka {} { "a": 1, "b": 2, "c": 3}).. Run (sesi) 
Dapatkan Contoh

  r. DB ( "database"). Tabel ( "meja"). Dapatkan ( "GUID"). Run (sesi) 
Peta Contoh (Func)

  r. Expr ([] antarmuka {} {1, 2, 3, 4, 5}). Peta (func (baris Term) antarmuka {} {
     kembali baris. Tambahkan (1)
 }). Run (sesi) 
Peta Contoh (implisit)

  r. Expr ([] antarmuka {} {1, 2, 3, 4, 5}). Peta (r. Row. Tambahkan (1)). Run (sesi) 
Antara (Opsional Args) Contoh

  r. DB ( "database"). Tabel ( "meja"). Antara (1, 10, r. BetweenOpts {
     Indeks: "num",
     RightBound: "tertutup",
 }). Run (sesi) 
 Argumen opsional

Seperti yang ditunjukkan di atas dalam contoh Antara argumen opsional dilewatkan ke fungsi sebagai struct. Setiap fungsi yang memiliki argumen opsional sebagai struct terkait. struct ini diberi nama dengan FunctionNameOpts Format, misalnya BetweenOpts adalah struct terkait untuk Antara.

 hasil

jenis hasil yang berbeda dikembalikan tergantung pada apa fungsi yang digunakan untuk mengeksekusi query.

Run mengembalikan kursor yang dapat digunakan untuk melihat semua baris yang dikembalikan.
RunWrite mengembalikan WriteResponse dan harus digunakan untuk query seperti Insert, Update, dll ...
Exec mengirimkan permintaan ke server dan menutup koneksi setelah membaca respon dari database. Jika Anda tidak ingin menunggu respon maka Anda dapat mengatur NoReply bendera.
Contoh:

  res, err:... = r DB ( "database") Tabel ( "tablename") Dapatkan (kunci) Run (sesi).
 jika berbuat salah! = nil {
     // error
 }
 menunda res. Tutup () // Selalu pastikan Anda menutup kursor untuk memastikan koneksi tidak bocor 
Kursor memiliki sejumlah metode yang tersedia untuk mengakses hasil query

Next mengambil dokumen berikutnya dari hasil set, menghalangi jika perlu.
All mengambil semua dokumen dari hasil set ke slice yang disediakan.
One mengambil dokumen pertama dari hasil set.
contoh:

  baris var antarmuka {}
 untuk res. Berikutnya (& row) {
     // Lakukan sesuatu dengan baris
 }
 jika res. Err ()! = nil {
     // error
 } 
  var baris [] antarmuka {}
 err:. = res Semua (& baris)
 jika berbuat salah! = nil {
     // error
 } 
  baris var antarmuka {}
 err: = res One (& baris).
 jika err == r. ErrEmptyResult {
     // Baris tidak ditemukan
 }
 jika berbuat salah! = nil {
     // error
 } 
 Encoding / Decoding

Ketika melewati struct ke Expr (Dan fungsi yang menggunakan Expr seperti Insert, Update) struct dikodekan ke dalam peta sebelum dikirim ke server. Setiap bidang diekspor ditambahkan ke peta kecuali

tag bidang ini adalah "-", atau
lapangan kosong dan tag yang menentukan "omitempty" pilihan.
Setiap bidang nama default pada peta adalah nama field tetapi dapat ditentukan dalam nilai tag struct lapangan. The "gorethink" kunci dalam nilai tag struct lapangan adalah nama kunci, diikuti oleh koma opsional dan pilihan. contoh:

  // Lapangan diabaikan oleh paket ini.
 Bidang int `gorethink:" - "`
 // Bidang muncul sebagai kunci "MYNAME".
 Bidang int `gorethink:" MYNAME "`
 // Bidang muncul sebagai kunci "MYNAME" dan
 // Lapangan dihilangkan dari objek jika nilainya kosong,
 // Seperti dijelaskan di atas.
 Bidang int `gorethink:" MYNAME, omitempty "`
 // Bidang muncul sebagai kunci "Field" (default), tetapi
 // Lapangan dilewati jika kosong.
 // Catatan terkemuka koma.
 Bidang int `gorethink:", omitempty "` 
CATATAN: Sangat disarankan bahwa tag struct digunakan secara eksplisit menentukan pemetaan antara tipe Go Anda dan bagaimana data disimpan oleh RethinkDB. Hal ini sangat penting ketika menggunakan Id lapangan sebagai default RethinkDB akan menciptakan lapangan bernama id sebagai kunci utama (perhatikan bahwa bidang RethinkDB adalah huruf kecil tapi versi Go dimulai dengan huruf kapital).

Ketika pengkodean peta dengan tombol non-string nilai kunci secara otomatis dikonversi ke string mana mungkin, namun disarankan agar Anda menggunakan string mana mungkin (misalnya map[string]T ).

Jika Anda ingin menggunakan json tag untuk GoRethink maka Anda dapat menghubungi SetTags("gorethink", "json") ketika memulai program Anda, ini akan menyebabkan GoRethink untuk memeriksa json tag setelah memeriksa gorethink tag. Secara default fitur ini dinonaktifkan. Fungsi ini juga akan membiarkan Anda mendukung setiap tag lain, pengemudi akan memeriksa tag di urutan yang sama seperti parameter.

 Referensi

Kadang-kadang Anda mungkin ingin menggunakan struct Go yang referensi dokumen di meja lain, bukan menciptakan struct baru yang hanya digunakan saat menulis untuk RethinkDB Anda dapat membubuhi keterangan struct Anda dengan pilihan tag referensi. Ini akan memberitahu GoRethink bahwa ketika pengkodean data Anda harus "memetik" ID bidang dari dokumen bersarang dan menggunakannya.

Ini semua cukup rumit jadi mudah-mudahan contoh ini akan membantu. Pertama mari kita asumsikan Anda memiliki dua jenis Author dan Book dan Anda ingin memasukkan buku baru ke dalam database Anda namun Anda tidak ingin menyertakan seluruh struct penulis dalam tabel buku. Seperti yang Anda lihat Author lapangan di Book struct memiliki beberapa tag tambahan, pertama kita telah menambahkan reference tag opsi yang memberitahu GoRethink untuk memetik lapangan dari Author struct bukannya memasukkan dokumen penulis seluruh. Kami juga memiliki gorethink_ref tag yang memberitahu GoRethink untuk mencari id lapangan di Author dokumen, tanpa tag ini GoRethink bukannya akan mencari author_id lapangan.

  Jenis Penulis struct {
     ID tali `gorethink:" id, omitempty "`
     String name `gorethink:" Nama "`
 }

 ketik Book struct {
     ID tali `gorethink:" id, omitempty "`
     Judul tali `gorethink:" title "`
     Penulis Penulis `gorethink:" author_id, referensi "gorethink_ref:" id "`
 } 
Data yang dihasilkan di RethinkDB harus terlihat seperti ini:

  {
     "Author_id": "c2182a10-6b9d-4ea1-a70c-d6649bb5f8d7",
     "Id": "eeb006d6-7fec-46c8-9d29-45b83f07ca14",
     "Title": "The Hobbit"
 } 
Jika Anda ingin membaca kembali buku dengan penulis termasuk maka Anda bisa menjalankan GoRethink query berikut:

  r. Tabel ( "buku"). Dapatkan ( "1"). Merge (func (p r. Term) antarmuka {} {
     Peta pulang [string] antarmuka {} {
         "Author_id":.. (. P Lapangan ( "author_id")) r Tabel ( "penulis") Dapatkan,
     }
 }). Run (sesi) 
 logging

Secara default pengemudi log kesalahan ketika gagal untuk terhubung ke database. Jika Anda ingin lebih verbose logging kesalahan Anda dapat menghubungi r.SetVerbose(true) .

Atau jika Anda ingin mengubah perilaku penebangan Anda dapat memodifikasi logger disediakan oleh github.com/Sirupsen/logrus . Misalnya kode berikut sepenuhnya menonaktifkan logger:

  r. Log. Out = ioutil. Buang 
 benchmark

Semua orang ingin benchmark proyek mereka untuk menjadi cepat. Dan sementara kita tahu bahwa rethinkDb dan driver gorethink cukup cepat, tujuan utama kami adalah untuk benchmark kami untuk menjadi benar. Mereka dirancang untuk memberikan Anda, pengguna, gambaran yang akurat dari menulis per detik (w / s). Jika Anda datang dengan tes yang akurat yang memenuhi tujuan ini, kirimkan permintaan tarik silakan.

Berkat @jaredfolkins untuk kontribusi.

Mengetik	Nilai
Nama model	MacBook Pro
Model Identifier	MacBookPro11,3
processor Nama	Intel Core i7
processor Kecepatan	2,3 GHz
Jumlah Prosesor	1
Total Jumlah Cores	4
L2 Cache (per Core)	256 KB
L3 Cache	6 MB
Ingatan	16 GB
  BenchmarkBatch200RandomWrites 20 557227775 ns / op
 BenchmarkBatch200RandomWritesParallel10 30 354465417 ns / op
 BenchmarkBatch200SoftRandomWritesParallel10 100 761639276 ns / op
 BenchmarkRandomWrites 100 10456580 ns / op
 BenchmarkRandomWritesParallel10 1000 1614175 ns / op
 BenchmarkRandomSoftWrites 3000 589660 ns / op
 BenchmarkRandomSoftWritesParallel10 10000 247588 ns / op
 BenchmarkSequentialWrites 50 24408285 ns / op
 BenchmarkSequentialWritesParallel10 1000 1755373 ns / op
 BenchmarkSequentialSoftWrites 3000 631211 ns / op
 BenchmarkSequentialSoftWritesParallel10 10000 263481 ns / op 
 contoh

Banyak fungsi memiliki contoh dan dapat dilihat di godoc itu, alternatif melihat beberapa fitur yang lebih lengkap contoh di wiki .

 Bacaan lebih lanjut

GoRethink Goes 1.0
Go, RethinkDB & Changefeeds
Membangun bot IRC di Go dengan changefeeds RethinkDB
 Lisensi

Hak Cipta 2013 Daniel Cannon

Berlisensi di bawah Apache, Versi 2.0 ( "Lisensi"); Anda mungkin tidak menggunakan file ini kecuali sesuai dengan Lisensi. Anda dapat memperoleh salinan Lisensi di

http://www.apache.org/licenses/LICENSE-2.0

Kecuali diwajibkan oleh hukum yang berlaku atau disetujui secara tertulis, perangkat lunak yang didistribusikan di bawah Lisensi didistribusikan pada "APA ADANYA", TANPA JAMINAN ATAU KONDISI APAPUN, baik tersurat maupun tersirat. Lihat Lisensi untuk izin bahasa tertentu dan keterbatasan di bawah Lisensi.

###NoSQL MongoDB

Golang tidak menyediakan interface generic untuk NoSQL, hal ini menjadikan driver tiap brand NoSQL untuk Golang bisa berbeda satu dengan lainnya.
Dari sekian banyak teknologi NoSQL yang ada, yang terpilih untuk dibahas di buku ini adalah MongoDB. Dan pada bab ini kita akan belajar cara berkomunikasi dengan MongoDB menggunakan driver mgo.
Persiapan

Ada beberapa hal yang perlu disiapkan sebelum mulai masuk ke bagian coding.
Instal mgo menggunakan go get.
go get gopkg.in/mgo.v2
Download driver mgo

Pastikan sudah terinstal MongoDB di komputer anda, dan jangan lupa untuk menjalankan daemon-nya. Jika belum, download dan install terlebih dahulu.
Instal juga MongoDB GUI untuk mempermudah browsing data. Bisa menggunakan MongoChef, Robomongo, atau lainnya.
Insert Data

Cara insert data lewat mongo tidak terlalu sulit. Yang pertama perlu dilakukan adalah import package yang dibutuhkan, dan juga menyiapkan struct model.
package main

import "fmt"
import "gopkg.in/mgo.v2"
import "os"

type student struct {
    Name  string `bson:"name"`
    Grade int    `bson:"Grade"`
}
Tag bson pada property struct dalam konteks mgo, digunakan sebagai penentu nama field ketika data disimpan kedalam collection. Jika sebuah property tidak memiliki tag bson, secara default nama field adalah sama dengan nama property hanya saja lowercase. Untuk customize nama field, gunakan tag bson.
Pada contoh di atas, property Name ditentukan nama field nya sebagai name, dan Grade sebagai Grade.
Selanjutnya siapkan fungsi untuk membuat session baru.
func connect() *mgo.Session {
    var session, err = mgo.Dial("localhost")
    if err != nil {
        os.Exit(0)
    }
    return session
}
Fungsi mgo.Dial() digunakan untuk membuat session baru (bertipe *mgo.Session). Fungsi tersebut memiliki sebuah parameter yang harus diisi, yaitu connection string dari server mongo yang akan diakses.
Secara default jenis konsistensi session yang digunakan adalah mgo.Primary. Anda bisa mengubahnya lewat method SetMode() milik struct mgo.Session. Lebih jelasnya silakan merujuk https://godoc.org/gopkg.in/mgo.v2#Session.SetMode.
Terkahir buat fungsi insert yang didalamnya berisikan kode untuk insert data ke mongodb, lalu implementasikan di main.
func insert() {
    var session = connect()
    defer session.Close()
    var collection = session.DB("belajar_golang").C("student")

    var err = collection.Insert(&student{"Wick", 2}, &student{"Ethan", 2})
    if err != nil {
        fmt.Println(err.Error())
    }
}

func main() {
    insert()
}
Session di mgo juga harus di close ketika sudah tidak digunakan, seperti pada instance connection di bab SQL. Statement defer session.Close() akan mengakhirkan proses close session dalam fungsi insert().
Struct mgo.Session memiliki method DB() yang digunakan untuk memilih database yang digunakan, dan bisa langsung di chain dengan fungsi C() untuk memilih collection.
Setelah mendapatkan instance collection-nya, digunakan method Insert() untuk insert data ke database. Method ini memiliki parameter variadic pointer data yang ingin di-insert.
Jalankan program tersebut, lalu cek menggunakan mongo GUI untuk melihat apakah data sudah masuk.
Insert mongo

###Membaca Data

method Find() milik tipe collection mgo.Collection digunakan untuk melakukan pembacaan data. Query selectornya dituliskan menggunakan bson.M lalu disisipkan sebagai parameter fungsi Find().
Untuk menggunakan bson.M, package gopkg.in/mgo.v2/bson harus di-import terlebih dahulu.
import "gopkg.in/mgo.v2/bson"
Setelah itu buat fungsi find yang didalamnya terdapat proses baca data dari database.
func find() {
    var session = connect()
    defer session.Close()
    var collection = session.DB("belajar_golang").C("student")

    var result = student{}
    var selector = bson.M{"name": "Wick"}
    var err = collection.Find(selector).One(&result)
    if err != nil {
        fmt.Println(err.Error())
    }

    fmt.Println("Name  :", result.Name)
    fmt.Println("Grade :", result.Grade)
}

func main() {
    find()
}
Variabel result di-inisialisasi menggunakan struct student. Variabel tersebut nantinya digunakan untuk menampung hasil pencarian data.
Tipe bson.M sebenarnya adalah alias dari map[string]interface{}, digunakan dalam penulisan selector.
Selector tersebut kemudian dimasukan sebagai parameter method Find(), yang kemudian di chain langsung dengan method One() untuk mengambil 1 baris datanya. Kemudian pointer variabel result disisipkan sebagai parameter method tersebut.
Pencarian data

###Update Data

Method Update() milik struct mgo.Collection digunakan untuk update data. Ada 2 parameter yang harus diisi:
Parameter pertama adalah query selector data yang ingin di update
Parameter kedua adalah data perubahannya
func update() {
    var session = connect()
    defer session.Close()
    var collection = session.DB("belajar_golang").C("student")

    var selector = bson.M{"name": "Wick"}
    var changes = student{"John Wick", 2}
    var err = collection.Update(selector, changes)
    if err != nil {
        fmt.Println(err.Error())
    }
}

func main() {
    update()
}
Bisa dicek lewat Mongo GUI apakah data sudah berubah.
Update data

### Menghapus Data

Cara menghapus document pada collection cukup mudah, tinggal gunakan method Remove() dengan isi parameter adalah query selector document yang ingin dihapus.
func remove() {
    var session = connect()
    defer session.Close()
    var collection = session.DB("belajar_golang").C("student")

    var selector = bson.M{"name": "John Wick"}
    var err = collection.Remove(selector)
    if err != nil {
        fmt.Println(err.Error())
    }
}

func main() {
    remove()
}
2 data yang sebelumnya sudah di-insert kini tinggal satu saja.
Menghapus data



