# Go dan Microservices
golang dan microservice memiliki konsep yang mirip yaitu servis yang kecil, yang difokuskan
untuk satu fungnsionalitas tertentu dan berkomunikasi melalui HTTP, TPC atau Message Queue
    Ketika mempublish API untuk konsumsi umum HTTP dan JSON sudah menjadi standar
tetapi ada juga alternatif lain untuk komunikasi inter-servis dengan menggunakan library
protobuf
            (Protokol buffer adalah cara untuk meng encoding struktur data dengan cara yang efisien
            dan juga degan format yang dapat di extend atau diperluas/diperbesar
            google menggunakan protocol buffers untuk hampir semua protokol internal RPC)

berikut ini adalah contoh golang micro-services yang menerima koneksi HTTP/JSON pada level API
dan kemudian memanfaatkan gRPC untuk komunikasi inter-servis.

![golang micro-services HTTP/JSON](https://cloud.githubusercontent.com/assets/739782/7439604/d1f324c2-f036-11e4-958a-6f6913049946.png)

gambar diatas adalah API endpoint menerima request http pada micro.demo:8080 dan kemudian menghasilkan beberapa 
RPC request kepada backend servis.
    NB: Data pada setiap servis disimpan dalam format JSON dan file biasa dimasukan pada direktori /data/
        pada kenyataannya setiap servis dapat memilih datastore nya sendiri - sendiri. Sebagai contoh
        Geo servis dapat menggunakan PostGis atau database lain untuk digunakan pada geospacial queries.

## Memahami Microservices
microservice adalah istilah untuk salah satu jenis arsitektur software atau aplikasi
yang bersifat micro(sangat kecil) yang difokuskan untuk satu fungsionalitas tertentu,
dan tidak sekedar dengan memisahkan berdasarkan user-role atau subdomain saja, tetapi 
aplikasi akan di breakdown lebih rinci lagi dari sisi fungsionalitasnya. Aplikasi 
akan dirancang agar setiap fungsi bekerja secara independen, dan setiap fungsi dapat 
menggunakan teknologi stack yang sesuai dengan kebutuhan, yang artinya akan terdapat
teknologi yang berbeda - beda dalam satu aplikasi besar.

Dengan pemisahan aplikasi berdasarkan fungsinya ini, pada akhirnya akan menemui keragaman
teknologi dalam sebuah satu layanan, misalkan dari layanan blog yang dapat dipecah menjadi 
beberapa fungsi utama misalnya : fungsi konten, user management, komentar, rangking, search
dan lainnya.
Pada bagian fungsi konten mungkin kita dapat menggunakan PHP dan mysql, tetapi pada komentar
kita menggunakan python dan mongodb, sedangkan di search menggunakan nodejs dan elasticsearch untuk 
menyimpan datanya. Disini terlihat bahwa setiap fungsi / permasalahan teknis dapat diselesaikan 
dengan cara dan teknologi yang berbeda - beda.  
Dalam konsep microservice, kita tidak hanya melakukan pemisahan di level aplikasi saja, tetapi
dari sisi infrastruktur juga, kita akan menemukan keragaman arsitektur, konfigurasi dan optimalisasi
sistem yang berbeda, dan sering juga ditemukan jumlah dan spesifikasi server yang tidak sama antara
service yang satu dengan yang lain.


Sumber : 	google.com 
			https://github.com/harlow/go-micro-services
			https://id.techinasia.com/talk/migrasi-microservice-tidak-optimal
			
## Arsitektur Microservices



## Go dan Microservices

Bagaimana serta peranti apa saja yang diperlukan untuk mengimplementasikan microservices di Go


## Microservices di Go Menggunakan GoKit



## Microservices di Go Menggunakan grpc

Microservices telah menjadi kata kunci sejak tahun lalu dan banyak arsitek yang mengeksplorasi manfaat yang mereka bisa dapatkan dengan menggunakan itu. Tapi harus dapat diekspos sebagai microservices, mencoba untuk sebanyak mungkin dan mengurangi jumlah proses pada sistem.
Saya merasa bahwa Go benar-benar baik dan menawarkan fleksibilitas yang lebih baik sebagai ukuran proses, tampaknya lebih rendah daripada bahasa lainnya. Saya melihat minimal overhead kurang. artcile ini bukan tentang mengapa MICROSERVICE & bagaimana seharusnya Anda memilih layanan Anda?saya ingin menempatkan bagaimana Anda dapat menulis webapplication berdasarkan MICROSERVICE di golang menggunakan grpc. Memilih dua layanan 1) layanan Katalog 2) Rekomendasi layanan sebagaimana microservices. webapplication yang dibangun menggunakan go. Webapplicaiton di golang, dua microservices di backend.<1>

