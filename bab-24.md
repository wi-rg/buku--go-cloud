# Go dan AppEngine

##Ikhtisar

Memanggil API Google menggunakan bahasa pemrograman Go . Ini menunjukkan Anda, bagaimana membangun sebuah aplikasi web App Engine sederhana menggunakan Go 

##tujuan

   * Membangun kerangka untuk aplikasi web Go App Engine Anda .
   * Memanggil API Google+ dari Go .
   * Menyebarkan aplikasi web Anda untuk App Engine .
   
##Prasyarat

Slahkan download dan Unzip Go App Engine

contoh aplikasi ini memiliki tujuan sederhana : menampilkan terbaru " Gopher " gambar dari aliran publik Google+ . Kode sumber tersedia di GitHub .

   * Hello Go App Engine
   * Hello Google+
   * ke App Engine
	
###Hello Go App Engine

Dalam materi ini , Anda akan membangun kerangka untuk aplikasi web Go App Engine Anda , direktori kerja Anda akan terlihat seperti ini :
~~~bash
google_appengine/
goplus/
    app.yaml
    hello.go
~~~

1. Download dan unzip Go App Engine SDK
2. Buat proyek Google Cloud Platform Console , menciptakan proyek App Engine Anda dan ysng akan digunakan nanti untuk mengkonfigurasi kredensial API Anda .

    a. Dalam Cloud Platform Konsol Google , klik Create Project .
    b. Pilih nama proyek dan ID . ID proyek menjadi bagian dari URL proyek App Engine Anda :
	   http://YOUR_PROJECT_ID.appspot.com . Anda akan menggunakan ID ini dalam pelajaran nanti .
	   
3. Membuat sebuah direktori bernama goplus baru dalam direktori kerja Anda . Direktori ini akan berisi file aplikasi App Engine Anda.
4. Di dalam direktori goplus , membuat file bernama app.yaml dengan isi sebagai berikut . File ini berisi konfigurasi aplikasi App Engine Anda.

~~~bash
application: goplus
version: 1
runtime: go
api_version: go1

handlers:
  - url: /.*
    script: _go_app
~~~

5. di dalam direktori goplus , membuat file bernama hello.go dengan isi sebagai berikut :

~~~bash
package goplus

import (
        "fmt"
        "net/http"
)

// init is called before the application starts.
func init() {
        // Register a handler for /hello URLs.
        http.HandleFunc("/", hello)
}

// hello is an HTTP handler that prints "Hello Gopher!"
func hello(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "Hello, Gopher!")
}
~~~

6. Mulai pembangunan server dengan menjalankan dev_appserver.py ( terletak di direktori google_appengine ) pada baris perintah , dengan direktori aplikasi Anda sebagai argumen pertama .
~~~bash
    /path/to/google_appengine/dev_appserver.py goplus
~~~
Anda dapat menggunakan -a 0.0.0.0 pilihan untuk mengikat server pengembangan ke semua alamat IP yang tersedia.

7. Buka http : // localhost : 8080 di browser Anda .

Anda harus melihat : Hello Gopher !

###Hello Google+

Sekarang Anda telah membuat aplikasi Go di App Engine, Anda Akan mencari arus Google+ publik untuk foto gopher.

####tujuan

   * Menyiapkan Google Library API Klien untuk Go
   * Buat proyek Google Cloud Platform Konsol dan mengaktifkan API Google+
   * Memanggil API Google+ dari Go.

Dalam rangka untuk membuat Google+ API panggilan, Anda harus terlebih dahulu mengimpor Google API Client Library untuk Go.

Google Libraries API Klien dapat ditemukan di bawah google.golang.org/api. Dengan konvensi paket Go diberi nama setelah URL repositori mereka, dan dipasang oleh alat go di direktori $ GOPATH.

   1. Download paket google-api-go-klien untuk Google+ API:
~~~bash
    / Path / ke / google_appengine / pergi get -v -d google.golang.org/api/plus/v1
~~~
    2. Aktifkan API Google+. Pilih proyek App Engine yang Anda buat dalam pelajaran sebelumnya. Klik Lanjutkan.

    3. Jika Anda belum melakukannya, buat key API aplikasi Anda dengan mengklik Tambahkan kredensial> API key. Berikutnya, mencari kunci API Anda di bagian kunci API.

    4. Dalam direktori goplus, membuat file bernama config.go dengan isi sebagai berikut, dan ganti YOUR_API_KEY dengan kunci api Anda yang dihasilkan.
