<div align="center">
   <h1>Module 3 Lesson 3 - Bekerja pada Local Repository</h1>
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
         <td>3.3.1 Staging Area</td>
         <td><a href="#331-staging-area">🔽</a></td>
      </tr>
      <tr>
         <td>3.3.2 Repo Status</td>
         <td><a href="#332-repo-status">🔽</a></td>
      </tr>
      <tr>
         <td>3.3.3 Commits dan Conventional Commits</td>
         <td><a href="#333-commits-dan-conventional-commits">🔽</a></td>
      </tr>
      <tr>
         <td>3.3.4 Push dan Pull request</td>
         <td><a href="#334-push-dan-pull-request">🔽</a></td>
      </tr>
   </tbody>
</table>
<p><br><br></p>
<h1 id="durasi-pembelajaran">Durasi Pembelajaran</h1>
<p><strong>60 menit</strong></p>
<br>
<h2 id="331-staging-area">3.3.1 Staging Area</h2>
<p>Staging area, juga dikenal sebagai &quot;index&quot; dalam beberapa konteks, adalah konsep penting dalam sistem kontrol versi terdistribusi seperti Git. Staging area merupakan langkah tambahan antara bekerja di direktori kerja (working directory) dan melakukan commit ke repositori lokal. Ini memungkinkan kamu untuk memilih perubahan mana yang akan dimasukkan ke dalam commit sebelum benar-benar melakukannya.</p>
<p>Berikut adalah penjelasan lebih lanjut mengenai staging area dalam konteks bekerja pada repositori lokal:</p>
<ol>
   <li>
      <p><strong>Working Directory (Direktori Kerja)</strong>:
         Direktori kerja adalah tempat di mana kamu melakukan perubahan pada berkas-berkas dalam proyek kamu. kamu bisa menambahkan, menghapus, atau mengubah isi berkas di sini sesuai kebutuhan.
      </p>
   </li>
   <li>
      <p><strong>Staging Area</strong>:
         Staging area adalah tempat di mana kamu mengumpulkan perubahan-perubahan yang ingin kamu sertakan dalam commit berikutnya. kamu dapat memilih perubahan-perubahan tertentu untuk disertakan, sambil mengabaikan perubahan lain yang mungkin tidak ingin kamu commit pada saat ini. Konsep ini memungkinkan kamu untuk mengontrol dengan cermat apa yang akan menjadi bagian dari histori versi kamu.
      </p>
   </li>
   <li>
      <p><strong>Commit</strong>:
         Commit adalah tindakan untuk menyimpan perubahan-perubahan yang ada dalam staging area ke dalam sejarah versi repositori kamu. Setiap commit memiliki pesan yang menjelaskan perubahan yang dilakukan, membuatnya lebih mudah untuk melacak sejarah proyek dan memahami evolusi kode.
      </p>
   </li>
</ol>
<p>Proses kerja dengan staging area dalam repositori lokal adalah sebagai berikut:</p>
<ol>
   <li>
      <p><strong>Mengedit Berkas</strong>:
         kamu melakukan perubahan pada berkas-berkas dalam direktori kerja sesuai kebutuhan.
      </p>
   </li>
   <li>
      <p><strong>Menambahkan Perubahan ke Staging Area</strong>:
         Setelah kamu puas dengan perubahan-perubahan yang kamu lakukan pada berkas, kamu harus menambahkan perubahan-perubahan ini ke staging area. Ini dapat dilakukan dengan menggunakan perintah <code>git add</code> di terminal. Misalnya, jika kamu ingin menambahkan semua perubahan pada berkas yang telah diubah, kamu dapat menjalankan <code>git add .</code>.
      </p>
   </li>
   <li>
      <p><strong>Memeriksa Staging Area</strong>:
         kamu dapat memeriksa perubahan-perubahan yang ada di staging area menggunakan perintah <code>git status</code>. Ini akan menampilkan perubahan-perubahan yang telah ditambahkan dan yang belum.
      </p>
   </li>
