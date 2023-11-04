<div align="center">
   <h1>Module 2 Lesson 6 - Manajemen Memori</h1>
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
         <td>2.6.1 Pengantar</td>
         <td><a href="#261-pengantar">🔽</a></td>
      </tr>
      <tr>
         <td>2.6.2 Memory Allocation</td>
         <td><a href="#262-memory-allocation">🔽</a></td>
      </tr>
      <tr>
         <td>2.6.3 Address Space</td>
         <td><a href="#263-address-space">🔽</a></td>
      </tr>
      <tr>
         <td>2.6.4 Dynamic Loading</td>
         <td><a href="#264-dynamic-loading">🔽</a></td>
      </tr>
      <tr>
         <td>2.6.5 Dynamic Linking</td>
         <td><a href="#265-dynamic-linking">🔽</a></td>
      </tr>
      <tr>
         <td>2.6.6 Swapping</td>
         <td><a href="#266-swapping">🔽</a></td>
      </tr>
   </tbody>
</table>
<p><br><br></p>
<h1 id="durasi-pembelajaran">Durasi Pembelajaran</h1>
<p><strong>75 menit</strong></p>
<br>
<h2 id="261-pengantar">2.6.1 Pengantar</h2>
<p>Sistem komputer modern telah mengalami evolusi yang luar biasa dalam hal kinerja dan kompleksitas. Salah satu komponen inti dalam operasi sistem komputer adalah manajemen memori. Manajemen memori berkaitan dengan cara pengaturan, pengalokasian, dan pembebasan memori dalam sistem komputer agar aplikasi dan proses dapat berjalan secara efisien dan aman.</p>
<p>Dalam pengertian sederhananya, memori mengacu pada ruang penyimpanan yang digunakan oleh program, data, dan instruksi yang sedang dieksekusi dalam sistem komputer. Meskipun memori bisa terlihat sebagai sumber daya yang cukup melimpah, namun dengan meningkatnya kompleksitas aplikasi dan sistem, serta adanya batasan fisik dalam perangkat keras, manajemen memori menjadi aspek yang sangat penting untuk diperhatikan.</p>
<p>Dalam pengelolaan memori, terdapat beberapa konsep dan teknik penting yang perlu dipahami, antara lain:</p>
<ol>
   <li>
      <p><strong>Model Memori:</strong> Model memori menjelaskan bagaimana memori diatur dalam sistem komputer. Ini melibatkan pembagian memori menjadi beberapa bagian seperti stack, heap, kode program, dan data, serta cara akses ke setiap bagian ini.</p>
   </li>
   <li>
      <p><strong>Pengalokasian Memori:</strong> Salah satu tugas utama manajemen memori adalah mengalokasikan ruang memori untuk setiap proses atau aplikasi. Pengalokasian dapat dilakukan menggunakan berbagai strategi seperti pengalokasian statis dan dinamis, yang melibatkan penentuan berapa banyak memori yang akan dialokasikan untuk setiap proses.</p>
   </li>
   <li>
      <p><strong>Fragmentasi:</strong> Fragmentasi adalah masalah yang terjadi ketika memori menjadi terfragmentasi, baik dalam bentuk fragmentasi internal (pada alokasi memori dinamis) atau fragmentasi eksternal (ketika tidak ada blok kontigu yang cukup besar untuk memenuhi permintaan alokasi).</p>
   </li>
   <li>
      <p><strong>Pembebasan Memori:</strong> Setelah memori tidak diperlukan lagi oleh suatu proses, penting untuk membebaskannya agar memori tersebut bisa digunakan oleh proses lain. Pembebasan memori yang tidak benar bisa menyebabkan kebocoran memori (memory leaks).</p>
   </li>
   <li>
      <p><strong>Virtual Memory:</strong> Konsep virtual memory memungkinkan sistem operasi untuk mengalokasikan ruang di disk sebagai perpanjangan memori fisik. Ini memungkinkan penggunaan aplikasi yang lebih besar dari memori fisik yang sebenarnya.</p>
   </li>
   <li>
      <p><strong>Paging dan Segmentasi:</strong> Dalam manajemen memori, terdapat konsep paging dan segmentasi yang membantu dalam alokasi dan penggunaan memori. Paging melibatkan pemecahan memori fisik menjadi blok-blok yang disebut &quot;frame&quot; dan memori logis dari proses dibagi menjadi &quot;page&quot;. Segmentasi, di sisi lain, membagi memori logis menjadi segmen-segmen yang mewakili bagian dari program atau data.</p>
   </li>
