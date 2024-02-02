<div align="center">
   <h1>Module 3 Lesson 1 - Pengenalan Version Control</h1>
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
         <td>3.1.1 Tentang Version Control</td>
         <td><a href="#211-apa-itu-sistem-operasi">🔽</a></td>
      </tr>
      <tr>
         <td>3.1.2 Metode Version Control</td>
         <td><a href="#212-tipe-tipe-sistem-operasi">🔽</a></td>
      </tr>
      <tr>
         <td>3.1.3 Manfaat Version Control System</td>
         <td><a href="#213-fitur-dari-sistem-operasi">🔽</a></td>
      </tr>
   </tbody>
</table>
<p><br><br></p>
<h1 id="durasi-pembelajaran">Durasi Pembelajaran</h1>
<p><strong>45 menit</strong></p>
<br>
<h2 id="311-tentang-version-control">3.1.1 Tentang Version Control</h2>
<p>Version Control (Pengendalian Versi) adalah praktik yang digunakan dalam pengembangan perangkat lunak dan pengelolaan proyek kolaboratif untuk melacak perubahan dalam kode sumber, dokumen, atau aset digital lainnya sepanjang waktu. Tujuan utama dari Version Control adalah mengontrol dan merekam setiap perubahan yang terjadi dalam proyek, memungkinkan tim pengembang bekerja secara efisien, berkolaborasi, dan memantau evolusi proyek.</p>
<p>Jika kamu seorang desainer grafis atau web dan ingin menyimpan setiap versi dari gambar atau tata letak yang pasti ingin kamu lakukan, Sistem Kontrol Versi (Version Control System/VCS) adalah hal yang sangat bijaksana untuk digunakan. Ini memungkinkan kamu mengembalikan file-file terpilih ke keadaan sebelumnya, mengembalikan seluruh proyek ke keadaan sebelumnya, membandingkan perubahan dari waktu ke waktu, melihat siapa yang terakhir memodifikasi sesuatu yang mungkin menyebabkan masalah, siapa yang memperkenalkan masalah dan kapan, dan banyak lagi. Menggunakan VCS juga umumnya berarti bahwa jika kamu membuat kesalahan atau kehilangan file, kamu dapat dengan mudah memulihkannya. Selain itu, kamu mendapatkan semua ini dengan biaya yang sangat rendah.</p>
<p><strong>Alat Version Control Populer</strong></p>
<ol>
   <li>
      <p><strong>Git:</strong> Sistem Version Control terdistribusi yang sangat populer. Git adalah dasar dari banyak platform pengelolaan kode seperti GitHub dan GitLab.</p>
   </li>
   <li>
      <p><strong>Subversion (SVN):</strong> Sistem Version Control terpusat yang lebih tua, yang masih digunakan dalam beberapa proyek.</p>
   </li>
   <li>
      <p><strong>Mercurial:</strong> Alternatif terhadap Git yang juga bersifat terdistribusi.</p>
   </li>
</ol>
<h2 id="312-metode-version-control">3.1.2 Metode Version Control</h2>
<p>Metode Version Control (VC) memungkinkan tim pengembang atau individu untuk bekerja sama dalam pengembangan perangkat lunak atau proyek-proyek lainnya dengan cara yang terorganisir, terdokumentasi, dan efisien.</p>
<h3 id="local-version-control-systems">Local Version Control Systems</h3>
<p>Metode kontrol versi pilihan banyak orang adalah dengan menyalin file ke direktori lain atau jika mereka sangat teliti, juga akan ada yang dilabeli dengan waktu pembuatan. Pendekatan ini sangat umum karena sangat sederhana, tetapi juga sangat rentan terhadap kesalahan. Mudah untuk lupa di direktori mana kamu berada dan secara tidak sengaja menulis ke file yang salah atau menyalin file yang sebenarnya tidak dimaksudkan.</p>
<p>Untuk mengatasi masalah ini, para programmer sudah lama mengembangkan VCS lokal yang memiliki basis data sederhana yang mencatat semua perubahan pada file yang berada di bawah kendali revisi.</p>
<div align= "center">
   <img src="https://git-scm.com/book/en/v2/images/local.png">
   <br>
   <figcaption>Local Version Control Systems</figcaption>
</div>
<br>
<h3 id="centralized-version-control-systems">Centralized Version Control Systems</h3>
<p>Sistem Kontrol Versi Terpusat atau Centralized Version Control Systems merupakan salah satu pengembangan version control dengan ide ketika developer melakukan kerja sama development dengan sistem yang berbeda(contoh sistem operasi). Dengan sistem, file-file development berada dalam satu server dan para developer dapat mengakses file tersebut dapat ada batasan perbedaan sistem.</p>
<p>Keuntungan dengan tipe version control ini semua developer atau kolaborator dapat mengetahui perubahan dan pembangunan dari development  tersebut.Pada sistem ini, Administrator dari development  ini dapat mengelola dan kontrol penuh terhadap developer  yang bekerja sehingga lebih mudah dalam mengelola semua perubahan dari proses development  tersebut.</p>
<div align= "center">
   <img src="https://git-scm.com/book/en/v2/images/centralized.png">
   <br>
   <figcaption>Centralized Version Control Systems</figcaption>
