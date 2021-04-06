<template>
  <div class="top">
    <div class="container">
      <h1>里親募集 掲示板</h1>
    </div>
    <!-- レコード表示 -->
    <div class="container" v-for="record in records" v-bind:key="record.id">
      <div class='col-xs-12'>
        <img class='float-left'
             style='padding:0;margin:0 15px 0 0;'
             v-bind:src="record.top_image_path"
             width="130"
             height="130">
      </div>
      <div class="row">
        <NuxtLink to="/detail">
          #{{ record.id }} {{ record.dog_name }}
        </NuxtLink>
      </div>
      <div class="row"><p>犬種 : {{ record.breed }}</p></div>
      <div class="row"><p>性別 : {{ record.gender }}</p></div>
      <div class="row"><p>自己紹介 : {{ record.introduction }}</p></div>
      <div class="row"><p>投稿日時 : {{ record.created_at }}</p></div>
      <br>
    </div>

    <!--  ページ数表示  -->
    <div class="container">
      <pagenation
        :showPages="showPages"
        :currentPage="currentPage"
        :totalCount="totalCount"
        :perPage="perPage"
        :totalPages="totalPages"
        @currentPage="doFetchIndexRecords"
      ></pagenation>
    </div>

    <!--  テンプレート終わり  -->
  </div>
</template>

<script>
import pagenation from "../components/pagenation";

const axios = require('axios');
process.env.DEBUG = 'nuxt:*' // nuxt.jsについてログ出力する

export default {
  components: {
    pagenation,
  },

  mounted() {
    this.doFetchIndexRecords(this.currentPage).then(this.doSetPagenation());
  },

  data() {
    return {
      records: [], // 投稿記事
      // show: true,

      //ページネーション設定
      currentPage: 1, //現在のページ（初期は1）
      showPages: 5, //ページネーションを何ページ表示するか（奇数でないとずれる）
      perPage: 20, //1ページの表示件数
      totalCount: Number, //取得したレコードの総件数
      totalPages: Number, //算出後の総ページ数
    }
  },

  computed: {
    // 表示対象の情報を返却する
    computedRecords() {
      return this.records
    },
  },

  methods: {
    // 1ページに表示する分、レコードを取得する
    doFetchIndexRecords(page) {
      return new Promise((resolve, reject) => {
        axios.get('/fosterparent/api/index', {
          params: {
            page: page,
          }
        }).then((response) => {
          if ((response.status != 200)) {
            throw new Error('レスポンスエラー')
          } else {
            this.records = response.data;
          }
        })
      })
    },

    // ページネーションのため
    // ・公開済記事総数を取得し、設定する。
    // ・総ページ数を算出し、設定する。
    doSetPagenation() {
      return new Promise((resolve, reject) => {
        axios.get('/fosterparent/api/pageCount').then((response) => {
          if ((response.status != 200)) {
            throw new Error('レスポンスエラー')
          } else {
            this.totalCount = response.data; // 公開済記事数 ex: 41
            this.totalPages = Math.ceil(this.totalCount / this.perPage); // 総ページ数 ex: 3
          }
        })
      })
    },
  },
}
</script>

<!-- cssはassetsから自動で読み込む-->
<style>
.top {
  overflow: hidden;
  width: 100%;
}

/*溢れた文字列を...として省略する設定*/
p {
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
</style>
