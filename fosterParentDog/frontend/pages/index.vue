<template>
  <div class="top">
    <div class="container">
      <h1>里親募集 掲示板</h1>
    </div>
    <!-- レコード表示 -->
    <div class="container" v-for="post in posts" v-bind:key="post.id">
      <div class='col-xs-12'>
        <img class='float-left'
             style='padding:0;margin:0 15px 0 0;'
             v-bind:src="post.top_image_path"
             width="130"
             height="130">
      </div>
      <div class="row">
        <nuxt-link :to="`/detail/${post.id}`">
          #{{ post.id }} {{ post.dog_name }}
        </nuxt-link>
      </div>
      <div class="row"><p>犬種 : {{ post.breed }}</p></div>
      <div class="row"><p>性別 : {{ $getLabel(genderMap, post.gender) }}</p></div>
      <div class="row"><p>自己紹介 : {{ post.introduction }}</p></div>
      <div class="row"><p>投稿日時 : {{ post.created_at }}</p></div>
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
        @currentPage="getIndex"
      ></pagenation>
    </div>

    <!--  テンプレート終わり  -->
  </div>
</template>

<script>
import pagenation from "../components/pagenation";
import genderMap from '@/assets/json/gender.json';

const axios = require('axios');

export default {
  components: {
    pagenation,
  },

  mounted() {
    this.getIndex(this.currentPage).then(this.generatePagination());
  },

  data() {
    return {
      genderMap: genderMap,

      posts: [], // 投稿記事
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
      return this.posts
    },
  },

  methods: {
    /**
     * 指定したページの投稿一覧を取得する
     * @param {int} pageNum
     */
    getIndex(page) {
      return new Promise((resolve, reject) => {
        axios.get('api/index', {
          params: {
            page: page,
          }
        }).then((response) => {
          if ((response.status !== 200)) {
            throw new Error('レスポンスエラー')
          } else {
            this.posts = response.data;
          }
        })
      })
    },

    /**
     * ページネーションのため
     * ・公開済記事総数を取得し、設定する。
     * ・総ページ数を算出し、設定する。
     * @param {int} postID
     */
    generatePagination() {
      return new Promise((resolve, reject) => {
        axios.get('api/published_post_count').then((response) => {
          if ((response.status !== 200)) {
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
