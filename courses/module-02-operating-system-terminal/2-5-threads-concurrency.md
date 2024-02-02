<div align="center">
   <h1>Module 2 Lesson 5 - Threads dan Concurrency</h1>
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
         <td>2.5.1 Pengantar</td>
         <td><a href="#251-pengantar">🔽</a></td>
      </tr>
      <tr>
         <td>2.5.2 Program, Processes, Threads</td>
         <td><a href="#252-program-processes-dan-threads">🔽</a></td>
      </tr>
      <tr>
         <td>2.5.3 Concurrency</td>
         <td><a href="#253-concurrency">🔽</a></td>
      </tr>
      <tr>
         <td>2.5.4 Parallelism</td>
         <td><a href="#254-parallelismi">🔽</a></td>
      </tr>
   </tbody>
</table>
<p><br><br></p>
<h1 id="durasi-pembelajaran">Durasi Pembelajaran</h1>
<p><strong>60 menit</strong></p>
<br>
<h2 id="251-pengantar">2.5.1 Pengantar</h2>
<p>Threads dan concurrency adalah konsep penting dalam pemrograman yang berkaitan dengan cara sebuah program dapat melakukan eksekusi beberapa tugas secara bersamaan. Dalam dunia komputasi, kemampuan untuk menjalankan beberapa tugas dalam waktu yang bersamaan sangatlah krusial, terutama dalam mengoptimalkan penggunaan sumber daya komputer yang tersedia. Threads dan concurrency memungkinkan aplikasi untuk menjadi lebih responsif, efisien, dan mampu mengatasi tugas-tugas yang kompleks.</p>
<p>Thread adalah unit dasar eksekusi dalam sebuah program. Setiap program memiliki minimal satu thread, yang biasanya disebut sebagai &quot;main thread&quot;. Threads memungkinkan eksekusi tugas-tugas secara konkuren (bersamaan), sehingga berbagai bagian dari program dapat dijalankan secara bersamaan untuk meningkatkan kinerja. Threads dapat digunakan untuk menangani tugas-tugas yang membutuhkan waktu lama untuk menyelesaikan, seperti mengunduh data dari jaringan, melakukan operasi I/O (Input/Output), dan tugas-tugas lain yang memerlukan waktu yang signifikan.</p>
<p>Concurrency adalah konsep yang berkaitan dengan manajemen dan koordinasi beberapa thread yang berjalan bersamaan. Meskipun threads dapat berjalan secara bersamaan, namun cara mereka berinteraksi dan berbagi sumber daya komputer dapat menjadi rumit. Konsep concurrency melibatkan pengaturan urutan eksekusi, sinkronisasi, dan manajemen data bersama agar tidak terjadi konflik atau ketidakseimbangan yang dapat mengakibatkan kesalahan atau performa yang buruk.</p>
<p>Dalam pengembangan perangkat lunak, penting untuk memahami cara bekerja threads dan concurrency, serta tantangan yang terkait dengannya. Beberapa tantangan umum meliputi race condition (kondisi perlomba-lomba), deadlocks (kemandekan), dan resource contention (persaingan sumber daya). Pemahaman yang baik tentang konsep ini akan memungkinkan pengembang untuk merancang aplikasi yang lebih efisien, aman, dan handal.</p>
<p>Dalam materi ini, kita akan menjelajahi konsep-konsep dasar threads dan concurrency, bagaimana cara membuat dan mengelola threads dalam berbagai bahasa pemrograman, serta teknik-teknik yang digunakan untuk menghindari masalah-masalah umum yang terkait dengan concurrent programming.</p>
<h2 id="252-program-processes-dan-threads">2.5.2 Program, Processes dan Threads</h2>
<p>Dalam dunia pemrograman dan sistem operasi, konsep program, proses, dan benang (threads) berperan penting dalam mengelola eksekusi tugas-tugas komputasi. Pemahaman tentang ketiga konsep ini menjadi krusial saat berbicara tentang threading dan konkurensi. Mari kita bahas satu per satu:</p>
<h3 id="program">Program</h3>
<p>Sebuah program adalah kumpulan instruksi yang ditulis dalam bahasa pemrograman tertentu untuk melakukan tugas-tugas tertentu. Program bisa menjadi sekumpulan fungsi, objek, atau modul-modul yang saling berinteraksi. Saat program dieksekusi, instruksi-instruksi tersebut diterjemahkan menjadi kode mesin yang dapat dipahami oleh komputer.</p>
<h3 id="processes">Processes</h3>
<p>Proses merupakan entitas eksekusi yang independen dalam sistem operasi. Setiap proses memiliki ruang alamat memorinya sendiri serta sumber daya yang terisolasi seperti CPU time, alamat memorinya, dan file yang dibukanya. Setiap program yang dieksekusi pada sistem berjalan sebagai proses yang terpisah. Proses dapat berinteraksi melalui komunikasi yang telah ditetapkan oleh sistem operasi, seperti antrean pesan atau soket.</p>
<h3 id="benang-threads">Benang (Threads)</h3>
<p>Benang (threads) adalah unit terkecil dari eksekusi dalam suatu proses. Satu proses dapat memiliki beberapa benang yang berbagi ruang alamat dan sumber daya dengan proses tersebut. Benang-benang ini dapat menjalankan tugas-tugas yang berbeda secara paralel atau konkuren. Benang-benang dalam satu proses dapat berkomunikasi dan berkoordinasi dengan mudah karena mereka berbagi memorinya.</p>
<h3 id="perbedaan-antara-proses-dan-benang">Perbedaan antara Proses dan Benang</h3>
<ul>
   <li>Proses memiliki ruang alamat memorinya sendiri, sedangkan benang dalam satu proses berbagi ruang alamat.</li>
   <li>Proses memiliki overhead komunikasi yang lebih besar karena memerlukan mekanisme eksternal (seperti komunikasi melalui antrean pesan) untuk berkomunikasi. Benang dalam satu proses dapat berkomunikasi langsung melalui berbagi memori.</li>
   <li>Pemutusan (crash) suatu proses tidak mempengaruhi proses lain. Namun, dalam satu proses, jika satu benang mengalami masalah, hal ini dapat mempengaruhi benang-benang lain dalam proses yang sama.</li>
