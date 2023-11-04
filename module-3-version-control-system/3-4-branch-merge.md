<div align="center">
   <h1>Module 3 Lesson 4 - Branch dan Merge</h1>
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
         <td>3.4.1 Pengenalan Git Branching</td>
         <td><a href="#341-pengenalan-git-branching">🔽</a></td>
      </tr>
      <tr>
         <td>3.4.2 Membuat Branch Baru</td>
         <td><a href="#342-membuat-branch-baru">🔽</a></td>
      </tr>
      <tr>
         <td>3.4.3 git checkout</td>
         <td><a href="#343-git-checkout">🔽</a></td>
      </tr>
      <tr>
         <td>3.4.4 git merge</td>
         <td><a href="#344-git-merge">🔽</a></td>
      </tr>
      <tr>
         <td>3.4.5 Branch Management Best Practices</td>
         <td><a href="#345-branch-management-best-practices">🔽</a></td>
      </tr>
   </tbody>
</table>
<p><br><br></p>
<h1 id="durasi-pembelajaran">Durasi Pembelajaran</h1>
<p><strong>60 menit</strong></p>
<br>
<h2 id="341-pengenalan-git-branching">3.4.1 Pengenalan Git Branching</h2>
<p>Dalam pengembangan perangkat lunak, terutama ketika bekerja dalam tim atau dengan kode yang kompleks, konsep <strong>branching</strong> (cabang) dan <strong>merging</strong> (penggabungan) menjadi sangat penting. Branching dan merging adalah strategi yang digunakan untuk mengelola kode sumber dalam repositori versi kontrol, seperti Git. Repositori versi kontrol memungkinkan para pengembang untuk bekerja secara bersamaan, mengelola perubahan, dan mengintegrasikan kode dengan efisien.</p>
<p><strong>Branching (Cabang)</strong></p>
<p>Ketika kamu bekerja pada proyek perangkat lunak, terkadang kamu perlu bekerja pada fitur baru, perbaikan bug, atau perubahan lainnya yang mungkin tidak siap untuk langsung dimasukkan ke dalam kode utama (main atau master). Inilah saatnya konsep branching menjadi penting. Branching memungkinkan kamu untuk membuat salinan terisolasi dari kode utama, yang disebut &quot;cabang&quot; (branch). Dengan cara ini, kamu dapat bekerja pada perubahan tersebut tanpa mengganggu kode utama atau pekerjaan rekan-rekan kamu.</p>
<p><strong>Merging (Penggabungan)</strong></p>
<p>Setelah kamu selesai dengan perubahan di cabang yang telah kamu buat, langkah selanjutnya adalah menggabungkan perubahan tersebut kembali ke dalam kode utama. Proses ini disebut merging. Penggabungan melibatkan mengambil perubahan dari satu cabang dan mengintegrasikannya ke dalam cabang lain, biasanya cabang utama. Tujuannya adalah untuk memasukkan perubahan dengan aman dan tanpa mengganggu kerja rekan-rekan yang mungkin telah bekerja pada cabang utama.</p>
<h2 id="342-membuat-branch-baru">3.4.2 Membuat Branch Baru</h2>
<p>Dalam sistem kontrol versi Git, &quot;branch&quot; adalah cabang dari riwayat pengembangan kode. Membuat branch baru memungkinkan pengembang untuk bekerja pada fitur atau perbaikan tertentu tanpa mengganggu kode yang ada di branch utama (biasanya disebut &quot;master&quot; atau &quot;main&quot;). Ini membantu dalam mengelola kode secara terpisah sebelum menggabungkannya kembali ke cabang utama.</p>
<h3 id="langkah-langkah-untuk-membuat-branch-baru">Langkah-langkah untuk Membuat Branch Baru</h3>
<p>Berikut adalah langkah-langkah umum untuk membuat branch baru dalam Git:</p>
<ol>
   <li>
      <p><strong>Memastikan Kode Terbaru:</strong> Pastikan kamu memiliki versi kode terbaru di cabang utama dengan menjalankan perintah:</p>
      <pre><code>git checkout main
git pull origin main
</code></pre>
   </li>
   <li>
      <p><strong>Membuat Branch Baru:</strong> Gunakan perintah <code>git branch</code> diikuti dengan nama branch yang ingin kamu buat. Contoh:</p>
      <pre><code>git branch fitur-baru
