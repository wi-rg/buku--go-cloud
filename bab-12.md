# Akses Basis Data

## Sekilas Tentang Basis Data



### Basis Data SQL



### Basis Data NoSQL


### Basis Data NewSQL



## Paket Standar Go untuk Akses Basis Data
Paket standar go untuk akses basis data adalah 'sql'. Paket sql menyediakan antarmuka umum di basisdata sql. Paket ini harus digunakan bersamaan dengan driver database. Untuk menggunakan paket sql ini maka harus dideklarasikan syntax :
import "database/sql".

### Variabel
var ErrNoRows = errors.New("sql: no rows in result set")
ErrNoRows dikembalikan oleh Scan ketika queryRow tidak kembali berturut-turut. Dalam kasus seperti itu, queryRow mengembalikan placeholder * nilai Row yang menunda kesalahan ini sampai Scan.
var ErrTxDone = errors.New("sql: Transaction has already been committed or rolled back")

### Fungsi Drivers
func Drivers() []string
Driver mengembalikan daftar diurutkan dari nama-nama driver terdaftar. 

### Fungsi Register
func Register(name string, driver driver.Driver)
Register membuat driver database yang tersedia dengan nama yang diberikan. Jika Register dipanggil dua kali dengan nama yang sama atau jika driver nol, itu buruk. 

### Tipe DB
type DB struct {
        // Mengandung filter atau field unexported
}
DB merupakan pegangan database yang mewakili kolam nol atau lebih yang mendasari koneksi. Ini aman untuk digunakan bersamaan oleh beberapa goroutines.
Paket sql menciptakan dan membebaskan koneksi secara otomatis; itu juga memelihara kolam bebas dari koneksi idle. Jika database memiliki konsep state per-koneksi, state tersebut hanya dapat diandalkan dan diamati dalam transaksi. Setelah DB.Begin dipanggil, Tx kembali terikat untuk koneksi tunggal. Setelah Komit atau Rollback dipanggil oleh transaksi, bahwa koneksi transaksi ini dikembalikan ke pool koneksi yang menganggur pada DB. Ukuran pool dapat dikontrol dengan SetMaxIdleConns.

### Fungsi Open
func Open(driverName, dataSourceName string) (*DB, error)
Open membuka database yang ditentukan oleh nama driver database dan nama sumber data spesifik-driver, biasanya terdiri dari setidaknya nama database dan informasi koneksi.
Sebagian besar pengguna akan membuka database melalui koneksi fungsi pembantu spesifik-driver yang mengembalikan *DB. Tidak ada driver database yang termasuk dalam Go perpustakaan standar. Lihat https://golang.org/s/sqldrivers untuk daftar driver pihak ketiga.
Open hanya dapat memvalidasi argumen tanpa membuat koneksi ke database. Untuk memverifikasi bahwa nama sumber data valid, disebut Ping.
DB yang kembali aman untuk digunakan bersamaan oleh beberapa goroutines dan memelihara pool sendiri koneksi idle. Dengan demikian, fungsi Terbuka harus disebut hanya sekali. Hal ini jarang diperlukan untuk menutup DB.

### Fungsi (*DB) Begin
func (db *DB) Begin() (*Tx, error)
Begin memulai transaksi. Tingkat isolasi tergantung pada driver.

### Fungsi (*DB) Close
func (db *DB) Close() error
Close menutup database, melepaskan sumber daya yang terbuka.
Sangat jarang untuk Menutup DB, sebagai penanganan DB dimaksudkan untuk menjadi berumur panjang dan dibagi antara banyak goroutines.

fungsi (*DB) Driver
func (db *DB) Driver() driver.Driver
Driver mengembalikan driver yang mendasari database.

### Fungsi (*DB) Exec
func (db *DB) Exec(query string, args ...interface{}) (Result, error)
Exec mengeksekusi query tanpa kembali ke baris manapun. Args adalah untuk setiap parameter placeholder dalam query.

