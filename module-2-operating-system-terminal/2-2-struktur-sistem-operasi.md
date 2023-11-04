<div align="center">
   <h1>Module 2 Lesson 2 - Struktur Sistem Operasi</h1>
   Author: <b>Samuel Pandohan Terampil Gultom</b>
   <br>
   Last Updated: <b>August 1, 2023</b>
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
         <td>2.2.1 Kernel</td>
         <td><a href="#221-kernel">🔽</a></td>
      </tr>
      <tr>
         <td>2.2.2 Shell</td>
         <td><a href="#222-shell">🔽</a></td>
      </tr>
      <tr>
         <td>2.2.3 Software</td>
         <td><a href="#223-software">🔽</a></td>
      </tr>
      <tr>
         <td>2.2.4 Hardware</td>
         <td><a href="#224-hardware">🔽</a></td>
      </tr>
   </tbody>
</table>
<p><br><br></p>
<h1 id="durasi-pembelajaran">Durasi Pembelajaran</h1>
<p><strong>60 menit</strong></p>
<br>
<h2 id="221-kernel">2.2.1 Kernel</h2>
<h3 id="definisi-kernel">Definisi Kernel</h3>
<p>Kernel merupakan komponen inti dari sebuah sistem operasi. Ia bertanggung jawab untuk mengatur interaksi antara perangkat keras (hardware) dan perangkat lunak (software), serta menyediakan layanan dasar yang diperlukan oleh aplikasi dan pengguna. Kernel bekerja sebagai jembatan antara aplikasi yang berjalan di atasnya dan perangkat keras yang ada di bawahnya. Mari kita jelaskan lebih lanjut tentang kernel dalam konteks struktur sistem operasi.</p>
<h3 id="fungsi-kernel">Fungsi Kernel</h3>
<p>Kernel adalah bagian terpenting dari sistem operasi yang bertanggung jawab untuk menjalankan tugas-tugas esensial dalam manajemen sistem. Fungsi utama kernel meliputi:</p>
<ol>
   <li>
      <p><strong>Manajemen Sumber Daya:</strong> Kernel mengelola dan mengalokasikan sumber daya perangkat keras seperti CPU, memori, perangkat masukan/keluaran, dan perangkat penyimpanan. Ia memastikan bahwa sumber daya ini digunakan secara efisien dan adil di antara berbagai aplikasi.</p>
   </li>
   <li>
      <p><strong>Manajemen Proses:</strong> Kernel mengelola proses-proses yang berjalan pada sistem. Ia mengendalikan pembuatan, penghentian, dan beralihnya proses-proses, serta menjaga pengaturan waktu (scheduling) untuk memastikan setiap proses mendapatkan akses ke CPU.</p>
   </li>
   <li>
      <p><strong>Manajemen Memori:</strong> Kernel mengatur penggunaan memori fisik dan virtual. Ia bertanggung jawab atas alokasi dan dealokasi memori bagi proses-proses, serta mengatasi konsep virtualisasi memori agar aplikasi dapat berjalan tanpa perlu mengetahui detail fisik dari memori yang digunakan.</p>
   </li>
   <li>
      <p><strong>Manajemen Sistem File:</strong> Kernel menyediakan akses ke sistem file, memungkinkan aplikasi dan pengguna untuk membaca, menulis, dan mengelola berkas dan direktori pada perangkat penyimpanan.</p>
   </li>
   <li>
      <p><strong>Manajemen Perangkat Masukan/Keluaran:</strong> Kernel mengelola interaksi dengan perangkat keras masukan dan keluaran seperti keyboard, mouse, printer, dan perangkat jaringan. Ia menyediakan antarmuka yang memungkinkan aplikasi berkomunikasi dengan perangkat ini.</p>
   </li>
   <li>
      <p><strong>Komunikasi Antar Proses:</strong> Kernel menyediakan mekanisme bagi proses-proses untuk berkomunikasi satu sama lain, baik dalam bentuk pertukaran data maupun sinkronisasi.</p>
   </li>
</ol>
<h3 id="struktur-kernel">Struktur Kernel</h3>
<p>Struktur internal kernel dapat berbeda-beda tergantung pada desain dan tipe sistem operasi. Namun, ada beberapa komponen umum dalam struktur kernel:</p>
<ol>
   <li>
      <p><strong>Monolithic Kernel:</strong> Kernel monolitik adalah tipe kernel yang memiliki semua fungsionalitas inti dalam satu ruang alamat. Semua layanan dan sistem panggilan dikelompokkan dalam satu entitas besar.</p>
   </li>
   <li>
      <p><strong>Microkernel:</strong> Microkernel memisahkan fungsionalitas inti ke dalam set servis yang lebih kecil. Hanya layanan paling esensial seperti manajemen proses dan komunikasi antar proses yang berada di dalam inti, sementara layanan lain berjalan di luar kernel.</p>
   </li>
   <li>
      <p><strong>Hybrid Kernel:</strong> Hybrid kernel mencoba menggabungkan keuntungan dari monolithic kernel dan microkernel. Beberapa layanan mungkin berada di dalam kernel, sementara yang lain berjalan di luar kernel dalam bentuk modul terpisah.</p>
   </li>
