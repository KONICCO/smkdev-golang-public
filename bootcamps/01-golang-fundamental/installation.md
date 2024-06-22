# Sejarah Golang

Setiap bahasa pemrograman punya keunggulan dan fungsinya masing-masing. Ada yang mengandalkan kesederhanaan, interaktivitas, kecepatan, maupun skala.Sebagai contoh, bahasa pemrograman C dan C++ unggul dari segi kecepatan, skala, dan kehandalan. Sementara itu, bahasa pemrograman Python dipilih karena sederhana. Adapun bahasa pemrograman Java yang dianggap mudah dipahami dan sangat scalable. Namun, terkadang software developer membutuhkan bahasa pemrograman yang bisa unggul dalam semua aspek tersebut.

Oleh karena itu, Google bersama Ken Thompson, Robert Griesemer, dan Rob Pike akhirnya meluncurkan bahasa pemrograman Go atau Golang pada 2009, seperti ditulis Computer Hope. Pasalnya, Golang dianggap unggul dari empat segi, yakni kecepatan/efisiensi, keandalan, skala, dan kesederhanaan. Namun, sebenarnya, apa itu Golang?

> Golang adalah bahasa pemrograman yang diketik secara statis dan menghasilkan kode biner mesin yang dikompilasi. Menariknya, bahasa pemrograman yang satu ini bersifat open source. Golang dihimpun dan ditulis menggunakan bahasa pemrograman C. Maka, tak heran jika banyak orang menganggap Golang adalah bahasa pemrograman C di abad ke-21. Selain itu, tak heran pula jika banyak orang jadi tertarik untuk belajar Golang. Kamu bisa menggunakan Golang untuk membuat berbagai program, seperti website, aplikasi, dan sebagainya.

# Keunggulan vs Kekurangan Golang

## Kelebihan Golang

Ada sejumlah keunggulan yang membuat bahasa pemrograman ini menjadi menarik di mata banyak perusahaan, terutama startup.Berikut diantaranya.

1. Mudah dipelajari
   Dibandingkan pesaingnya, gaya sintaks yang dimiliki Golang lebih kecil sehingga lebih mudah dipelajari. Kamu pun tidak perlu menggunakan banyak waktu untuk mencari istilah-istilah yang sulit dimengerti. Kemudahan ini bahkan juga dapat dirasakan oleh programmer yang menggunakan gaya sintaks berbeda sekalipun.
2. Lebih cepat
   Golang dikompilasi ke dalam kode mesin sehingga dapat melampaui bahasa pemrograman lain yang bekerja dengan virtual runtime. Program-program di dalamnya juga bekerja cepat, dengan API yang dapat mengkompilasi dalam hitungan detik. Ini menjadikan Golang sebagai bahasa pemrograman yang lebih cepat.
3. Memperbaiki kekurangan dari bahasa pemrograman yang sudah ada
   Golang dilengkapi dengan sejumlah fungsi mutakhir sehingga dapat mengatasi masalah pada bahasa pemrograman lainnya, seperti:
   - kurangnya dukungan komputasi paralel
   - kurangnya dukungan multicore
   - pengelolaan ketergantungan yang buruk
   - sistem tipe yang rumit
   - pengelolaan memori yang rumit

## Kekurangan Golang

Seperti bahasa pemrograman lainnya, Golang juga memiliki sejumlah kekurangan. Berikut diantaranya :

1. Interface terlalu implisit
   Interface merupakan batas bersama ketika dua atau lebih komponen dalam komputer bertukar informasi. Golang memang dilengkapi dengan interface, tapi sifatnya yang implisit dapat membuat pengguna kesulitan untuk membedakan isi struct (komposit data). Kamu hanya bisa mengetahuinya begitu program telah dikompilasi.
2. Tidak bisa menggunakan fungsi yang sama untuk koleksi data berbeda
   Berbeda dengan Java, penggunaan kode pada Golang tidak bisa dilakukan secara berulang. Walaupun fungsi-fungsi yang dimilikinya terbilang canggih, kode-kode yang dapat digunakan pada satu jenis koleksi data ternyata tidak dapat digunakan untuk kelompok data yang lain.

