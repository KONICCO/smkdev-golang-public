<div align="center">
   <h1>Module 2 Lesson 4 - Navigasi dan Manajemen Berkas</h1>
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
         <td>2.4.1 Manipulasi Berkas dan Direktori</td>
         <td><a href="#241-manipulasi-berkas-dan-direktori">🔽</a></td>
      </tr>
      <tr>
         <td>2.4.2 Perizinan Berkas dan Direktori</td>
         <td><a href="#242-perizinan-berkas-dan-direktori">🔽</a></td>
      </tr>
      <tr>
         <td>2.4.3 Piping dan Redirection</td>
         <td><a href="#243-piping-dan-redirection">🔽</a></td>
      </tr>
      <tr>
         <td>2.4.4 Absolute dan Relative Paths</td>
         <td><a href="#244-absolute-dan-relative-paths">🔽</a></td>
      </tr>
   </tbody>
</table>
<p><br><br></p>
<h1 id="durasi-pembelajaran">Durasi Pembelajaran</h1>
<p><strong>60 menit</strong></p>
<br>
<h2 id="241-manipulasi-berkas-dan-direktori">2.4.1 Manipulasi Berkas dan Direktori</h2>
<p>Navigasi dan manajemen berkas adalah proses penting dalam penggunaan sistem operasi yang memungkinkan kamu untuk bekerja dengan berkas dan direktori (folder) pada komputer. Manipulasi berkas dan direktori melibatkan tindakan seperti membuat, menghapus, memindahkan, menyalin, mengganti nama, dan mengatur izin akses terhadap berkas dan direktori. Berikut ini adalah beberapa operasi dasar yang terkait dengan manipulasi berkas dan direktori:</p>
<h3 id="membuat-berkas-atau-direktori">Membuat Berkas atau Direktori</h3>
<ul>
   <li>Membuat Berkas: kamu dapat membuat berkas kosong dengan menggunakan perintah seperti <code>touch</code> (Unix/Linux) atau <code>type nul &gt; filename</code> (Windows).</li>
   <li>Membuat Direktori: Untuk membuat direktori baru, kamu dapat menggunakan perintah seperti <code>mkdir</code> (Unix/Linux) atau <code>mkdir foldername</code> (Windows).</li>
</ul>
<h3 id="menghapus-berkas-atau-direktori">Menghapus Berkas atau Direktori</h3>
<ul>
   <li>Menghapus Berkas: kamu dapat menghapus berkas dengan menggunakan perintah <code>rm</code> (Unix/Linux) atau <code>del filename</code> (Windows).</li>
   <li>Menghapus Direktori: Menghapus direktori kosong dapat dilakukan menggunakan <code>rmdir</code> (Unix/Linux) atau <code>rmdir /s foldername</code> (Windows). Untuk menghapus direktori bersama dengan isinya, gunakan <code>rm -r</code> (Unix/Linux) atau <code>rmdir /s /q foldername</code> (Windows).</li>
</ul>
<h3 id="memindahkan-dan-menyalin-berkas-atau-direktori">Memindahkan dan Menyalin Berkas atau Direktori</h3>
<ul>
   <li>Memindahkan Berkas/Direktori: Perintah <code>mv</code> (Unix/Linux) atau <code>move</code> (Windows) digunakan untuk memindahkan berkas atau direktori dari satu lokasi ke lokasi lain.</li>
   <li>Menyalin Berkas/Direktori: Untuk menyalin berkas atau direktori, gunakan <code>cp</code> (Unix/Linux) atau <code>copy</code> (Windows).</li>
</ul>
<h3 id="mengganti-nama-berkas-atau-direktori">Mengganti Nama Berkas atau Direktori</h3>
<ul>
   <li>Mengganti Nama Berkas: kamu dapat mengganti nama berkas menggunakan perintah <code>mv</code> (Unix/Linux) atau <code>ren</code> (Windows).</li>
   <li>Mengganti Nama Direktori: Gunakan perintah <code>mv</code> (Unix/Linux) atau <code>ren</code> (Windows) untuk mengganti nama direktori juga.</li>
