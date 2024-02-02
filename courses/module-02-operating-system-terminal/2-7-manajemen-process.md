<div align="center">
   <h1>Module 2 Lesson 7 - Manajemen Proses</h1>
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
         <td>2.7.1 Pengantar</td>
         <td><a href="#271-pengantar">🔽</a></td>
      </tr>
      <tr>
         <td>2.7.2 Arsitektur Proses</td>
         <td><a href="#272-arsitektur-proses">🔽</a></td>
      </tr>
      <tr>
         <td>2.7.3 Process Control Blocks</td>
         <td><a href="#273-process-control-blocks">🔽</a></td>
      </tr>
      <tr>
         <td>2.7.4 Process States dan Creation</td>
         <td><a href="#274-process-states-dan-creation">🔽</a></td>
      </tr>
   </tbody>
</table>
<p><br><br></p>
<h1 id="durasi-pembelajaran">Durasi Pembelajaran</h1>
<p><strong>60 menit</strong></p>
<br>
<h2 id="271-pengantar">2.7.1 Pengantar</h2>
<p>Selamat datang dalam modul Sistem Operasi dan Terminal! Dalam modul ini, kita akan membahas konsep-konsep penting dalam manajemen proses dalam sistem operasi. Proses merupakan inti dari pengoperasian sistem komputer, dan pemahaman yang kuat tentang bagaimana sistem operasi mengelola dan mengendalikan proses sangatlah krusial.</p>
<h3 id="1-arsitektur-proses">1. Arsitektur Proses**</h3>
<p>Arsitektur proses merupakan dasar dari pemahaman kita tentang bagaimana proses diimplementasikan dan dijalankan dalam sistem operasi. Kami akan membahas bagaimana sistem operasi memandu eksekusi proses, bagaimana alokasi sumber daya dilakukan untuk setiap proses, dan bagaimana proses berinteraksi satu sama lain. Dalam pembahasan ini, kamu akan mempelajari konsep seperti <strong>konteks proses</strong>, <strong>area data proses</strong>, dan <strong>peralihan antar proses</strong>.</p>
<h3 id="2-process-control-blocks-pcbs">2. Process Control Blocks (PCBs)**</h3>
<p>Process Control Blocks (PCBs) adalah struktur data yang sangat penting dalam manajemen proses. PCB mengandung informasi yang diperlukan oleh sistem operasi untuk mengendalikan dan melacak setiap proses secara efisien. Kami akan menggali lebih dalam tentang informasi apa yang biasanya ada dalam PCB, seperti <strong>identifikasi proses</strong>, <strong>status proses</strong>, <strong>prioritas</strong>, dan <strong>counter waktu</strong>. Pemahaman tentang PCB akan membantu kamu menggambarkan bagaimana sistem operasi memelihara informasi tentang semua proses yang sedang berjalan.</p>
<h3 id="3-process-states-dan-creation">3. Process States dan Creation**</h3>
<p>Proses dapat berada dalam berbagai status atau keadaan, seperti berjalan, siap, atau tertunda. Kami akan membahas dengan rinci tentang <strong>berbagai status proses</strong> dan bagaimana transisi antara status ini terjadi. Selain itu, kamu akan memahami bagaimana proses diciptakan dan dimulai oleh sistem operasi. Konsep seperti <strong>forking</strong> dan <strong>executing</strong> akan menjadi bagian dari pembahasan ini.</p>
<p>Dalam modul ini, kami akan memberikan penjelasan yang komprehensif tentang konsep-konsep ini dengan contoh-contoh yang jelas dan relevan. Pemahaman yang baik tentang manajemen proses dalam sistem operasi akan memberi kamu wawasan yang kuat tentang bagaimana sistem komputer bekerja secara lebih mendalam. Mari mulai perjalanan kita dalam memahami inti dari operasi sistem komputer melalui pemahaman yang lebih mendalam tentang manajemen proses.</p>
<h2 id="272-arsitektur-proses">2.7.2 Arsitektur Proses</h2>
<p>Dalam konteks sistem operasi, sebuah proses adalah sebuah entitas yang merepresentasikan program yang sedang dieksekusi. Proses adalah unit dasar dari eksekusi dalam lingkungan sistem operasi. Manajemen proses merupakan bagian kritis dari sistem operasi yang bertanggung jawab atas pengaturan, koordinasi, dan eksekusi proses.</p>
<h3 id="arsitektur-proses">Arsitektur Proses</h3>
<p>Arsitektur proses merujuk pada struktur internal dari sebuah proses. Ini mencakup elemen-elemen penting yang memungkinkan sistem operasi mengelola proses dengan efisien. Berikut adalah komponen utama dalam arsitektur proses:</p>
<ol>
   <li>
      <p><strong>Program Counter (PC):</strong> Program Counter adalah register khusus yang menyimpan alamat instruksi berikutnya yang akan dieksekusi dalam proses. Saat sebuah instruksi dieksekusi, nilai PC akan diperbarui untuk menunjuk ke instruksi berikutnya.</p>
   </li>
   <li>
      <p><strong>Registers:</strong> Registers adalah penyimpanan cepat yang terdapat dalam unit pemrosesan pusat (CPU) yang digunakan untuk menyimpan data sementara dan hasil perhitungan dalam proses.</p>
   </li>
   <li>
      <p><strong>Stack:</strong> Stack adalah area memori yang digunakan untuk menyimpan data sementara, seperti variabel lokal dan alamat pengembalian, dalam panggilan fungsi. Ini memungkinkan manajemen yang efisien dari pemanggilan dan pengembalian fungsi.</p>
   </li>
   <li>
      <p><strong>Data Section:</strong> Bagian data dari memori proses digunakan untuk menyimpan variabel global dan data statis lainnya yang diperlukan oleh proses.</p>
   </li>
   <li>
      <p><strong>Code Section:</strong> Bagian kode dari memori proses berisi instruksi-instruksi program yang akan dieksekusi. Ini mencakup kode dari program asli dan kode yang dihasilkan oleh kompilator.</p>
   </li>
   <li>
      <p><strong>Heap:</strong> Heap adalah area memori dinamis di mana alokasi memori dilakukan untuk data yang diciptakan secara dinamis pada saat runtime, seperti alokasi memori untuk struktur data dinamis.</p>
   </li>
   <li>
      <p><strong>File Descriptors:</strong> File descriptors adalah mekanisme yang digunakan untuk mewakili file yang terbuka dalam proses. Setiap file descriptor memiliki nomor unik yang mengidentifikasikan file tersebut.</p>
   </li>
