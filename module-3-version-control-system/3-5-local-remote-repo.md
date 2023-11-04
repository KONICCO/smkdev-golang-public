<div align="center">
   <h1>Module 3 Lesson 5 - Local vs Remote Repo</h1>
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
         <td>3.5.1 Local Repository</td>
         <td><a href="#351-local-repository">🔽</a></td>
      </tr>
      <tr>
         <td>3.5.2 Remote Repository</td>
         <td><a href="#352-remote-repository">🔽</a></td>
      </tr>
      <tr>
         <td>3.5.3 Tabel Perbedaan</td>
         <td><a href="#353-tabel-perbedaan">🔽</a></td>
      </tr>
      <tr>
         <td>3.5.4 Konflik dan Resolusi Konflik</td>
         <td><a href="#354-konflik-dan-resolusi-konflik">🔽</a></td>
      </tr>
   </tbody>
</table>
<p><br><br></p>
<h1 id="durasi-pembelajaran">Durasi Pembelajaran</h1>
<p><strong>60 menit</strong></p>
<br>
<h2 id="351-local-repository">3.5.1 Local Repository</h2>
<p>Local repository adalah salinan dari repository yang berada di komputer lokal kamu. Ini adalah tempat di mana kamu menyimpan semua perubahan, versi, dan riwayat proyek kamu. Repo lokal memungkinkan kamu bekerja secara offline dan mengelola proyek tanpa koneksi internet. kamu dapat melakukan perubahan, menambahkan file, menghapus file, dan mengubah versi proyek kamu di repo lokal. kamu juga dapat membuat cabang (branch) baru dan menggabungkan (merge) cabang tersebut ke dalam cabang utama (biasanya cabang utama disebut &quot;master&quot; atau &quot;main&quot;). Repo lokal sangat berguna saat kamu sedang mengembangkan proyek secara mandiri atau ketika kamu ingin mencoba perubahan sebelum mempublikasikannya ke repo jarak jauh (remote repositories).</p>
<p><strong>Keuntungan Local Repository:</strong></p>
<ul>
   <li><strong>Kecepatan Akses:</strong> Karena repository berada di komputer lokal kamu, akses ke file dan sejarah versi sangat cepat.</li>
   <li><strong>Kontrol Penuh:</strong> kamu memiliki kendali penuh atas repository ini dan tidak perlu bergantung pada koneksi internet atau server eksternal.</li>
   <li><strong>Isolasi Pekerjaan:</strong> kamu dapat bekerja di repository ini tanpa mempengaruhi versi yang dikerjakan oleh orang lain.</li>
</ul>
<h2 id="352-remote-repository">3.5.2 Remote Repository</h2>
<p>Di sisi lain, remote repository adalah repositori yang berada di server jarak jauh, biasanya di layanan seperti GitHub, GitLab, atau Bitbucket. Remote repository adalah tempat di mana kamu menyimpan proyek kamu secara terpusat dan memungkinkan kolaborasi dengan orang lain. Ini adalah tempat di mana kamu dapat berbagi perubahan dan memperbarui proyek kamu dengan orang lain melalui internet. Ketika kamu ingin membagikan proyek kamu dengan orang lain atau mengunduh perubahan yang dilakukan oleh orang lain, kamu akan berinteraksi dengan remote repository. Remote repository juga berfungsi sebagai cadangan (backup) proyek kamu sehingga jika terjadi kerusakan atau kehilangan pada komputer lokal kamu, data proyek kamu tetap aman.</p>
<p>Perbedaan antara local repository dan remote repository adalah bahwa local repository berada di komputer lokal kamu dan digunakan untuk mengelola dan mengubah proyek secara lokal, sementara remote repository adalah salinan proyek yang disimpan di server jarak jauh dan digunakan untuk berbagi, berkolaborasi, dan membuat cadangan proyek kamu.</p>
<p><strong>Keuntungan Remote Repository:</strong></p>
<ul>
   <li><strong>Kolaborasi:</strong> Memungkinkan anggota tim bekerja bersama pada proyek yang sama tanpa batasan geografis.</li>
   <li><strong>Cadangan:</strong> Memberikan cadangan sentral untuk kode kamu, sehingga mencegah kehilangan data jika terjadi masalah pada komputer lokal.</li>
   <li><strong>Pelacakan Perubahan:</strong> Menyimpan histori perubahan dan versi dari kode, memungkinkan kamu untuk melacak evolusi proyek dari waktu ke waktu.</li>