</ol>
<p>Kernel adalah komponen terpenting dalam struktur sistem operasi. Ia mengatur interaksi antara perangkat keras dan perangkat lunak, serta menyediakan layanan dasar untuk aplikasi dan pengguna. Ada berbagai jenis struktur kernel, seperti monolithic, microkernel, dan hybrid, yang digunakan dalam berbagai sistem operasi.</p>
<h2 id="222-shell">2.2.2 Shell</h2>
<p>Shell adalah antarmuka yang menghubungkan pengguna dengan sistem operasi pada komputer. Ini adalah lingkungan di mana pengguna dapat berinteraksi dengan sistem operasi melalui perintah-perintah teks atau skrip. Shell berfungsi sebagai perantara antara pengguna dan inti (kernel) sistem operasi.</p>
<p>Shell dapat membantu pengguna untuk menjalankan perintah-perintah sistem, mengelola berkas dan direktori, menjalankan program, serta melakukan berbagai tugas administratif lainnya pada sistem operasi. Shell juga memungkinkan pengguna untuk melakukan otomatisasi tugas dengan membuat skrip atau file batch yang berisi kumpulan perintah.</p>
<h3 id="jenis-jenis-shell">Jenis-Jenis Shell</h3>
<p>Ada beberapa jenis shell yang tersedia dalam sistem operasi. Beberapa jenis shell yang umum dijumpai adalah:</p>
<ol>
   <li>
      <p><strong>Bourne Shell (sh)</strong>: Ini adalah shell yang lebih awal dan sederhana, yang memberikan dasar untuk shell lainnya. Shell ini memiliki banyak variasi, termasuk Bourne-Again Shell (bash), yang banyak digunakan pada sistem Linux.</p>
   </li>
   <li>
      <p><strong>C Shell (csh)</strong>: Shell ini memiliki sintaks yang lebih mirip dengan bahasa pemrograman C. Meskipun memiliki fitur-fitur menarik, C Shell tidak sepopuler bash.</p>
   </li>
   <li>
      <p><strong>Korn Shell (ksh)</strong>: Shell ini merupakan perluasan dari Bourne Shell dan C Shell. Ia menyediakan banyak fitur tambahan dan peningkatan dalam hal scripting.</p>
   </li>
   <li>
      <p><strong>Bash (Bourne-Again Shell)</strong>: Ini adalah salah satu shell paling umum digunakan pada sistem Linux dan macOS. Bash menggabungkan fitur-fitur dari Bourne Shell, C Shell, dan Korn Shell.</p>
   </li>
   <li>
      <p><strong>Zsh (Z Shell)</strong>: Zsh adalah shell yang canggih dengan banyak fitur keren, seperti penyelesaian otomatis yang lebih cerdas dan dukungan untuk tema-tema visual.</p>
   </li>
</ol>
<h3 id="fungsi-shell-dalam-struktur-sistem-operasi">Fungsi Shell dalam Struktur Sistem Operasi</h3>
<ol>
   <li>
      <p><strong>Interaksi Pengguna</strong>: Shell memungkinkan pengguna untuk berinteraksi dengan sistem operasi melalui perintah-perintah teks. Pengguna dapat menjalankan program, mengelola berkas, dan menjalankan tugas-tugas lainnya melalui shell.</p>
   </li>
   <li>
      <p><strong>Manajemen Berkas dan Direktori</strong>: Pengguna dapat menggunakan perintah-perintah shell untuk membuat, menghapus, memindahkan, dan mengelola berkas dan direktori.</p>
   </li>
   <li>
      <p><strong>Eksekusi Program</strong>: Shell memungkinkan pengguna untuk menjalankan program atau skrip dengan memberikan perintah eksekusi.</p>
   </li>
   <li>
      <p><strong>Pengelolaan Tugas</strong>: Pengguna dapat menjalankan perintah-perintah latar belakang (background) atau latar depan (foreground) dalam shell. Ini memungkinkan pengguna untuk menjalankan beberapa tugas secara bersamaan.</p>
   </li>
   <li>
      <p><strong>Otomatisasi Tugas</strong>: Dengan membuat skrip shell, pengguna dapat mengotomatisasi tugas-tugas yang perlu diulang secara teratur.</p>
   </li>
   <li>
      <p><strong>Pengalihan Input dan Output</strong>: Shell memungkinkan pengguna untuk mengalihkan input dan output dari dan ke berkas atau perangkat lainnya.</p>
   </li>
   <li>
      <p><strong>Variabel Lingkungan</strong>: Shell dapat digunakan untuk mengatur variabel lingkungan yang memengaruhi perilaku program dan skrip yang dijalankan.</p>
   </li>
   <li>
      <p><strong>Pengaturan Lingkungan</strong>: Shell memungkinkan pengguna untuk mengatur pengaturan lingkungan mereka, seperti variabel PATH, yang mengarahkan shell ke lokasi program.</p>
   </li>