</code></pre>
   </li>
   <li>
      <p><strong>Beralih ke Branch Baru:</strong> Kamu perlu beralih ke branch baru untuk mulai bekerja padanya:</p>
      <pre><code>git checkout fitur-baru
</code></pre>
      <p>Atau, langkah 2 dan 3 bisa digabung menjadi satu dengan perintah:</p>
      <pre><code>git checkout -b fitur-baru
</code></pre>
   </li>
   <li>
      <p><strong>Mengembangkan Kode:</strong> Lakukan perubahan kode yang diperlukan di branch baru. Kamu dapat melakukan komit seperti biasa dengan <code>git commit</code>.</p>
   </li>
   <li>
      <p><strong>Menggabungkan Kode Kembali:</strong> Setelah selesai dengan perubahan di branch baru, kamu bisa menggabungkan perubahan tersebut kembali ke cabang utama (misalnya, <code>main</code>). Ini disebut &quot;merge&quot; atau &quot;pull request&quot; tergantung pada alur kerja yang kamu gunakan.</p>
   </li>
   <li>
      <p><strong>Menghapus Branch (Opsional):</strong> Setelah branch telah digabungkan dan tidak diperlukan lagi, kamu bisa menghapusnya untuk menjaga kebersihan repositori:</p>
      <pre><code>git branch -d fitur-baru
</code></pre>
   </li>
</ol>
<p><strong>Keuntungan Membuat Branch:</strong></p>
<ul>
   <li>
      <p><strong>Isolasi Perubahan:</strong> Membuat branch memungkinkan isolasi perubahan terhadap kode utama, sehingga menghindari potensi kerusakan pada kode yang bekerja.</p>
   </li>
   <li>
      <p><strong>Pengembangan Paralel:</strong> Tim pengembang bisa bekerja secara paralel pada berbagai fitur tanpa mengganggu satu sama lain.</p>
   </li>
   <li>
      <p><strong>Uji Coba Fitur:</strong> Branch baru bisa digunakan untuk menguji fitur baru sebelum menggabungkannya ke cabang utama.</p>
   </li>
   <li>
      <p><strong>Pelacakan Pekerjaan:</strong> Branch dapat digunakan untuk melacak pekerjaan yang sedang dikerjakan, memungkinkan manajemen yang lebih baik.</p>
   </li>
   <li>
      <p><strong>Riwayat Perubahan:</strong> Branch mempertahankan riwayat perubahan, yang memudahkan melihat apa yang telah diubah.</p>
   </li>
</ul>
<p>Menggunakan branch dalam Git adalah praktik yang sangat umum dalam pengembangan perangkat lunak, karena membantu dalam mengatur dan mengelola perubahan kode dengan lebih efektif. Proses ini bervariasi tergantung pada alur kerja dan alat yang digunakan, seperti GitHub, GitLab, atau Bitbucket.</p>
<h2 id="343-git-checkout">3.4.3 git checkout</h2>
<p>Ketika kamu bekerja dengan proyek di Git, kamu seringkali akan bekerja di berbagai <em>branch</em>. Branching memungkinkan kamu untuk menjaga kode yang sedang dikembangkan tetap terpisah dari kode utama (biasanya disebut sebagai <em>master</em> atau <em>main</em> branch). Untuk pindah dari satu branch ke branch lain, kamu dapat menggunakan perintah <code>git checkout</code>.</p>
<p><strong>Contoh:</strong></p>
<pre><code class="language-bash">git checkout nama_branch
</code></pre>
<p>Misalnya, jika kamu ingin pindah ke branch dengan nama &quot;feature-xyz&quot;:</p>
<pre><code class="language-bash">git checkout feature-xyz
</code></pre>
<h3 id="git-checkout-untuk-merge"><strong>Git Checkout untuk Merge</strong></h3>
<p>Setelah kamu selesai mengembangkan suatu fitur atau perbaikan bug di branch yang terpisah, langkah selanjutnya adalah menggabungkan perubahan tersebut kembali ke branch utama, biasanya dengan menggunakan <em>merge</em>.</p>
<p><strong>Contoh:</strong></p>
<ol>
   <li>
      <p>Pertama, pastikan kamu berada di branch target (biasanya branch utama):</p>
      <pre><code class="language-bash">git checkout main
</code></pre>
   </li>
   <li>
      <p>Kemudian, gunakan perintah <code>git merge</code> untuk menggabungkan perubahan dari branch lain ke branch target:</p>
      <pre><code class="language-bash">git merge nama_branch