</ol>
<p>Staging area membantu kamu mengelola perubahan dengan lebih efisien dan memberikan fleksibilitas dalam memilih perubahan mana yang akan dimasukkan dalam commit. Hal ini juga memungkinkan kamu untuk memisahkan perubahan logis menjadi beberapa commit yang terpisah, sehingga histori proyek tetap bersih dan mudah dipahami.</p>
<h2 id="332-repo-status">3.3.2 Repo Status</h2>
<p>Dalam konteks pengembangan perangkat lunak dengan menggunakan sistem kontrol versi seperti Git, &quot;Repo Status&quot; mengacu pada informasi tentang perubahan-perubahan yang terjadi pada repositori lokal kamu. Repositori lokal adalah salinan dari proyek kamu yang disimpan di komputer kamu sendiri. Ketika kamu bekerja pada proyek dan melakukan perubahan pada file-file dalam repositori tersebut, status repositori akan memberi kamu gambaran tentang status perubahan-perubahan tersebut.</p>
<p>Repo status memberikan informasi penting kepada kamu, seperti file mana yang telah dimodifikasi, file yang baru ditambahkan, dan file yang dihapus. Ini memungkinkan kamu untuk melacak dan mengelola perubahan-perubahan tersebut sebelum kamu memutuskan untuk mengirimkan (commit) perubahan-perubahan tersebut ke repositori pusat atau berbagi dengan tim kamu.</p>
<p>Berikut adalah beberapa komponen utama dari repo status dalam kerangka kerja Git:</p>
<ol>
   <li>
      <p><strong>Perubahan yang Dimodifikasi (Modified Changes):</strong> Ini mencakup file-file yang telah kamu ubah tetapi belum di-commit. Git akan memberi tahu kamu tentang perubahan ini dan memungkinkan kamu untuk memutuskan apakah perubahan ini akan dimasukkan ke dalam commit berikutnya.</p>
   </li>
   <li>
      <p><strong>File yang Baru Ditambahkan (Staged Changes):</strong> Setelah kamu mengubah file, kamu dapat memilih file-file tertentu yang ingin kamu masukkan ke dalam commit berikutnya. Proses ini disebut &quot;staging&quot;. File-file yang telah di-staging akan masuk ke dalam repo status sebagai perubahan yang siap untuk di-commit.</p>
   </li>
   <li>
      <p><strong>File yang Dihapus (Deleted Changes):</strong> Jika kamu menghapus file dari proyek kamu, Git akan mendeteksinya dan menampilkan informasi ini dalam repo status. Ini memberi kamu kesempatan untuk memutuskan apakah kamu ingin menghapus file ini dari repositori secara permanen dalam commit berikutnya.</p>
   </li>
   <li>
      <p><strong>Perubahan yang Tidak Terlacak (Untracked Changes):</strong> Ini merujuk pada file-file yang ada dalam direktori proyek kamu tetapi belum ditambahkan ke dalam repositori. Git tidak melacak perubahan pada file-file ini secara default. kamu harus menjalankan perintah &quot;git add&quot; untuk mengawasi perubahan pada file yang tidak terlacak ini.</p>
   </li>
   <li>
      <p><strong>Status Branch Aktif:</strong> Repo status juga akan memberi tahu kamu cabang mana yang sedang aktif. Cabang ini adalah cabang yang kamu kerjakan dan akan menerima commit berikutnya.</p>
   </li>
