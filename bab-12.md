# Akses Basis Data

## Sekilas Tentang Basis Data



### Basis Data SQL



### Basis Data NoSQL


### Basis Data NewSQL



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



## Go dan NoSQL: RethinkDB