</ol>
<p>Dalam pengantar materi ini, kita akan menjelajahi konsep-konsep tersebut lebih mendalam, melihat bagaimana sistem operasi mengelola memori untuk menjaga kinerja, keamanan, dan efisiensi dalam operasi komputer sehari-hari. Dengan pemahaman yang lebih baik tentang manajemen memori, kamu akan dapat mengaplikasikan prinsip-prinsip ini dalam pengembangan perangkat lunak dan pemecahan masalah terkait kinerja sistem.</p>
<h2 id="262-memory-allocation">2.6.2 Memory Allocation</h2>
<p>Dalam komputasi, manajemen memori merujuk pada proses mengelola penggunaan dan alokasi memori sistem untuk program yang berjalan. Salah satu aspek penting dari manajemen memori adalah <em>memory allocation</em> atau alokasi memori, yaitu cara komputer mengatur dan menyediakan ruang memori kepada program-program yang berjalan.</p>
<h3 id="tujuan-memory-allocation">Tujuan Memory Allocation</h3>
<p>Tujuan utama dari memory allocation adalah untuk memastikan efisiensi dan efektivitas penggunaan memori yang terbatas pada sistem. Hal ini mencakup pemberian memori yang cukup kepada program-program, menghindari tumpang-tindih dan konflik alokasi, serta memastikan penggunaan yang aman dan terkoordinasi.</p>
<h3 id="tipe-tipe-alokasi-memori">Tipe-tipe Alokasi Memori</h3>
<h4 id="static-memory-allocation">Static Memory Allocation</h4>
<p>Memori dialokasikan saat program mulai dijalankan dan tetap dialokasikan selama program berjalan. Ini cocok untuk variabel yang memiliki ukuran tetap dan dikenal sebelum eksekusi.</p>
<h4 id="dynamic-memory-allocation">Dynamic Memory Allocation</h4>
<p>Memori dialokasikan saat program berjalan. Ini memungkinkan alokasi memori yang fleksibel sesuai dengan kebutuhan, seperti array dengan ukuran yang ditentukan oleh pengguna atau struktur data yang dinamis.</p>
<h3 id="metode-alokasi-memori">Metode Alokasi Memori</h3>
<h4 id="stack-allocation">Stack Allocation</h4>
<p>Digunakan untuk menyimpan data lokal dalam fungsi. Memori dialokasikan dalam bentuk tumpukan (stack) yang secara otomatis dikelola oleh sistem. Alokasi dan dealokasi terjadi secara otomatis ketika fungsi masuk dan keluar.</p>
<h4 id="heap-allocation">Heap Allocation</h4>
<p>Digunakan untuk menyimpan data yang ukurannya tidak diketahui pada waktu kompilasi atau perlu alokasi yang tahan lama. Alokasi memori di heap memerlukan tindakan eksplisit untuk alokasi dan dealokasi, seperti menggunakan fungsi <code>malloc</code>, <code>calloc</code>, <code>realloc</code>, dan <code>free</code> dalam bahasa C.</p>
<h3 id="masalah-dalam-memory-allocation">Masalah dalam Memory Allocation</h3>
<ol>
   <li>
      <p><strong>Fragmentasi:</strong> Terjadi ketika ada potongan-potongan kecil memori yang tidak dapat digunakan karena terpisah oleh bagian lain yang dialokasikan. Terdapat dua jenis fragmentasi: <em>internal fragmentation</em> (penyia-nyiaan internal) terjadi ketika ada bagian memori yang diambil oleh program tetapi tidak digunakan sepenuhnya, sedangkan <em>external fragmentation</em> (penyia-nyiaan eksternal) terjadi ketika terdapat ruang kosong di antara blok-blok alokasi yang tidak cukup besar untuk digunakan oleh alokasi berikutnya.</p>
   </li>
   <li>
      <p><strong>Memory Leaks:</strong> Terjadi ketika program gagal untuk melepaskan memori yang sebelumnya dialokasikan, sehingga memori tidak dapat digunakan lagi oleh program lain.</p>
   </li>