</ol>
<p>Untuk melihat repo status, kamu dapat menjalankan perintah <code>git status</code> di terminal di dalam direktori repositori kamu. Ini akan menampilkan ringkasan perubahan-perubahan yang terjadi dalam repo lokal kamu, serta memberi kamu petunjuk tentang apa yang harus dilakukan selanjutnya, seperti menambahkan perubahan ke staging area atau melakukan commit.</p>
<p>Repo status adalah alat yang sangat penting untuk mengelola perubahan dalam proyek kamu dan membantu kamu menjaga repositori lokal kamu tetap rapi sebelum kamu mengirimkan perubahan-perubahan tersebut ke repositori pusat atau berbagi dengan tim kamu.</p>
<h2 id="333-commits-dan-conventional-commits">3.3.3 Commits dan Conventional Commits</h2>
<p>Dalam pengembangan perangkat lunak, penggunaan sistem kontrol versi seperti Git sangat penting untuk mengelola perubahan pada kode sumber. Saat bekerja dengan repositori lokal, konsep &quot;Commits&quot; dan &quot;Conventional Commits&quot; memegang peranan penting dalam menjaga sejarah perubahan dan kolaborasi tim yang efisien.</p>
<h3 id="commits">Commits</h3>
<p><strong>Commits</strong> adalah tindakan merekam perubahan pada repositori Git. Setiap commit merepresentasikan satu atau lebih perubahan yang telah dilakukan pada berkas-berkas dalam repositori. Commits memungkinkan kamu untuk:</p>
<ol>
   <li>
      <p><strong>Mengelola Sejarah Perubahan:</strong> Setiap commit memiliki pesan yang menjelaskan apa yang telah diubah atau ditambahkan. Ini membantu dalam melacak evolusi kode dan memahami mengapa perubahan tertentu dilakukan.</p>
   </li>
   <li>
      <p><strong>Pemulihan Kode:</strong> Jika terjadi kesalahan atau perlu kembali ke versi sebelumnya, commit memungkinkan kamu untuk dengan mudah mengembalikan kode ke status sebelumnya.</p>
   </li>
   <li>
      <p><strong>Kolaborasi Tim:</strong> Tim yang bekerja pada proyek yang sama dapat berbagi perubahan melalui commit. Setiap anggota tim dapat melihat perubahan yang dilakukan oleh anggota lain.</p>
   </li>
   <li>
      <p><strong>Pengujian dan Peninjauan Kode:</strong> Commits dapat diuji dan ditinjau sebelum dimasukkan ke dalam cabang utama (misalnya, <em>main</em> atau <em>master</em>) melalui <em>pull request</em> atau mekanisme serupa.</p>
   </li>
</ol>
<h3 id="conventional-commits">Conventional Commits</h3>
<p><strong>Conventional Commits</strong> adalah suatu konvensi atau pendekatan dalam menulis pesan commit yang terstruktur dan konsisten. Tujuannya adalah untuk meningkatkan kejelasan dan membantu otomatisasi dalam proses penggabungan kode. Struktur pesan commit dalam pendekatan ini biasanya terdiri dari tiga bagian: tipe, cakupan, dan deskripsi.</p>
<p>Contoh struktur pesan commit dalam Conventional Commits:</p>

<pre><code class="language-bash">&lt;type&gt;(&lt;scope&gt;): &lt;description&gt;
[optional body]
[optional footer]
</code></pre>

<ul>
   <li><strong>Tipe (Type):</strong> Menunjukkan jenis perubahan yang dilakukan, seperti &quot;feat&quot; (fitur baru), &quot;fix&quot; (perbaikan bug), &quot;chore&quot; (tugas umum), dll.</li>
   <li><strong>Cakupan (Scope):</strong> Menunjukkan bagian kode atau komponen yang terpengaruh oleh perubahan.</li>
   <li><strong>Deskripsi (Description):</strong> Deskripsi singkat tentang apa yang telah dilakukan dalam commit.</li>
   <li><strong>Body:</strong> Bagian opsional untuk memberikan penjelasan lebih lanjut tentang perubahan.</li>
   <li><strong>Footer:</strong> Bagian opsional untuk menghubungkan commit dengan nomor <em>issue</em> atau tindakan lain.</li>
</ul>
<p>Contoh pesan commit menggunakan Conventional Commits:</p>

<pre><code>feat(user-auth): add user registration functionality
This commit adds the user registration API endpoint and updates the database schema accordingly.
Closes #123
</code></pre>