</ol>
<p><strong>Transisi Proses:</strong>
   Proses dapat berada dalam beberapa status yang merefleksikan tahapan eksekusinya. Beberapa status umum meliputi &quot;Running&quot; (sedang dieksekusi), &quot;Ready&quot; (siap dieksekusi), dan &quot;Blocked&quot; (terhalang). Transisi antara status ini terjadi saat proses di-schedule oleh manajer proses sistem operasi.
</p>
<p><strong>Koordinasi Proses:</strong>
   Manajer proses sistem operasi bertanggung jawab atas koordinasi antara berbagai proses. Ini termasuk pembagian waktu pemrosesan CPU, alokasi sumber daya, dan penjadwalan proses.
</p>
<p><strong>Komunikasi Proses:</strong>
   Proses mungkin perlu berkomunikasi satu sama lain, baik dalam lingkungan yang serempak maupun asinkron. Sistem operasi menyediakan mekanisme seperti shared memory, pipes, sockets, dan lainnya untuk mendukung komunikasi antarproses.
</p>
<p>Arsitektur proses adalah dasar dari manajemen proses dalam sistem operasi. Ini mencakup struktur internal dari proses, termasuk program counter, registers, stack, data section, code section, heap, dan file descriptors. Memahami arsitektur proses memungkinkan pengembang dan administrator sistem operasi mengelola proses dengan efisien dan mengoptimalkan kinerja sistem.</p>
<h2 id="273-process-control-blocks">2.7.3 Process Control Blocks</h2>
<p>Process Control Block (PCB) merupakan struktur data yang digunakan dalam sistem operasi untuk menyimpan informasi penting mengenai suatu proses yang sedang berjalan di sistem. PCB berfungsi sebagai representasi dari proses dalam sistem operasi dan berisi berbagai informasi yang diperlukan untuk mengelola dan mengawasi proses tersebut.</p>
<h3 id="informasi-yang-disimpan-dalam-pcb">Informasi yang Disimpan dalam PCB</h3>
<p>PCB mengandung informasi penting yang diperlukan oleh sistem operasi untuk mengontrol dan mengawasi eksekusi suatu proses. Beberapa informasi yang biasanya disimpan dalam PCB antara lain:</p>
<ol>
   <li>
      <p><strong>Process Identifier (PID):</strong> Sebuah unik identifier numerik yang digunakan untuk mengidentifikasi proses secara unik di dalam sistem.</p>
   </li>
   <li>
      <p><strong>Program Counter (PC):</strong> Nilai counter yang menunjukkan alamat instruksi berikutnya yang akan dieksekusi dalam program.</p>
   </li>
   <li>
      <p><strong>Registers:</strong> Nilai-nilai dari register prosesor (seperti nilai akumulator, counter, dll.) yang diperlukan untuk menjaga konteks proses.</p>
   </li>
   <li>
      <p><strong>State Information:</strong> Informasi mengenai status proses, seperti apakah proses sedang berjalan, siap dieksekusi, atau sedang dalam keadaan tidur.</p>
   </li>
   <li>
      <p><strong>Priority:</strong> Prioritas proses, digunakan dalam penjadwalan proses untuk menentukan urutan eksekusi.</p>
   </li>
   <li>
      <p><strong>Memory Management Information:</strong> Informasi mengenai lokasi memori yang digunakan oleh proses, termasuk base register, limit register, dll.</p>
   </li>
   <li>
      <p><strong>Open Files:</strong> Daftar file yang sedang dibuka oleh proses.</p>
   </li>
   <li>
      <p><strong>Parent and Child Process Information:</strong> Informasi mengenai hubungan antara proses induk dan proses anak.</p>
   </li>
   <li>
      <p><strong>Accounting Information:</strong> Data mengenai penggunaan sumber daya oleh proses, seperti jumlah waktu CPU yang digunakan, jumlah instruksi dieksekusi, dll.</p>
   </li>