</code></pre>
   </li>
</ol>
<p><strong>Catatan:</strong> Kadang-kadang, proses merge dapat menghasilkan konflik jika ada perubahan yang bertentangan antara dua branch. Dalam hal ini, kamu perlu mengatasi konflik secara manual sebelum melanjutkan proses merge.</p>
<p>Perintah <code>git checkout</code> memainkan peran penting dalam pengelolaan <em>branch</em> dan proses <em>merge</em> dalam Git. Ini memungkinkan kamu untuk berpindah antara <em>branch</em> yang berbeda dan juga memfasilitasi penggabungan perubahan dari satu <em>branch</em> ke <em>branch</em> lainnya. Dengan memahami cara menggunakan perintah ini, kamu dapat lebih efektif dalam mengelola pengembangan perangkat lunak menggunakan Git.</p>
<h2 id="344-git-merge">3.4.4 git merge</h2>
<p>Merge adalah proses menggabungkan perubahan dari satu branch ke branch lain. Setelah kamu telah selesai mengembangkan suatu fitur atau perbaikan dalam sebuah branch, kamu dapat menggabungkannya kembali ke branch utama agar perubahan tersebut tersedia untuk pengguna lain atau pada lingkungan produksi.</p>
<h3 id="menggunakan-merge">Menggunakan Merge</h3>
<p>Misalkan kamu memiliki branch fitur-baru yang berisi perubahan yang ingin digabungkan ke branch master. Berikut adalah langkah-langkah untuk melakukan merge:</p>
<ol>
   <li>Pastikan kamu berada di branch target (dalam hal ini, branch master) dengan menggunakan perintah <code>git checkout master</code>.</li>
   <li>Jalankan perintah merge dengan menggunakan <code>git merge</code> diikuti dengan nama branch yang ingin digabungkan, yaitu <code>fitur-baru</code>:</li>
</ol>
<pre><code>git merge fitur-baru
</code></pre>
<ol>
   <li>Git akan mencoba menggabungkan perubahan dari branch fitur-baru ke branch master. Jika tidak ada konflik, Git akan otomatis melakukan merge. Jika ada konflik, kamu harus memecahkannya secara manual.</li>
</ol>
<h3 id="resolusi-konflik">Resolusi Konflik</h3>
<p>Konflik terjadi ketika Git tidak dapat secara otomatis menggabungkan perubahan dari dua branch. Ini mungkin terjadi jika kedua branch mengubah bagian yang sama dari kode sumber. Untuk memecahkan konflik, kamu perlu membuka file yang konflik, memodifikasi kode sumber secara manual untuk menggabungkan perubahan, dan kemudian menyimpannya. Setelah semua konflik dipecahkan, kamu dapat melanjutkan merge dengan menjalankan <code>git merge --continue</code>.</p>
<p>Merge memungkinkan perubahan yang dikembangkan dalam branch terpisah untuk diintegrasikan kembali ke branch utama dengan pengawasan yang lebih baik terhadap perubahan tersebut.</p>
<h2 id="345-branch-management-best-practices">3.4.5 Branch Management Best Practices</h2>
<p>Branching dan merging adalah konsep kunci dalam pengembangan perangkat lunak menggunakan sistem kontrol versi seperti Git. Mereka memungkinkan tim pengembang untuk bekerja secara terpisah pada fitur-fitur atau perbaikan bug yang berbeda dan kemudian menggabungkan perubahan mereka ke dalam kode induk (main branch). Namun, untuk menjaga kelancaran kolaborasi dan mencegah konflik yang merugikan, ada beberapa praktik terbaik yang perlu diikuti. Berikut adalah beberapa best practices pada materi branch dan merge:</p>
<p><strong>1. Gunakan Branching Strategy yang Tepat:</strong></p>
<ul>
   <li>Pilih strategi branching yang sesuai dengan ukuran dan dinamika tim kamu. Beberapa strategi umum meliputi Git Flow, GitHub Flow, dan GitLab Flow.</li>
   <li>Pastikan setiap fitur, perbaikan bug, atau tugas memiliki branch terpisah untuk mengisolasi perubahan tersebut.</li>
</ul>
<p><strong>2. Beri Nama Branch dengan Jelas:</strong></p>
<ul>
   <li>Gunakan nama branch yang deskriptif dan menggambarkan tujuan perubahan tersebut.</li>
   <li>Hindari penggunaan nama branch yang umum seperti &quot;fix&quot;, &quot;feature&quot;, atau &quot;update&quot; tanpa konteks tambahan.</li>
