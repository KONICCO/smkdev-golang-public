___

# Apa itu Variable ?

`Variabel` adalah salah satu konsep fundamental dalam pemrograman, termasuk di bahasa pemrograman Go (Golang).

Dalam Go, `Variabel` digunakan untuk `menyimpan` dan `mengelola` data yang dapat berubah selama eksekusi program. jika anda masih belum paham Mari kita analogikan lebih dalam mengenai `variabel` di Go.

###### **analogi**:
bayangkan Anda memiliki sebuah `kotak ajaib` yang dapat Anda `beri nama` dan dapat `menyimpan apa saja`. Kotak ajaib ini disebut “Variabel” dalam bahasa pemrograman Go.

##### Tetapi, apa sebenarnya `variabel` di dalam bahasa pemrogramman itu ?
`Variabel` seperti `wadah` atau `tempat` kita menyimpan data. Sebagai contoh, kita bisa menyimpan data berbentuk integer, string, atau bahkan potongan kode program di dalamnya.


___

> ###  jika anda sudah paham `Variable secara definisi`, Sekarang, saatnya kita belajar bagaimana cara `mendeklarasikan Variabel`.
>> di Go Ada berbagai macam cara untuk mendeklarasikan `Variabel`.
>> Mari kita lihat satu per satu.


### • deklarasi menggunakan `var`

**format** ⇒ `var variableName dataType`

_contoh penggunaan_:
```go
var name string

var age int
```

_contoh kasus_:
```go
package main

import "fmt"

func main() {
    var name string
    var age int

    fmt.Print("Enter Your Name: ")
    fmt.Scanln(&name)
    fmt.Print("Enter Your Age: ")
    fmt.Scanln(&age)

    fmt.Println("Your Name is: ", name)
    fmt.Println("Your Age is: ", age)

}
```


___

### • deklarasi menggunakan `var` dan lansung mengisinya dengan `value`

**format** ⇒ `var variableName dataType = InitialValue`

_contoh penggunaan_:
```go
var name string = "Eko Edy"

var age int = 20
```

_contoh kasus_:
```go
package main

import "fmt"

func main() {
    var name string = "Eko Edy"
    var age int = 20

    fmt.Println("Your Name is: ", name)
    fmt.Println("Your Age is: ", age)

}
```


___

### • deklarasi menggunakan `:=` atau biasa di sebut `short declaration`

**format** ⇒ `variableName := initialValue`

_contoh penggunaan_:
```go
name := "Eko Edy"

age := 20
```

_contoh kasus_:
```go
package main

import "fmt"

func main() {
    name := "Eko Edy"
    age := 20

    fmt.Println("Your Name is: ", name)
    fmt.Println("Your Age is: ", age)

}
```


___

### • deklarasi beberapa variable atau biasa di sebut `Multiple declarations` 

**format** ⤵️ 

```go
var (
    variableName1 dataType1
    variableName2 dataType2
    ..bisa di tambahkan lagi.. 
)
```

_contoh penggunaan_:
```go
var (
 name string

 age int
)
```

_contoh kasus_:
```go
package main

import "fmt"

func main() {
    var (
    name string

    age int
    )

    fmt.Print("Enter Your Name: ")
    fmt.Scanln(&name)
    fmt.Print("Enter Your Age: ")
    fmt.Scanln(&age)

    fmt.Println("Your Name is: ", name)
    fmt.Println("Your Age is: ", age)

}
```


___

### • deklarasi beberapa variable atau biasa di sebut `Multiple declarations` dan lansung mengisinya dengan `value`

**format** ⤵️ 

```go
var (
    variableName1 dataType1 = initialValue1
    variableName2 dataType2 = initialValue2
    ..bisa di tambahkan lagi.. 
)
```

_contoh penggunaan_:
```go
var (
 name string = "Eko Edy"

 age int = 20
)
```

_contoh kasus_:
```go
package main

import "fmt"

func main() {
    var (
    name string = "Eko Edy"

    age int = 20
    )

    fmt.Println("Your Name is: ", name)
    fmt.Println("Your Age is: ", age)

}
```


___

> ### jika anda sudah paham cara `Mendeklarasikan Variable`, di Go. Sekarang, saatnya kita belajar bagaimana cara `mendeklarasikan Konstanta` di Go.