</ul>
<h2 id="353-tabel-perbedaan">3.5.3 Tabel Perbedaan</h2>
<table>
   <thead>
      <tr>
         <th><strong>Aspek</strong></th>
         <th><strong>Local Repository</strong></th>
         <th><strong>Remote Repository</strong></th>
      </tr>
   </thead>
   <tbody>
      <tr>
         <td><strong>Lokasi</strong></td>
         <td>Terletak di komputer lokal pengembang.</td>
         <td>Terletak di server jarak jauh (umumnya hosting service).</td>
      </tr>
      <tr>
         <td><strong>Akses</strong></td>
         <td>Hanya dapat diakses oleh pengembang pada komputer tertentu.</td>
         <td>Dapat diakses oleh tim pengembang dari berbagai lokasi.</td>
      </tr>
      <tr>
         <td><strong>Kontrol Versi</strong></td>
         <td>Kontrol versi hanya berlaku secara lokal.</td>
         <td>Kontrol versi bersifat kolaboratif dan dapat diakses oleh semua anggota tim.</td>
      </tr>
      <tr>
         <td><strong>Kepemilikan</strong></td>
         <td>Milik penuh oleh pengembang lokal.</td>
         <td>Dapat dimiliki oleh individu atau organisasi yang memiliki akses.</td>
      </tr>
      <tr>
         <td><strong>Sinkronisasi</strong></td>
         <td>Memerlukan tindakan manual untuk menggabungkan perubahan dengan repository lainnya.</td>
         <td>Sinkronisasi sering terjadi secara otomatis melalui operasi push dan pull.</td>
      </tr>
      <tr>
         <td><strong>Ketersediaan Data</strong></td>
         <td>Hanya tersedia pada komputer lokal yang memiliki salinan repository.</td>
         <td>Tersedia dan dapat diakses oleh siapa pun dengan hak akses.</td>
      </tr>
      <tr>
         <td><strong>Kecepatan Akses</strong></td>
         <td>Akses cepat karena berjalan di komputer lokal.</td>
         <td>Akses bisa lebih lambat tergantung pada koneksi internet.</td>
      </tr>
      <tr>
         <td><strong>Kerahasiaan</strong></td>
         <td>Data hanya ada di komputer pengembang lokal, lebih aman secara fisik.</td>
         <td>Memerlukan pengaturan keamanan untuk melindungi data dari akses tidak sah.</td>
      </tr>
      <tr>
         <td><strong>Kolaborasi</strong></td>
         <td>Tidak mendukung kolaborasi tim secara langsung.</td>
         <td>Mendukung kolaborasi tim dengan berbagi kode dan pekerjaan.</td>
      </tr>
      <tr>
         <td><strong>Cadangan (Backup)</strong></td>
         <td>Memerlukan upaya terpisah untuk mencadangkan data lokal.</td>
         <td>Data sering dicadangkan oleh penyedia layanan hosting.</td>
      </tr>
      <tr>
         <td><strong>Risiko Kehilangan</strong></td>
         <td>Lebih berisiko jika terjadi kerusakan pada komputer lokal tanpa cadangan.</td>
         <td>Lebih aman karena data disimpan di server jarak jauh.</td>
      </tr>
   </tbody>