~~~bash
package goplus

// apiKey is the Google API key of the project copied from the Google API Console.
var apiKey = "YOUR_API_KEY"
~~~
	5. Dalam direktori goplus , mengedit file app.yaml dan mengganti aplikasi : goplus dengan ID proyek Anda dalam materi sebelumnya .
~~~bash
application: YOUR_PROJECT_ID
~~~
	6. Dalam direktori goplus , membuat file bernama gopher.go dengan konten berikut :
~~~bash
package goplus

import (
        "net/http"

        // Import appengine urlfetch package, that is needed to make http call to the api
        "appengine"
        "appengine/urlfetch"

        // Import google api go client library
        "google.golang.org/api/googleapi/transport"
        // Import Google+ package, the package will be named "plus"
        "google.golang.org/api/plus/v1"
)

// gopherFallback is the official gopher URL (in case we don't find any in the Google+ stream)
const gopherFallback = "http://golang.org/doc/gopher/gophercolor.png"

// init is called before the application start
func init() {
        // Register a handler for /gopher URLs.
        http.HandleFunc("/gopher", gopher)
}

// gopher is an HTTP handler that searches Google+ for an activity
// with a Gopher photo and redirects to the image thumbnail.
func gopher(w http.ResponseWriter, r *http.Request) {
        // Create appengine context, needed to do urlfetch call
        c := appengine.NewContext(r)

        // Create a new http client, supplying the API key we
        // generated to identify our application, and the urlfetch
        // transport necessary to make HTTP calls on App Engine
        transport := &transport.APIKey{
                Key:       apiKey,
                Transport: &urlfetch.Transport{Context: c}}
        client := &http.Client{Transport: transport}

        // Create the plus service object with which we can make an API call
        service, err := plus.New(client)
        if err != nil {
                // error creating the Google+ client
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
        }

        // Search the stream for "gopher" related activities
        result, err := service.Activities.Search("#gopher").Do()
        if err != nil {
                // error searching Google+ for gopher
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
        }

        // Iterate through the activities search result until we find a photo or album attachment
        for _, item := range result.Items {
                for _, att := range item.Object.Attachments {
                        switch att.ObjectType {
                        case "photo":
                                // Redirect to the gopher thumbnail
                                http.Redirect(w, r, att.Image.Url, 302)
                                return
                        case "album":
                                http.Redirect(w, r, att.Thumbnails[0].Image.Url, 302)
                                return
                        }
                }
        }
        // Fallback on the official gopher if we didn't find any gophers in the stream
        http.Redirect(w, r, gopherFallback, http.StatusFound)
}
~~~

Fungsi gopher dalam beberapa hal :
        a. Mengatur layanan plus menggunakan kunci pengembang ,
        b. Memanggil metode Cari dengan kata kunci " gopher " ,
        c. Memindai hasilnya , mencari foto atau album lampiran , dan
        d. Mengarahkan klien ke URL gambar pertama ditemukan .

    Google -api -go - client menyediakan satu set jenis Go dengan metode untuk berinteraksi dengan API Google+ seperti yang Anda lakukan dengan paket asli . jenis ini dihasilkan secara otomatis menggunakan Google API Discovery Service .

    7. Kunjungi http : // localhost : 8080 / gopher di browser Anda .
	
###Deploy ke App Engine

Sekarang Anda dapat melepaskan aplikasi Anda di dunia dengan mengerahkan ke App Engine .

    1. Mendeploykan aplikasi Anda dengan menjalankan appcfg.py , yang terletak di direktori SDK , dengan perintah update dan direktori aplikasi Anda sebagai argumen terakhir .
~~~bash
    /path/to/google_appengine/appcfg.py pembaruan goplus
~~~
    2. Memuat http://YOUR_APP_ID.appspot.com/gopher , Google+ .

Sekarang Anda atau teman Anda dapat dengan mudah menanamkan gambar terkait Gopher di halaman web mereka menggunakan cuplikan HTML ini .
~~~bash
< Img src = " YOUR_APP_ID.appspot.com/gopher " >
~~~

Sumber : https://cloud.google.com/appengine/training/go-plus-appengine/hello-appengine