</ul>
<h3 id="konkurensi-dan-manfaat-threads">Konkurensi dan Manfaat Threads</h3>
<p>Konkurensi adalah kemampuan untuk menjalankan tugas-tugas secara bersamaan, dan benang memungkinkan pengembang untuk mencapai konkurensi dalam program mereka. Penggunaan benang dapat memberikan manfaat seperti:</p>
<ul>
   <li>Peningkatan responsivitas: Program dapat merespons input atau peristiwa secara cepat tanpa terkunci dalam tugas tunggal.</li>
   <li>Pemanfaatan multi-core: Dalam komputer modern yang memiliki banyak inti (core), benang-benang dapat dijalankan secara paralel di inti yang berbeda, mempercepat pemrosesan.</li>
   <li>Pemisahan tugas: Benang memungkinkan pemisahan tugas yang berbeda dalam satu proses, meningkatkan efisiensi dan organisasi.</li>
</ul>
<p>Namun, penggunaan benang juga memperkenalkan tantangan seperti sinkronisasi dan perlombaan sumber daya (race conditions), yang harus dikelola dengan hati-hati.</p>
<p>Dalam rangka memahami threading dan konkurensi, penting untuk memiliki pemahaman yang kokoh tentang program, proses, dan benang. Ini akan membantu dalam merancang sistem yang efisien, responsif, dan aman.</p>
<h2 id="253-concurrency">2.5.3 Concurrency</h2>
<p>Concurrency merujuk pada kemampuan sistem untuk mengatur dan menjalankan beberapa tugas atau proses secara bersamaan dalam waktu yang hampir bersamaan. Dalam konteks pemrograman, concurrency berfokus pada kemampuan menjalankan beberapa alur eksekusi (threads) secara bersamaan dalam satu program. Hal ini sangat bermanfaat untuk mengoptimalkan penggunaan sumber daya komputer dan meningkatkan responsifitas program.</p>
<h3 id="thread">Thread</h3>
<p>Thread adalah unit dasar eksekusi dalam sebuah program. Program biasanya memiliki satu thread utama yang menjalankan kode secara berurutan. Namun, dengan menggunakan konsep multithreading, program dapat memiliki beberapa thread yang berjalan secara paralel. Thread-thread ini dapat berbagi sumber daya dan informasi program.</p>
<h3 id="concurrency-vs-parallelism">Concurrency vs Parallelism</h3>
<p>Concurrency dan parallelism sering kali digunakan secara bergantian, tetapi sebenarnya memiliki perbedaan yang penting. Concurrency berfokus pada pengaturan eksekusi beberapa tugas sehingga terlihat seolah-olah berjalan bersamaan. Parallelism, di sisi lain, melibatkan eksekusi tugas-tugas secara benar-benar bersamaan pada multiple core atau prosesor yang berbeda.</p>
<h3 id="manfaat-concurrency">Manfaat Concurrency</h3>
<ol>
   <li><strong>Responsifitas:</strong> Dengan menggunakan thread, program bisa tetap merespons input dari pengguna atau berkomunikasi dengan perangkat eksternal sementara tugas berat dijalankan di background.</li>
   <li><strong>Pemanfaatan Sumber Daya:</strong> Dalam sistem dengan banyak core atau prosesor, concurrency memungkinkan aplikasi menggunakan sumber daya tersebut secara lebih efisien.</li>
   <li><strong>Kemudahan Pembangunan:</strong> Menggunakan thread dapat mempermudah pemrograman aplikasi yang kompleks dengan membaginya menjadi tugas-tugas yang lebih kecil dan independen.</li>