</ul>
<h3 id="izin-akses-dan-pengaturan-properti">Izin Akses dan Pengaturan Properti</h3>
<ul>
   <li>Izin Akses: Sistem operasi memberikan izin akses untuk berkas dan direktori kepada pengguna. Perintah seperti <code>chmod</code> (Unix/Linux) digunakan untuk mengubah izin akses pada sistem Unix/Linux.</li>
   <li>Pengaturan Properti: kamu dapat mengatur properti berkas atau direktori, seperti waktu modifikasi dan atribut lainnya dengan perintah seperti <code>chown</code> (Unix/Linux) atau menggunakan tampilan properti di lingkungan Windows.</li>
</ul>
<h3 id="menampilkan-isi-direktori">Menampilkan Isi Direktori</h3>
<ul>
   <li>kamu dapat melihat isi dari sebuah direktori dengan menggunakan perintah <code>ls</code> (Unix/Linux) atau <code>dir</code> (Windows).</li>
</ul>
<h3 id="path-dan-working-directory">Path dan Working Directory</h3>
<ul>
   <li>Path: Path adalah urutan direktori yang mengarah ke suatu berkas atau direktori. Absolute path adalah path lengkap dari root direktori, sedangkan relative path adalah path terhadap working directory saat ini.</li>
   <li>Working Directory: Working directory adalah direktori saat ini di mana kamu sedang berada dalam lingkungan terminal atau file manager.</li>
</ul>
<h2 id="242-perizinan-berkas-dan-direktori">2.4.2 Perizinan Berkas dan Direktori</h2>
<p>Dalam lingkup ini, perizinan berkas dan direktori adalah konsep yang penting untuk memahami bagaimana sistem operasi mengontrol akses pengguna terhadap berkas dan direktori tertentu. Berikut adalah penjelasan mengenai perizinan berkas dan direktori:</p>
<h3 id="perizinan-berkas">Perizinan Berkas</h3>
<p>Perizinan berkas adalah mekanisme yang digunakan oleh sistem operasi untuk mengatur hak akses pengguna terhadap suatu berkas. Setiap berkas memiliki metadata perizinan yang menentukan jenis akses yang diizinkan untuk setiap kelompok pengguna (biasanya pemilik, grup, dan lainnya). Terdapat tiga jenis izin utama yang dapat diberikan kepada suatu berkas:</p>
<ul>
   <li><strong>Read (Baca):</strong> Izin ini memungkinkan pengguna untuk membaca isi berkas.</li>
   <li><strong>Write (Tulis):</strong> Izin ini memungkinkan pengguna untuk mengubah atau menulis ke dalam berkas.</li>
   <li><strong>Execute (Jalankan):</strong> Izin ini memungkinkan pengguna untuk menjalankan atau melaksanakan berkas jika berkas tersebut merupakan program atau skrip.</li>
</ul>
<p>Kombinasi dari tiga izin di atas dinyatakan dalam bentuk kode numerik atau simbolik, seperti &quot;rw-r--r--&quot; (644) yang berarti pemilik memiliki izin baca dan tulis, sedangkan grup dan lainnya hanya memiliki izin baca.</p>
<h3 id="perizinan-direktori">Perizinan Direktori</h3>
<p>Perizinan direktori mengatur hak akses pengguna terhadap suatu direktori atau folder. Perbedaan utama antara perizinan direktori dengan perizinan berkas adalah bahwa izin pada direktori juga mempengaruhi akses terhadap berkas-berkas yang ada di dalamnya. Tiga jenis izin yang sama (baca, tulis, dan jalankan) juga diterapkan pada direktori, tetapi interpretasinya sedikit berbeda:</p>
<ul>
   <li><strong>Read (Baca) pada Direktori:</strong> Izin ini memungkinkan pengguna untuk melihat daftar berkas yang ada dalam direktori.</li>
   <li><strong>Write (Tulis) pada Direktori:</strong> Izin ini memungkinkan pengguna untuk menambahkan atau menghapus berkas dalam direktori.</li>
   <li><strong>Execute (Jalankan) pada Direktori:</strong> Izin ini memungkinkan pengguna untuk masuk ke dalam direktori dan menjalankan perintah di dalamnya, asalkan izin ini ada pada semua direktori yang berada di atasnya dalam hierarki struktur direktori.</li>
