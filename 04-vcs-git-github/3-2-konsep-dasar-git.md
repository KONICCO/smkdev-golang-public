<div align="center">
   <h1>Module 3 Lesson 2 - Konsep Dasar Git</h1>
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
         <td>3.2.1 Sejarah Git</td>
         <td><a href="#331-sejarah-git">🔽</a></td>
      </tr>
      <tr>
         <td>3.2.2 Instalasi</td>
         <td><a href="#332-instalasi">🔽</a></td>
      </tr>
      <tr>
         <td>3.2.3 Inisiasi repo</td>
         <td><a href="#333-inisiasi-user">🔽</a></td>
      </tr>
      <tr>
         <td>3.2.4 Git Command</td>
         <td><a href="#334-git-command">🔽</a></td>
      </tr>
   </tbody>
</table>
<p><br><br></p>
<h1 id="durasi-pembelajaran">Durasi Pembelajaran</h1>
<p><strong>60 menit</strong></p>
<br>
<h2 id="321-sejarah-git">3.2.1 Sejarah Git</h2>
<p>Kernel Linux adalah proyek perangkat lunak sumber terbuka dengan cakupan yang cukup luas. Selama tahun-tahun awal perawatan kernel Linux (1991–2002), perubahan pada perangkat lunak tersebut disebarluaskan dalam bentuk &quot;patch&quot; dan file arsip. Pada tahun 2002, proyek kernel Linux mulai menggunakan DVCS (Distributed Version Control System) berlisensi tertutup yang disebut BitKeeper.</p>
<div align= "center">
   <img src="https://pi.tedcdn.com/r/talkstar-photos.s3.amazonaws.com/uploads/f69760c8-2dea-43ce-9594-f02ca4175946/LinusTorvalds_2016-embed.jpg?u%5Br%5D=2&u%5Bs%5D=0.5&u%5Ba%5D=0.8&u%5Bt%5D=0.03&quality=82w=640">
   <br>
   <figcaption>Linus Torvalds. Pembuat Distributed Version Control Systems Terkenal, Git</figcaption>
</div>
<br>
<p>Pada tahun 2005, hubungan antara komunitas yang mengembangkan kernel Linux dan perusahaan komersial yang mengembangkan BitKeeper mengalami keretakan, dan status gratis alat ini dicabut. Hal ini mendorong komunitas pengembangan Linux (terutama Linus Torvalds, pencipta Linux) untuk mengembangkan alat mereka sendiri berdasarkan pengalaman yang mereka dapatkan saat menggunakan BitKeeper. Beberapa tujuan dari sistem baru ini adalah sebagai berikut:</p>
<ul>
   <li>Kecepatan</li>
   <li>Desain yang sederhana</li>
   <li>Dukungan kuat untuk pengembangan non-linear (ribuan cabang paralel)</li>
   <li>Sepenuhnya terdistribusi (Distributed Version Control System)</li>
   <li>Dapat mengelola proyek besar seperti kernel Linux secara efisien (kecepatan dan ukuran data)</li>