</ol>
<p>Dalam struktur sistem operasi, shell berfungsi sebagai antarmuka antara pengguna dan sistem operasi. Ia memungkinkan pengguna untuk menjalankan perintah-perintah sistem, mengelola berkas, menjalankan program, dan melakukan tugas-tugas lainnya melalui antarmuka teks. Ada berbagai jenis shell yang tersedia, masing-masing dengan fitur dan sintaks yang berbeda.</p>
<h2 id="223-software">2.2.3 Software</h2>
<p>Sistem operasi adalah perangkat lunak yang berperan penting dalam mengelola dan mengoordinasikan semua sumber daya perangkat keras (hardware) serta menyediakan layanan bagi perangkat lunak aplikasi. Dalam konteks struktur sistem operasi, perangkat lunak atau software dapat dikelompokkan menjadi beberapa lapisan atau komponen yang bekerja bersama untuk menjalankan fungsi-fungsi tertentu. Berikut ini penjelasan mengenai peran dan jenis-jenis software dalam struktur sistem operasi:</p>
<h3 id="1-kernel-inti-sistem-operasi">1. Kernel (Inti Sistem Operasi)</h3>
<p>Kernel adalah bagian terdalam dari sistem operasi yang bertanggung jawab untuk mengatur akses langsung ke perangkat keras dan menyediakan layanan dasar yang diperlukan oleh aplikasi. Kernel juga mengelola alokasi sumber daya seperti memori, prosesor, dan input/output. Terdapat dua jenis utama kernel, yaitu:</p>
<ul>
   <li><strong>Monolithic Kernel:</strong> Semua fungsi dan layanan terkait sistem operasi berada dalam satu unit monolitik.</li>
   <li><strong>Microkernel:</strong> Kernel ini hanya menyediakan layanan inti seperti manajemen tugas dan komunikasi antarproses. Layanan lain seperti sistem berkas dan jaringan ditempatkan di lapisan terpisah yang berjalan di atas mikrokernel.</li>