### Fungsi (*DB) Ping
func (db *DB) Ping() error
Ping memverifikasi koneksi ke database yang masih hidup, membuat sambungan jika perlu.

### Fungsi (*DB) Prepare
func (db *DB) Prepare(query string) (*Stmt, error)
Prepare menciptakan sebuah pernyataan untuk query kemudian atau eksekusi. Beberapa pertanyaan atau eksekusi dapat dijalankan secara bersamaan dari pernyataan returned. Pemanggil harus memanggil pernyataan method Close ketika pernyataan tidak lagi diperlukan.

### Fungsi (*DB) Query
func (db *DB) Query(query string, args ...interface{}) (*Rows, error)
Query eksekusi yang mengembalikan baris, biasanya SELECT. Args untuk setiap parameter placeholder dalam query.

### Fungsi (*DB) queryRow
func (db *DB) QueryRow(query string, args ...interface{}) *Row
QueryRow mengeksekusi sebuah query yang diperkirakan akan kembali paling banyak pada satu baris. QueryRow selalu mengembalikan nilai non-nol. Kesalahan ditangguhkan sampai metode Scan Row ini dipanggil.

### Fungsi (*DB) SetConnMaxLifetime
func (db *DB) SetConnMaxLifetime(d time.Duration)
SetConnMaxLifetime menetapkan jumlah maksimum waktu sambungan dapat digunakan kembali.
Koneksi mungkin akan ditutup dengan malas sebelum digunakan kembali. 
Jika d <= 0, koneksi digunakan kembali selamanya.

### Fungsi (*DB) SetMaxIdleConns
func (db *DB) SetMaxIdleConns(n int)
SetMaxIdleConns menetapkan jumlah maksimum koneksi di kolam koneksi idle.
Jika MaxOpenConns lebih besar dari 0 tetapi kurang dari MaxIdleConns baru maka MaxIdleConns baru akan berkurang sesuai dengan batas MaxOpenConns
Jika n <= 0, tidak ada koneksi idle dipertahankan.

### Fungsi (*DB) SetMaxOpenConns
func (db *DB) SetMaxOpenConns(n int)
SetMaxOpenConns menetapkan jumlah maksimum koneksi terbuka ke database.
Jika MaxIdleConns lebih besar dari 0 dan MaxOpenConns baru kurang dari MaxIdleConns, maka MaxIdleConns akan berkurang sesuai dengan batas MaxOpenConns baru
Jika n <= 0, maka tidak ada batasan pada jumlah koneksi terbuka. default adalah 0 (unlimited).

### Fungsi (*DB) Stats
func (db *DB) Stats() DBStats
Stats mengembalikan statistik database.

### Tipe DBStats
type DBStats struct {
        // OpenConnections adalah jumlah koneksi terbuka ke database.
        OpenConnections int
}
DBStats berisi statistik database.

### Tipe NullBool
type NullBool struct {
        Bool  bool
        Valid bool // valid adalah benar jika Bool tidak NULL
}
NullBool merupakan bool yang mungkin nol. NullBool mengimplementasikan antarmuka Scanner sehingga dapat digunakan sebagai tujuan scan, mirip dengan NullString.

### Fungsi (*NullBool) Scan
func (n *NullBool) Scan(value interface{}) error
Scan mengimplementasikan antarmuka Scanner. 

### Fungsi (NullBool) Value
func (n NullBool) Value() (driver.Value, error)
Value mengimplementasikan interface driver Valuer.

### Tipe NullFloat64
type NullFloat64 struct {
        Float64 float64
        Valid   bool  // valid adalah benar jika Float64 tidak NULL
     }
NullFloat64 merupakan float64 yang mungkin nol. NullFloat64 mengimplementasikan antarmuka Scanner sehingga dapat digunakan sebagai tujuan scan, mirip dengan NullString.

### Fungsi (*NullFloat64) Scan
func (n *NullFloat64) Scan(value interface{}) error
Scan mengimplementasikan antarmuka Scanner.

### Fungsi (NullFloat64) Value
func (n NullFloat64) Value() (driver.Value, error)
Value mengimplementasikan interface driver Valuer.

