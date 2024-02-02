<div align="center">
   <h1>Module 1 Lesson 1 - Protokol dan Standar Internet</h1>
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
         <td>1.1.1 Pengenalan</td>
         <td><a href="#111-pengenalan">🔽</a></td>
      </tr>
      <tr>
         <td>1.1.2 Konsep dan Terminologi</td>
         <td><a href="#112-konsep-dan-terminologi">🔽</a></td>
      </tr>
      <tr>
         <td>1.1.3 IPv4 vs IPv6</td>
         <td><a href="#113-ipv4-vs-ipv6">🔽</a></td>
      </tr>
   </tbody>
</table>
<p><br><br></p>
<h1 id="durasi-pembelajaran">Durasi Pembelajaran</h1>
<p><strong>45 menit</strong></p>
<br>
<h2 id="111-pengenalan">1.1.1 Pengenalan</h2>
<p>Bersiaplah untuk menjelajahi keajaiban yang luar biasa sebagai programmer! Bayangkan saja, internet seperti jaring laba-laba yang sangat rumit yang menjalin seluruh dunia dalam satu koneksi yang tak terlihat. Ia adalah jembatan ajaib yang menghubungkan kita dengan pengetahuan, hiburan, dan kesempatan tak terbatas dengan hanya satu sentuhan jari.</p>
<div align= "center">
   <img src="https://1.bp.blogspot.com/-586iidfLH6I/YLhKcj0RT7I/AAAAAAAAKsI/Lcl71t0_LVgTxXJNwPsxA3Rh9fhDTldFQCLcBGAsYHQ/s1054/Subnet-mask-diagram.webp" alt="cara-internet-bekerja">
   <br>
   <figcaption>Cara Internet Bekerja</figcaption>