</table>
<p>Penting untuk memahami perbedaan antara local dan remote repository dalam konteks pengembangan perangkat lunak. Pilihan antara keduanya tergantung pada kebutuhan proyek dan tim yang terlibat.</p>
<p>Dalam pengembangan perangkat lunak modern, kombinasi antara Local dan Remote Repository sangat penting. kamu bekerja pada Local Repository untuk mengembangkan fitur baru atau memperbaiki bug, dan kemudian menghubungkan dan berkolaborasi dengan kontributor lain melalui Remote Repository.</p>
<h2 id="354-konflik-dan-resolusi-konflik">3.5.4 Konflik dan Resolusi Konflik</h2>
<p>Dalam pengembangan perangkat lunak menggunakan sistem pengontrol versi, seperti Git, seringkali terjadi situasi di mana terjadi konflik antara perubahan yang dilakukan pada repositori lokal (local repository) dan repositori jarak jauh (remote repository). Konflik ini bisa terjadi ketika terdapat perubahan pada berkas yang sama oleh dua atau lebih kontributor, baik di repositori lokal maupun remote. Resolusi konflik adalah proses mengatasi perbedaan ini dan menggabungkan perubahan dengan benar.</p>
<p><strong>Konflik pada Materi Local vs Remote Repo:</strong></p>
<ol>
   <li>
      <p><strong>Perubahan Tidak Sinkron:</strong> Konflik sering terjadi ketika repositori lokal dan remote tidak lagi sinkron. Ini bisa terjadi ketika terdapat perubahan di repositori lokal yang belum diunggah ke remote, atau sebaliknya.</p>
   </li>
   <li>
      <p><strong>Perubahan Bersamaan:</strong> Konflik terjadi saat dua atau lebih kontributor membuat perubahan pada berkas yang sama pada waktu yang bersamaan. Ketika salah satu kontributor mengunggah perubahan ke remote, yang lain harus mengatasi konflik saat mereka mencoba untuk mengunggah perubahan mereka.</p>
   </li>
   <li>
      <p><strong>Perbedaan Implementasi:</strong> Konflik juga bisa terjadi ketika dua kontributor membuat perubahan yang bertentangan atau mengubah bagian yang sama dari kode. Ini bisa menghasilkan perbedaan dalam implementasi yang sulit untuk digabungkan.</p>
   </li>
</ol>
<p><strong>Resolusi Konflik pada Materi Local vs Remote Repo:</strong></p>
<ol>
   <li>
      <p><strong>Pemahaman Konflik:</strong> Langkah pertama adalah memahami sumber konflik dan perubahan yang terjadi. Hal ini melibatkan pemeriksaan kode yang bertabrakan dan pemahaman tentang bagaimana perubahan ini bisa saling mempengaruhi.</p>
   </li>
   <li>
      <p><strong>Penggabungan Manual:</strong> Jika ada konflik, Git tidak akan secara otomatis menggabungkan perubahan tersebut. Ini berarti kamu harus melakukan penggabungan manual dengan mengedit berkas yang terlibat. Pada bagian berkonflik, Git akan menunjukkan markah yang memungkinkan kamu untuk memilih perubahan mana yang harus tetap dan mana yang harus dihapus.</p>
   </li>
   <li>
      <p><strong>Pengujian Kembali:</strong> Setelah konflik diatasi dan perubahan digabungkan, penting untuk menguji kembali perangkat lunak untuk memastikan bahwa tidak ada masalah baru yang muncul akibat penggabungan tersebut.</p>
   </li>
   <li>
      <p><strong>Komit dan Unggah:</strong> Setelah penggabungan berhasil, kamu harus melakukan commit pada perubahan dan mengunggahnya kembali ke remote repository. Pastikan untuk memberikan pesan commit yang menjelaskan tentang penggabungan dan perubahan yang dilakukan.</p>
   </li>
   <li>
      <p><strong>Koordinasi Tim:</strong> Penting untuk berkoordinasi dengan anggota tim lainnya saat mengatasi konflik, terutama jika perubahan yang bersangkutan cukup kompleks. Komunikasi yang baik akan membantu mencegah konflik berlarut-larut dan memastikan bahwa semua kontributor memahami perubahan yang dilakukan.</p>
   </li>
</ol>
<p>Konflik adalah bagian alami dari pengembangan perangkat lunak dalam lingkungan kolaboratif. Resolusi konflik melibatkan pemahaman, penggabungan manual, pengujian, dan koordinasi untuk memastikan bahwa perubahan yang dilakukan oleh anggota tim dapat diintegrasikan dengan lancar ke dalam repositori. Dengan praktek terbaik dan komunikasi yang baik, konflik dapat diatasi dengan efisien tanpa mengganggu kelancaran pengembangan perangkat lunak.</p>
