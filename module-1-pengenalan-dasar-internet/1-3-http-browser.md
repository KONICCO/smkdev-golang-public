<div align="center">
   <h1>Module 1 Lesson 3 - HTTP & Browser</h1>
   Author: <b>Dheo Fuady Samuel Pandohan Terampil Gultom</b>
   <br>
   Last Updated: <b>August 4, 2023</b>
</div>
<br><br><br>
<h1 id="detail-pembelajaran">Detail Pembelajaran</h1>
<table>
   <thead>
      <tr>
         <th>Daftar Pembelajaran</th>
         <th>Halaman</th>
      </tr>
   </thead>
   <tbody>
      <tr>
         <td>1.3.1 HTTP vs HTTPS</td>
         <td><a href="#131-http-vs-https">🔽</a></td>
      </tr>
      <tr>
         <td>1.3.2 SSL/TLS</td>
         <td><a href="#132-ssltls">🔽</a></td>
      </tr>
      <tr>
         <td>1.3.3 Fungsi Utama</td>
         <td><a href="#133-fungsi-utama">🔽</a></td>
      </tr>
      <tr>
         <td>1.3.4 Struktur Browser</td>
         <td><a href="#134-struktur-browser">🔽</a></td>
      </tr>
   </tbody>
</table>
<p><br><br></p>
<h1 id="durasi-pembelajaran">Durasi Pembelajaran</h1>
<p><strong>45 menit</strong></p>
<br>
<h2 id="131-http-vs-https">1.3.1 HTTP vs HTTPS</h2>
<p>HTTP (Hypertext Transfer Protocol) dan HTTPS (Hypertext Transfer Protocol Secure) adalah protokol yang digunakan untuk mengirimkan data melalui internet. Keduanya berperan penting dalam menghubungkan pengguna dengan situs web dan memastikan keamanan komunikasi.</p>
<p>Mari kita bayangkan internet sebagai jaringan besar yang menghubungkan banyak komputer di seluruh dunia. Setiap komputer di internet memiliki alamat yang unik, seperti alamat rumah kita sendiri. Ketika kita ingin mengakses situs web, kita menggunakan protokol HTTP atau HTTPS untuk mengirim permintaan dari komputer kita ke server yang menyimpan situs web tersebut.</p>
<div align= "center">
   <img src="https://www.hostinger.co.id/tutorial/wp-content/uploads/sites/11/2022/06/perbedaan-http-dan-https-1024x735.webp">
   <br>
   <figcaption>Perbedaan HTTP dengan HTTPS</figcaption>