</ol>
<p><strong>Strategi Alokasi Memori:</strong></p>
<ol>
   <li>
      <p><strong>First-Fit:</strong> Memori dialokasikan pada blok pertama yang cukup besar untuk memenuhi kebutuhan alokasi.</p>
   </li>
   <li>
      <p><strong>Best-Fit:</strong> Memori dialokasikan pada blok terkecil yang dapat memuat alokasi dengan ukuran tertentu.</p>
   </li>
   <li>
      <p><strong>Worst-Fit:</strong> Memori dialokasikan pada blok terbesar yang dapat memuat alokasi dengan ukuran tertentu. </p>
   </li>
   <li>
      <p><strong>Buddy Allocation:</strong> Digunakan untuk menghindari fragmentasi. Memori dibagi menjadi blok-blok berukuran 2^n, dan ketika alokasi tidak cocok, blok-blok dapat digabungkan kembali.</p>
   </li>
</ol>
<h2 id="263-address-space">2.6.3 Address Space</h2>
<p>Address space (ruang alamat) dalam konteks manajemen memori mengacu pada kisaran alamat yang digunakan oleh sebuah program atau sistem operasi untuk mengakses lokasi-lokasi tertentu dalam memori fisik. Setiap program yang berjalan di dalam sistem memiliki address space sendiri-sendiri yang menyediakan alamat-alamat yang digunakan untuk mengakses instruksi dan data di memori.</p>
<h3 id="komponen-komponen-address-space">Komponen-Komponen Address Space</h3>
<ol>
   <li>
      <p><strong>Kode (Text Segment)</strong>: Bagian ini berisi instruksi-instruksi program yang dieksekusi oleh CPU. Ini adalah bagian dari memori di mana kode program ditempatkan dan biasanya bersifat read-only, artinya instruksi tidak dapat diubah saat program berjalan.</p>
   </li>
   <li>
      <p><strong>Data</strong>: Bagian ini berisi variabel-variabel global, variabel statik, dan variabel lain yang disimpan selama program berjalan. Bagian data ini dapat dibagi lagi menjadi dua sub-bagian:</p>
      <ul>
         <li><strong>Initialized Data Segment</strong>: Berisi variabel yang diberikan nilai awal sebelum program dimulai.</li>
         <li><strong>Uninitialized Data Segment (BSS)</strong>: Berisi variabel yang diinisialisasi dengan nilai nol.</li>
      </ul>
   </li>
   <li>
      <p><strong>Heap</strong>: Heap adalah area memori yang digunakan untuk alokasi dinamis, seperti ketika program meminta memori saat berjalan. Ini adalah tempat di mana data struktur seperti array atau objek yang ukurannya tidak diketahui sebelumnya bisa ditempatkan.</p>
   </li>
   <li>
      <p><strong>Stack</strong>: Stack digunakan untuk menyimpan data sementara yang berkaitan dengan pemanggilan fungsi dan alokasi variabel lokal. Data dalam stack diatur dalam bentuk tumpukan, di mana data yang terakhir dimasukkan (pushed) adalah data pertama yang diambil (popped).</p>
   </li>