</ul>
<h3 id="pengaturan-perizinan">Pengaturan Perizinan</h3>
<p>Pengaturan perizinan pada berkas dan direktori biasanya dilakukan melalui perintah-perintah khusus di dalam terminal atau melalui antarmuka grafis. Dalam lingkungan berbasis Unix/Linux, perintah seperti <code>chmod</code> digunakan untuk mengubah perizinan, sedangkan perintah seperti <code>chown</code> digunakan untuk mengubah kepemilikan.</p>
<p>Pemahaman mengenai perizinan berkas dan direktori penting untuk menjaga keamanan data dan sistem. Dengan menggunakan perizinan dengan bijak, sistem operasi dapat melindungi berkas-berkas sensitif dari akses yang tidak sah, serta memastikan bahwa pengguna memiliki akses yang sesuai sesuai dengan tugas mereka.</p>
<h2 id="243-piping-dan-redirection">2.4.3 Piping dan Redirection</h2>
<p>Dalam sistem operasi berbasis teks, seperti Unix dan Linux, pengguna sering kali berurusan dengan perintah berkas dan manipulasi teks melalui terminal. Pada materi navigasi dan manajemen berkas, dua konsep penting yang perlu dipahami adalah piping (pipa) dan redirection (pengalihan).</p>
<h3 id="piping-pipa">Piping (Pipa)</h3>
<p>Piping adalah teknik yang memungkinkan output dari suatu perintah digunakan sebagai input untuk perintah lainnya tanpa harus menyimpan output di dalam berkas sementara. Tanda Piping (<code>|</code>) digunakan untuk menghubungkan perintah-perintah ini. Dengan Piping, kamu dapat menggabungkan serangkaian perintah untuk melakukan tugas yang lebih kompleks.</p>
<p>Contoh penggunaan Piping:</p>
<pre><code class="language-bash">cat file.txt | grep &quot;kata kunci&quot;
</code></pre>
<p>Dalam contoh di atas, perintah <code>cat</code> membaca isi dari berkas <code>file.txt</code>, kemudian outputnya digunakan sebagai input untuk perintah <code>grep</code>, yang akan mencari baris yang mengandung &quot;kata kunci&quot;.</p>
<h3 id="redirection-pengalihan">Redirection (Pengalihan)</h3>
<p>Redirection adalah teknik yang memungkinkan kamu mengarahkan input atau output dari suatu perintah ke atau dari berkas atau perangkat lain. Ada beberapa operator redirection yang umum digunakan:</p>
<ul>
   <li><code>&gt;</code> : Mengarahkan output ke berkas (mengganti isi berkas jika berkas sudah ada).</li>
   <li><code>&gt;&gt;</code> : Mengarahkan output dan menambahkannya ke akhir berkas (tanpa mengganti isi berkas).</li>
   <li><code>&lt;</code> : Menggunakan berkas sebagai input untuk perintah.</li>
   <li><code>2&gt;</code> : Mengarahkan output kesalahan (stderr) ke berkas.</li>
   <li><code>2&gt;&amp;1</code> : Mengarahkan output dan output kesalahan ke tempat yang sama.</li>