</div>
<br>
<p>Mari kita mulai dengan gambaran besarnya. Bayangkan internet sebagai lanskap yang luas, dipenuhi dengan ratusan juta rumah dan bangunan yang saling terhubung oleh jalan raya dan jembatan. Setiap rumah adalah komputer atau perangkat yang terhubung ke internet, sementara jalan raya adalah jalur komunikasi yang memungkinkan informasi bergerak dengan cepat.</p>
<p>Sekarang, mari kita bayangkan kamu adalah seorang petualang internet yang ingin mencari informasi baru. Pertama, kamu membuka browser webmu dan memasukkan alamat situs web yang ingin kamu kunjungi. Pikirkan situs web ini sebagai tujuan akhir petualanganmu, misalnya sebuah toko baju di pusat perbelanjaan yang besar.</p>
<p>Tapi bagaimana kamu bisa sampai ke tujuanmu? Nah, di sinilah penyedia layanan internet (ISP) berperan. ISP adalah seperti pemandu yang membantu kamu menavigasi jalan raya internet. Ketika kamu memasukkan alamat situs web, ISP akan mengirimkan permintaanmu ke server <strong>DNS (Domain Name System)</strong> yang bertugas menerjemahkan alamat situs web menjadi alamat IP yang dapat dimengerti oleh komputer.</p>
<p>Setelah kamu memiliki alamat IP yang benar, perjalananmu dimulai. Informasi yang kamu cari, seperti halaman web atau gambar, terdiri dari paket-paket data kecil. Pikirkan paket-paket data ini sebagai truk-truk pengangkut yang membawa barang-barang pentingmu. Paket-paket data ini akan meluncur di jalan raya internet, melewati jaringan-jaringan yang berbeda, termasuk serat optik bawah laut dan kabel yang menghubungkan benua-benua.</p>
<p>Namun, jangan khawatir! Di sepanjang perjalanan, ada &quot;pengatur lalu lintas&quot; yang disebut router. Router ini bertugas mengarahkan paket-paket datamu ke tujuan yang benar, memastikan mereka tiba dengan aman dan tepat waktu. Mereka seperti pandu yang membimbing kamu melalui jalan-jalan buntu dan jembatan yang rusak.</p>
<p>Akhirnya, setelah perjalanan yang luar biasa ini, paket-paket datamu mencapai tujuan akhir: server situs web yang kamu cari. Server ini seperti pintu gerbang yang membuka akses ke informasi yang kamu inginkan. Ketika server menerima paket-paket datamu, ia menyusunnya kembali menjadi satu kesatuan yang lengkap, seperti merakit puzzle.</p>
<p>Dan voila! Halaman web yang kamu inginkan terbentang di layar komputermu. Kamu telah berhasil menemukan harta karun internet! Semua ini terjadi dalam hitungan detik, berkat kerja keras dan kerjasama tak terlihat dari jaringan global yang luar biasa.</p>
<p>Jadi, internet adalah dunia ajaib yang terjalin oleh jaringan kompleks dan misterius. Ia menghubungkan kita dengan informasi, hiburan, dan peluang di seluruh dunia. Jadi, bersiaplah untuk terus menjelajahi dan mengeksplorasi keindahan dan keajaiban internet yang tak terbatas!</p>
<h2 id="112-konsep-dan-terminologi">1.1.2 Konsep dan Terminologi</h2>
<p>Bayangkan internet sebagai sebuah kota yang tak terbatas dengan jalan-jalan yang tak terhitung jumlahnya dan bangunan-bangunan yang penuh dengan kejutan. Setiap bangunan adalah sebuah website yang menawarkan pengalaman yang berbeda. Kamu bisa berjalan-jalan dari satu website ke website lainnya, seperti berkeliling di kota yang penuh warna-warni.</p>
<p>Nah, sekarang bayangkan alamat website sebagai nomor rumah yang unik. Ketika kamu ingin mengunjungi suatu website, kamu cukup mengetikkan alamatnya di browsermu. Seperti memasukkan nomor rumah yang tepat ke dalam GPS, browsermu akan membawamu langsung ke halaman yang kamu cari. Kamu bisa menjelajahi berbagai jenis website, mulai dari toko online yang menawarkan barang-barang menarik hingga situs berita yang memberikan informasi terkini.</p>
<p>Tapi bagaimana informasi bisa bergerak dari satu tempat ke tempat lainnya dengan cepat? Di sinilah jaringan internet berperan. Jaringan ini seperti jalan-jalan raya yang menghubungkan seluruh kota. Informasi dalam bentuk paket-paket data berjalan melalui jaringan ini dengan kecepatan cahaya, seperti mobil-mobil yang melaju di jalan raya.</p>
<p>Untuk memastikan paket-paket datamu mencapai tujuan dengan tepat, ada &quot;pengatur lalu lintas&quot; yang disebut router. Mereka bertugas mengarahkan paket-paket data ke tempat yang tepat, seperti petunjuk lalu lintas yang membimbingmu saat berkendara di jalan raya. Dengan bantuan router, informasi yang kamu cari bisa sampai ke komputermu dengan cepat dan akurat.</p>
<h3 id="mari-kita-kupas-konsep-dan-terminologi-yang-mendasari-ini">Mari kita kupas konsep dan terminologi yang mendasari ini.</h3>
<ol>
   <li><strong>Jaringan</strong><br>Pada dasarnya, internet adalah jaringan komputer yang saling terhubung di seluruh dunia. Komputer-komputer ini bisa menjadi server (pusat informasi) atau klien (komputer individu). Mereka berkomunikasi melalui kabel, serat optik, dan teknologi nirkabel seperti Wi-Fi. Analoginya seperti jalan-jalan raya yang menghubungkan berbagai tempat di dunia, memungkinkan aliran informasi yang lancar.</li>
   <li><strong>Protokol</strong><br>Untuk memungkinkan komunikasi yang efisien di jaringan ini, digunakan protokol. Protokol adalah aturan yang ditetapkan untuk pertukaran data. Salah satu protokol utama yang digunakan adalah TCP/IP (Transmission Control Protocol/Internet Protocol). Protokol ini memecah data menjadi paket-paket kecil yang dapat dikirim melalui jaringan dan memastikan mereka tiba dengan aman dan tepat waktu.</li>
   <li><strong>ISP (Internet Service Provider)</strong><br>Penyedia layanan internet yang menghubungkan pengguna ke internet. Mereka menyediakan akses ke jaringan global melalui teknologi seperti DSL, kabel, atau serat optik. ISP juga memberikan alamat IP (Internet Protocol) kepada pengguna, yang berfungsi sebagai identifikasi unik untuk setiap perangkat yang terhubung ke internet.</li>
   <li><strong>DNS (Domain Name System)</strong><br>Sistem yang menerjemahkan alamat web yang mudah diingat (seperti <a href="http://www.example.com">www.example.com</a>) menjadi alamat IP yang dapat dimengerti oleh komputer. DNS bekerja seperti buku telepon raksasa yang menghubungkan nama domain dengan alamat IP, memungkinkan kita mengakses situs web hanya dengan memasukkan alamat domain yang mudah diingat.</li>
   <li><strong>Browser</strong><br>Browser web adalah alat yang kita gunakan untuk menjelajahi internet. Mereka mengambil informasi dari server web dan menampilkannya dalam bentuk halaman web yang dapat kita lihat dan interaksi. Beberapa browser populer termasuk Google Chrome, Mozilla Firefox, dan Safari.</li>
   <li><strong>HTTP dan HTTPS</strong><br>Protokol HTTP (Hypertext Transfer Protocol) adalah aturan yang digunakan untuk mentransfer data melalui internet. Ketika kita mengunjungi sebuah situs web, alamat yang kita ketikkan diawali dengan &quot;http://&quot; (contoh: <code>http://www.example.com</code>). Versi yang lebih aman dari protokol ini adalah HTTPS (HTTP Secure), yang menggunakan enkripsi SSL/TLS untuk melindungi data yang ditransfer antara pengguna dan situs web. Alamat HTTPS diawali dengan &quot;https://&quot; (contoh: <code>https://www.example.com</code>).</li>
   <li><strong>HTML dan URL</strong><br>HTML (Hypertext Markup Language) adalah bahasa pemrograman yang digunakan untuk membuat halaman web. Dalam HTML, kita mengatur tata letak, teks, gambar, dan elemen lainnya. URL (Uniform Resource Locator) adalah alamat yang digunakan untuk mengidentifikasi sumber daya di internet, seperti halaman web, gambar, atau dokumen. URL juga berfungsi sebagai jalan menuju sumber daya tersebut.</li>