</div>
<br>
<p>HTTP adalah protokol yang digunakan secara umum. Ia mengatur cara komunikasi antara browser (seperti Chrome, Firefox, atau Safari) yang kita gunakan dan server situs web. Ketika kita memasukkan URL (Uniform Resource Locator) ke dalam browser, contohnya <code>http://www.example.com</code>, browser kita akan mengirim permintaan HTTP ke server yang menyimpan situs web &quot;example.com&quot;. Permintaan ini berisi informasi seperti tipe permintaan (GET, POST, PUT, DELETE), path atau alamat file yang diminta, dan informasi tambahan seperti cookie.</p>
<p>Saat server menerima permintaan HTTP, ia akan mengirimkan kembali respons yang berisi informasi yang diminta. Respons ini bisa berisi halaman web yang diminta, gambar, file teks, atau data lainnya. Respons ini juga dapat mencakup kode status HTTP yang memberikan informasi tentang apakah permintaan berhasil (kode 200), tidak ditemukan (kode 404), atau ada masalah lainnya.</p>
<p>Namun, meskipun HTTP sangat umum digunakan, ia memiliki kelemahan dari segi keamanan. Karena data yang dikirimkan melalui protokol HTTP tidak dienkripsi, data tersebut dapat dengan mudah diintersepsi oleh pihak yang tidak sah. Misalnya, jika kita mengisi formulir pribadi yang tidak dienkripsi dengan HTTP, data yang kita masukkan bisa diketahui oleh orang lain.</p>
<p>Inilah mengapa HTTPS hadir untuk mengatasi kelemahan tersebut. HTTPS menggunakan <strong>SSL (Secure Sockets Layer)</strong> atau <strong>TLS (Transport Layer Security)</strong> untuk mengenkripsi data yang dikirimkan antara browser dan server. Enkripsi ini menjaga kerahasiaan dan integritas data, sehingga pihak yang tidak sah sulit untuk mengakses atau mengubah informasi yang dikirimkan.</p>
<p>Saat kita mengakses situs web yang menggunakan HTTPS, URL akan dimulai dengan <code>https://</code>, dan browser akan menunjukkan ikon gembok atau indikator lain yang menandakan koneksi aman. Data yang dikirimkan melalui HTTPS tetap aman bahkan jika diintersepsi, karena hanya penerima yang memiliki kunci enkripsi yang benar dapat membacanya.</p>
<p>Jadi, dalam istilah sederhana, HTTP adalah protokol dasar yang digunakan untuk mengirimkan permintaan dan menerima respons antara browser dan server, sementara HTTPS adalah varian yang aman dari HTTP yang menggunakan enkripsi untuk melindungi data yang dikirimkan. Penting untuk memahami perbedaan antara keduanya dan memastikan kita menggunakan HTTPS ketika mengakses situs web yang membutuhkan keamanan, seperti situs perbankan atau toko online, untuk melindungi privasi dan keamanan kita.</p>
<h2 id="132-ssltls">1.3.2 SSL/TLS</h2>
<p>SSL dan TLS adalah protokol yang bertanggung jawab untuk menyediakan enkripsi dan keamanan dalam komunikasi antara klien (seperti browser) dan server di internet. Mereka membentuk lapisan keamanan di atas protokol HTTP dan mengamankan data yang dikirimkan antara pengguna dan situs web.</p>
<h3 id="ssl-secure-sockets-layer">SSL (Secure Sockets Layer)</h3>
<p>SSL adalah protokol keamanan yang pertama kali dikembangkan oleh Netscape pada tahun 1995. Awalnya, SSL digunakan untuk melindungi transaksi finansial online, seperti transaksi kartu kredit. Namun, seiring berjalannya waktu, SSL juga digunakan untuk mengamankan komunikasi dalam berbagai konteks, termasuk e-mail, VPN, dan lainnya.</p>
<p>Proses SSL dimulai dengan negosiasi kunci enkripsi antara klien dan server. Setelah kunci enkripsi disepakati, data yang dikirimkan antara klien dan server akan dienkripsi menggunakan algoritma enkripsi simetris. SSL juga memvalidasi identitas server melalui penggunaan sertifikat digital yang dikeluarkan oleh otoritas sertifikat terpercaya. Namun, karena ada kerentanan keamanan yang ditemukan dalam versi SSL yang lebih tua, protokol ini secara bertahap digantikan oleh TLS.</p>
<h3 id="tls-transport-layer-security">TLS (Transport Layer Security)</h3>
<p>TLS merupakan penerus SSL dan digunakan secara luas saat ini. Protokol ini memiliki tujuan yang sama dengan SSL, yaitu menyediakan keamanan dan enkripsi dalam komunikasi internet. Seperti SSL, TLS juga melibatkan proses negosiasi kunci enkripsi dan validasi identitas server melalui sertifikat digital. TLS mendukung berbagai algoritma enkripsi yang lebih kuat, termasuk AES (Advanced Encryption Standard) dan RSA (Rivest-Shamir-Adleman), yang menjadikannya lebih aman daripada SSL.</p>
<p>TLS juga telah mengalami beberapa pembaruan dan versi yang berbeda, seperti TLS 1.0, TLS 1.1, TLS 1.2, dan TLS 1.3. Setiap versi TLS menyediakan perbaikan keamanan dan peningkatan kinerja dibandingkan dengan versi sebelumnya.</p>
<h3 id="perbedaan-antara-ssl-dan-tls">Perbedaan antara SSL dan TLS</h3>
<p>Perbedaan utama antara SSL dan TLS adalah versi dan evolusi protokol. TLS dirancang sebagai pengganti SSL dan telah mengalami pembaruan dan perbaikan berkelanjutan untuk mengatasi kelemahan yang ditemukan dalam SSL.</p>
<p>Secara konseptual, keduanya berfungsi sama, yaitu menyediakan enkripsi dan keamanan dalam komunikasi internet. Namun, karena perkembangan yang lebih lanjut dan peningkatan keamanan, disarankan untuk menggunakan TLS daripada SSL.</p>
<p>Dalam prakteknya, ketika kita berbicara tentang HTTPS (Hypertext Transfer Protocol Secure), protokol yang digunakan adalah TLS, bukan SSL. Sebagai pengguna, kita sering mendengar istilah HTTPS karena itu adalah implementasi keamanan yang umum digunakan dalam browsing web modern.
   Jadi, secara ringkas, SSL dan TLS adalah protokol keamanan yang digunakan untuk menyediakan enkripsi dan keamanan dalam komunikasi internet. TLS adalah pengganti SSL yang lebih aman dan cenderung digunakan secara luas saat ini.