</ol>
<h3 id="virtual-address-space-dan-physical-memory">Virtual Address Space dan Physical Memory</h3>
<p>Penting untuk memahami bahwa address space yang digunakan oleh program tidak selalu sesuai dengan lokasi fisik di memori utama (physical memory). Dalam sistem manajemen memori modern, konsep virtual memory digunakan. Virtual memory memungkinkan program untuk berjalan seolah-olah memiliki lebih banyak memori daripada yang sebenarnya tersedia. Sistem operasi akan melakukan pemetaan antara alamat virtual yang digunakan oleh program dengan alamat fisik di memori fisik. Ini memungkinkan program untuk mengakses memori yang lebih besar daripada yang tersedia dalam fisik memori RAM sebenarnya.</p>
<p>Dalam manajemen memori, address space merujuk pada rentang alamat yang digunakan oleh program atau sistem operasi untuk mengakses memori. Ini terdiri dari kode, data, heap, dan stack, dan dapat dikelola melalui konsep virtual memory untuk mengakses memori fisik yang lebih besar daripada yang tersedia secara fisik.</p>
<h2 id="264-dynamic-loading">2.6.4 Dynamic Loading</h2>
<p>Dynamic loading (pemuatan dinamis) adalah konsep dalam manajemen memori komputer di mana sebuah program atau aplikasi hanya memuat bagian-bagian tertentu dari kode atau data yang diperlukan saat dieksekusi, daripada memuat seluruh program ke dalam memori pada saat dimulainya. Pendekatan ini membantu mengoptimalkan penggunaan memori dan mempercepat waktu pemuatan program.</p>
<h3 id="cara-kerja">Cara Kerja</h3>
<p>Ketika suatu program dimulai, dynamic loading memungkinkan sistem operasi untuk memuat hanya bagian-bagian tertentu yang dibutuhkan oleh program pada tahap awal. Ketika bagian-bagian lainnya diperlukan selama eksekusi, sistem operasi akan memuatnya secara dinamis ke dalam memori. Ini mengurangi waktu yang dibutuhkan untuk memulai program, karena hanya bagian penting pertama kali dimuat.</p>
<h3 id="keuntungan">Keuntungan</h3>
<ol>
   <li>
      <p><strong>Efisiensi Memori:</strong> Dynamic loading memungkinkan penggunaan memori yang lebih efisien karena hanya bagian-bagian yang benar-benar diperlukan yang dimuat ke dalam memori. Ini memungkinkan lebih banyak program berjalan pada saat yang bersamaan tanpa menghabiskan terlalu banyak memori fisik.</p>
   </li>
   <li>
      <p><strong>Waktu Pemuatan Lebih Cepat:</strong> Karena hanya bagian-bagian awal program yang dimuat saat pertama kali dieksekusi, waktu yang dibutuhkan untuk memulai program menjadi lebih cepat.</p>
   </li>
   <li>
      <p><strong>Fleksibilitas:</strong> Dynamic loading memungkinkan program untuk memuat modul atau fungsi hanya saat dibutuhkan. Ini berguna dalam skenario di mana beberapa bagian dari program hanya akan digunakan dalam kondisi atau situasi tertentu.</p>
   </li>