![Repository microservices grpc](images/microservice-components.png)

Web Server (main.go): Anda dapat membuat server web di golang dengan sangat sedikit garis ... Di bawah ini adalah kode ..

~~~bash
http.HandleFunc(“/catalog”, GetProductCatalog)
http.Handle(“/”, http.FileServer(http.Dir(“public”)))
http.ListenAndServe(“:8080”, nil)
~~~

 proto3 adalah versi yang digunakan dalam file penyangga protokol grpc .Three diciptakan dan dimasukkan ke dalam direktori "proto".
Kode go sesuai dihasilkan dengan menjalankan perintah "protoc -aku ../protos ../protos/*.proto -go_out = plugin = grpc: basicwebapp". Saya telah memeriksa di file yang dihasilkan juga untuk memudahkan pemahaman.

CatalogServer / CatalogServiceServer:
 Proses go adalah CatalogServicesServer dan melayani "layanan katalog". Anda harus dapat menavigasi ke folder "CatalogServer" dalam kode sumber di github
menjalankan "pergi membangun CatalogServiceServer.go" di commandline. Jika Anda berhasil maka catalogserviceserver.exe akan dihasilkan. menjalankannya untuk memulai server. Tujuan utama dari MICROSERVICE ini adalah untuk memberikan produk.

~~~bash
    service CatalogService {

    rpc GetProductCatalog(CatalogRequest) returns (CatalogResponse) {}

    }

    message Category {

    string categoryName=1;
    }

    message CatalogResponse {

    repeated Product products =1;

    }

    message CatalogRequest {

    Category category=1;
    }
~~~

File penyangga protokol yang menyatakan metode rpc dan serialisasi format dalam folder "proto" ... di bawah ini adalah inti dari itu ..

~~~bash    service RecommendationService {

    rpc GetRecommendations(Product) returns (RecommendationResponse);
    }
    message RecommendationResponse{

    message Recommendation{

    int32 rating =1;

    int32 productid=2;
    }
    repeated Recommendation result=1;

    }
~~~


Complile server rekomendasi dan menjalankannya sebelum menjalankan webserver ..

Komunikasi antara server grpc dan webserver:

1. Jalankan main.go

2. Jenis localhost: 8080 di browser .. Anda harus dapat melihat halaman html dengan tombol submit seperti di bawah ini.

![Repository microservices grpc2](images/microservices.png)

Dengan mengklik tombol submit akan dikirimkan ke server. Path target "katalog" sehingga akan ditangani oleh fungsi callback getproductcatalog. Cari kode di bawah ini di fungsi yang bertanggung jawab untuk memanggil MICROSERVICE ..
~~~bash
    conn, err := grpc.Dial(catalogServerAddress, opts…) //DIaling in to the catalogServerAddress “port is hardcoded at 50052”

    if err !=nil {
    log.Fatal(“fail to dial %v”, err)
    }

    defer conn.Close()

    client := pb.NewCatalogServiceClient(conn) // Create new Catalogserviceclient stub.
    stream, err := client.GetProductCatalog(context.Background(), &pb.CatalogRequest{&pb.Category{CategoryName:”Computer”}}) // Invoke the remote method “GetProductCatalog” and get the response into stream variable..
	~~~
	
Sumber : <1>https://devicharan.wordpress.com/2015/04/12/microservices-based-webapplication-in-golang-using-grpc/
source code :  https://github.com/devicharan/basicwebapp
[grpc](http://grpc.io)

introduction :

bab 1. Taufik Nur Rakhaman/145410215 ## Memahami Microservices
bab 2. Angga Ibnu Saputra/115410162 ## Arsitektur Microservices dan Microservices di Go Menggunakan grpc
bab 3. Fendi Djatmiko/145410203 ## Microservices di Go Menggunakan GoKitGo
bab 4. Yandri W.Sironga/125410215 ## go dan Microservices