</ol>
<p><strong>Fungsi PCB dalam Manajemen Proses:</strong></p>
<ol>
   <li>
      <p><strong>Process Scheduling:</strong> PCB memungkinkan sistem operasi untuk mengelola antrian proses yang siap dieksekusi dan memutuskan urutan eksekusi berdasarkan informasi prioritas dan status.</p>
   </li>
   <li>
      <p><strong>Context Switching:</strong> Saat sistem beralih dari satu proses ke proses lainnya, informasi konteks proses saat ini disimpan dalam PCB dan digantikan dengan informasi konteks proses berikutnya.</p>
   </li>
   <li>
      <p><strong>Process Termination:</strong> Saat proses selesai dieksekusi atau dihentikan, informasi yang diperlukan untuk membersihkan sumber daya dan menghapus PCB dari sistem disimpan dalam PCB.</p>
   </li>
   <li>
      <p><strong>Resource Management:</strong> PCB menyimpan informasi mengenai alokasi sumber daya kepada proses, seperti memori, file yang sedang dibuka, dan status lainnya.</p>
   </li>
   <li>
      <p><strong>Communication and Synchronization:</strong> PCB bisa digunakan untuk menyimpan informasi yang diperlukan untuk sinkronisasi dan komunikasi antara proses, seperti mutex, semaphore, dan lainnya.</p>
   </li>
