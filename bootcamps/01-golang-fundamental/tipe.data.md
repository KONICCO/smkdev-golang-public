___

# Apa itu Tipe Data?

`Tipe data` adalah konsep fundamental dalam pemrograman yang menentukan jenis nilai yang dapat disimpan dalam variabel, bagaimana nilai itu diinterpretasikan, dan operasi apa yang dapat dilakukan terhadap nilai tersebut. Dalam bahasa pemrograman Go (Golang), `tipe data` memainkan peran penting dalam mendefinisikan struktur dan perilaku pada data.

###### **Mengapa Tipe Data Di Go (Golang) Penting ?**

- **Keamanan**: Tipe data memastikan bahwa variabel hanya menyimpan nilai-nilai yang diizinkan, mencegah kesalahan dan bug.
- **Optimisasi**: Kompilator atau biasa di kenal dengan `compiler` dapat melakukan optimisasi yang lebih baik ketika tipe data diketahui, meningkatkan kinerja program.
- **Keterbacaan**: Tipe data membantu programmer memahami apa yang diharapkan dari variabel dan fungsi, meningkatkan keterbacaan dan maintainability kode.

###### **Tipe Data Di Go (Golang) Menggunakan `Static Data Type`**

`Static Data Type` berarti tipe data untuk setiap variabel ditentukan pada saat deklarasi dan tidak dapat berubah selama eksekusi program. Hal ini memastikan bahwa setiap variabel harus memiliki tipe data yang tetap dan diketahui sejak awal.

Contohnya, jika kita mendeklarasikan sebuah variabel x sebagai integer, maka x hanya dapat menyimpan nilai-nilai integer sepanjang eksekusi program.

```go
var x int

x = 10  // valid
x = "Hello"  // error: cannot use "Hello" (type untyped string) as type int
```

Hal ini berbeda dengan bahasa pemrograman yang menggunakan Dynamic Data Type, di mana tipe data variabel dapat berubah-ubah selama waktu eksekusi.


___

> ###  jika anda sudah paham `Definisi Tipe Data dan seberapa pentingnya Tipe Data`, Sekarang, saatnya kita belajar macam-macam Tipe Data pada Go (Golang)
>> di Go Ada berbagai macam Tipe Data
>> Mari kita lihat Tabel di bawah ini.

#### **signed integer**
| Tipe Data | Ukuran (bit) | Rentang Nilai                                              |
|-----------|--------------|------------------------------------------------------------|
| int8      | 8            | -128 `to` 127                                              |
| int16     | 16           | -32,768 `to` 32,767                                        |
| int32     | 32           | -2,147,483,648 `to` 2,147,483,647                          |
| int64     | 64           | -9,223,372,036,854,775,808 `to` 9,223,372,036,854,775,807  |

#### **unsigned integer**
| Tipe Data | Ukuran (bit) | Rentang Nilai                          |
|-----------|--------------|----------------------------------------|
| uint8     | 8            | 0 `to` 255                             |
| uint16    | 16           | 0 `to` 65,535                          |
| uint32    | 32           | 0 `to` 4,294,967,295                   |
| uint64    | 64           | 0 `to` 18,446,744,073,709,551,615      |

#### **alias integer**
| Tipe Data | Ukuran (bit)         | Rentang Nilai                          | alias untuk     |
|-----------|----------------------|----------------------------------------|-----------------|
| byte      | 8                    | 0 to 255                               | uint8           |
| rune      | 32                   | -2,147,483,648 to 2,147,483,647        | int32           |
| int       | Platform dependent   | Platform dependent                     | Minimal int32   |
| uint      | Platform dependent   | Platform dependent                     | Minimal uint32  |

#### **floating point**
| Tipe Data  | Ukuran (bit) | Rentang Nilai                                          |
|------------|--------------|--------------------------------------------------------|
| float32    | 32           | ±1.18E-38 to ±3.4E38                                   |
| float64    | 64           | ±2.23E-308 to ±1.8E308                                 |
| complex64  | 64           | complex numbers with float32 real and imaginary parts. |
| complex128 | 128          | complex numbers with float64 real and imaginary parts. |


#### **boolean**
| Tipe Data | Nilai                          |
|-----------|--------------------------------|
| bool      | true or false                  |

#### **string**
| Tipe Data | Nilai                                                                                                                 | Rentang Nilai                                                             |
|-----------|-----------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------|
| string    | Nilai data String di Go-Lang selalu diawali dengan karakter “ (petik dua) dan diakhiri dengan karakter “ (petik dua)  | Jumlah karakter di dalam String bisa string kosong sampai tidak terhingga |
