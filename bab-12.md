# Akses Basis Data

## Sekilas Tentang Basis Data



### Basis Data SQL



### Basis Data NoSQL


### Basis Data NewSQL



## Paket Standar Go untuk Akses Basis Data




## Driver Go untuk Akses Basis Data
tipe Conn
type Conn interface {
        // Menyiapkan kembali sebuah pernyataan, terikat pada hubungan ini.
         Prepare(query string) (Stmt, error)
        // Menutup yang tidak valid dan berpotensi berhenti saat apapun
        // prepared statement dan transaksi, menandai ini
        // koneksi seperti tidak lagi digunakan.
        //
        // Karena paket sql memelihara pool bebas dari
        // koneksi dan hanya panggilan Tutup ketika ada surplus dari
        // koneksi idle, seharusnya tidak perlu driver untuk
        // melakukan caching koneksi mereka sendiri.
        Close() error

        // Mulailah start dan returns sebuah transaksi baru.
        Begin() (Tx, error)
}
Conn adalah koneksi ke database. Hal ini tidak digunakan secara bersamaan oleh beberapa goroutines. Conn diasumsikan untuk menjadi stateful.

tipe Driver
type Driver interface {
        // Buka kembali sambungan baru ke database.
        // Nama adalah string dalam format sebuah driver-spesifik.
        //
        // Dapat dibuka kembali koneksi cache (yang sebelumnya
        // ditutup), namun hal ini tidak perlu; paket sql
        // memelihara pool dari koneksi idle agar efisien saat digunakan kembali.
        //
        // Returned connection hanya digunakan oleh satu goroutine pada satu
        // waktu.
        Open(name string) (Conn, error)
}
Driver adalah interface yang harus diimplementasikan oleh driver database.

tipe Execer
type Execer interface {
        Exec(query string, args []Value) (Result, error)
}
Execer adalah sebuah antarmuka opsional yang dapat diimplementasikan oleh Conn.
Jika sebuah Conn tidak mengimplementasikan Execer, paket sql ini DB. Exec akan mempersiapkan query pertama, mengeksekusi pernyataan, dan kemudian tutup pernyataan tersebut.
Exec dapat kembali ke ErrSkip.

tipe NotNull
type NotNull struct {
        Converter ValueConverter
}
NotNull adalah tipe yang mengimplementasikan ValueConverter oleh pelarangan nilai nol tetapi sebaliknya mendelegasikan ke ValueConverter lain.

func (NotNull) ConvertValue
func (n NotNull) ConvertValue(v interface{}) (Value, error)

tipe Null
type Null struct {
        Converter ValueConverter
}
Null adalah tipe yang mengimplementasikan ValueConverter dengan memungkinkan nilai-nilai nol tetapi sebaliknya mendelegasikan ke ValueConverter lain.

func (Null) ConvertValue
func (n Null) ConvertValue(v interface{}) (Value, error)

tipe Queryer
type Queryer interface {
        Query(query string, args []Value) (Rows, error)
}
Queryer adalah sebuah antarmuka opsional yang dapat diimplementasikan oleh Conn.
Jika Conn tidak mengimplementasikan Queryer, paket sql ini DB. Query akan mempersiapkan query pertama, mengeksekusi pernyataan, dan kemudian menutup pernyataan tersebut.
Query dapat kembali ErrSkip.

tipe Result
type Result interface {
        // LastInsertId mengembalikan ID otomatis yang dihasilkan database ini
        // setelah, misalnya, INSERT ke tabel dengan
        // kunci primer.
        LastInsertId() (int64, error)
        // RowsAffected mengembalikan jumlah baris yang dipengaruhi oleh
        // query.
        RowsAffected() (int64, error)
}
Hasilnya adalah hasil dari eksekusi query.