</ol>
<h3 id="contoh">Contoh</h3>
<p>Misalkan kamu memiliki program pengolahan gambar yang memiliki beberapa filter yang berbeda. Dengan dynamic loading, kamu dapat memuat hanya filter yang akan digunakan oleh pengguna pada saat itu. Jika pengguna memilih untuk menerapkan suatu filter, hanya saat itu filter tersebut akan dimuat ke dalam memori. Ini menghemat memori dan memungkinkan program berjalan lebih efisien.</p>
<p>Dynamic loading adalah pendekatan dalam manajemen memori yang memungkinkan program memuat bagian-bagian tertentu dari kode atau data secara dinamis saat diperlukan. Ini membantu mengoptimalkan penggunaan memori, mempercepat waktu pemuatan program, dan memberikan fleksibilitas dalam menjalankan program.</p>
<h2 id="265-dynamic-linking">2.6.5 Dynamic Linking</h2>
<p>Dynamic linking adalah suatu konsep dalam manajemen memori yang terkait dengan cara menghubungkan program dengan pustaka (library) eksternal secara dinamis saat program dijalankan, bukan pada saat kompilasi. Dalam konteks sistem operasi dan manajemen memori, dynamic linking memungkinkan program untuk mengakses kode dan data yang ada di dalam pustaka eksternal tanpa harus menggandakan kode tersebut ke dalam setiap program yang menggunakannya.</p>
<h3 id="cara-kerja-dynamic-linking">Cara Kerja Dynamic Linking</h3>
<ol>
   <li>
      <p><strong>Kompilasi</strong>: Saat mengembangkan program, kita menggunakan pustaka-pustaka eksternal (seperti pustaka matematika, I/O, dll.) yang disediakan oleh sistem operasi atau pihak ketiga. Kode dalam program kita merujuk pada fungsi-fungsi yang terdapat dalam pustaka-pustaka ini.</p>
   </li>
   <li>
      <p><strong>Linking</strong>: Ketika program dikompilasi, referensi-referensi ini ditandai sebagai &quot;tunduk pada dynamic linking&quot;. Ini berarti bahwa program hanya mengetahui bahwa ada referensi ke fungsi-fungsi dalam pustaka tersebut, namun tidak mengandung kode penuh dari pustaka tersebut.</p>
   </li>
   <li>
      <p><strong>Dynamic Linker</strong>: Saat program dijalankan, dynamic linker (bagian dari sistem operasi) mengambil peran untuk menghubungkan program dengan pustaka-pustaka yang diperlukan. Dynamic linker menemukan lokasi pustaka dalam sistem dan memuatnya ke dalam memori.</p>
   </li>
   <li>
      <p><strong>Eksekusi</strong>: Setelah pustaka terhubung secara dinamis, program dapat menjalankan fungsi-fungsi dari pustaka ini seperti biasa. Kode pustaka akan digunakan oleh beberapa program berbeda yang telah melakukan dynamic linking terhadapnya.</p>
   </li>
</ol>
<h3 id="keuntungan-dynamic-linking">Keuntungan Dynamic Linking</h3>
<ol>
   <li>
      <p><strong>Penghematan Ruang</strong>: Karena kode pustaka tidak perlu disalin ke setiap program yang menggunakannya, ini menghemat ruang di dalam media penyimpanan dan memori fisik.</p>
   </li>
   <li>
      <p><strong>Perubahan Tertangguh</strong>: Jika ada pembaruan atau perbaikan dalam pustaka, perubahan tersebut hanya perlu dilakukan di satu tempat, yaitu dalam pustaka itu sendiri. Semua program yang menghubungkannya secara dinamis akan secara otomatis mendapatkan perubahan ini tanpa harus dikompilasi ulang.</p>
   </li>
   <li>
      <p><strong>Pembaruan Aplikasi yang Mudah</strong>: Pembaruan aplikasi hanya berfokus pada kode program aplikasi itu sendiri, tanpa harus khawatir tentang pustaka-pustaka yang dihubungkannya.</p>
   </li>
</ol>
<h3 id="kerugian-dynamic-linking">Kerugian Dynamic Linking</h3>
<ol>
   <li>
      <p><strong>Ketergantungan Eksternal</strong>: Program yang mengandalkan dynamic linking memerlukan ketersediaan pustaka yang diperlukan saat program dijalankan. Jika pustaka hilang atau tidak kompatibel, program tidak dapat berjalan.</p>
   </li>
   <li>
      <p><strong>Overhead Waktu Eksekusi</strong>: Proses dynamic linking memerlukan waktu tambahan saat program dijalankan. Dynamic linker harus menemukan pustaka yang diperlukan, memuatnya ke dalam memori, dan menghubungkannya dengan program.</p>
   </li>