</p>
<h2 id="133-fungsi-utama-browser">1.3.3 Fungsi Utama Browser</h2>
<p>Browser, atau peramban web, adalah perangkat lunak yang kita gunakan setiap hari untuk mengakses dan menjelajahi internet. Tanpa browser, sulit bagi kita untuk mengakses situs web, mencari informasi, berkomunikasi secara online, atau menjalankan aplikasi web. Browser memiliki peran krusial dalam membawa dunia internet ke layar komputer atau perangkat seluler kita.</p>
<p>Namun, browser bukan hanya alat untuk menampilkan halaman web. Mereka memiliki berbagai fitur dan fungsionalitas yang memungkinkan kita untuk berpindah antara situs web, menyimpan tautan yang penting, menjaga keamanan dan privasi, menyesuaikan pengaturan, dan banyak lagi. Pemahaman tentang fungsi utama browser akan membantu kita memaksimalkan pengalaman menjelajah web kita dan menggunakan internet dengan lebih efisien.</p>
<p>Dalam penjelasan berikutnya, kita akan melihat lebih dekat pada fungsi-fungsi utama browser, termasuk rendering dan tampilan halaman web, navigasi web, bookmark dan manajemen tautan, pencarian web, keamanan dan privasi, ekstensi dan add-on, serta pengaturan kustomisasi. Dengan memahami peran yang dimainkan oleh setiap fungsi ini, kita akan memiliki pemahaman yang lebih baik tentang bagaimana browser membantu kita menjelajahi dan berinteraksi dengan internet. Berikut adalah penjelasan yang lebih terperinci dan terstruktur tentang fungsi utama browser:</p>
<h3 id="rendering-dan-tampilan-halaman-web">Rendering dan Tampilan Halaman Web</h3>
<p>Browser bertanggung jawab untuk mengambil halaman web dari server dan menampilkannya di jendela browser. Proses ini melibatkan rendering, di mana browser menerjemahkan kode HTML (Hypertext Markup Language), CSS (Cascading Style Sheets), dan JavaScript yang ada pada halaman web untuk menghasilkan tampilan visual yang dapat kita lihat. Browser memastikan agar elemen-elemen halaman web, seperti teks, gambar, video, dan elemen interaktif lainnya, ditampilkan dengan benar dan sesuai dengan desain yang dimaksudkan oleh pembuat halaman web.</p>
<h3 id="navigasi-web">Navigasi Web</h3>
<p>Browser menyediakan antarmuka yang memungkinkan kita untuk berpindah antara halaman web. Kita dapat menggunakan tombol &quot;Back&quot; untuk kembali ke halaman sebelumnya, atau tombol &quot;Forward&quot; untuk kembali ke halaman yang telah kita kunjungi sebelumnya. Browser juga menyediakan bilah alamat tempat kita dapat memasukkan URL atau melakukan pencarian untuk menemukan halaman web yang ingin kita kunjungi. Selain itu, browser juga dapat memiliki fitur tab yang memungkinkan kita membuka beberapa halaman web secara bersamaan dalam jendela browser yang sama.</p>
<h3 id="bookmark-dan-manajemen-tautan">Bookmark dan Manajemen Tautan</h3>
<p>Browser memungkinkan kita untuk menyimpan bookmark atau tautan ke halaman web yang sering kita kunjungi. Ini memudahkan kita untuk mengakses situs web favorit tanpa harus memasukkan alamat URL secara manual setiap kali. Kita dapat menyimpan bookmark dalam folder yang terorganisir dan mengatur bookmark sesuai dengan kategori atau preferensi kita. Browser juga menyediakan fitur manajemen tautan, yang memungkinkan kita untuk mengelola dan menghapus tautan yang sudah disimpan.</p>
<h3 id="pencarian-web">Pencarian Web</h3>
<p>Browser seringkali memiliki kotak pencarian terintegrasi yang memungkinkan kita untuk melakukan pencarian web langsung dari bilah alamat. Ketika kita memasukkan kata kunci atau pertanyaan di kotak pencarian, browser akan mengirimkan permintaan ke mesin pencari seperti Google atau Bing. Hasil pencarian kemudian ditampilkan di halaman hasil pencarian, di mana kita dapat memilih tautan yang sesuai dengan kebutuhan kita.</p>
<h3 id="keamanan-dan-privasi">Keamanan dan Privasi</h3>
<p>Browser memiliki peran penting dalam menjaga keamanan dan privasi pengguna saat menjelajah internet. Mereka dilengkapi dengan berbagai fitur keamanan, seperti pengingat password, yang membantu kita mengelola kata sandi dan menghindari penggunaan kata sandi yang lemah. Browser juga dapat mengenali dan memperingatkan pengguna tentang situs web berbahaya atau mencurigakan yang mungkin mencoba melakukan penipuan atau mencuri informasi pribadi. Beberapa browser juga menyediakan fitur blokir iklan untuk mengurangi gangguan iklan yang tidak diinginkan. Selain itu, browser memungkinkan kita untuk mengatur preferensi privasi, seperti pengelolaan cookie, penghapusan riwayat penelusuran, dan kontrol atas data pengguna yang disimpan oleh situs web.</p>
<h3 id="ekstensi-dan-add-on">Ekstensi dan Add-on</h3>
<p>Banyak browser modern menyediakan ekosistem ekstensi atau add-on yang memperluas fungsionalitas browser. Ekstensi adalah program kecil yang dapat diinstal ke dalam browser untuk menambahkan fitur khusus, seperti pemblokir iklan, pengelola kata sandi, alat terjemahan, pengedit gambar, dan banyak lagi. Ini memungkinkan pengguna untuk menyesuaikan browser mereka sesuai dengan kebutuhan dan preferensi mereka.</p>
<h3 id="kinerja-dan-pengaturan">Kinerja dan Pengaturan</h3>
<p>Browser juga menawarkan berbagai pengaturan dan opsi kustomisasi yang mempengaruhi kinerja dan perilaku browser. Pengguna dapat mengatur preferensi pencarian, mengubah tampilan dan tema browser, mengelola bahasa yang didukung, mengatur pengaturan tab, dan mengelola penyimpanan cache untuk meningkatkan kecepatan dan efisiensi penjelajahan.</p>
<p>Keseluruhan, browser berperan penting dalam menghubungkan kita dengan internet dan memberikan pengalaman penjelajahan web yang nyaman. Dengan memastikan tampilan yang baik, kemudahan navigasi, manajemen tautan, pencarian web, keamanan, privasi, ekstensi, dan pengaturan yang fleksibel, browser memungkinkan kita untuk menjelajahi web dengan mudah dan efektif sesuai dengan kebutuhan kita.</p>
<h2 id="134-struktur-browser">1.3.4 Struktur Browser</h2>
<p>Browser terdiri dari beberapa komponen yang bekerja bersama-sama untuk menghadirkan pengalaman penjelajahan web yang lancar. Berikut adalah beberapa komponen utama yang membentuk struktur browser:
   (Tampilkan ilustrasi browser) 
