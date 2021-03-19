<template>
  <div>
    <div class="container">
      <h1>里親募集 掲示板</h1>
    </div>

    <div class="container">
      <div class="one-contents">
        <div class='col-xs-12'>
          <img class='float-left'
               style='padding:0;margin:0 15px 0 0;'
               src="../assets/image/dog-sample1.png"
               width="130"
               height="130">
        </div>
        <p class="mb-0">No125 はなこ</p>
        <p class="mb-0">犬種: 雑種</p>
        <p class="mb-0">性別: メス</p>
        <p class="mb-0">名前: はなこ</p>
        <p class="mb-0">内容: 元気で可愛いお婆ちゃん犬です。 シニアではありますが、…</p>
        <p class="mb-0">
          <nuxt-link to="./detail">詳細をみる</nuxt-link>
        </p>
      </div>
    </div>
    <br>

    <div class="container">
      <div class="one-contents">
        <div class='col-xs-12'>
          <img class='float-left'
               style='padding:0;margin:0 15px 0 0;'
               src="../assets/image/dog-sample2.png"
               width="130"
               height="130">
        </div>
        <p class="mb-0">No124 たろう</p>
        <p class="mb-0">犬種: 和犬系の雑種</p>
        <p class="mb-0">性別: オス</p>
        <p class="mb-0">名前: たろう</p>
        <p class="mb-0">内容: 元気な男の子です。 人に対してはさほど警戒せず、…</p>
        <p class="mb-0">
          <nuxt-link to="./detail">詳細をみる</nuxt-link>
        </p>
      </div>
    </div>
    <br>

    <div class="container">
      <p class="text-center"> << 1 2 3 4 5 >> </p>
    </div>
  </div>
</template>

<script>
const axios = require('axios');
process.env.DEBUG = 'nuxt:*' // nuxt.jsについてログ出力する
console.log("test...............")
axios.get('/user?ID=12345')
  .then(function (response) {
    // handle success
    console.log(response);
  })
  .catch(function (error) {
    // handle error
    console.log(error);
  })
  .then(function () {
    // always executed
  });

</script>

<!-- cssはassetsから自動で読み込む-->
<style></style>
<!--<template>-->
<!--  <div>-->
<!--    <compo></compo>-->
<!--  </div>-->
<!--</template>-->

<!--<script>-->
<!--process.env.DEBUG = 'nuxt:*' // nuxt.jsについてログ出力する-->

<!--import Compo from '../components/Top'-->

<!--export default {-->
<!--  components: {-->
<!--    compo: Compo,-->
<!--  },-->
<!--}-->
<!--</script>-->
