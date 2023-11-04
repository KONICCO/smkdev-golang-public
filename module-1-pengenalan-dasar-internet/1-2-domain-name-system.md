<div align="center">
   <h1>Module 1 Lesson 2 - Domain Name System</h1>
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
         <td>1.2.1 Pengenalan DNS</td>
         <td><a href="#121-pengenalan-dns">🔽</a></td>
      </tr>
      <tr>
         <td>1.2.2 Bagaimana itu bekerja?</td>
         <td><a href="#122-bagaimana-itu-bekerja">🔽</a></td>
      </tr>
      <tr>
         <td>1.2.3 Struktur DNS</td>
         <td><a href="#123-struktur-dns">🔽</a></td>
      </tr>
      <tr>
         <td>1.2.4 Relasi DNS dan Back-End</td>
         <td><a href="#124-relasi-dns-dan-back-end">🔽</a></td>
      </tr>
   </tbody>
</table>
<p><br><br></p>
<h1 id="durasi-pembelajaran">Durasi Pembelajaran</h1>
<p><strong>45 menit</strong></p>
<br>
<h2 id="121-pengenalan-dns">1.2.1 Pengenalan DNS</h2>
<p>Sistem Nama Domain (DNS) adalah sistem yang digunakan untuk mengonversi alamat IP numerik menjadi nama domain yang lebih mudah diingat dan lebih manusiawi, serta sebaliknya. Ini adalah bagian kritis dari infrastruktur internet yang memungkinkan pengguna untuk mengakses situs web, mengirim email, dan melakukan berbagai tindakan lainnya dengan menggunakan nama domain yang lebih mudah diingat daripada alamat IP yang kompleks.</p>
<p>Berikut adalah beberapa komponen utama dari Pengenalan DNS:</p>
<ol>
   <li>
      <p><strong>Nama Domain</strong><br>Nama domain adalah alamat yang diberikan kepada sumber daya atau layanan di internet. Misalnya, &quot;<a href="http://www.contoh.com">www.contoh.com</a>&quot; atau &quot;mail.contoh.net&quot;. Nama domain biasanya terdiri dari dua bagian utama: nama tingkat atas (Top-Level Domain atau TLD) seperti .com, .org, .net, dan nama domain yang lebih spesifik seperti &quot;contoh&quot;.</p>
   </li>
   <li>
      <p><strong>Server DNS</strong><br>Server DNS adalah server khusus yang bertugas untuk menghubungkan nama domain dengan alamat IP. Ketika Kamu memasukkan alamat web di browser Kamu, server DNS akan mencari alamat IP yang sesuai dengan nama domain tersebut. Ada berbagai jenis server DNS, termasuk server DNS otoritatif (yang menyimpan informasi domain) dan server DNS resolver (yang meminta informasi dari server otoritatif).</p>
   </li>
   <li>
      <p><strong>Rekaman DNS</strong><br>Rekaman DNS adalah entri dalam basis data DNS yang mengaitkan informasi tertentu dengan suatu nama domain. Beberapa jenis rekaman DNS meliputi:</p>
      <ul>
         <li>Rekaman A: Menghubungkan alamat IP dengan nama domain.</li>
         <li>Rekaman CNAME: Menghubungkan satu nama domain dengan nama domain lain (alias).</li>
         <li>Rekaman MX: Menentukan server email yang bertanggung jawab untuk nama domain.</li>
         <li>Rekaman NS: Menunjukkan server DNS otoritatif untuk domain.</li>
         <li>Rekaman PTR: Digunakan untuk melakukan kebalikan DNS, yaitu mengonversi alamat IP menjadi nama domain.</li>
      </ul>
   </li>
   <li>
      <p><strong>Proses Resolusi DNS</strong><br>Proses resolusi DNS terjadi ketika Kamu memasukkan alamat domain di browser atau aplikasi. Langkah-langkah utama dalam proses ini adalah:</p>
      <ul>
         <li>Cache: Jika permintaan telah dilakukan sebelumnya, server DNS resolver dapat menggunakan cache untuk mengambil informasi yang tersimpan sebelumnya.</li>
         <li>Pertanyaan: Jika informasi tidak ada dalam cache, resolver akan mengirimkan permintaan ke server DNS otoritatif yang berkaitan dengan domain yang dicari.</li>
         <li>Resolusi: Server DNS otoritatif mengembalikan rekaman DNS yang sesuai, yang kemudian dikirimkan kembali ke resolver.</li>
         <li>Hasil akhir: Resolver menyimpan hasil dalam cache untuk penggunaan mendatang dan mengirimkan informasi IP ke aplikasi atau perangkat pengguna, memungkinkan akses ke layanan yang diminta.</li>
      </ul>
   </li>