>> #### Note:
>> - `Konstanta` memiliki nilai tetap yang tidak dapat berubah setelah pertama kali nilainya ditetapkan.
>> - `Konstanta` harus memiliki nilai saat dideklarasikan.


### • deklarasi menggunakan `const`

**format** ⇒ `const constantName typeData = constantValue`

_contoh penggunaan_:
```go
const phi float64 = 3.14
```

_contoh kasus_:
```go
package main

import "fmt"

// Deklarasi konstanta
const phi float64 = 3.14

func main() {
    var radius float64

    // Meminta input dari pengguna untuk radius
    fmt.Print("Masukkan radius lingkaran: ")
    fmt.Scanln(&radius)

    // Menghitung keliling lingkaran
    circumference := 2 * phi * radius
    fmt.Printf("Keliling lingkaran dengan radius %.2f adalah %.2f\n", radius, circumference)

    // Menghitung luas lingkaran
    area := phi * radius * radius
    fmt.Printf("Luas lingkaran dengan radius %.2f adalah %.2f\n", radius, area)
}

```


___

### • deklarasi beberapa konstanta atau biasa di sebut `Multiple constants`

**format** ⤵️ 

```go
const (
    constanName1 dataType1 = constanValue1
    constantName2 dataType2 = constantValue2
    ..bisa di tambahkan lagi.. 
)
```

_contoh penggunaan_:
```go
const (
    celsiusToFahrenheitRatio float64 = 9.0 / 5.0
    celsiusToKelvinOffset    float64 = 273.15
    ..bisa di tambahkan lagi.. 
)
```

_contoh kasus_:
```go
package main

import "fmt"

// Deklarasi Multi konstanta untuk konversi suhu
const (
    celsiusToFahrenheitRatio float64 = 9.0 / 5.0
    celsiusToKelvinOffset    float64 = 273.15
)

func main() {
    var temperatureInCelsius float64

    // Meminta input dari pengguna untuk suhu dalam Celsius
    fmt.Print("Masukkan suhu dalam derajat Celsius: ")
    fmt.Scanln(&temperatureInCelsius)

    // Menghitung konversi suhu dari Celsius ke Fahrenheit dan Kelvin
    temperatureInFahrenheit := temperatureInCelsius*celsiusToFahrenheitRatio + 32
    temperatureInKelvin := temperatureInCelsius + celsiusToKelvinOffset

    // Menampilkan hasil konversi
    fmt.Printf("%.2f derajat Celsius setara dengan %.2f derajat Fahrenheit\n", temperatureInCelsius, temperatureInFahrenheit)
    fmt.Printf("%.2f derajat Celsius setara dengan %.2f Kelvin\n", temperatureInCelsius, temperatureInKelvin)
}

```


___

> ### jika anda sudah paham cara Mendeklarasikan `Variable` dan juga `Konstanta`, di Go. Sekarang, saatnya kita belajar bagaimana cara `Penamaan` "Variable" dan "Konstanta" di Go.
### penamaan variabel yang umum digunakan dalam bahasa pemrograman Go:

- Hindari nama yang ambigu atau sulit dipahami.
- Pilih nama yang menggambarkan tujuan variabel sehingga orang lain dapat dengan mudah memahaminya saat membaca kode.
- Gunakan gaya penulisan camelCase. 
    - **contoh** ⇒ `firstName` `lastName` `fullName`
- Gunakan singkatan yang umum
    - **penjelasan** ⇒ Jika ada singkatan yang umum digunakan dalam pemrograman, seperti “URL” atau “ID”, Anda dapat menggunakannya untuk membuat kode Anda lebih mudah dipahami oleh pengembang lain.
- Gunakan bahasa Inggris
    - **penjelasan** ⇒ Dalam menulis variabel, yang terbaik adalah menggunakan bahasa Inggris untuk memastikan konsistensi dalam kode dan mempermudah komunikasi antar programmer.
- Hindari penamaan yang berlebihan
    - **penjelasan** ⇒ Jika konteksnya sudah jelas, hindari menambahkan informasi yang berlebihan pada nama variabel. Gunakan nama yang singkat dan bermakna untuk membuat kode lebih bersih dan mudah dibaca.

___