</p>
<h3 id="user-interface-antarmuka-pengguna">User Interface (Antarmuka Pengguna)</h3>
<p>Antarmuka pengguna adalah tampilan visual browser yang kita lihat saat menggunakannya. Ini termasuk bilah alamat, tombol navigasi (seperti Back dan Forward), bilah tab, tombol refresh, dan fitur-fitur lainnya yang memungkinkan kita berinteraksi dengan browser. User interface juga dapat mencakup elemen desain dan tema yang dapat disesuaikan agar sesuai dengan preferensi pengguna.</p>
<h3 id="rendering-engine-mesin-rendering">Rendering Engine (Mesin Rendering)</h3>
<p>Rendering engine adalah komponen inti browser yang bertanggung jawab untuk mengambil kode HTML, CSS, dan JavaScript dari halaman web yang diminta dan mengubahnya menjadi tampilan visual yang kita lihat. Mesin rendering menerjemahkan instruksi-instruksi dalam kode tersebut untuk menentukan bagaimana halaman web akan ditampilkan, diatur, dan berinteraksi dengan pengguna. Contoh mesin rendering populer adalah WebKit (digunakan oleh Safari) dan Blink (digunakan oleh Chrome dan Opera).</p>
<h3 id="networking-jaringan">Networking (Jaringan)</h3>
<p>Komponen jaringan browser mengelola proses pengiriman dan penerimaan data melalui jaringan internet. Ini termasuk mengirimkan permintaan HTTP atau HTTPS ke server yang menyimpan halaman web, menerima respons dari server, dan mengelola koneksi untuk mentransfer data dengan aman dan efisien.</p>
<h3 id="javascript-engine">JavaScript Engine</h3>
<p>JavaScript engine adalah komponen yang mengeksekusi dan menjalankan kode JavaScript yang ada pada halaman web. Browser menggunakan JavaScript engine untuk menginterpretasikan dan menjalankan skrip JavaScript, yang memberikan kemampuan interaktif pada halaman web. Contoh JavaScript engine yang terkenal adalah V8 (digunakan oleh Chrome) dan SpiderMonkey (digunakan oleh Firefox).</p>
<h3 id="cookie-manager">Cookie Manager</h3>
<p>Cookie manager adalah komponen yang mengelola cookie yang disimpan pada komputer pengguna oleh situs web yang dikunjungi. Cookie adalah file teks kecil yang digunakan untuk menyimpan informasi tentang preferensi, sesi login, atau jejak aktivitas pengguna. Cookie manager memastikan bahwa cookie disimpan, dihapus, atau dikirim ke situs web yang sesuai sesuai dengan pengaturan privasi pengguna.</p>
<h3 id="browser-extensions-ekstensi-browser">Browser Extensions (Ekstensi Browser)</h3>
<p>Browser extensions, atau ekstensi browser, adalah program tambahan yang dapat diinstal ke dalam browser untuk menambahkan fitur dan fungsionalitas tambahan. Ekstensi dapat memberikan fitur seperti blokir iklan, alat terjemahan, manajer kata sandi, integrasi media sosial, dan banyak lagi. Mereka memperluas kemampuan dasar browser dan memberikan penggunaan yang lebih kustom.</p>
<p>Keseluruhan, struktur browser terdiri dari antarmuka pengguna yang terlihat, mesin rendering yang mengubah kode web menjadi tampilan visual, komponen jaringan yang mengelola pengiriman dan penerimaan data, mesin JavaScript yang menjalankan kode skrip, cookie manager yang mengelola cookie, dan ekstensi browser yang memperluas fungsionalitas dasar. Semua komponen ini bekerja bersama-sama untuk memberikan pengalaman penjelajahan web yang efisien, interaktif, dan aman kepada pengguna.</p>