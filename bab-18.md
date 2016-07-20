# Middleware
###1.1  Pengertian Middleware
Middleware adalah perangkat lunak komputer yang menyediakan layanan untuk aplikasi perangkat lunak di luar yang tersedia dari sistem operasi. Hal ini dapat digambarkan sebagai "perangkat lunak lem". Middleware memudahkan pengembang perangkat lunak untuk melakukan komunikasi dan input / output, sehingga mereka dapat fokus pada tujuan khusus dari aplikasi mereka. Middleware adalah perangkat lunak yang menghubungkan komponen perangkat lunak atau aplikasi perusahaan. Middleware adalah lapisan perangkat lunak yang terletak di antara sistem operasi dan aplikasi pada setiap sisi jaringan komputer terdistribusi. Biasanya, mendukung kompleks, aplikasi bisnis perangkat lunak yang didistribusikan.
###1.2  Tujuan
Perangkat Middleware memiliki beberapa tujuan, diantaranya adalah :
a.       Menyediakan fasilitas bagi programmer untuk dapat mendistribusikan objek yang digunakan pada beberapa proses yang berbeda.
b.      Dapat berjalan dalam satu mesin ataupun di beberapa mesin yang terhubung dengan jaringan.
jika boleh diperjelas, tujuan dari Middleware ialah sebagai interkoneksi interkoneksi beberapa aplikasi dan masalah interoperabilitas. Middleware sangat dibutuhkan untuk bermigrasi dari aplikasi mainframe ke aplikasi client/server dan juga untuk menyediakan komunikasi antar platform yang berbeda.
###1.3  Manfaat Middleware
Sebuah Abstraksi Middleware diciptakan sebagai perantara antara Sistem Operasi dengan Software Apliskasi yang terdistribusi pastinya memiliki manfaat yang besar :
a.       2 buah platform/aplikasi dapat dijalankan secara bersamaan pada sistem yang terdistribusi
b.      memungkinkan satu aplikasi berkomunikasi dengan lainnya walaupun berjalan pada platform yang berbeda
c.       Transparansi di seluruh jaringan sehingga menyediakan interaksi dengan layanan atau aplikasi lain
d.      Independen dari layanan jaringan
e.       Handal dan selalu tersedia

![Repository downloads Go](images/middleware.png)

Middleware sebagai perangkat yang dirancang untuk mendukung Enterprise Arsitektur (EA) sebagai sistem yang tersebar dan saling berhubungan, Ia memiliki beberapa layanan yang bisa digunakan dan dimanfaatkan.
Contoh layanan Middleware :

####Transaction Monitor
1. Produk pertama yang disebut middleware.
2. Menempati posisi antara permintaan dari program client dan database, untuk menyakinkan bahwa semua transaksi ke database terlayani dengan baik

####Messaging Middleware
![Repository downloads Go](images/layanan.png)
1. Menyimpan data dalam suatu antrian message jika mesin tujuan sedang mati atau overloaded
2. Berisi business logic yang merutekan message ke ujuan sebenarnya dan memformat ulang data lebih tepat
3. Sama seperti sistem messaging email, kecuali messaging middleware digunakan untuk mengirim data antar aplikasi

#####Produk Messaging Middleware
Produk utama messaging (pengiriman pesan) untuk pengaturan komunikasi asinkronus antar aplikasi adalah MQSeries dari IBM. MQSeries telah dipasangkan pada semua platform server. Microsoft memperkenalkan sistem messagingnya sendiri yang digabungkan dengan Component Object Model(COM), yaitu Microsoft Message QueueServer (MSMQ). MSMQ dan MQSeries menawarkan fungsi yang sama.

#####Distributed Object Middleware
menurut terminologinya, sebuah Object yang terdistribusikan oleh layanan Middleware, ini merupakan layanan utama yang dimiliki oleh Middleware. layanan ini dibagi menjadi beberapa :
Contoh: RPC (Remote Procedure Calls), CORBA (Common Object Request Broker Architecture) dan DCOM/COM (Distributed Component Object Model)

#####Middleware basis data
menyediakan antarmuka antara sebuah query dengan beberapa database yang terdistribusi
Contoh: JDBC, ODBC, dan ADO.NET

#####Application Server Middleware
J2EE Application Server, Oracle Application Server
- Lebih detailnya untuk keterangan Middleware sebagai Application Server ialah Sebuah Web-based Application server, yang menyediakan antarmuka untuk berbagai aplikasi,digunakan sebagai middleware antara browser dan aplikasi.
- J2EE adalah contoh application serverA wide range of server-side processing has been supported by appservers(i.e.;J2EE).