</ol>
<h3 id="tantangan-concurrency">Tantangan Concurrency</h3>
<p>Concurrency membawa tantangan tersendiri:</p>
<ol>
   <li><strong>Race Condition:</strong> Ketika dua thread atau lebih berusaha mengakses dan memodifikasi data bersamaan, dapat terjadi kondisi di mana hasil akhir tidak terduga akibat urutan eksekusi yang tidak terprediksi.</li>
   <li><strong>Deadlock:</strong> Situasi di mana beberapa thread saling menunggu satu sama lain dan tidak ada thread yang dapat melanjutkan eksekusi.</li>
   <li><strong>Starvation:</strong> Sebuah thread mungkin tidak mendapatkan kesempatan untuk dieksekusi jika terus-menerus mendapatkan prioritas lebih rendah dibandingkan thread lain.</li>
   <li><strong>Synchronization Overhead:</strong> Mengelola sinkronisasi antara thread-thread dapat memakan waktu dan sumber daya.</li>
</ol>
<h3 id="cara-menggunakan-concurrency">Cara Menggunakan Concurrency</h3>
<p>Pada pemrograman, beberapa cara untuk menggunakan concurrency melalui thread antara lain:</p>
<ol>
   <li><strong>Pemrograman Berbasis Thread:</strong> Membuat dan mengatur thread secara langsung menggunakan API yang disediakan oleh bahasa pemrograman (seperti Java, Python, C++).</li>
   <li><strong>Task Parallelism:</strong> Menggunakan konsep task untuk mendelegasikan tugas-tugas ke sistem, yang kemudian akan mengatur bagaimana tugas-tugas tersebut dijalankan secara konkuren.</li>
   <li><strong>Asynchronous Programming:</strong> Menggunakan konsep async/await atau promise untuk mengatasi tugas-tugas yang mengandalkan I/O yang lambat tanpa harus memblokir eksekusi program.</li>