<p>Jenis-jenis Conventional Commits yang umum digunakan meliputi:</p>
<ol>
   <li>
      <p><strong>feat</strong>: Digunakan ketika menambahkan fitur baru ke dalam proyek. Contoh: <code>feat: tambahkan fitur login dengan OAuth</code>.</p>
   </li>
   <li>
      <p><strong>fix</strong>: Digunakan ketika memperbaiki sebuah bug atau masalah yang ada dalam kode. Contoh: <code>fix: perbaiki bug tampilan pada halaman utama</code>.</p>
   </li>
   <li>
      <p><strong>chore</strong>: Digunakan untuk perubahan yang tidak berkaitan dengan kode sumber utama, seperti pembaruan konfigurasi, peningkatan alat, atau tugas administratif lainnya. Contoh: <code>chore: perbarui dependensi versi XYZ</code>.</p>
   </li>
   <li>
      <p><strong>docs</strong>: Digunakan ketika melakukan perubahan pada dokumentasi, baik itu dokumentasi kode, dokumentasi pengguna, atau yang lainnya. Contoh: <code>docs: perbarui panduan instalasi</code>.</p>
   </li>
   <li>
      <p><strong>style</strong>: Digunakan untuk perubahan yang berkaitan dengan gaya kode, seperti perubahan tata letak, indentasi, atau format lainnya. Contoh: <code>style: perbaiki indentasi pada file X</code>.</p>
   </li>
   <li>
      <p><strong>refactor</strong>: Digunakan ketika melakukan perubahan dalam kode yang tidak memperbaiki bug atau menambahkan fitur baru, tetapi melakukan restrukturisasi atau perubahan dalam desain yang lebih baik. Contoh: <code>refactor: optimalkan fungsi hitung() untuk performa lebih baik</code>.</p>
   </li>
   <li>
      <p><strong>test</strong>: Digunakan ketika menambahkan atau mengubah pengujian (testing), baik itu pengujian unit, integrasi, atau lainnya. Contoh: <code>test: tambahkan pengujian integrasi untuk komponen Y</code>.</p>
   </li>
   <li>
      <p><strong>perf</strong>: Digunakan ketika melakukan perubahan untuk meningkatkan performa. Contoh: <code>perf: optimalkan algoritma pencarian</code>.</p>
   </li>
   <li>
      <p><strong>ci</strong>: Digunakan untuk perubahan pada konfigurasi atau skrip yang berkaitan dengan alat otomatisasi (Continuous Integration) dan penerapan (Continuous Deployment). Contoh: <code>ci: perbarui skrip otomasi deploy</code>.</p>
   </li>
   <li>
      <p><strong>build</strong>: Digunakan ketika melakukan perubahan pada konfigurasi build, seperti penambahan atau penghapusan dependensi build. Contoh: <code>build: perbarui versi kompiler ke 2.0</code>.</p>
   </li>
   <li>
      <p><strong>revert</strong>: Digunakan ketika mengembalikan perubahan sebelumnya dengan mengembalikan keadaan kode ke komit sebelumnya. Contoh: <code>revert: revert perubahan pada file Z</code>.</p>
   </li>
</ol>
<p>Dengan mengikuti konvensi Conventional Commits, kamu dapat dengan mudah melihat jenis perubahan yang terjadi dalam setiap komit hanya dari pesan komitnya. Ini sangat membantu dalam memahami sejarah pengembangan proyek, menghasilkan catatan rilis yang terstruktur, dan memudahkan kerjasama antara anggota tim pengembangan.</p>
<p>Keuntungan dari menggunakan Conventional Commits adalah:</p>
<ul>
   <li>Mudah dipahami dan dikelola oleh manusia.</li>
   <li>Dapat digunakan untuk menghasilkan catatan perubahan yang terstruktur dan mudah dipahami oleh pengguna atau pelanggan.</li>
   <li>Mempermudah otomatisasi dalam menghasilkan catatan perubahan, memicu alur kerja otomatis, dan menghasilkan versi perangkat lunak.</li>