</ul>
<p>Sejak lahirnya pada tahun 2005, Git telah berkembang dan matang menjadi mudah digunakan dan tetap mempertahankan kualitas awal ini. Git sangat cepat, sangat efisien dalam mengelola proyek-proyek besar, dan memiliki sistem percabangan yang luar biasa untuk pengembangan non-linear.</p>
<h2 id="322-instalasi">3.2.2 Instalasi</h2>
<p>Sebelum kita mengenal dan menggunakan Git sendiri, kita perlu menginstall Git ke dalam komputer kita. Namun, apabila sebelumnya sudah terinstal ada baiknya untuk mengupdate Git tersebut ke versi terbaru.</p>
<h3 id="git---windows">Git - Windows</h3>
<p>Jika kamu pengguna sistem operasi Windows, kamu bisa langsung menginstal Git sendiri melalui <a href="https://git-scm.com/download/win">https://git-scm.com/download/win</a> dan proses download akan segera dilakukan secara otomatis.</p>
<h3 id="git---linux">Git - Linux</h3>
<p>Bagi kamu pengguna Linux dan distribusi nya, kamu bisa menginstalnya dengan beberapa cara. Namun untuk saat ini, kita mencoba untuk menginstal Git melalui command line agar kita semakin familiar dengan interaksi terminal (lihat lesson 2: interaksi terminal pada module Operating System dan Terminal). Jika kamu menggunakan distribusi Fedora dan sejenis nya yang menggunakan package dnf, kamu bisa mengikuti langkah ini:</p>
<pre><code class="language-bash">sudo dnf install git-all
</code></pre>
<p>Jika kamu menggunakan distribusi Debian dan sejenisnya dengan package apt,  kamu bisa menginstall nya seperti berikut:</p>
<pre><code class="language-bash">sudo apt install git-all
</code></pre>
<p>Agar proses penginstalan berjalan dengan lancar, ada baiknya untuk kamu mengupdate dan mengupgrade package pada distribusi Linux yang kamu miliki. Jika proses instalasi telah selesai, kamu bisa cek hasil instalasi dengan melakukan interaksi terminal dengan command:</p>
<pre><code class="language-bash">git --version
</code></pre>
<h2 id="323-inisiasi-user">3.2.3 Inisiasi User</h2>
<p>Setelah proses instalasi telah selesai, kita lanjut untuk melakukan inisiasi user pada local repository, yakni dari komputer kita sendiri. Ada beberapa command yang perlu kamu perhatikan agar proses inisiasi user berjalan lancar. Ini akan membantu kamu ketika akan lanjut untuk proses development yang kolaboratif dan menuntut kamu untuk melakukan dokumentasi dari proses penyimpanan source code pada remote repository.</p>
<p>Lakukan inisiasi secara global untuk username. Untuk username sendiri usahakan sama dengan username yang ada/akan ada pada GitHub kamu. Kamu bisa lakukan proses inisiasi global username kamu dengan command ini:</p>
<pre><code class="language-bash">git config --global user.name “&lt;username github kamu&gt;”
</code></pre>
<p>Lakukan inisiasi secara global untuk user email. Untuk user email sendiri usahakan sama dengan user email yang ada/akan ada pada GitHub kamu. Kamu bisa lakukan proses inisiasi global user email kamu dengan command ini:</p>
<pre><code class="language-bash">git config --global user.email “&lt;email github kamu&gt;”
</code></pre>
<p>Proses inisiasi user telah selesai. Kita akan lanjut untuk berkenalan dengan beberapa command pada git yang biasa dan fundamental digunakan dalam proses development dan dokumentasi source code kamu pada local repositories maupun remote repositories</p>
<h2 id="324-git-command">3.2.4 Git Command</h2>
<p>Pada lesson kali ini, akan ada beberapa command yang akan kita pelajari dan akan membantu kamu dalam berinteraksi dengan local repositories dan remote repositories. Berikut command pada git yang akan kita pelajari:</p>
<ul>
   <li><strong>git init</strong>: command ini akan membantu kamu untuk menyimpan/mengunduh remote repositories yang ada ke local repositories kamu</li>
   <li><strong>git clone</strong>: command ini akan membantu kamu untuk menyimpan/mengunduh remote repositories yang ada ke local repositories kamu</li>
   <li><strong>git add</strong>: command ini akan membantu kamu untuk memantau file source code dan perubahan yang ada pada local repositories berdasarkan history yang telah ada.</li>
   <li><strong>git commit</strong>: command ini akan membantu kamu untuk menyimpan history dan juga perubahan dari file source code yang telah dilakukan di local repositories</li>
   <li><strong>git push</strong>: command ini akan membantu kamu untuk menyimpan perubahan dan file source code ke remote repositories, seperti GitHub, Gitlab, atau sebagainya.</li>
   <li><strong>git pull</strong>: command ini akan membantu kamu untuk menyimpan/mengunduh perubahan yang ada pada remote repositories ke local repositories. Proses ini merupakan keterbalikan dari git pull</li>
   <li><strong>git branch</strong>: command ini akan membantu kamu untuk membuat salinan atau cabang dari proyek Git yang sedang berjalan. Fitur ini akan membantu kamu dalam proses klasifikasi pengembangan dan sejenisnya.</li>
   <li><strong>git log</strong>: command ini akan membantu kamu untuk melihat catatan atau riwayat perubahan yang telah dilakukan pada suatu repositori</li>
</ul>
<br>
<div align= "center">
   <img src="https://www.cs.swarthmore.edu/~adanner/help/git/git-repos.svg">
   <br>
   <figcaption>Diagram Proses Local Repositories ke Remote Repositories</figcaption>
</div>