</ol>
<p>Dynamic linking adalah pendekatan dalam manajemen memori di mana program dihubungkan dengan pustaka-pustaka eksternal saat program dijalankan, menghemat ruang dan memudahkan pembaruan. Namun, ini juga memiliki beberapa keterbatasan terkait ketergantungan pada pustaka eksternal dan overhead waktu eksekusi.</p>
<h2 id="266-swapping">2.6.6 Swapping</h2>
<p>Swapping adalah salah satu teknik dalam manajemen memori komputer yang digunakan untuk mengalokasikan dan mengelola ruang memori secara efisien. Konsep swapping melibatkan pemindahan data antara memori fisik (RAM) dan memori sekunder (misalnya, hard disk) untuk memberikan ruang yang cukup bagi program yang sedang berjalan. Ini memungkinkan sistem operasi untuk menjalankan lebih banyak program daripada kapasitas memori fisik yang tersedia.</p>
<h3 id="cara-kerja-swapping">Cara Kerja Swapping</h3>
<p>Proses swapping dilakukan ketika memori fisik tidak lagi mencukupi untuk menampung semua program dan data yang aktif dalam sistem. Berikut adalah langkah-langkah umum yang terjadi saat proses swapping dilakukan:</p>
<ol>
   <li>
      <p><strong>Deteksi Kekurangan Memori:</strong> Sistem operasi terus memantau penggunaan memori oleh berbagai program. Ketika kapasitas memori fisik mendekati penuh, sistem operasi mendeteksi kekurangan memori.</p>
   </li>
   <li>
      <p><strong>Pemilihan Proses untuk Swapping:</strong> Sistem operasi akan memilih salah satu atau beberapa proses yang sedang berjalan sebagai kandidat untuk di-swap. Proses ini mungkin sedang idle atau memiliki prioritas rendah.</p>
   </li>
   <li>
      <p><strong>Pemindahan Data:</strong> Data dari proses yang dipilih akan dipindahkan dari memori fisik ke memori sekunder (misalnya, hard disk). Ini dapat melibatkan pindahnya area memori yang digunakan oleh proses tersebut, termasuk kode program, data, dan status register.</p>
   </li>
   <li>
      <p><strong>Alokasi Memori untuk Proses Baru:</strong> Setelah memori fisik dibebaskan melalui proses swapping, sistem operasi dapat mengalokasikan memori untuk menjalankan proses baru atau memuat kembali proses yang telah di-swap jika diperlukan.</p>
   </li>
   <li>
      <p><strong>Pemulihan Proses Swapped:</strong> Jika proses yang di-swap perlu dilanjutkan, data yang diperlukan akan dikembalikan dari memori sekunder ke memori fisik. Proses ini juga dikenal sebagai &quot;swap in.&quot;</p>
   </li>
</ol>
<h3 id="keuntungan-dan-kekurangan-swapping">Keuntungan dan Kekurangan Swapping</h3>
<p><strong>Keuntungan:</strong></p>
<ul>
   <li><strong>Peningkatan Kapasitas:</strong> Swapping memungkinkan sistem operasi untuk menjalankan lebih banyak program daripada kapasitas memori fisik yang terbatas.</li>
   <li><strong>Efisiensi Pemanfaatan Memori:</strong> Program yang sedang aktif tetap dapat berjalan, meskipun sebagian data dipindahkan ke memori sekunder.</li>
   <li><strong>Prioritasi Penggunaan Memori:</strong> Sistem operasi dapat mengatur prioritas pemindahan data berdasarkan pentingnya proses.</li>
</ul>
<p><strong>Kekurangan:</strong></p>
<ul>
   <li><strong>Kinerja Rendah:</strong> Pemindahan data antara memori fisik dan memori sekunder memakan waktu, yang dapat mengakibatkan penurunan kinerja saat swapping terjadi terlalu sering.</li>
   <li><strong>Fragmentasi Memori:</strong> Penggunaan berulang swapping dapat menyebabkan fragmentasi memori, yang pada akhirnya dapat membatasi kemampuan sistem operasi dalam mengalokasikan memori secara efisien.</li>
   <li><strong>Resiko Kehilangan Data:</strong> Terdapat risiko data yang belum tersimpan di memori sekunder jika terjadi kegagalan sistem.</li>
</ul>
<p>Swapping adalah teknik yang digunakan dalam manajemen memori untuk mengatasi keterbatasan memori fisik dengan memindahkan data antara memori fisik dan memori sekunder. Meskipun dapat memungkinkan sistem operasi untuk menjalankan lebih banyak program, penggunaan yang berlebihan dapat berdampak negatif pada kinerja dan efisiensi sistem.</p>