</ol>
<p>DNS memainkan peran kunci dalam menghubungkan alamat IP dengan nama domain, menjadikannya salah satu fondasi penting dari pengalaman internet yang kita nikmati hari ini.</p>
<h2 id="122-bagaimana-itu-bekerja">1.2.2 Bagaimana itu bekerja?</h2>
<p>Ketika pengguna memasukkan nama domain ke dalam bilah alamat browser web, mereka akan dibawa ke situs yang ingin mereka kunjungi. Namun, tugas yang tampak instan ini sebenarnya melibatkan beberapa langkah yang dikenal sebagai proses pencarian DNS atau proses resolusi DNS.</p>
<p>Berikut contoh dari proses resolusi DNS yang biasanya terjadi untuk menggambarkan bagaimana DNS bekerja. Kamu ingin mengunjungi situs web Hostinger, jadi Kamu memasukkan nama domain hostinger.com ke dalam bilah alamat browser web. Apa yang Kamu lakukan di sini adalah mengirim permintaan DNS. Selanjutnya, komputer Kamu akan memeriksa apakah sudah menyimpan entri DNS dari domain yang Kamu masukkan secara lokal. Catatan DNS adalah alamat IP yang sesuai dengan nama domain lengkap.</p>
<p>Pertama, komputer Kamu akan mencari dalam berkas host dan cache-nya. Berkas host adalah berkas teks biasa yang memetakan nama host ke alamat IP dalam sistem operasi, sedangkan cache adalah data sementara yang disimpan oleh komponen perangkat keras atau perangkat lunak. Alamat IP yang cocok untuk layanan DNS umumnya ditemukan dalam cache browser Kamu atau cache penyedia layanan internet (ISP). Namun, jika tidak ada alamat IP yang cocok ditemukan dalam berkas host dan cache Kamu, langkah-langkah tambahan lainnya akan ditambahkan ke dalam proses resolusi DNS.</p>
<div align= "center">
   <img src="https://www.hostinger.com/tutorials/wp-content/uploads/sites/2/2023/01/how-does-dns-work-1536x884.webp">
   <br>
   <figcaption>Cara Domain Name System (DNS) Bekerja</figcaption>
</div>
<h3 id="berikut-adalah-penjelasan-lengkap-tentang-bagaimana-dns-bekerja">Berikut adalah penjelasan lengkap tentang bagaimana DNS bekerja</h3>
<ol>
   <li>
      <p><strong>Permintaan Resolusi</strong>: Ketika Kamu memasukkan alamat URL (Uniform Resource Locator) ke dalam peramban web Kamu (seperti <a href="http://www.contoh.com">www.contoh.com</a>), peramban akan mengirimkan permintaan resolusi DNS ke server DNS.</p>
   </li>
   <li>
      <p><strong>Cache Lokal</strong>: Peramban Kamu pertama-tama akan memeriksa cache lokalnya untuk melihat apakah resolusi DNS untuk alamat tersebut telah diambil sebelumnya. Jika ada, dan informasi masih valid, maka peramban akan menggunakan alamat IP yang telah di-cache.</p>
   </li>
   <li>
      <p><strong>Server DNS Tepat</strong>: Jika informasi tidak ada dalam cache atau sudah kadaluarsa, peramban akan mengirimkan permintaan ke server DNS yang ditetapkan oleh penyedia layanan internet (ISP) Kamu. Server DNS ini dikenal sebagai &quot;server DNS primer&quot;.</p>
   </li>
   <li>
      <p><strong>Rekursi atau Iterasi</strong>: Server DNS primer bisa saja menyelesaikan permintaan resolusi itu sendiri, atau ia bisa meneruskan permintaan ke server DNS lain dalam proses yang disebut rekursi atau iterasi. Dalam rekursi, server DNS primer akan berusaha mengambil informasi langsung dari server DNS lain hingga mendapatkan jawaban yang valid. Dalam iterasi, server DNS primer akan meminta informasi langkah demi langkah dari server DNS lain.</p>
   </li>
   <li>
      <p><strong>Pohon Nama Domain</strong>: Server DNS bekerja berdasarkan hierarki pohon nama domain. Jika server DNS primer tidak memiliki informasi yang diperlukan, ia akan melakukan permintaan ke server DNS yang lebih tinggi dalam hierarki, seperti server otoritatif untuk domain yang diminta.</p>
   </li>
   <li>
      <p><strong>Server DNS Otoritatif</strong>: Server DNS otoritatif adalah server yang menyimpan catatan resmi untuk nama domain tertentu. Server ini memiliki informasi yang tepat tentang alamat IP yang terkait dengan nama domain yang diminta.</p>
   </li>
   <li>
      <p><strong>Pengembalian Jawaban</strong>: Setelah server DNS primer atau server DNS otoritatif menemukan informasi yang benar, alamat IP yang sesuai dikirim kembali ke peramban Kamu.</p>
   </li>
   <li>
      <p><strong>Cache DNS</strong>: Informasi resolusi DNS yang diterima dari server DNS akan disimpan dalam cache DNS, baik pada peramban Kamu maupun pada server DNS Kamu. Ini membantu mempercepat proses resolusi DNS di masa mendatang.</p>
   </li>
   <li>
      <p><strong>Akses ke Sumber Daya</strong>: Setelah peramban Kamu menerima alamat IP yang sesuai dari resolusi DNS, ia dapat menggunakannya untuk menghubungi server web yang dituju dan mengambil halaman web yang diminta.</p>
   </li>