### Tipe NullInt64
type NullInt64 struct {
        Int64 int64
        Valid bool // valid adalah benar jika Int64 tidak NULL
    }
NullInt64 merupakan int64 yang mungkin nol. NullInt64 mengimplementasikan antarmuka Scanner sehingga dapat digunakan sebagai tujuan scan, mirip dengan NullString.

### Fungsi (*NullInt64) Scan
func (n *NullInt64) Scan(value interface{}) error
Scan mengimplementasikan antarmuka Scanner.

### Fungsi (NullInt64) Value
func (n NullInt64) Value() (driver.Value, error)
value mengimplementasikan interface driver Valuer.

### Tipe NullString
type NullString struct {
        String string
        Valid  bool // valid adalah benar jika String tidak NULL
     }
NullString merupakan string yang mungkin nol. NullString mengimplementasikan antarmuka Scanner sehingga dapat digunakan sebagai tujuan scan:
var s NullString
err := db.QueryRow("SELECT name FROM foo WHERE id=?", id).Scan(&s)
...
if s.Valid {
   // menggunakan s.String
} else {
   // nilai NULL
}

### Fungsi (*NullString) Scan
func (ns *NullString) Scan(value interface{}) error
Scan mengimplementasikan antarmuka Scanner.

### Fungsi (NullString) Value
func (ns NullString) Value() (driver.Value, error)
Value mengimplementasikan interface driver valuer.

### Tipe RawBytes
type RawBytes []byte
RawBytes adalah sepotong byte yang memegang referensi ke memori yang dimiliki oleh database itu sendiri. Setelah Scan menjadi RawBytes, slice ini hanya berlaku sampai panggilan Next, Scan, atau Close.

### Tipe Result
type Result interface {
        // LastInsertId mengembalikan integer yang dihasilkan oleh database
        // dalam menanggapi perintah. Biasanya ini akan terjadi dari
        // "Auto increment" kolom saat memasukkan baris baru. Tidak semua
        // database mendukung fitur ini, dan sintaks seperti
        // pernyataan dapat bervariasi.
        LastInsertId() (int64, error)
        // RowsAffected mengembalikan jumlah baris yang dipengaruhi oleh
        // update, insert, atau delete. Tidak setiap basisdata atau driver database
        // dapat mendukung ini.
        RowsAffected() (int64, error)
}
Sebuah Result merangkum perintah SQL dieksekusi.

### Tipe Row
type Row struct {
        // Mengandung filter atau field unexported
}
Row adalah hasil dari panggilan queryRow untuk memilih satu baris.

### Fungsi (*Row) Scan
func (r *Row) Scan(dest ...interface{}) error
salinan scan kolom dari baris yang cocok ke dalam nilai-nilai yang ditunjuk oleh dest. Lihat dokumentasi pada Rows.Scan untuk rincian. Jika lebih dari satu baris sesuai permintaan, Scan menggunakan baris pertama dan membuang sisanya. Jika tidak ada baris sesuai permintaan Scan kembali ErrNoRows.

### Tipe Rows
type Rows struct {
        // Mengandung filter atau field unexported
}
Rows adalah hasil dari query. Kursor yang dimulai sebelum baris pertama dari hasil set. Gunakan Next untuk memajukan melalui baris:
rows, err := db.Query("SELECT ...")
...
defer rows.Close()
for rows.Next() {
    var id int
    var name string
    err = rows.Scan(&id, &name)
    ...
}
err = rows.Err() // get any error encountered during iteration
...

### Fungsi (*Rows) Close
func (rs *Rows) Close() error
Close menutup Baris, mencegah penghitungan lebih lanjut. Jika Berikutnya mengembalikan false, Rows ditutup secara otomatis dan akan cukup untuk memeriksa hasil Err. Close adalah idempoten dan tidak mempengaruhi hasil Err.

### Fungsi (*Rows) Columns
func (rs *Rows) Columns() ([]string, error)
Columns mengembalikan nama kolom. Kolom mengembalikan sebuah kesalahan jika baris ditutup, atau jika baris dari queryRow dan ada kesalahan akan ditangguhkan.

