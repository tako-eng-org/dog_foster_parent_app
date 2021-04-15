<template>
  <div class="top">
    <div class="container">
      <h1>里親募集 掲示板</h1>
    </div>
    <!-- レコード表示 -->
    <div class="container" v-for="post in posts" v-bind:key="post.id">
      <ImageObj :imagepath="post.top_image_path"/>
      <div class="row">
        <nuxt-link :to="`/detail/${post.id}`">
          #{{ post.id }} {{ post.dog_name }}
        </nuxt-link>
      </div>
      <div class="row"><p>犬種 : {{ post.breed }}</p></div>
      <div class="row"><p>性別 : {{ $getLabel($GENDER, post.gender) }}</p></div>
      <div class="row"><p>自己紹介 : {{ post.introduction }}</p></div>
      <div class="row"><p>投稿日時 : {{ post.created_at }}</p></div>
      <br>
    </div>

    <!--  ページ数表示  -->
    <div class="container">
      <Pagenation
        :showPages="showPages"
        :currentPage="currentPage"
        :totalCount="totalCount"
        :perPage="perPage"
        :totalPages="totalPages"
        @currentPage="getIndex"
      ></Pagenation>
    </div>

  </div>
</template>

<script>
import axios from 'axios';
import Pagenation from "../components/Pagenation";
import ImageObj from "@/components/Image.vue";

export default {
  components: {
    Pagenation,
    ImageObj,
  },

  mounted() {
    this.getIndex(this.currentPage).then(this.generatePagination());
    console.log()
  },

  data() {
    return {
      posts: [], // 投稿記事

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
            // TODO エラー時メッセージをわかりやすい文へ
            throw new Error(response.statusText)
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
            throw new Error(response.statusText)
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