</ol>
<p>Dalam intinya, DNS memungkinkan Kamu untuk mengakses sumber daya di internet dengan menggunakan nama yang mudah diingat, tanpa perlu mengkhawatirkan tentang alamat IP yang sebenarnya di belakangnya. Proses ini memastikan bahwa internet tetap dapat diakses dengan cara yang lebih manusiawi dan efisien.</p>
<h2 id="123-struktur-dns">1.2.3 Struktur DNS</h2>
<p>Sistem Nama Domain (DNS) adalah sistem yang digunakan di Internet untuk mengonversi nama domain yang mudah diingat menjadi alamat IP numerik yang diperlukan untuk mengidentifikasi dan lokalisasi sumber daya online seperti situs web, server email, atau layanan jaringan lainnya. Struktur DNS adalah cara di mana informasi mengenai nama domain diatur dan disimpan dalam hierarki yang terstruktur. Ini melibatkan beberapa komponen dan tingkatan yang bekerja bersama untuk mengelola proses penamaan dan pencarian dalam jaringan.</p>
<p>Berikut adalah penjelasan mengenai struktur DNS:</p>
<ol>
   <li>
      <p><strong>Root Domain</strong>: Ini adalah tingkat tertinggi dalam hierarki DNS. Root domain tidak memiliki nama domain spesifik; itu hanya mewakili titik awal di mana semua nama domain dimulai. Root domain diwakili oleh titik (.) di akhir alamat domain. Misalnya, &quot;<a href="http://www.example.com">www.example.com</a>.&quot; titik di akhir menunjukkan akar domain.</p>
   </li>
   <li>
      <p><strong>Top-Level Domain (TLD)</strong>: Ini adalah tingkat di bawah root domain. TLD adalah ekstensi domain paling kanan dalam sebuah nama domain. Contoh TLD yang umum adalah .com, .net, .org, .gov, .edu, dan lain-lain. Setiap TLD dikelola oleh organisasi atau lembaga yang berbeda.</p>
   </li>
   <li>
      <p><strong>Second-Level Domain (SLD)</strong>: Ini adalah bagian di antara TLD dan nama spesifik dari sebuah situs web atau sumber daya. Dalam &quot;<a href="http://www.example.com">www.example.com</a>,&quot; &quot;example&quot; adalah SLD. Ini adalah bagian yang dapat dipersonalisasi oleh pemilik domain.</p>
   </li>
   <li>
      <p><strong>Subdomain</strong>: Ini adalah tingkat yang lebih rendah dari nama domain yang dapat ditambahkan untuk mengatur atau membedakan bagian-bagian dari sebuah situs web. Sebagai contoh, &quot;blog.example.com&quot; adalah subdomain dari &quot;example.com.&quot;</p>
   </li>
   <li>
      <p><strong>Hostnames</strong>: Ini adalah bagian dari subdomain yang merujuk pada sumber daya spesifik di dalam domain. Dalam &quot;<a href="http://www.example.com">www.example.com</a>,&quot; &quot;www&quot; adalah hostname yang merujuk pada server web.</p>
   </li>
   <li>
      <p><strong>Zone</strong>: Ini adalah wilayah administratif dalam struktur DNS yang berisi informasi tentang domain dan subdomain tertentu. Setiap zona biasanya dikelola oleh satu entitas yang bertanggung jawab atas pengaturan DNS untuk domain tersebut.</p>
   </li>
   <li>
      <p><strong>DNS Records</strong>: Informasi tentang nama domain dan bagaimana mereka harus diarahkan ke alamat IP disimpan dalam catatan DNS. Beberapa jenis catatan DNS meliputi:</p>
      <ul>
         <li><strong>A Record</strong>: Menghubungkan sebuah nama domain ke alamat IP.</li>
         <li><strong>CNAME Record</strong>: Menciptakan alias untuk sebuah nama domain, sering digunakan untuk membuat subdomain atau mengarahkan ke alamat lain.</li>
         <li><strong>MX Record</strong>: Menentukan server email yang bertanggung jawab untuk domain.</li>
         <li><strong>TXT Record</strong>: Digunakan untuk menyimpan informasi teks terkait domain, seperti verifikasi kepemilikan atau konfigurasi lainnya.</li>
         <li><strong>NS Record</strong>: Menunjukkan server nama yang memiliki otoritas atas zona DNS.</li>
      </ul>
   </li>