</div>
<br>
<p>Kekurangan dari sistem version control  ini karena semua development files yang telah tersimpan dalam satu server tersebut bisa saja tidak dapat diakses(contoh yang memungkinkan server maintenance) dan jika ada developer ingin menyimpan hasil perubahan tertentu akan sama sekali tidak tersimpan sehingga akan hanya tersimpan pada lokal. Sistem ini juga bergantung dengan basis data pusat, seperti storage. Maka, jika storage mengalami kerusakan, maka semua perubahan dan inisiasi juga hilang.</p>
<h3 id="distributed-version-control-systems">Distributed Version Control Systems</h3>
<p>Hingga saat ini, Sistem Kontrol Versi Terdistribusi atau Distributed Version Control Systems merupakan sistem yang terbaik daripada sistem-sistem sebelumnya. Disini, file dan juga perubahannya disimpan dalam bentuk repositori. Sehingga apabila adanya server maintenance,  repositori dapat disalin kembali dan bisa melanjutkan pengembagannya. History  dari pengembangan tersebut akan tercermin kepada repositori yang sudah ada.</p>
<div align= "center">
   <img src="https://git-scm.com/book/en/v2/images/distributed.png">
   <br>
   <figcaption> Distributed Version Control Systems</figcaption>
</div>
<br>
<h2 id="313-manfaat-version-control-system">3.1.3 Manfaat Version Control System</h2>
<p>Version Control System (VCS) adalah sistem yang digunakan untuk mengelola perubahan pada dokumen, kode sumber, atau berkas-berkas lainnya dalam pengembangan perangkat lunak atau proyek kolaboratif. VCS memiliki peran penting dalam memfasilitasi pengelolaan dan pemantauan perubahan, serta meningkatkan efisiensi dan kualitas pengembangan perangkat lunak. Berikut ini adalah beberapa manfaat utama dari penggunaan Version Control System:</p>
<h3 id="pemantauan-perubahan">Pemantauan Perubahan</h3>
<p>VCS memungkinkan pengguna untuk melacak setiap perubahan yang terjadi pada berkas-berkas proyek. Setiap perubahan akan direkam dalam bentuk revisi atau commit, yang mencakup informasi tentang perubahan yang dilakukan, siapa yang melakukan perubahan, dan kapan perubahan tersebut terjadi. Ini memungkinkan pengembang untuk memahami kronologi perubahan dan menelusuri kembali jika diperlukan.</p>
<h3 id="kolaborasi-tim">Kolaborasi Tim</h3>
<p>Dalam pengembangan perangkat lunak atau proyek kolaboratif, berbagai anggota tim biasanya bekerja pada berkas-berkas yang sama. VCS memungkinkan banyak pengembang untuk bekerja bersama pada proyek yang sama tanpa mengalami konflik atau kebingungan tentang versi mana yang saat ini digunakan. Setiap pengembang dapat bekerja pada cabang (branch) yang terpisah dan kemudian menggabungkan perubahan mereka melalui proses yang disebut merging.</p>
<h3 id="pengelolaan-konflik">Pengelolaan Konflik</h3>
<p>Ketika beberapa pengembang bekerja pada berkas yang sama secara bersamaan, ada kemungkinan terjadi konflik saat mereka mencoba menggabungkan perubahan mereka. VCS menyediakan alat untuk mengatasi konflik dengan cara membandingkan perubahan yang dilakukan oleh berbagai pengembang dan memungkinkan pengguna untuk memilih perubahan mana yang harus diterima.</p>
<h3 id="pemulihan-dan-reversi">Pemulihan dan Reversi</h3>
<p>Jika terjadi kesalahan atau perubahan yang tidak diinginkan, VCS memungkinkan pengembang untuk membatalkan perubahan tersebut atau kembali ke revisi sebelumnya. Ini membantu dalam mengelola risiko dan memastikan bahwa perubahan yang dilakukan tidak merusak fungsionalitas atau integritas proyek.</p>
<h3 id="pelacakan-bug">Pelacakan Bug</h3>
<p>Dalam pengembangan perangkat lunak, seringkali diperlukan pelacakan bug atau masalah lainnya yang mungkin muncul selama pengembangan. Dengan VCS, pengembang dapat melihat perubahan yang terjadi sebelum dan setelah bug ditemukan, membantu dalam mengidentifikasi penyebab dan solusi potensial.</p>
<h3 id="penyimpanan-aman">Penyimpanan Aman</h3>
<p>VCS menyediakan penyimpanan aman untuk berkas-berkas proyek. Revisi-revisi dan perubahan-perubahan disimpan dalam repositori, yang dapat diakses dan dikelola oleh tim pengembang yang memiliki izin akses. Hal ini mengurangi risiko kehilangan data atau kerusakan berkas.</p>
<h3 id="uji-coba-dan-eksperimen">Uji Coba dan Eksperimen</h3>
<p>VCS memungkinkan pengembang untuk membuat cabang-cabang baru dari proyek utama, yang dapat digunakan untuk uji coba atau eksperimen tanpa mempengaruhi proyek utama. Setelah pengujian selesai, cabang tersebut dapat digabungkan kembali ke proyek utama atau dihapus.</p>
<h3 id="dokumentasi">Dokumentasi</h3>
<p>Setiap commit pada VCS dilengkapi dengan pesan komit yang menjelaskan perubahan yang dilakukan. Ini membantu dalam menciptakan dokumentasi otomatis tentang perubahan-perubahan yang terjadi pada proyek dari waktu ke waktu.</p>
<p>Dengan manfaat-manfaat di atas, penggunaan Version Control System menjadi sangat penting dalam pengembangan perangkat lunak modern dan proyek-proyek kolaboratif. Ini membantu meningkatkan efisiensi, kualitas, dan keteraturan pengembangan serta memfasilitasi kerja tim secara lebih efektif.</p>