</ol>
<h2 id="113-ipv4-vs-ipv6">1.1.3 IPv4 vs IPv6</h2>
<p>Mari kita masuki duel antara IPv4 dan IPv6 dalam pertarungan konsep dan kepraktisan di dunia internet! Dalam satu sudut, kita memiliki IPv4, sang veteran yang telah bertahan lama dalam pertempuran ini. Di sudut lain, hadirlah IPv6, sang pewaris yang siap menggantikan posisi sang veteran. Ayo kita simak!</p>
<h3 id="ipv4-internet-protocol-version-4">IPv4 (Internet Protocol version 4)</h3>
<p>IPv4 adalah versi protokol internet yang telah digunakan sejak awal internet. Bayangkan IPv4 sebagai alamat rumah dalam kota internet. Alamat ini terdiri dari serangkaian angka (misalnya, <code>192.168.0.1</code>) yang unik untuk setiap perangkat yang terhubung ke internet. Namun, dengan cepat kita menyadari bahwa jumlah alamat IPv4 terbatas. Dalam dunia yang semakin terhubung, ketersediaan alamat IPv4 semakin menipis.</p>
<h3 id="ipv6-internet-protocol-version-6">IPv6 (Internet Protocol version 6)</h3>
<p>Ini dia pewaris yang siap mengambil alih. IPv6 adalah generasi berikutnya dari protokol internet. Konsepnya cukup sederhana, namun konsekuensinya luar biasa. IPv6 menggunakan sistem penomoran yang jauh lebih besar daripada IPv4, menghasilkan alamat yang terdiri dari angka dan huruf yang lebih panjang (misalnya, <code>2001:0db8:85a3:0000:0000:8a2e:0370:7334</code>). Jumlah alamat yang tersedia dengan IPv6 hampir tak terbatas. Bayangkan seperti memiliki kota yang tak terbatas, di mana setiap perangkat bisa mendapatkan alamat unik tanpa batasan.</p>
<p>Namun, peralihan dari IPv4 ke IPv6 tidaklah mudah. Sang veteran IPv4 telah membangun infrastruktur internet yang luas selama bertahun-tahun, dan sebagian besar perangkat masih menggunakan alamat IPv4. Untuk menghadapinya, terdapat mekanisme seperti <strong>NAT (Network Address Translation)</strong> yang mengizinkan beberapa perangkat menggunakan satu alamat IPv4. Namun, NAT memiliki kelemahan dan dapat mempengaruhi performa serta fungsionalitas jaringan.</p>
<p>IPv6 menawarkan beberapa keuntungan dibandingkan IPv4. Selain dari jumlah alamat yang tak terbatas, IPv6 juga memiliki fitur-fitur keamanan yang lebih baik dan dukungan yang kuat untuk kualitas layanan yang ditingkatkan. Namun, peralihan ke IPv6 memerlukan upaya kolaboratif dari penyedia layanan internet, produsen perangkat, dan pengguna. Hal ini melibatkan pengaturan jaringan yang kompleks dan perangkat yang mendukung IPv6.</p>
<p>Jadi, pertarungan antara IPv4 dan IPv6 adalah tentang melanjutkan evolusi internet. IPv4 telah memberikan fondasi kuat untuk pertumbuhan internet, tetapi batasannya semakin terasa. IPv6 hadir sebagai solusi masa depan yang mampu menampung pertumbuhan pesat perangkat yang terhubung. Melalui upaya dan adaptasi yang tepat, kita dapat menghadapi peralihan ini dan memastikan keberlanjutan dunia yang terhubung ini.</p>