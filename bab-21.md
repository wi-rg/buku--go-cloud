# Framework Aplikasi Web
Setelah beberapa tahun membangun server soket di Go, Saya tiba-tiba Memiliki kebutuhan untuk membangun beberapa API berbasis REST HTTP untuk Javascript dan Objective-C klien. Memiliki latar belakang di Rails, dan memiliki harapan yang cukup tinggi untuk tujuan umum MVC framework. Namun kerangka seperti Sinatra di Ruby keseimbangan yang bagus kesederhanaan untuk sederhana poin REST End. Aku memilih untuk Lanjutkan dengan Go untuk esta, Karena kinerja dan statis mengetik, saya pikir akan menang pada proyek-proyek jangka panjang tidak perlu melakukan banyak template dan bentuk web dalam pengembangan aplikasi gaya tradisional.
##Beego

Beego terlihat menjadi kerangka MVC penuh fitur, mungkin pada tingkat yang sama seperti Rails. Itu tidak membuat Upaya sama yang kecil. Ini perpustakaan logging, ORM dan Web kerangka.
komunitas:

Masyarakat Beego cukup besar. Namun Muncul perusahaan menggunakannya adalah Cina, jadi pasti ada hambatan bahasa untuk komunitas ini. Mungkin dapat skala untuk banyak lalu lintas sebagai situs Cina secara besar-besaran besar menggunakan framework.
##Best Feature:

Karena ini adalah kerangka kerja MVC penuh baku, Anda Tidak perlu menjelajahi internet untuk ton perpustakaan. Banyak pertanyaan keluar jendela segera: seperti apa kerangka yang digunakan dan bagaimana struktur aplikasi. Ini adalah jenis fitur yang menghemat banyak waktu pada awal sebuah proyek baru.
##Martini

Martini, terinspirasi oleh Sinatra, memiliki nuansa kerangka kerja yang sangat ringan. Ini menangani hal-hal dasar seperti routing, penanganan eksepsi, dan metode umum untuk melakukan middleware. Awalnya ada beberapa reaksi di masyarakat Go karena memiliki banyak teknik refleksi untuk pembersihan struktur API dari routing. Hal ini dapat melakukan hal-hal keren seperti dinamis menyuntikkan data ke set yang berbeda dari penangan berdasarkan jenis (lihat fitur terbaik untuk contoh). Ini 'ajaib' yang umum di kerangka Ruby, telah mendapat banyak sambutan hangat di masyarakat Go. Begitu banyak sehingga yang penulis bekerja pada kerangka sederhana disebut Negroni. Namun, di review saya, itu tidak Memberikan nilai sebanyak Martini tidak.
komunitas:

Meskipun komunitas bertubuh kecil, masih Sepertinya sangat aktif dan ada Sekitar 20 atau lebih plugin aktif, juga saya punya satu Pengintegrasian ke sumber GetSentry. Mengingat Bahwa mantra dari kerangka kerja ini adalah untuk menjadi kecil, itu Masuk Akal tidak ada banyak hal yang perlu untuk add ons.
##Best Feature:

Karena framework itu dapat menggunakan refleksi dinamis memasukkan data ke dalam fungsi handler, Tergantung pada kebutuhan Anda lihat contoh di bawah ini. Anda bahkan dapat secara dinamis menambahkan layanan baru.

~~~bash
m.Get("/", func(res http.ResponseWriter, req *http.Request) { // res and req are injected by Martini
  res.WriteHeader(200) // HTTP 200
})
m.Get("/", func() {
  // show something
})
m.Get("/", func(c *Context, r render.Render) {
  // show something
})
~~~

##gorila

Mungkin yang terbesar dan terpanjang framework web berjalan Gi, Gorilla adalah kerangka modular yang Bisa Punya sebanyak atau sesedikit mungkin bagi pengguna. Misalnya, dalam satu proyek kami hanya Sesi Mengambil keluar paket dan digunakan kembali untuk Tupoksi. Saya pikir Gorilla bagus penyebab banyak komponen yang dapat digunakan kembali Dengan perpustakaan net / http lurus.
komunitas:

Gorilla mungkin memiliki komunitas terbesar berbicara Inglés semua kerangka. Ini menunjukkan, yang volume geser posting blog dan middleware tersedia di GitHub.
##Best Feature:

Saya pikir ada banyak hype sekitar menggunakan web Days ini soket untuk secara dinamis memperbarui aplikasi dalam kerangka seperti Meteor.js. Gorilla memiliki soket web luar kotak, sehingga Anda dapat menghubungkan kode yang sama persis untuk REST endpoint sebagai soket web Anda tanpa menggunakan layanan pihak ketiga seperti Pusher.
##GoCraft

Saya pikir salah satu masalah terbesar dengan net / http adalah tidak ada cara untuk lulus konteks rantai handler Anda. Jadi Anda tidak dapat berbagi data atau transaksi Mudah Antara middleware dan penangan. Kerangka kerja ini AIMS untuk menjadi super-minim dan memecahkan hanya masalah esta. Mari saya tunjukkan beberapa contoh cepat

Mari kita mengambil contoh khas, mana beberapa bagian dari tengah kewenangan pengguna dan kemudian pengguna menempatkan objek ke dalam sesi.

~~~bash
router := web.New(YourContext{})
router.Middleware((*YourContext).UserAuthRequired)
router.Put("/users/:id", (*YourContext).UsersUpdate)

func (c *YourContext) UserAuthRequired(rw web.ResponseWriter, r *web.Request, next web.NextMiddlewareFunc) {
    //Auth the user and put the data into the context
    user := userFromSession(r)  // Pretend like this is defined. It reads a session cookie and returns a *User or nil.
    if user != nil {
        c.User = user
        next(rw, r)
    } 
}

func (c *YourContext) Root(rw web.ResponseWriter, req *web.Request) {
    if c.User != nil {
        fmt.Fprint(rw, "Hello,", c.User.Name)
    } else {
        fmt.Fprint(rw, "Hello, anonymous person")
    }
}
~~~

##Net / HTTP

Jika Anda membaca [Go] milis, mungkin akan Dikatakan ini adalah satu-satunya kerangka yang Anda butuhkan. Dalam beberapa hal yang benar. Kami membangun server XMPP Dengan hanya Seluruh net / http, dan bekerja dengan baik. Namun aplikasi web kompleks Cenderung perlu middleware. Ada beberapa proyek yang menarik seperti menempatkan manakah Anda memungkinkan untuk mencampur dan mencocokkan middleware dari kerangka web lain GO Dengan standar net / http.
masyarakat

Masyarakat jelas cukup besar sebagai pengguna dapat menggunakan kembali bit dari banyak proyek lainnya. Namun memiliki antarmuka yang sangat terbatas, dan itu tidak benar-benar menentukan cara standar Memperluas middleware. Routing ITS cukup lemah di luar kotak sehingga Anda menggunakan kerangka dengan Biasanya itu.
###kesimpulan

Kami mencoba beberapa kerangka ini untuk mendapatkan merasa untuk itu. Kami menggunakan Martini untuk SISA API kami. Sementara kita menyukainya, saya pikir di masa depan saya akan mencoba menggunakan GoCraft Sebaliknya itu adalah sebagai bahkan lebih ringan namun memecahkan masalah konteks yang sama itu Martini memiliki. Mungkin kita akan mengadaptasi beberapa middleware Martini sebagai percobaan. Awalnya kami mencoba menggunakan sistem admin kami untuk Gorilla Yang saya pikir esta akhirnya menjadi kesalahan besar. Itu benar-benar Kekurangan Rails adalah bagaimana dogmatis, dan banyak fitur kecil untuk membuat benar-benar dasar cookie cutter aplikasi web. Untuk lurus ke atas bentuk tradisi aplikasi CRUD berdasarkan, saya pikir kita akan bersandar berat terhadap Beego di masa depan, karena benar-benar terasa seperti MVC fitur set lengkap.