# Instalasi Golang

Sebelum Belajar bahasa Golang yang dilakukan pertama kali yaitu menginstall Golang di Komputer yang kita pakai. Golang dapat di install di semua sistem Operasi baik Linux, Windows maupun MacOS. Ketika anda menginstall Golang saat ini sudah mencapai versi go1.21.

## Windows

Berikut cara bagaimana proses instalasi Go pada Windows:

1. Kunjungi situs resmi Go di [https://golang.org/dl/](https://golang.org/dl/) untuk mengunduh file instalasi Go yang sesuai dengan arsitektur sistem operasi Windows kamu (32-bit atau 64-bit).
2. Setelah unduhan selesai, buka file instalasi yang diunduh dan ikuti petunjuk instalasi yang ditampilkan di layar.
3. Pilih direktori instalasi Go. Direktori default biasanya adalah **"C:\Go"**, tetapi kamu dapat memilih direktori yang sesuai dengan preferensi kamu.
4. Pada langkah "Custom Setup", pastikan opsi **"Add Go to your PATH environment variable"** dicentang. Hal ini akan memungkinkan kamu untuk menjalankan perintah Go dari baris perintah tanpa harus menentukan jalur lengkap ke direktori instalasi Go.
5. Lanjutkan dengan instalasi dan tunggu hingga selesai.
6. Setelah instalasi selesai, buka command prompt (cmd.exe) dan ketik perintah
   ```go
   go version
   ```
   untuk memverifikasi instalasi Go. Jika instalasi berhasil, kamu akan melihat versi Go yang terpasang.

## **Linux**

Untuk menginstall Go di Linux dapat menggunakan 2 cara, yang pertama kita dapat menginstall secara default dengan versi tertentu, yang kedua dapat menggunakan GVM(Golang Version Manager). Dengan menggunakan GVM maka kita dapat menentukan Versi Golang yang ingin kita gunakan sebagai pengaturan default Go.Sebagai media belajar sebaiknya untuk pertama dapat menginstall Golang default, secara umum.

Berikut cara bagaimana proses instalasi Go pada Linux:

1. Buka terminal pada Linux.
2. Kunjungi situs resmi Go di [https://golang.org/dl/](https://golang.org/dl/) untuk mengunduh file instalasi Go yang sesuai dengan arsitektur sistem operasi Linux kamu (misalnya, amd64 untuk arsitektur 64-bit).
3. Buka terminal dan navigasikan ke direktori tempat file instalasi Go diunduh.
4. Ekstrak file instalasi Go dengan menjalankan perintah berikut (menggantikan **"go1.x.x.linux-amd64.tar.gz"** dengan nama file yang diunduh):

   ```bash
   tar -C /usr/local -xzf go1.x.x.linux-amd64.tar.gz
   ```

5. Setelah ekstraksi selesai, tambahkan direktori Go ke PATH dengan menambahkan baris berikut pada file `/.profile` atau `/.bashrc` (sesuaikan dengan shell yang kamu gunakan):

   ```bash
   export PATH=$PATH:/usr/local/go/bin
   export GOPATH=$HOME/go
   export PATH=$PATH:$GOPATH/bin:$GOROOT/bin
   ```

6. Muat ulang file profil dengan menjalankan perintah:

   ```bash
   source ~/.profile
   # atau
   source ~/.bashrc
   ```

7. Verifikasi instalasi Go dengan menjalankan perintah `go version` pada terminal. Jika instalasi berhasil, kamu akan melihat versi Go yang terpasang.

## **Mac OS**

1. Download file Installer
   Installer Golang pada Mac OS dengan ekstensi .pkg di [https://golang.org/dl/](https://golang.org/dl/). Silahkan sesuaikan dengan bit komputer yang anda gunakan.
2. Jalankan file Install
   Silahkan install installer go dengan ekstensi .pkg dan lakukan next saja jika terdapat keterangan install.
3. Export PATH dan GOPATH
   Seperti d Linux, di Mac OS juga harus mengekspor agar folder go dapat dikenali oleh sistem Operasi.
   ```bash
   export GOROOT=/usr/local/go
   export GOPATH=$HOME/go
   export PATH=$PATH:$GOPATH/bin:$GOROOT/bin
   ```
4. Periksa Versi Golang
   ```go
   go version
   ```

# **Code Editor**

Ada beberapa code editor yang populer dan sering digunakan oleh pengembang Go. Berikut ini beberapa contoh code editor yang dapat digunakan untuk pengembangan aplikasi dengan Go beserta tautan unduhnya:

1. Visual Studio Code (VS Code):  
   Link unduh: [https://code.visualstudio.com/download  
   ](https://code.visualstudio.com/download)
   Catatan: VS Code memiliki ekstensi resmi untuk Go yang menyediakan fitur-fitur seperti autocompletion, debugging, dan integrasi dengan alat-alat Go.
2. GoLand:  
   Link unduh: [https://www.jetbrains.com/go/download  
   ](https://www.jetbrains.com/go/download)
   Catatan: GoLand adalah IDE yang dikembangkan oleh JetBrains dan menawarkan fitur-fitur canggih untuk pengembangan aplikasi Go.
3. Sublime Text:  
   Link unduh: [https://www.sublimetext.com  
   ](https://www.sublimetext.com)
   Catatan: Sublime Text adalah editor teks yang populer dan dapat diperluas dengan paket Go yang tersedia melalui package manager Package Control.
4. Atom:  
   Link unduh: [https://atom.io  
   ](https://atom.io)
   Catatan: Atom adalah editor teks yang dapat disesuaikan dan memiliki paket Go yang memungkinkan pengembangan aplikasi Go yang efisien.
5. Visual Studio:  
   Link unduh: [https://visualstudio.microsoft.com/downloads/  
   ](https://visualstudio.microsoft.com/downloads/)
   Catatan: Visual Studio (versi Community) menyediakan dukungan untuk bahasa Go melalui ekstensi resmi.

Pastikan untuk memilih code editor yang sesuai dengan preferensi pribadi kamu dan yang menyediakan fitur-fitur yang kamu butuhkan untuk pengembangan aplikasi dengan Go.

# **Hello World !!**

Setelah kamu menyiapkan instalasi yang dibutuhkan untuk pengembangan Go, berikut ini penjelasan singkat tentang proses Hello World dalam Go:

1. Buka code editor pilihan kamu (misalnya, Visual Studio Code) dan buat file baru dengan ekstensi **.go**. Misalnya, beri nama file **hello.go**.
2. Di dalam file **hello.go**, ketikkan kode berikut:

```go
package main
import ("fmt")

func main() {
	fmt.Println("Hello, World!")
}
```

Kode ini akan mencetak pesan "Hello, World!" ke output.

3. Simpan file **hello.go**.
4. Buka command prompt atau terminal, lalu arahkan ke direktori tempat kamu menyimpan file **hello.go**.
5. Jalankan perintah berikut untuk menjalankan program Go:

   ```go
   go run hello.go
   ```

6. Outputnya akan menampilkan "Hello, World!".
   ```bash
   Hello, World!
   ```

Dalam program Hello World di atas, beberapa konsep penting dalam Go dapat ditemukan:

- Pernyataan `package main` menandakan bahwa kita berada dalam package utama yang akan menjadi titik masuk (entry point) program.
- Pernyataan `import "fmt"` mengimpor paket "fmt" yang berisi fungsi-fungsi dasar untuk input/output. Fungsi `fmt.Println()` digunakan untuk mencetak teks ke output.
- Fungsi `main()` adalah fungsi utama yang akan dieksekusi saat program dijalankan.
- `fmt.Println("Hello, World!")` digunakan untuk mencetak keluaran Hello, World! ke terminal.