### Fungsi (*Rows) Err
func (rs *Rows) Err() error
Err mengembalikan kesalahan, jika ada yang ditemui selama iterasi. Err dapat dipanggil setelah Close eksplisit atau implisit.

### Fungsi (*Rows) Next
func (rs *Rows) Next() bool
Next mempersiapkan baris hasil selanjutnya untuk dibaca dengan method Scan. Ia mengembalikan true jika sukses, atau false jika tidak ada hasil baris berikutnya atau kesalahan yang terjadi saat mempersiapkan itu. Err harus dikonsultasikan untuk membedakan antara dua kasus.
Setiap pemanggilan untuk Scan, bahkan yang pertama, harus didahului dengan panggilan untuk Next.

### Fungsi (*Rows) Scan
func (rs *Rows) Scan(dest ...interface{}) error
Salinan scan kolom pada baris saat ini ke dalam nilai-nilai yang ditunjuk oleh dest. Jumlah nilai di dest harus sama dengan jumlah kolom di Rows.
Scan mengkonversi kolom yang membaca dari database ke jenis Go umum dan jenis khusus yang disediakan oleh paket sql:
*string
*[]byte
*int, *int8, *int16, *int32, *int64
*uint, *uint8, *uint16, *uint32, *uint64
*bool
*float32, *float64
*interface{}
*RawBytes
any type implementing Scanner (see Scanner docs)
Dalam kasus yang paling sederhana, jika jenis value dari kolom sumber adalah integer, bool atau string jenis T dan dest adalah tipe *T, Scan hanya memberikan nilai melalui pointer.
Scan juga mengubah antara string dan tipe numerik, selama tidak ada informasi akan hilang. Sementara Scan stringifies semua nomor dipindai dari kolom database yang numerik ke *string, scan menjadi tipe numerik diperiksa untuk overflow. Misalnya, float64 dengan nilai 300 atau string dengan nilai "300" dapat di Scan ke uint16, tapi tidak ke uint8, meskipun float64 (255) atau "255" dapat di scan ke uint8 a. Satu pengecualian adalah bahwa scan dari beberapa nomor float64 ke string mungkin kehilangan informasi ketika stringifying. Secara umum, memindai floating point kolom ke *float64.
Jika argumen dest telah dengan tipe *[]byte, scan akan menyimpan dalam argumen salinan data yang sesuai. Salinan dimiliki oleh pemanggil dan dapat dimodifikasi dan diselenggarakan tanpa batas. Salinan dapat dihindari dengan menggunakan argumen tipe * RawBytes instead.
Jika argumen memiliki tipe *interface{}, salinan nilai scan yang diberikan oleh driver yang mendasari tanpa konversi. Ketika memindai dari nilai jenis []byte ke *interface{}, salinan slice dibuat dan pemanggil memiliki hasilnya.
Values tipe time.Time dapat dipindai ke nilai-nilai dari tipe *time.Time, *interface{}, *string, atau *[]byte. Ketika mengkonversi ke dua terakhir, time.Format3339Nano digunakan.
Values tipe bool dapat di scan ke dalam jenis *bool, *interface{}, *string, *[]byte, atau *RawBytes.
Untuk scanning ke *bool, sumber mungkin benar, salah, 1, 0, atau input string yang parseable oleh strconv.ParseBool.

### Tipe Scanner
type Scanner interface {
        // scan memberikan nilai dari database driver.
        //
        // Nilai src akan menjadi salah satu dari jenis berikut:
        //
        // int64
        // float64
        // bool
        // [] Byte
        // string yang
        // time.Time
        // Nol - untuk nilai-nilai NULL
        //
        // Kesalahan harus dikembalikan jika nilai tidak dapat disimpan
        // tanpa kehilangan informasi.
        Scan(src interface{}) error
}
Scanner adalah sebuah antarmuka yang digunakan oleh Scan.