</ul>
<p>Contoh penggunaan redirection:</p>
<pre><code class="language-bash">ls &gt; daftar_berkas.txt
</code></pre>
<p>Perintah di atas akan mengarahkan output dari perintah <code>ls</code> ke dalam berkas <code>daftar_berkas.txt</code>.</p>
<h3 id="kombinasi-piping-dan-redirection">Kombinasi Piping dan Redirection</h3>
<p>kamu juga dapat menggabungkan teknik piping dan redirection untuk menciptakan alur kerja yang lebih kompleks. Misalnya:</p>
<pre><code class="language-bash">cat file.txt | grep &quot;kata kunci&quot; &gt; hasil_pencarian.txt
</code></pre>
<p>Dalam contoh di atas, output dari perintah <code>grep</code> akan diarahkan ke berkas <code>hasil_pencarian.txt</code>.</p>
<h2 id="244-absolute-dan-relative-paths">2.4.4 Absolute dan Relative Paths</h2>
<p>Saat bekerja dengan berkas dan direktori (folder), kamu akan sering kali perlu merujuk atau mengacu pada lokasi berkas dalam struktur sistem berkas. Dalam konteks ini, konsep <em>Absolute Paths</em> dan <em>Relative Paths</em> (jalur absolut dan jalur relatif) adalah konsep penting untuk memahami bagaimana berkas dan direktori diakses dan dikelola.</p>
<h3 id="absolute-paths-jalur-absolut">Absolute Paths (Jalur Absolut)</h3>
<p>Jalur absolut adalah alamat lengkap yang menunjukkan lokasi berkas atau direktori mulai dari akar sistem berkas. Jalur ini memberikan rute tepat untuk mencapai berkas atau direktori, tidak peduli dari mana kamu memulai. Jalur absolut umumnya dimulai dengan nama drive (pada Windows) atau &quot;/&quot; (pada sistem Unix-like, termasuk Linux dan macOS).</p>
<p>Contoh jalur absolut pada sistem Windows:</p>
<pre><code>C:\Users\Username\Documents\file.txt
</code></pre>
<p>Contoh jalur absolut pada sistem Unix-like:</p>
<pre><code>/home/username/Documents/file.txt
</code></pre>
<h3 id="relative-paths-jalur-relatif">Relative Paths (Jalur Relatif)</h3>
<p>Jalur relatif adalah alamat yang mengacu pada lokasi berkas atau direktori relatif terhadap lokasi saat ini. Dalam hal ini, &quot;lokasi saat ini&quot; adalah direktori tempat perintah atau skrip yang sedang dijalankan berada. Jalur relatif berguna ketika kamu ingin merujuk pada berkas atau direktori yang terletak dalam hierarki yang sama atau terkait dengan berkas atau direktori saat ini.</p>
<p>Ada dua jenis notasi umum dalam jalur relatif:</p>
<ul>
   <li>
      <p><code>./</code>: Mengacu pada direktori saat ini. Misalnya, <code>./file.txt</code> merujuk pada berkas &quot;file.txt&quot; dalam direktori saat ini.</p>
   </li>
   <li>
      <p><code>../</code>: Mengacu pada direktori di atasnya. Misalnya, <code>../folder/file.txt</code> merujuk pada berkas &quot;file.txt&quot; dalam direktori &quot;folder&quot; yang berada satu tingkat di atas direktori saat ini.</p>
   </li>
</ul>
<p>Contoh jalur relatif:</p>
<ul>
   <li>Jika perintah atau skrip dijalankan dari direktori <code>/home/username/</code>, maka <code>./Documents/file.txt</code> akan merujuk pada berkas &quot;file.txt&quot; dalam direktori &quot;Documents&quot; di dalam direktori saat ini.</li>
   <li>Jika perintah atau skrip dijalankan dari direktori <code>/home/username/Pictures/</code>, maka <code>../Documents/file.txt</code> akan merujuk pada berkas &quot;file.txt&quot; dalam direktori &quot;Documents&quot; yang berada satu tingkat di atasnya.</li>
</ul>
<p>Pemahaman tentang jalur absolut dan relatif sangat penting saat berurusan dengan navigasi dan manajemen berkas. Jalur absolut memberikan alamat lengkap dari akar sistem berkas, sementara jalur relatif memberikan alamat yang relatif terhadap lokasi saat ini. Pemilihan jalur yang tepat dapat membantu dalam menghindari kesalahan dan mempermudah pengelolaan berkas dan direktori.</p>