<template>
  <div class="hello">
    <p>
      Bu proje kartaca çekirdekten yetişenler programının 2. aşaması kapsamında adaylara verilen görevlerinden biridir.
      <br>
      <a href="https://kartaca.com/cekirdekten-yetisenler-programi/" target="_blank" rel="noopener">Kartaca Çekirdekten Yetişenler Programı</a>.
    </p>
    <h3>Seçilen: Görev 1</h3>
    <p>Ön tarafta (FE) Vue.js, arka tarafta (BE) Golang kullanarak yazılmış bir proje istenmektedir.</p>
    <p>Gereksinimler:</p>
    <ul>
      <li>-Arayüz, kullanıcı girişi (login) ve post'ların olduğu bir dashboard olarak düşünülebilir.</li>
      <li>-Kimliklendirme (authentication) JWT ile sağlanmalı, kullanıcı bilgileri SQLite gibi bir veritabanı üzerinden, statik olarak eklenmiş şekilde sunulmalıdır.</li>
      <li>-Dashboard'un en üst kısmında bir post ekleme kutucuğu olmalı ve buradan post eklendiğinde herkes yeni post'u görecek şekilde tasarlanmalıdır.</li>
      <li>-Herhangi bir post'a basınca bir popup açılmalı ve içerisinde post detayları (kim, ne zaman vs. gibi bilgiler ile) belirtilmelidir. Post'un içeriğini Twitter post'u gibi düşünebilirsiniz.</li>
    </ul>
    <h3>Eksiklikler</h3>
    <ul>
      <li>-Authentication için JWT kullanılıyor fakat kullanım şeklinin best practice olduğuna emin değilim.</li>
      <li>-Şifre bilgileri unencrypted olarak saklanıyor.</li>
      <li>-Postlara tıklanılınca herhangi bir popup açılmıyor.</li>
      <li>-Post gönderme tarihi objelerde mevcut fakat sayfada render olmuyor.</li>
    </ul>
    <hr>
    <h3>Yararlanınlan Kaynaklar</h3>
    <ul>
      <li>><a href="https://vuejs.org/v2/guide/" target="_blank" rel="noopener">Vue JS Guide</a></li>
      <li>><a href="https://godoc.org" target="_blank" rel="noopener">Documentation for Go packages</a></li>
      <li>><a href="https://www.youtube.com/watch?v=Wy9q22isx3U" target="_blank" rel="noopener">Vue JS Crash Course</a></li>
      <li>><a href="https://www.youtube.com/watch?v=SqrbIlUwR0U" target="_blank" rel="noopener">Go / Golang Crash Course</a></li>
      <li>><a href="https://www.youtube.com/watch?v=SonwZ6MF5BE" target="_blank" rel="noopener">Golang REST API With Mux</a></li>
      <li>><a href="https://github.com/mattn/go-sqlite3/blob/b4f5cc77d1cca1470922e916c9f775ef17d2d78f/sqlite3_test.go" target="_blank" rel="noopener">Golang Sqlite3 Example</a></li>
      <li>><a href="https://www.youtube.com/watch?v=HV7DtH3J2PU" target="_blank" rel="noopener">Animated Login Form Using Only HTML & CSS</a></li>
    </ul>
  </div>
</template>

<script>
export default {
  name: 'about',
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
div[class="hello"]{
  text-align:left
}
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
a {
  color: #DAA520;
  font-weight: bold;
}
a:hover {
  color: rgb(107, 77, 0);
  font-weight: bold;
}
</style>