### Tipe Stmt
type Stmt struct {
        // Mengandung filter atau field unexported
}
Stmt adalah prepared statement. Sebuah Stmt aman untuk digunakan bersamaan oleh beberapa goroutines.

### Fungsi (*Stmt) Close
func (s *Stmt) Close() error
Close menutup pernyataan itu.

### Fungsi (*Stmt) Exec
func (s *Stmt) Exec(args ...interface{}) (Result, error)
Exec mengeksekusi prepared statement dengan argumen yang diberikan dan mengembalikan Result summarizing yang merupakan efek dari pernyataan.

### Fungsi (*Stmt) Query
func (s *Stmt) Query(args ...interface{}) (*Rows, error)
Permintaan mengeksekusi pernyataan prepared query dengan argumen yang diberikan dan mengembalikan hasil query sebagai *Rows.

### Fungsi (*Stmt) QueryRow
func (s *Stmt) QueryRow(args ...interface{}) *Row
QueryRow mengeksekusi pernyataan prepared query dengan argumen yang diberikan. Jika kesalahan terjadi selama eksekusi pernyataan, kesalahan yang akan dikembalikan oleh pemanggil untuk Scan pada *Row kembali, yang selalu non-nol. Jika query memilih tidak ada baris, yang *Row maka scan akan kembali ErrNoRows. Jika tidak, *Row scan akan memindai baris yang dipilih pertama dan membuang sisanya.
Contoh penggunaan:
var name string
err := nameByUseridStmt.QueryRow(id).Scan(&name)

### Tipe Tx
type Tx struct {
        // Mengandung filter atau field unexported
}
Tx adalah transaksi database yang ada didalam proses.
Sebuah transaksi harus diakhiri dengan panggilan untuk Commit atau Rollback.
Setelah panggilan ke Komit atau Rollback, semua operasi pada transaksi gagal dengan ErrTxDone.
Laporan dipersiapkan untuk transaksi dengan memanggil transaksi Prepare atau Stmt metode ditutup oleh panggilan untuk commit atau Rollback.

### Fungsi (*Tx) Commit
func (tx *Tx) Commit() error
Commit melakukan transaksi.

### Fungsi (*Tx) Exec
func (tx *Tx) Exec(query string, args ...interface{}) (Result, error)
Exec mengeksekusi query yang tidak mengembalikan baris. Sebagai contoh: INSERT dan UPDATE.

### Fungsi (*Tx) Prepare
func (tx *Tx) Prepare(query string) (*Stmt, error)
Prepared menciptakan sebuah pernyataan untuk digunakan dalam transaksi.
Pernyataan itu kembali beroperasi dalam transaksi dan tidak lagi dapat digunakan sekali transaksi telah dilakukan atau digulung kembali.
Untuk menggunakan pernyataan prepared ada pada transaksi ini, lihat Tx.Stmt.

### Fungsi (*Tx) Query
func (tx *Tx) Query(query string, args ...interface{}) (*Rows, error)
Permintaan mengeksekusi query yang mengembalikan baris, biasanya SELECT.

### Fungsi (*Tx) QueryRow
func (tx *Tx) QueryRow(query string, args ...interface{}) *Row
QueryRow mengeksekusi query yang diperkirakan akan kembali paling banyak satu baris. QueryRow selalu mengembalikan nilai non-nol. Kesalahan ditangguhkan sampai metode Pindai Row ini dipanggil.

### Fungsi (*Tx) Rollback
func (tx *Tx) Rollback() error
Rollback membatalkan transaksi.

### Fungsi (*Tx) Stmt
func (tx *Tx) Stmt(stmt *Stmt) *Stmt
Stmt mengembalikan sebuah pernyataan khusus transaksi dari pernyataan yang ada.
Contoh:
updateMoney, err := db.Prepare("UPDATE balance SET money=money+? WHERE id=?")
...
tx, err := db.Begin()
...
res, err := tx.Stmt(updateMoney).Exec(123.45, 98293203)
Pernyataan itu kembali beroperasi dalam transaksi dan tidak lagi dapat digunakan sekali transaksi telah dilakukan atau digulung kembali.