</ul>
<p>Penerapan Conventional Commits juga dapat mendukung alat-alat seperti <a href="https://semver.org/">Semantic Versioning (SemVer)</a>, yang membantu dalam memberi label versi perangkat lunak berdasarkan jenis perubahan yang dilakukan.</p>
<h2 id="334-push-dan-pull-request">3.3.4 Push dan Pull request</h2>
<p>Dalam pengembangan perangkat lunak, penggunaan sistem kontrol versi sangat penting untuk mengelola perubahan pada kode sumber. Git adalah salah satu sistem kontrol versi yang populer, dan dalam bekerja dengan Git, terdapat konsep &quot;push&quot; dan &quot;pull request&quot; yang berkaitan dengan mengirimkan dan menerima perubahan antara repositori lokal dan repositori jarak jauh (remote repository).</p>
<p><strong>1. Push:</strong>
   Push adalah aksi mengirimkan perubahan yang telah dilakukan di repositori lokal ke repositori jarak jauh (remote repository), seperti GitHub atau GitLab. Dengan melakukan push, perubahan yang telah kamu komit di repositori lokal akan terlihat dan tersimpan di remote repository. Langkah-langkah umum dalam melakukan push adalah sebagai berikut:
</p>
<ul>
   <li>
      <p><strong>Commit Perubahan:</strong> Pertama, kamu perlu melakukan komit (commit) untuk menyimpan perubahan di repositori lokal. Ini dapat dilakukan dengan perintah <code>git commit -m &quot;Pesan komit&quot;</code>.</p>
   </li>
   <li>
      <p><strong>Push ke Remote:</strong> Setelah kamu melakukan komit, perubahan tersebut masih hanya ada di repositori lokal. Untuk mengirimnya ke remote repository, gunakan perintah <code>git push nama_remote nama_cabang</code>. Misalnya, <code>git push origin main</code> akan mengirim perubahan dari cabang &quot;main&quot; ke remote repository dengan nama &quot;origin&quot;.</p>
   </li>
</ul>
<p><strong>2. Pull Request (PR):</strong>
   Pull Request adalah konsep yang terkait dengan kerjasama dalam pengembangan perangkat lunak yang menggunakan Git. Ini lebih umum dalam kerja tim dan pengembangan open source. PR memungkinkan kontributor untuk mengajukan perubahan mereka dari cabang di repositori fork (duplikat repositori) ke repositori asli (upstream repository). Langkah-langkah dalam membuat pull request adalah sebagai berikut:
</p>
<ul>
   <li>
      <p><strong>Fork Repositori:</strong> Buatlah salinan (fork) dari repositori asli ke akun GitHub/GitLab kamu.</p>
   </li>
   <li>
      <p><strong>Clone Repositori Fork:</strong> Lakukan cloning (pengunduhan) dari repositori fork ke lokal menggunakan perintah <code>git clone URL_fork</code>.</p>
   </li>
   <li>
      <p><strong>Buat Cabang dan Lakukan Perubahan:</strong> Buatlah cabang baru di repositori fork untuk fitur/perubahan yang ingin kamu buat. Lakukan perubahan pada cabang ini dengan melakukan komit.</p>
   </li>
   <li>
      <p><strong>Push Cabang ke Repositori Fork:</strong> Setelah kamu puas dengan perubahan di cabang kamu, lakukan push cabang tersebut ke repositori fork di akun kamu di GitHub/GitLab.</p>
   </li>
   <li>
      <p><strong>Buat Pull Request:</strong> Di halaman repositori asli, kamu dapat membuat Pull Request. Ini berfungsi sebagai permintaan kepada pemilik repositori asli untuk memasukkan perubahan dari cabang kamu ke repositori asli. Pada tahap ini, tim pengembang bisa melakukan diskusi dan mengevaluasi perubahan kamu sebelum diintegrasikan.</p>
   </li>
   <li>
      <p><strong>Review dan Merge:</strong> Pemilik repositori asli atau pemilik proyek akan melakukan review terhadap perubahan kamu. Jika perubahan dianggap baik, mereka akan menggabungkan (merge) perubahan kamu ke repositori asli.</p>
   </li>
</ul>
<p>Push dan Pull Request adalah konsep penting dalam pengelolaan perubahan kode sumber dalam pengembangan perangkat lunak menggunakan Git. Push digunakan untuk mengirimkan perubahan dari repositori lokal ke repositori jarak jauh, sementara Pull Request digunakan untuk mengajukan perubahan dari repositori fork ke repositori asli dengan tujuan untuk diintegrasikan setelah melalui proses review.</p>