tipe Rows
type Rows interface {
        // Kolom yang mengembalikan nama-nama dari kolom. Jumlah
        // kolom hasilnya disimpulkan dari panjang
        // slice. Jika nama kolom tertentu tidak diketahui, sebuah string kosong
        // yang harus dikembalikan untuk masuk.
        Columns() []string
        // Tutup menutup baris iterator.
        Close() error
        // Berikutnya dipanggil untuk mengisi baris berikutnya dari data ke dalam
        // slice yang disediakan. Slice yang disediakan akan sama
        // ukurannya dengan columns() lebar.
        //
        // Slice dest dapat hanya diisi dengan
        // sebuah driver tipe Value, tetapi tidak termasuk string.
        // Semua nilai-nilai string harus dikonversi ke [] byte.
        //
        // Berikutnya harus kembali io. EOF ketika tidak ada lagi baris.
        Next(dest []Value) error
}
Rows adalah iterator atas hasil query dieksekusi.

tipe RowsAffected
type RowsAffected int64
RowsAffected mengimplementasikan Result untuk INSERT atau UPDATE operasi yang bermutasi ke sejumlah baris.

func (RowsAffected) lastInsertId
func (RowsAffected) LastInsertId() (int64, error)

func (RowsAffected) RowsAffected
func (v RowsAffected) RowsAffected() (int64, error)

tipe Stmt
type Stmt interface {
        // Close menutup pernyataan. 
        //
        // Sampai Go 1.1, Stmt yang tidak akan ditutup jika digunakan
        // oleh beberapa pertanyaan.
        Close() error

        // NumInput mengembalikan jumlah parameter placeholder.
        //
        // Jika NumInput mengembalikan > = 0, paket sql akan melakukan cek kewajaran
        // menghitung argumen dari pemanggil dan mengembalikan kesalahan ke pemanggil
        // sebelum pernyataan Exec atau Query methods dipanggil.
        //
        // NumInput juga dapat kembali -1, jika driver tidak mengetahui
        // jumlah placeholder. Dalam hal ini, paket sql
        // tidak akan memeriksa kewajaran jumlah argumen Exec atau Query.
        NumInput() int

        // Exec mengeksekusi sebuah query yang tidak dapat kembali ke baris, seperti
        // INSERT atau UPDATE.
        Exec(args []Value) (Result, error)

        // Query mengeksekusi sebuah query yang dapat kembali ke baris, seperti
        // SELECT.
        Query(args []Value) (Rows, error)
}
Stmt adalah sebuah prepared statement. Hal ini terikat dengan Conn dan tidak digunakan oleh beberapa goroutines bersamaan.

tipe Tx
type Tx interface {
        Commit() error
        Rollback() error
}
Tx adalah sebuah transaksi.

tipe Value
type Value interface{}
Value adalah nilai yang driver harus mampu menangani. Itu baik nol atau turunan dari salah satu tipe:
int64
float64
bool
[]byte
string [*] di manapun kecuali dari Rows. Next.
time.Time

tipe ValueConverter
type ValueConverter interface {
        // ConvertValue converts a value to a driver Value.
        ConvertValue(v interface{}) (Value, error)
}
ValueConverter adalah antarmuka yang menyediakan metode ConvertValue.
Berbagai implementasi dari ValueConverter disediakan oleh paket driver untuk menyediakan implementasi yang konsisten dari konversi antara driver. ValueConverters memiliki beberapa kegunaan:

* Mengkonversi dari tipe Value yang disediakan oleh paket sql  menjadi tipe kolom tertentu pada tabel database dan memastikan itu  cocok, seperti memastikan suatu int64 tertentu sesuai dalam kolom uint16 tabel.

* Mengkonversi nilai seperti yang diberikan dari database ke dalam salah satu  tipe driver Value.

* Dengan paket sql, untuk mengkonversi dari tipe driver Value tipe pengguna di scan.

tipe Valuer
type Valuer interface {
        // Value returns a driver Value.
        Value() (Value, error)
}
Valuer adalah antarmuka menyediakan method Value.
Tipe mengimplementasikan antarmuka Valuer mampu mengkonversi diri untuk Value driver.




## Go dan SQL: PostgreSQL



## Go dan NoSQL: RethinkDB




