# Operator 

Setelah anda memahami penggunaan variable dan type data. Chapter kali ini akan membahas mengenai macam operator yang bisa digunakan di Go. Secara umum terdapat 3 kategori operator: aritmatika, perbandingan, dan logika.

# Aritmatika

Kita bisa menggunakan tipe data number dengan operasi matematika sebagai berikut:

| Operator / Syntax | Keterangan           |
|-------------------|----------------------|
| +                 | Penambahan           |
| -                 | Pengurangan          |
| *                 | Perkalian            |
| /                 | Pembagian            |
| %                 | Mendapatkan sisa pembagian |

## Contoh :

```go
func main() {
    fmt.Println(1 + 5)
    fmt.Println(2 - 1)
    fmt.Println(2 * 8)
    fmt.Println(4 / 2)
    fmt.Println(10 % 3) // sisa pembagian 10 dengan 3 adalah 1
}
```
# Augmented Assignment

Ini merupakan operasi yang digunakan untuk mempersingkat operasi matematika.

| Operator / Syntax | Operasi Matematika | Keterangan |
|-------------------|-------------------|------------|
| a += 10           | a = a + 10         | Penambahan |
| a -= 10           | a = a - 10         | Pengurangan|
| a *= 10           | a = a * 10         | Perkalian  |
| a /= 10           | a = a / 10         | Pembagian  |
| a %= 10           | a = a % 10         | Mendapatkan sisa pembagian |

## Contoh :

```go
func main() {
    var a = 10
    a += 10 // a = 10 + 10 = 20
    fmt.Println(a)
}
```

> go run main.go

Output:
```
20
```

# Unary Operator

Operasi ini juga digunakan untuk mempersingkat operasi matematika khusus.

| Operator / Syntax | Operasi Matematika | Keterangan              |
|-------------------|-------------------|-------------------------|
| a++               | a = a + 1         | increment (naik 1 angka) |
| a--               | a = a - 1         | decrement (turun 1 angka)|
| -a                | a = -a            | mengubah jadi negatif (negasi) |
| +a                | a = +a            | mengubah jadi positif    |

## Contoh :

```go
func main() {
    var a = 10
    a++ // a = 10 + 1 = 11

    fmt.Println(a)

    var b = 10
    fmt.Println(-b)
}
```

> go run main.go

Output:
```
11
-10
```

# Operator Perbandingan

Operator perbandingan digunakan untuk menentukan kebenaran suatu kondisi. Hasil dari operator ini berupa nilai boolean, yaitu `true` atau `false`.

## Tabel Operator Perbandingan

| Operator / Syntax | Keterangan                                   |
|-------|----------------------------------------------|
| `==`  | Apakah nilai kiri sama dengan nilai kanan?  |
| `!=`  | Apakah nilai kiri tidak sama dengan nilai kanan? |
| `<`   | Apakah nilai kiri lebih kecil daripada nilai kanan? |
| `<=`  | Apakah nilai kiri lebih kecil atau sama dengan nilai kanan? |
| `>`   | Apakah nilai kiri lebih besar dari nilai kanan? |
| `>=`  | Apakah nilai kiri lebih besar atau sama dengan nilai kanan? |

## Contoh :

```go
package main

import (
    "fmt"
)

func main() {
    var value = 1+1
    var isEqual = (value == 2)

    fmt.Printf("nilai %d (%t) \n", value, isEqual)
}
```
Pada kode di atas, terdapat statement operasi aritmatika yang hasilnya ditampung oleh variabel `value`. Selanjutnya, variabel tersebut dibandingkan dengan angka 2 untuk dicek apakah nilainya sama. Jika iya, maka hasilnya adalah `true`, jika tidak maka `false`.
# Operator Logika

Operator ini digunakan untuk mencari benar tidaknya kombinasi data bertipe bool (bisa berupa variabel bertipe bool, atau hasil dari operator perbandingan).

Beberapa operator logika standar yang bisa digunakan:

| Operator / Syntax | Keterangan         |
|-------|--------------------|
| &&    | kiri dan kanan     |
| \|\| | kiri atau kanan    |
| !     | negasi / nilai kebalikan |

## Contoh :

```go
var left = false
var right = true

var leftAndRight = left && right
fmt.Printf("left && right =(%t) \n", leftAndRight)

var leftOrRight = left || right
fmt.Printf("left || right =(%t) \n", leftOrRight)

var leftReverse = !left
fmt.Printf("!left =(%t) \n", leftReverse) 
```
Berikut penjelasan statemen operator logika pada kode di atas.

`leftAndRight` bernilai `false`, karena hasil dari `false` dan `true` adalah `false`.

`leftOrRight` bernilai `true`, karena hasil dari `false` atau `true` adalah `true`.

`leftReverse` bernilai `true`, karena `negasi (atau lawan dari) false` adalah `true`.