Driver : Paket driver mendefinisikan interface yang harus dilaksanakan oleh database driver seperti yang digunakan oleh paket sql.


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

### Tipe Conn
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

### Tipe Driver
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

### Tipe Execer
type Execer interface {
        Exec(query string, args []Value) (Result, error)
}
Execer adalah sebuah antarmuka opsional yang dapat diimplementasikan oleh Conn.
Jika sebuah Conn tidak mengimplementasikan Execer, paket sql ini DB. Exec akan mempersiapkan query pertama, mengeksekusi pernyataan, dan kemudian tutup pernyataan tersebut.
Exec dapat kembali ke ErrSkip.

### Tipe NotNull
type NotNull struct {
        Converter ValueConverter
}
NotNull adalah tipe yang mengimplementasikan ValueConverter oleh pelarangan nilai nol tetapi sebaliknya mendelegasikan ke ValueConverter lain.

### Fungsi (NotNull) ConvertValue
func (n NotNull) ConvertValue(v interface{}) (Value, error)

### Tipe Null
type Null struct {
        Converter ValueConverter
}
Null adalah tipe yang mengimplementasikan ValueConverter dengan memungkinkan nilai-nilai nol tetapi sebaliknya mendelegasikan ke ValueConverter lain.

### Fungsi (Null) ConvertValue
func (n Null) ConvertValue(v interface{}) (Value, error)

### Tipe Queryer
type Queryer interface {
        Query(query string, args []Value) (Rows, error)
}
Queryer adalah sebuah antarmuka opsional yang dapat diimplementasikan oleh Conn.
Jika Conn tidak mengimplementasikan Queryer, paket sql ini DB. Query akan mempersiapkan query pertama, mengeksekusi pernyataan, dan kemudian menutup pernyataan tersebut.
Query dapat kembali ErrSkip.

### Tipe Result
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

### Tipe Rows
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

### Tipe RowsAffected
type RowsAffected int64
RowsAffected mengimplementasikan Result untuk INSERT atau UPDATE operasi yang bermutasi ke sejumlah baris.
func (RowsAffected) lastInsertId
func (RowsAffected) LastInsertId() (int64, error)
func (RowsAffected) RowsAffected
func (v RowsAffected) RowsAffected() (int64, error)

### Tipe Stmt
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

### Tipe Tx
type Tx interface {
        Commit() error
        Rollback() error
}
Tx adalah sebuah transaksi.

### Tipe Value
type Value interface{}
Value adalah nilai yang driver harus mampu menangani. Itu baik nol atau turunan dari salah satu tipe:
int64
float64
bool
[]byte
string [*] di manapun kecuali dari Rows. Next.
time.Time

### Tipe ValueConverter
type ValueConverter interface {
        // ConvertValue converts a value to a driver Value.
        ConvertValue(v interface{}) (Value, error)
}
ValueConverter adalah antarmuka yang menyediakan metode ConvertValue.
Berbagai implementasi dari ValueConverter disediakan oleh paket driver untuk menyediakan implementasi yang konsisten dari konversi antara driver. ValueConverters memiliki beberapa kegunaan:
* Mengkonversi dari tipe Value yang disediakan oleh paket sql  menjadi tipe kolom tertentu pada tabel database dan memastikan itu  cocok, seperti memastikan suatu int64 tertentu sesuai dalam kolom uint16 tabel.
* Mengkonversi nilai seperti yang diberikan dari database ke dalam salah satu  tipe driver Value.
* Dengan paket sql, untuk mengkonversi dari tipe driver Value tipe pengguna di scan.

### Tipe Valuer
type Valuer interface {
        // Value returns a driver Value.
        Value() (Value, error)
}
Valuer adalah antarmuka menyediakan method Value.
Tipe mengimplementasikan antarmuka Valuer mampu mengkonversi diri untuk Value driver.

## Go dan SQL: PostgreSQL



## Go dan NoSQL: RethinkDB