</ol>
<p>Struktur DNS ini berfungsi sebagai dasar untuk navigasi dan komunikasi dalam jaringan. Ketika Kamu memasukkan nama domain di peramban web Kamu, sistem DNS bekerja di belakang layar untuk menerjemahkan nama domain tersebut menjadi alamat IP yang dapat dimengerti oleh komputer, sehingga Kamu dapat mengakses sumber daya yang diinginkan.</p>
<h2 id="124-relasi-dns-dan-back-end">1.2.4 Relasi DNS dan Back-End</h2>
<p><strong>Back-End</strong> dalam konteks ini mengacu pada komponen-komponen sistem yang tersembunyi dari pengguna akhir, yang bertanggung jawab untuk memproses dan menyimpan data, serta melakukan berbagai fungsi penting dalam aplikasi atau situs web. Ini bisa mencakup server, database, logika bisnis, dan berbagai komponen lain yang bekerja di latar belakang untuk menjalankan aplikasi.</p>
<p><strong>Hubungan antara DNS dan Back-End:</strong></p>
<ol>
   <li>
      <p><strong>Resolusi DNS:</strong> Ketika pengguna memasukkan nama domain (seperti <a href="http://www.contoh.com">www.contoh.com</a>) dalam peramban web mereka, peramban akan mengirim permintaan DNS ke server DNS lokal atau server DNS penyedia layanan internet (ISP). Server DNS ini bertindak sebagai bagian dari Back-End dan bertanggung jawab untuk mencari dan mengembalikan alamat IP yang terkait dengan nama domain tersebut.</p>
   </li>
   <li>
      <p><strong>Tautan dengan Server Web (HTTP/HTTPS):</strong> Setelah alamat IP ditemukan, peramban akan menggunakan alamat IP tersebut untuk berkomunikasi dengan server web yang di-hosting di alamat IP tersebut. Ini melibatkan pertukaran data antara Back-End server web dan perangkat akhir pengguna.</p>
   </li>
   <li>
      <p><strong>Database DNS:</strong> Server DNS menggunakan basis data DNS untuk menyimpan informasi tentang nama domain dan alamat IP yang terkait. Database ini merupakan bagian dari Back-End DNS dan diatur oleh otoritas DNS yang berwenang.</p>
   </li>
   <li>
      <p><strong>Kontrol DNS:</strong> Pengaturan DNS yang lebih kompleks, seperti penerusan (forwarding), penyesuaian zona, dan penanganan DNSSEC (DNS Security Extensions), diatur di Back-End oleh administrator jaringan atau sistem yang berwenang.</p>
   </li>
</ol>
<p>Jadi, DNS berfungsi sebagai penghubung antara nama domain yang mudah diingat dengan alamat IP yang diperlukan oleh Back-End sistem untuk mengarahkan koneksi dan memberikan konten yang diminta kepada pengguna. Back-End sendiri mengatur semua proses dan komponen teknis yang memungkinkan sebuah situs web atau aplikasi berjalan dengan lancar dan menyediakan layanan kepada pengguna akhir.</p>