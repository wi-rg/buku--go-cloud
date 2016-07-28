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
Pengenalan Basisdata NoSQL NoSQL adalah sebuah konsep mengenai penyimpanan data non-relasional. Berbeda dengan model basis data relasional yang selama ini digunakan, NoSQL menggunakan beberapa metode yang berbeda-beda. Metode ini bergantung dari jenis database yang digunakan. Karena NoSQL sendiri merupakan konsep database, dan pada implementasinya, banyak jenis-jenis dari NoSQL ini.

NoSQL sangat berguna pada data-data yang terus-menerus berkembang, dimana data tersebut sangat kompleks sehingga sebuah database relational tidak lagi bisa mengakomodir. Salah satu bentuknya adalah ketika suatu data saling berhubungan satu sama lain, maka akan muncul proses duplikasi data. Dimana data saling memanggil ke beberapa permintaan, tambahan data baru, perubahan data, dan lain-lain dengan key yang sama. Karena faktor hubungan antar data yang sama terjadi terus-menerus, mendorong faktor redudansi data, data menjadi berlipat-lipat(terjadi kesamaan data pada banyak tempat) dan pada akhirnya akan menyebabkan crash pada database berkonsep RDBMS (Relational database management system).

Contoh nyata penjelasan diatas bisa dilihat pada situs jejaring sosial, FACEBOOK.COM. Pernahkah terpikirkan 'sebesar apa server penyimpanan facebook dengan anggota ribuan orang tersebut?' 'Bagaimana kalau setengah dari member facebook adalah orang-orang yang narsis sehingga dibutuhkan storage yang sangat besar hanya untuk menyimpan foto seluruh member?' 'Berapa besar ukuran yang dibutuhkan facebook pada servernya?'.

Untuk menanggulangi hal tersebut-lah kemudian dikenal suatu cara penataan Database yang disebut Basisdata NoSQL.

MACAM-MACAM DATABASE NOSQL Dalam NoSQL, terdapat berbagai macam bentuk database dengan schema dan konsepnya sendiri. Sama halnya dengan database SQL yang seperti teman-teman kenal seperti MySQL, SQL Server, Oracle, PostgreSQL, dan lain-lain. Dalam NoSQL, terdapat MongoDB, CouchDB, Riak, dan lain-lain. Berikut ini adalah beberapa contoh dari database NoSQL yang sudah cukup banyak dikenal dan digunakan oleh kebanyakan orang, namun kami akan bahas beberapa saja, diantaranya yaitu:

MongoDB MongoDB adalah database open source yang sudah dikembangkan sejak Oktober 2007 dan dipublikasikan pada februari 2009. MongoDB merupakan salah satu database yang berbasis document (Document-Oriented Database). Dari segi performa MongoDB 4 kali lebih cepat dibandingkan MySQL. Dikarenakan konsep MongoDB berbasis document maka MongoDB tidak memiliki table, kolom maupun baris namun yang ada adalah collection (tabel) dan document (record). Data modelnya disebut dengan BSON yang strukturnya mirip dengan JSON.
Konsep key-value yang ada pada MongoDB, setiap document otomatis memiliki index id yang unik yang membantu mempercepat proses pencarian data secara global.

"Pada MongoDb kita tidak mengenal istilah RELASI karena kita bepikir dalam kerangka dokumen. Semua yang kita butuhkan diletakkan didalam sebuah dokumen."

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
untuk contoh Pembahasan, Cara Penggunaan dan Tambahan Topik dapat dilihat langsung pada: https://en.wikipedia.org/wiki/MongoDB

untuk contoh tambahan Penggunaan Dokumen MongoDB dapat dilihat langsung pada: https://docs.mongodb.com/ dan www.tutorialspoint.com/mongodb/

untuk Mendapatkan https://www.mongodb.com/