</ol>
<p>Concurrency adalah konsep yang penting dalam pemrograman modern untuk mengoptimalkan penggunaan sumber daya komputer dan meningkatkan responsifitas aplikasi. Namun, diperlukan perhatian khusus dalam mengatasi tantangan-tantangan yang muncul akibat penggunaan thread secara bersamaan.</p>
<h2 id="254-parallelism">2.5.4 Parallelism</h2>
<p>Parallelisme adalah konsep dalam komputasi di mana beberapa tugas atau operasi dilakukan secara bersamaan dalam rangka meningkatkan kinerja dan efisiensi. Dalam konteks threads dan concurrency, parallelisme berarti menjalankan beberapa tugas atau operasi pada saat yang bersamaan dengan tujuan memanfaatkan sumber daya yang ada secara lebih efisien.</p>
<h3 id="parallelisme-dalam-threads-dan-concurrecy">Parallelisme dalam Threads dan Concurrecy</h3>
<p>Parallelisme melibatkan menjalankan beberapa thread atau tugas pada saat yang bersamaan dengan tujuan memanfaatkan kekuatan pemrosesan yang ada dalam multi-core atau multi-processor sistem. Ini dapat meningkatkan kinerja dan mengurangi waktu eksekusi tugas-tugas yang besar dan kompleks.</p>
<h3 id="keuntungan-parallelisme">Keuntungan Parallelisme</h3>
<ul>
   <li>
      <p><strong>Kinerja yang Lebih Baik</strong>: Dengan menjalankan tugas-tugas secara bersamaan, waktu eksekusi dapat dikurangi secara signifikan. Ini penting dalam aplikasi yang membutuhkan pemrosesan cepat, seperti komputasi ilmiah atau pemrosesan data besar.</p>
   </li>
   <li>
      <p><strong>Memanfaatkan Sumber Daya</strong>: Dalam sistem dengan beberapa core atau prosesor, parallelisme memungkinkan pemanfaatan sumber daya secara efisien. Jika sebuah tugas hanya menggunakan satu core, sumber daya yang tersedia di core lainnya tidak terbuang percuma.</p>
   </li>
   <li>
      <p><strong>Responsif</strong>: Dalam aplikasi yang melibatkan interaksi pengguna, menggunakan parallelisme dapat memastikan bahwa tampilan antarmuka pengguna tetap responsif saat tugas-tugas berat sedang dieksekusi di latar belakang.</p>
   </li>
</ul>
<h3 id="tantangan-parallelisme">Tantangan Parallelisme</h3>
<ul>
   <li>
      <p><strong>Sinkronisasi</strong>: Mengatur bagaimana thread-thread berkomunikasi dan berbagi data adalah tantangan utama. Tanpa sinkronisasi yang baik, dapat terjadi kondisi perlombaan (race conditions) atau masalah lain yang menyebabkan ketidakpastian dalam eksekusi.</p>
   </li>
   <li>
      <p><strong>Deadlock</strong>: Deadlock adalah situasi di mana beberapa thread saling menunggu satu sama lain, sehingga tidak ada yang dapat melanjutkan eksekusi. Hal ini dapat mengakibatkan aplikasi menggantung dan tidak merespons.</p>
   </li>
   <li>
      <p><strong>Overhead</strong>: Membuat dan mengelola thread-thread memerlukan sumber daya dan waktu tertentu. Dalam beberapa kasus, biaya ini mungkin melebihi manfaat dari parallelisme.</p>
   </li>
</ul>
<h3 id="penggunaan-yang-tepat">Penggunaan yang Tepat</h3>
<p>Parallelisme cocok untuk tugas-tugas yang dapat dipecah menjadi bagian-bagian independen yang dapat dieksekusi secara bersamaan. Misalnya, dalam pemrosesan data, mencari elemen terbesar dalam sebuah daftar dapat dipecah menjadi beberapa tugas yang masing-masing mencari elemen dalam subdaftar yang berbeda.</p>
<p>Parallelisme dalam konteks threads dan concurrecy adalah tentang menjalankan beberapa tugas atau operasi pada saat yang bersamaan untuk meningkatkan kinerja dan efisiensi. Ini memungkinkan pemanfaatan sumber daya yang ada secara lebih efektif, tetapi juga memerlukan manajemen yang hati-hati terkait sinkronisasi dan penanganan masalah potensial seperti deadlock.</p>