</ul>
<p><strong>3. Kerjakan Perubahan dalam Branch Kecil:</strong></p>
<ul>
   <li>Pecah tugas besar menjadi tugas kecil yang dapat diselesaikan dalam branch terpisah.</li>
   <li>Branch yang lebih kecil meminimalkan risiko konflik dan memudahkan debugging.</li>
</ul>
<p><strong>4. Lindungi Branch Default:</strong></p>
<ul>
   <li>Lindungi branch default (misalnya, main atau master) dari perubahan langsung untuk mencegah perubahan yang tidak diotorisasi.</li>
   <li>Setiap perubahan harus dimasukkan melalui pull request (atau merge request) untuk melalui proses review.</li>
</ul>
<p><strong>5. Review Code Secara Teratur:</strong></p>
<ul>
   <li>Setiap perubahan harus melalui proses code review sebelum digabungkan.</li>
   <li>Review code membantu memastikan kualitas dan keseragaman kode serta mengidentifikasi potensi masalah.</li>
</ul>
<p><strong>6. Resolve Konflik dengan Bijak:</strong></p>
<ul>
   <li>Jika terjadi konflik saat melakukan merge, segera atasi konflik tersebut.</li>
   <li>Hindari mengabaikan atau menimpa perubahan dari branch lain tanpa pertimbangan.</li>
</ul>
<p><strong>7. Automatisasi dengan Continuous Integration (CI):</strong></p>
<ul>
   <li>Terapkan praktik CI dengan alat seperti Jenkins, Travis CI, atau GitHub Actions.</li>
   <li>CI memastikan bahwa perubahan baru diuji secara otomatis sebelum digabungkan, mencegah perubahan yang rusak masuk ke branch utama.</li>
</ul>
<p><strong>8. Gunakan Rebase dengan Bijak:</strong></p>
<ul>
   <li>Saat menggabungkan perubahan dari branch lain, pertimbangkan untuk menggunakan rebase daripada merge tradisional.</li>
   <li>Rebase membantu menjaga sejarah commit yang lebih bersih dan lebih mudah diikuti.</li>
</ul>
<p><strong>9. Hapus Branch yang Tidak Diperlukan:</strong></p>
<ul>
   <li>Setelah branch telah digabungkan dan tidak lagi diperlukan, hapus branch tersebut untuk menjaga kerapihan repository.</li>
   <li>Hati-hati saat menghapus branch, pastikan tidak ada perubahan penting yang hilang.</li>
</ul>
<p><strong>10. Komunikasikan dengan Tim:</strong></p>
<ul>
   <li>Selalu komunikasikan perubahan yang kamu kerjakan dengan anggota tim.</li>
   <li>Diskusikan apabila terdapat perubahan besar yang berdampak pada pengembangan secara keseluruhan.</li>
</ul>
<h3 id="konflik-dan-manajemen-konflik">Konflik dan Manajemen Konflik</h3>
<p>Ketika menggabungkan perubahan, terkadang bisa terjadi <strong>konflik</strong>. Konflik muncul ketika ada perubahan yang bertentangan antara dua cabang yang sedang digabungkan. Ini bisa terjadi misalnya ketika dua pengembang mengubah bagian yang sama dari kode. Manajemen konflik melibatkan penyelesaian perbedaan ini dengan memilih versi mana yang harus dipertahankan atau bagaimana menggabungkan perubahan tersebut.</p>
<p>Dalam pengembangan perangkat lunak modern, pemahaman tentang konsep branching dan merging adalah kunci untuk kolaborasi yang efisien dan pengelolaan perubahan yang baik. Dengan menggunakan alat versi kontrol seperti Git, tim pengembang dapat bekerja secara paralel tanpa mengorbankan stabilitas kode sumber utama. Dalam materi ini, kami akan menjelaskan lebih detail tentang praktik terbaik dalam pembuatan cabang, penggabungan, penanganan konflik, dan strategi branching yang umum digunakan dalam pengembangan perangkat lunak.</p>
<p>Dengan mengikuti praktik-praktik terbaik ini, tim pengembang dapat bekerja lebih efisien, mencegah masalah konflik yang berpotensi merugikan, dan memastikan bahwa pengembangan perangkat lunak berjalan dengan lancar.</p>