</ul>
<h3 id="2-device-drivers-driver-perangkat">2. Device Drivers (Driver Perangkat)</h3>
<p>Device drivers adalah software yang memungkinkan sistem operasi berinteraksi dengan perangkat keras. Mereka menerjemahkan perintah dari sistem operasi menjadi perintah yang dimengerti oleh perangkat keras, sehingga memungkinkan aplikasi dan sistem operasi untuk berkomunikasi dengan perangkat-perangkat seperti printer, keyboard, mouse, dan lainnya.</p>
<h3 id="3-system-libraries-perpustakaan-sistem">3. System Libraries (Perpustakaan Sistem)</h3>
<p>Perpustakaan sistem adalah kumpulan kode yang menyediakan berbagai fungsi umum kepada aplikasi. Ini termasuk fungsi-fungsi untuk operasi input/output, manipulasi string, alokasi memori, dan lain-lain. Penggunaan perpustakaan ini menghindarkan pengulangan kode dan mempermudah pengembangan aplikasi.</p>
<h3 id="4-shell-lingkungan-shell">4. Shell (Lingkungan Shell)</h3>
<p>Shell adalah antarmuka yang memungkinkan pengguna berinteraksi dengan sistem operasi melalui perintah-perintah teks atau grafis. Shell menerima perintah dari pengguna, mengeksekusi perintah tersebut, dan menampilkan hasilnya. Ada berbagai jenis shell, seperti shell berbasis baris perintah (command-line) dan shell berbasis GUI (Graphical User Interface).</p>
<h3 id="5-utilities-utilitas">5. Utilities (Utilitas)</h3>
<p>Utilities adalah program-program kecil yang membantu dalam melakukan tugas-tugas spesifik seperti manajemen berkas, pemrosesan teks, pemadaman sistem, pengelolaan proses, dan lain-lain. Contoh utilitas termasuk perintah <code>ls</code>, <code>cp</code>, <code>mv</code>, <code>rm</code>, dan lainnya dalam lingkungan berbasis baris perintah.</p>
<h3 id="6-application-programs-program-aplikasi">6. Application Programs (Program Aplikasi)</h3>
<p>Program aplikasi adalah perangkat lunak yang dibuat untuk tujuan tertentu, seperti pengolah kata, lembar kerja, aplikasi desain grafis, permainan, dan sebagainya. Aplikasi ini berjalan di atas sistem operasi dan menggunakan layanan-layanan yang disediakan oleh sistem operasi melalui antarmuka aplikasi (API).</p>
<p>Dalam struktur sistem operasi, perangkat lunak memiliki peran penting untuk mengelola sumber daya perangkat keras dan memberikan layanan kepada aplikasi. Dari kernel yang menjadi inti, hingga aplikasi yang digunakan oleh pengguna akhir, setiap komponen software dalam sistem operasi bekerja bersama untuk menjalankan sistem dengan efisien dan efektif.</p>
<h2 id="224-hardware">2.2.4 Hardware</h2>
<p>Hardware dalam konteks sistem operasi merujuk pada komponen fisik yang membentuk infrastruktur komputer, termasuk perangkat keras utama dan perangkat keras pendukung. Sistem operasi (OS) adalah perangkat lunak yang mengelola sumber daya perangkat keras dan menyediakan antarmuka antara perangkat lunak aplikasi dengan perangkat keras. Pengetahuan tentang hardware sangat penting dalam memahami cara kerja sistem operasi, karena OS harus berinteraksi secara efektif dengan perangkat keras untuk menjalankan tugasnya.</p>
<p>Berikut adalah beberapa komponen perangkat keras utama yang relevan dalam konteks struktur sistem operasi:</p>
<ol>
   <li>
      <p><strong>CPU (Central Processing Unit)</strong>:
         CPU adalah otak komputer yang menjalankan instruksi-instruksi program. Sistem operasi berinteraksi dengan CPU untuk mengatur penjadwalan tugas, alokasi sumber daya, dan mengawasi eksekusi program. Konsep seperti mode pengguna dan mode kernel (mode supervisor) juga terkait dengan CPU, di mana mode kernel memungkinkan sistem operasi menjalankan instruksi dengan hak akses lebih tinggi daripada mode pengguna.
      </p>
   </li>
   <li>
      <p><strong>Memory (RAM)</strong>:
         Memori adalah tempat penyimpanan sementara untuk program dan data yang sedang diakses oleh CPU. Sistem operasi bertanggung jawab atas manajemen memori, termasuk alokasi, dealokasi, dan pengaturan tindakan saat memori penuh. Konsep manajemen memori, seperti virtual memory, paging, dan segmentasi, penting dalam memahami bagaimana sistem operasi mengelola memori fisik dan logis.
      </p>
   </li>
   <li>
      <p><strong>Perangkat Penyimpanan (Storage Devices)</strong>:
         Sistem operasi harus berinteraksi dengan perangkat penyimpanan seperti hard drive, solid-state drive (SSD), dan perangkat penyimpanan lainnya. Ini melibatkan manajemen berkas dan sistem berkas yang memungkinkan pembuatan, penghapusan, pengaturan izin akses, dan organisasi data di perangkat penyimpanan.
      </p>
   </li>
   <li>
      <p><strong>Perangkat Input dan Output (I/O Devices)</strong>:
         OS harus berkomunikasi dengan perangkat input (keyboard, mouse) dan output (layar, printer, speaker) untuk mengizinkan pengguna berinteraksi dengan sistem. Manajemen I/O melibatkan penjadwalan tugas yang menunggu operasi I/O, buffering data, dan mengoptimalkan kinerja perangkat I/O.
      </p>
   </li>
   <li>
      <p><strong>Perangkat Jaringan</strong>:
         Dalam lingkungan jaringan, sistem operasi juga berperan dalam mengelola koneksi jaringan, mengatur protokol komunikasi, dan memastikan keamanan dan integritas data dalam transfer data melalui jaringan.
      </p>
   </li>
   <li>
      <p><strong>Interrupts</strong>:
         Sistem operasi harus merespons interrupts, yang adalah sinyal dari perangkat keras kepada CPU bahwa suatu peristiwa telah terjadi, seperti tombol ditekan atau transfer data selesai. Interrupts memungkinkan OS untuk mengalihkan perhatian CPU dari tugas yang sedang berjalan untuk menangani peristiwa penting.
      </p>
   </li>
   <li>
      <p><strong>BIOS/UEFI</strong>:
         Basic Input/Output System (BIOS) atau Unified Extensible Firmware Interface (UEFI) adalah program perangkat keras yang memulai sistem saat dinyalakan. Sistem operasi berinteraksi dengan BIOS/UEFI selama proses booting.
      </p>
   </li>
</ol>
<p>Memahami bagaimana struktur sistem operasi berinteraksi dengan perangkat keras membantu dalam mengoptimalkan kinerja sistem dan memahami aspek-aspek fundamental dalam pengembangan, pemecahan masalah, dan administrasi sistem operasi.</p>