</ol>
<p>Process Control Blocks (PCBs) adalah struktur data yang vital dalam manajemen proses dalam sistem operasi. Mereka mengandung informasi yang diperlukan untuk mengontrol, mengawasi, dan menjalankan proses secara efisien. PCBs memainkan peran penting dalam menjaga integritas konteks proses saat terjadi pergantian antar proses dan dalam menyediakan informasi yang dibutuhkan oleh sistem operasi untuk mengelola sumber daya dan interaksi antar proses.</p>
<h2 id="274-process-states-dan-creation">2.7.4 Process States dan Creation</h2>
<p>Dalam sistem operasi, sebuah proses merupakan sebuah unit eksekusi yang independen, termasuk program yang sedang dieksekusi bersama dengan sumber daya yang terkait seperti memori, file, dan penghitung. Manajemen proses merujuk pada cara sistem operasi mengelola dan mengawasi proses-proses ini selama berjalannya program di dalam sistem.</p>
<h3 id="process-states-keadaan-proses">Process States (Keadaan Proses)</h3>
<p>Setiap proses dalam sistem operasi dapat berada dalam salah satu dari beberapa keadaan proses yang berbeda, yang menggambarkan fase di mana proses tersebut berada selama eksekusinya. Berikut adalah beberapa keadaan proses umum:</p>
<ol>
   <li>
      <p><strong>New (Baru):</strong> Proses berada dalam keadaan ini saat sedang dibuat. Pada titik ini, sistem operasi telah mengalokasikan sumber daya awal untuk proses tersebut, tetapi proses ini belum dimulai eksekusinya.</p>
   </li>
   <li>
      <p><strong>Ready (Siap):</strong> Setelah proses dibuat dan diberi sumber daya awal, ia berpindah ke keadaan siap. Ini berarti proses siap untuk dieksekusi oleh CPU, tetapi CPU belum mengalokasikan waktunya.</p>
   </li>
   <li>
      <p><strong>Running (Berjalan):</strong> Proses berada dalam keadaan ini saat CPU sedang mengeksekusi instruksi-instruksi dari programnya. Pada satu waktu, hanya ada satu proses yang sedang berjalan pada satu CPU.</p>
   </li>
   <li>
      <p><strong>Blocked (Terblokir):</strong> Proses dapat berpindah ke keadaan terblokir ketika ia menunggu sumber daya eksternal seperti masukan dari pengguna atau hasil dari operasi I/O. Proses dalam keadaan ini tidak dapat berjalan sampai kondisi yang memblokirnya terpenuhi.</p>
   </li>
   <li>
      <p><strong>Terminated (Berakhir):</strong> Proses berpindah ke keadaan ini setelah menyelesaikan eksekusi programnya atau karena alasan lain seperti kesalahan fatal. Pada titik ini, sumber daya yang dialokasikan untuk proses tersebut akan dilepaskan kembali ke sistem.</p>
   </li>
</ol>
<h3 id="creation-pembuatan-proses">Creation (Pembuatan Proses)</h3>
<p>Proses dibuat oleh sistem operasi untuk menjalankan program. Langkah-langkah utama dalam pembuatan proses adalah sebagai berikut:</p>
<ol>
   <li>
      <p><strong>Fork:</strong> Proses baru dapat diciptakan dengan menggandakan proses yang sudah ada melalui operasi &quot;fork&quot;. Proses baru ini disebut &quot;proses anak&quot;, sedangkan proses yang menggandakannya disebut &quot;proses induk&quot;. Proses anak memiliki salinan kode, data, dan sumber daya dari proses induk.</p>
   </li>
   <li>
      <p><strong>Exec:</strong> Setelah proses anak tercipta melalui operasi fork, langkah selanjutnya adalah menggantikan kode dan data proses anak dengan kode dan data program yang baru akan dieksekusi. Operasi ini disebut &quot;exec&quot;.</p>
   </li>
   <li>
      <p><strong>Wait:</strong> Proses induk dapat menunggu proses anaknya selesai eksekusi dengan menggunakan operasi &quot;wait&quot;. Ini memungkinkan proses induk untuk menangani hasil dari proses anak setelah eksekusinya selesai.</p>
   </li>
</ol>
<p>Manajemen proses dalam sistem operasi melibatkan pengelolaan berbagai keadaan proses dan proses penciptaan yang diperlukan untuk menjalankan program-program. Dengan mengelola keadaan proses dengan tepat, sistem operasi memungkinkan eksekusi yang efisien dan penggunaan sumber daya yang optimal dalam lingkungan multi-tugas.</p>