Cassandra Apache adalah sebuah aplikasi database berbasis Bigtabel’s Data. Dikutip dari situs resminya, Database Apache Cassandra adalah pilihan yang tepat ketika Anda membutuhkan skalabilitas dan ketersediaan tinggi tanpa mengorbankan kinerja. Skalabilitas linear dan terbukti toleransi kesalahan pada perangkat keras komoditas atau infrastruktur cloud membuatnya menjadi platform yang sempurna untuk misi-data penting. Dukungan Cassandra untuk mereplikasi di beberapa pusat data yang terbaik di kelasnya, memberikan latency rendah untuk pengguna Anda dan ketenangan pikiran mengetahui bahwa Anda dapat bertahan hidup padam daerah. ColumnFamily Cassandra data model menawarkan kenyamanan indeks kolom dengan kinerja log-terstruktur update, dukungan yang kuat untuk pandangan terwujud, dan kuat built-in caching. Cassandra Dikembangkan oleh APACHE. Aplikasi Inilah yang digunakan facebook untuk penyimpanan miliayaran data hingga saat ini. (http://cassandra.apache.org/)
Secara garis besar, Basis data NoSQL Cassandra dapat menangani data dalam jumlah yang besar (Big Data). Cassandra secara otomatis mereplikasi Data ke simpul (node) yang mendukung replikasi di beberapa pusat data yang terkait. Dengan arsitektur desentralisasi seperti ini risiko kegagalan penyimpanan data dapat meminimalkan secara default.

Beberapa Fitur Utama Cassandra:
    
    - Desentralisasi: Setiap node di cluster memiliki peran yang sama. Tidak ada single point of failure. Data didistribusikan lintas cluster (sehingga setiap node berisi data yang berbeda), namun tidak ada master karena setiap node dapat melayani permintaan apapun yang sama. Mendukung replikasi dan replikasi di multi data center: Strategi replikasi dapat dikonfigurasi. Cassandra dirancang sebagai sistem terdistribusi, untuk penyebaran sejumlah besar node di beberapa pusat data. Fitur utama dari arsitektur terdistribusi Cassandra adalah secara khusus dirancang untuk penyebaran di multi-pusat data, untuk redundansi, untuk failover dan pemulihan dari bencana.

    - Skalabilitas: Membaca dan menulis throughput yang baik meningkat secara linear ketika mesin baru ditambahkan, tanpa downtime atau dampak gangguan pada aplikasi.

    - Fault-tolerant: Data secara otomatis direplikasi ke beberapa node untuk toleransi kesalahan. Replikasi di beberapa pusat data didukung. Node gagal dapat diganti tanpa downtime.

    - Penyetelan secara konsisten: Menulis dan membaca menawarkan tingkat penyetelan secara konsisten, sepanjang jalan dari “menulis tidak pernah gagal” untuk “blok untuk semua replika agar dapat dibaca”, dengan tingkat kuorum di tengah.

    - Dukungan MapReduce: Cassandra memiliki kemampuan untuk di integrasi dengan Hadoop, berikut dukungan MapReduce. Juga mendukung Apache Pig dan Apache Hive.

    - Bahasa query: Memperkanalkan CQL (Cassandra Query Language) diperkenalkan, merupakan SQL-like alternatif terhadap antarmuka RPC tradisional. Bahasa driver tersedia untuk Java (JDBC), Python (DBAPI2) dan Node.js (Helenus). 



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
Paket standar go untuk akses basis data adalah 'sql'. Paket sql menyediakan antarmuka umum di basisdata sql. Paket ini harus digunakan bersamaan dengan driver database. Untuk menggunakan paket sql ini maka harus dideklarasikan syntax :
import "database/sql".



## Driver Go untuk Akses Basis Data
Ada banyak sekali driver go untuk akses basis data, seperti :
- Apache Phoenix/Avatica
- Couchbase N1QL
- MySQL
- DB2
- Firebird SQL
Pendeklarasian driver ini dengan menggunakan sintax : 
import "database/sql/driver"

### Variabel
a. var Bool boolType
Bool adalah ValueConverter yang mengubah nilai-nilai masukan untuk variabel bool.
Aturan konversi :
- Boolean dikembalikan tanpa mengalami perubahan
- Untuk tipe integer,
      1 = benar
      0 = salah,
      bilangan bulat lainnya adalah kesalahan
- Untuk string dan [] byte, aturan yang sama seperti strconv.ParseBool
- Semua jenis lain kesalahan

b. var DefaultParameterConverter defaultConverter
DefaultParameterConverter adalah implementasi default ValueConverter yang digunakan saat Stmt tidak mengimplementasikan ColumnConverter.
DefaultParameterConverter mengembalikan argumen langsung jika IsValue (arg). Jika tidak, argumen mengimplementasikan Valuer, metode Nilai yang digunakan untuk mengembalikan nilai a. Sebagai fallback, jenis yang mendasari diberikan argumen yang digunakan untuk mengubahnya menjadi Nilai: tipe integer yang mendasari dikonversi ke int64, float ke float64, dan string ke [] byte. Jika argumen adalah pointer nol, ConvertValue mengembalikan nilai nol. Jika argumen adalah pointer non-nol, maka dereferenced dan ConvertValue disebut rekursif. Jenis lain adalah kesalahan.

c. var ErrBadConn = errors.New("driver: bad connection")
ErrBadConn harus dikembalikan oleh driver untuk sinyal ke paket sql bahwa driver.Conn dalam keadaan buruk (seperti server yang sebelumnya telah menutup sambungan) dan paket sql harus mencoba lagi pada sambungan baru.
Untuk mencegah duplikasi operasi, ErrBadConn TIDAK harus dikembalikan jika ada kemungkinan bahwa database server mungkin telah melakukan operasi tersebut. Bahkan jika server mengirim kembali kesalahan, Anda tidak harus kembali ErrBadConn.

d. var ErrSkip = errors.New("driver: skip fast-path; continue as if unimplemented")
ErrSkip dapat dikembalikan dengan metode beberapa antarmuka opsional untuk menunjukkan pada runtime bahwa jalur cepat tidak tersedia dan paket sql harus meneruskan seolah antarmuka opsional tidak diterapkan. ErrSkip hanya didukung keberadaan secara eksplisit didokumentasikan.

e. var Int32 int32Type
Int32 adalah Nilai Converter yang mengubah nilai-nilai masukan untuk int64 dan menghargai batasan nilai int32.

f. var ResultNoRows noRows
ResultNoRows adalah Result standar driver untuk kembali ketika perintah DDL (seperti CREATE TABLE) berhasil. Kesalahan untuk  lastInsertId dan RowsAffected.

g. var String stringType
String adalah ValueConverter yang mengubah input ke string. Jika nilai sudah string atau [] byte, maka tidak mengalami perubahan. Jika nilai dari jenis lain, konversi ke string dilakukan dengan fmt.Sprintf ("% v", v).

### Fungsi IsScanValue
func IsScanValue(v interface{}) bool
IsScanValue merupakan jenis pemindaian nilai valid. Tidak seperti IsValue, IsScanValue tidak mengizinkan tipe string.

### Fungsi IsValue
func IsValue(v interface{}) bool
Is Value reports whether v adalah jenis Nilai parameter yang valid. Tidak seperti IsScanValue, IsValue memungkinkan tipe string.

### Tipe ColumnConverter
type ColumnConverter interface {
   // ColumnConverter mengembalikan ValueConverter untuk diberikan
         // Indeks kolom. Jika jenis kolom tertentu tidak diketahui
         // Atau tidak harus ditangani secara khusus, DefaultValueConverter
         // Dapat dikembalikan.
         ColumnConverter (idx int) ValueConverter
}
ColumnConverter mungkin opsional dilaksanakan oleh Stmt jika pernyataan sadar akan jenis kolom sendiri dan dapat mengkonversi dari jenis ke driver value.






## Go dan SQL: PostgreSQL

###Tentang PostgreSQL
PostgreSQL merupakan Sebuah Obyek-Relasional Data Base Management System (ORDBMS) yang dikembangkan oleh Berkeley Computer Science Department. System yang ditawarkan PostgreSQL diharapkan sanggup dan dapat mencukupi untuk kebutuhan proses aplikasi data masa depan. PostgreSQL juga menawarkan tambahan-tambahan yang cukup signifikan yaitu class, inheritance, type, dan function. Tambahan keistimewaan lain yang tidak dimiliki database management system yang lain berupa constraint, triggers, rule, dan transaction integrity, dengan adanya feature (keistimewaan) tersebut maka para pemakai dapat dengan mudah mengimplementasikan dan menyampaikan sistem ini. Sejak tahun 1996 PostgreSQL mengalami kemajuan yang sangat berarti, berbagai keistimewaan dari PostgreSQL sanggup membuat database ini melebihi database lain dari berbagai sudut pandang. Pada awal pembuatannya di University of California Berkeley (1977-1985) postgresl masih mempunyai banyak kekurangan bila dibandingkan dengan database yang lain, namun seiring dengan berjalannya waktu tepatnya pada tahun 1996 PostgresSQL berubah menjadi sebuah database yang menawarkan standar melebihi standar ANSI-SQL92 dan sanggup memenuhi permintaan dunia open source akan server database SQL. Standar ANSI-SQL92 merupakan standar yang ditetapkan untuk sebuah database berskala besar seperti Oracle, Interbase, DB2 dan yang lainnya. Kelebihan PostgreSQL Berbeda dengan database lain, PostgreSQL menyediakan begitu banyak dokumentasi yang disertakan pada berbagai distribusi Linux, sehingga para pembaca bisa dengan mudah mempelajari bahkan mengimplementasikannya. Tidak hanya itu berbagai dokumentasi yang bertebaran di Internet maupun mailing list yang semuanya dapat kita ambil dan pelajari. PostgreSQL memiliki keluwesan dan kinerja yang tinggi, artinya sesuai dengan niatan awal para pembuat PostgreSQL bahwa database yang mereka buat harus melebihi database lain dan ini terbukti pada arsitekturnya. Dengan arsitektur yang luwes maka sebuah user PostgreSQL mampu mendefenisikan sendiri SQL-nya, inilah yang membuat database PostgreSQL berbeda dengan sistem relasional standar. Di samping mendefenisikan sendiri SQL-nya, PostgreSQL juga memungkinkan setiap user untuk membuat sendiri object file yang dapat diterapkan untuk mendefenisikan tipe data, fungsi dan bahasa pemrograman yang baru sehingga PostgreSQL sangat mudah dikembangkan maupun di implementasikan pada tingkat user. PostgreSQL versi 7.0.x dan versi di atasnya menyertakan dokumentasi maupun berbagai macam contoh pembuatan fungsi maupun sebuah prosedur. Dengan keluwesan dan fitur yang dimilikinya, PostgreSQL patut bahkan melebihi jika disandingkan dengan database yang berskala besar lainnya. Jika kita menggunkan sebuah database , tentunya tak lepas dari tujuan dan maksud apa yang ingin dicapai serta kelebihan yang bagaimana yang kita inginkan.

####PostgreSQL juga mendukung beberapa fitur database modern, antar lain;
- complex queries
- foreign keys triggers
- views
- transactional integrity
- multiversion concurrency control

####Selain itu PostgreSQL juga dapat di extend sesuai kebutuhan pengguna melalui beberapa metode dengan menambangkan obyek baru, seperti
- Penambahan Tipe Data
- Penambahan Fungsi
- Penambahan Operator
- Penambahan Fungsi Aggregate
- Metode Index
- Bahasa prosedural

###Pengguna PostgreSQL
- Yahoo! untuk analisa prilaku pengguna web, menyimpan 2 petabyte data dan mengklaim sebagai gudang data terbesar. Menggunakan versi PostgreSQL yang dimodifikasi, dengan engine penyimpanan berbasis kolom yang sepenuhnya berbeda.[6][7]
- MySpace, situs jejaring sosial populer, menggunakan basisdata Aster nCluster untuk gudang data, dibangun diatas PostgreSQL tanpa modifikasi.[8][9]
- OpenStreetMap, proyek kolaboratif untuk menciptakan peta dunia yang bebas sunting.[10]
- Afilias, register domain untuk .org, .info, dan sebagainya.[11]
- Sony Online multiplayer online game.[12]
- BASF, platform belanja untuk portal agribisnisnya.[13]
- hi5.com portal jejaring sosial.[14]
- Skype aplikasi VoIP, basisdata pusat bisnis.[15]
- Sun xVM, perangkat lunak virtualisasi dan otomasi datacenter milik Sun.
- PostgreSQL adalah tujuan umum dan sistem manajemen database relasional obyek, yang paling canggih dari sistem database open source. Database postgreSQL dikembangkan berdasarkan POSTGRES 4.2 di departemen Ilmu Komputer, Berkeley University of California, inilah salah satu sebab menjadi database open source.

Database postgreSQL  sebagai database open source dirancang untuk berjalan pada platform UNIX-like. Namun, database postgreSQL kemudian selain sebagai database open source juga dirancang untuk menjadi portabel sehingga dapat berjalan di berbagai platform seperti Mac OS X, Solaris dan Windows.

Database postgreSQL adalah perangkat lunak bebas dan open source. Kode sumbernya tersedia di bawah lisensi database postgreSQL  lisensi open source liberal. Anda bebas untuk memakai, memodifikasi serta mendistribusikan database postgreSQL dalam wujud apa pun.

Database postgreSQL memerlukan upaya sangat minimal, sebagai database open source  ini dipertahankan karena stabilitasnya. Oleh karena itu, jika Anda mengembangkan aplikasi berbasis database postgreSQL yang notabene adalah database open source, postgreSQL tutorial memastikan biaya total kepemilikan yang rendah bila dibandingkan dengan sistem manajemen database lain, ataupun database berbayar yang lain.
Menyoroti  Fitur Database PostgreSQL. 

####Database postgreSQL sebagai database open source mempunyai banyak feature canggih yang sebagaimana system manajemen database perusahaan lain tawarkan, seperti:
- jenis-jenis User-defined 
- Tabel inheritance
- mekanisme penguncian Canggih
- Foreign key referential integrity
- Views, rules, sub-select
- Transaksi bersarang / nested transaction (savepoints)
- Multi-version concurrency control (MVCC)
- Asynchronous replication

####Versi terbaru dari database postgreSQL sebagai database open source mendukung fitur berikut:
- Microsoft Windows versi Server
- Tabel spasi
- Point-in-time recovery
Dan fitur baru lainnya ditambahkan dalam setiap rilis baru dari database postgreSQL sebagai database open source yang paling stabil saat ini.

#####Apa yang membuat Database PostgreSQL sebagai pelopor database open source lebih menonjol dari database lainnya?

Database postgreSQL adalah sistem manajemen database open source pertama yang mengimplementasikan fitur multi-versi-concurrency control (MVCC), bahkan sebelum Oracle yang notabene adalah database berbayar. Fitur MMVC dikenal sebagai isolasi snapshot di Oracle.

Database postgreSQL sebagai database open source adalah obyek sistem manajemen database relasional dengan tujuan umum. Hal ini memungkinkan Anda untuk menambahkan fungsi kustom untuk dikembangkan dengan menggunakan bahasa pemrograman yang berbeda seperti C / C + +, Java, dll

Database postgreSQL sebagai database open source dirancang untuk bisa diperluas. Di database postgreSQL  Anda dapat menentukan jenis data Anda sendiri, jenis indeks, bahasa fungsional, dll. Pada database open source  ini, jika Anda tidak menyukai setiap bagian dari sistem, Anda selalu dapat mengembangkan custom plugin untuk meningkatkan itu untuk memenuhi kebutuhan Anda misalnya, menambahkan optimizer baru.

Jika anda memerlukan dukungan apapun, komunitas yang aktif tersedia untuk membantu. Anda selalu dapat menemukan jawaban dari masyarakat PostgreSQL untuk masalah yang mungkin Anda miliki ketika bekerja dengan database postgreSQL sebagai database open source  Banyak perusahaan menawarkan layanan dukungan komersial dalam kasus Anda memerlukan lebih.



## Go dan NoSQL: RethinkDB
Go dan NoSQL: RethinkDB
GoRethink adalah salah satu driver klien pihak ketiga yang paling populer dan terpelihara dengan baik untuk RethinkDB.
Untuk baru-baru ini diperbarui driver untuk membuat kompatibel dengan RethinkDB 1,16, menambahkan dukungan untuk changefeeds. Pada fitur concurrency asli, bahasa ini membuatnya mudah untuk mengkonsumsi changefeeds dalam aplikasi realtime Go.
RethinkDB changefeeds ini menyediakan cara untuk berlangganan aliran update database realtime. 
Disini kami mengambil kode Go untuk dijadikan contoh,  seperti berikut.

type Issue struct {
    Description, Type string
}

db, err := r.Connect(r.ConnectOpts{Address: "localhost:28015"})
if err != nil {
    log.Fatal("Database connection failed:", err)
}

issues, _ := r.Db("rethinkdb").Table("current_issues").Filter(
    r.Row.Field("critical").Eq(true)).Changes().Field("new_val").Run(db)

go func() {
    var issue Issue
    for issues.Next(&issue) {
      if issue.Type != "" {
        log.Println(issue.Description)
      }
    }
}()

The ReQL ini menggunakan Filter untuk mencocokkan masalah di mana properti penting membawa nilai  yang sebenarnya.
The changefeed query ini hanya memancarkan dokumen yang cocok dengan filternya.
Ketika menerima output dari changefeed, kita langsung dapat membungkus di goroutine, sehingga asynchronous dapat beroperasi di latar belakang bukan menghalangi eksekusi. 
Menggunakan goroutines untuk pemrograman asynchronous dapat menyederhanakan arsitektur aplikasi realtime yang ada.

Membuat bot IRC
GoIRC terdapat perpustakaan yang membuatnya relatif mudah untuk membuat IRC bot sederhana. Kode berikut ini kami ambil di buku refrensi, dimana ini untuk menghubungkan ke server IRC dan menginstruksikan bot untuk bergabung dengan saluran tertentu :

ircConf := irc.NewConfig("mybot")
ircConf.Server = "localhost:6667" 
bot := irc.Client(ircConf)

bot.HandleFunc("connected", func(conn *irc.Conn, line *irc.Line) {
    log.Println("Connected to IRC server")
    conn.Join("#mychannel")
})

Jika terdapat masalah IRC bot memberikan pemberitahuan, disini hanya menambahkan beberapa baris saja, dari kode sebelumnya :

issues, _ := r.Db("rethinkdb").Table("current_issues").Filter(
    r.Row.Field("critical").Eq(true)).Changes().Field("new_val").Run(db)

go func() {
    var issue Issue
    for issues.Next(&issue) {
        if issue.Type != "" {
            text := strings.Split(issue.Description, "\n")[0]
            message := fmt.Sprintf("(%s) %s ...", issue.Type, text)
            bot.Privmsg("#mychannel", message)
        }
    }
}()

Disini juga terdapat beberapa perintah dasar dari pengguna. kami ingin program untuk terus berjalan sampai pengguna di channel IRC memberitahu bot untuk berhenti.

quit := make(chan bool, 1)
...
bot.HandleFunc("privmsg", func(conn *irc.Conn, line *irc.Line) {
    log.Println("Received:", line.Nick, line.Text())
    if strings.HasPrefix(line.Text(), config.IRC.Nickname) {
        command := strings.Split(line.Text(), " ")[1]
        switch command {
        case "quit":
            log.Println("Received command to quit")
            quit <- true
        }
        ...
    }
})

disini menggunakan pernyataan switch supaya lebih mudah untuk memperkenalkan perintah baru di masa depan dengan menambahkan kasus tambahan yang cocok dengan string lainnya. Untuk saat ini masih sangat sederhana. 





 Membangun aplikasi web realtime dengan Go dan RethinkDB

IRC ini cukup baik untuk digunakan latihan, tapi disini kami belom mengerti untuk membangun aplikasi web realtime dengan driver Go.
Tapi kami disini mengambil sebuah contoh dari buku referensi yang kami baca, untuk membangun sebuah versi Go dari aplikasi monitoring klaster sederhana.

server, _ := socketio.NewServer(nil)

conn, _ := r.Connect(r.ConnectOpts{Address: "localhost:28015"})
stats, _ := r.Db("rethinkdb").Table("stats").Filter(
    r.Row.Field("id").AtIndex(0).Eq("cluster")).Changes().Run(conn)

go func() {
    var change r.WriteChanges
    for stats.Next(&change) {
        server.BroadcastTo("monitor", "stats", change.NewValue)
    }
}()

































