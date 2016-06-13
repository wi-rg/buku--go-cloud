# Go dan Microservices


## Memahami Microservices



## Arsitektur Microservices
Secara sederhana, arsitektur aplikasi Microservice ini menggunakan desain yang memecah aplikasi berdasarkan fungsinya secara spesifik. Tidak sekedar dengan memisahkan berdasarkan user-role atau subdomain saja, tetapi aplikasi akan di breakdown lebih rinci lagi dari sisi fungsionalitasnya. Aplikasi akan dirancang agar setiap fungsi bekerja secara independent. Dan setiap fungsi dapat menggunakan teknologi stack yang sesuai dengan kebutuhan, walaupun itu artinya akan terdapat teknologi yang berbeda-beda dalam satu aplikasi besar

(images/Graph.png)

source: https://www.nginx.com/blog/introduction-to-microservices/

Dengan pemisahan aplikasi berdasarkan fungsi-nya ini, pada akhirnya kita akan menemui keragaman teknologi dalam sebuah satu layanan digital. Misalkan dari layanan blog yang telah dicontohkan diatas, kita bisa coba pecahkan aplikasi blog tersebut menjadi fungsi konten, user management, komentar, rangking, search dan lainnya.

Pada bagian fungsi konten mungkin kita menggunakan PHP dan mysql, tetapi pada komentar kita menggunakan Python dan Mongodb, sedangkan di search menggunakan nodejs dan elasticsearch untuk penyimpanan datanya. Disini terlihat bahwa setiap fungsi / permasalahan teknis dapat diselesaikan dengan cara dan teknologi yang berbeda-beda.

Dalam konsep microservices, kita tidak hanya melalukan pemisahan di level aplikasi saja, tetapi dari sisi infrastruktur pula. Kita akan menemukan keragaman arsitektur infra, konfigurasi dan optimalisasi sistem yang berbeda, dan sering pula ditemukan jumlah dan spesifikasi server yang tidak sama antara service yang satu dengan yang lain.
## Go dan Microservices

Bagaimana serta peranti apa saja yang diperlukan untuk mengimplementasikan microservices di Go


## Microservices di Go Menggunakan GoKit



## Microservices di Go Menggunakan grpc

[grpc](http://grpc.io)


