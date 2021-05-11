<template>
  <div class="top">
    <h1>里親募集 掲示板</h1>
    <!--  投稿記事 一覧表示  -->
    <div class="container" v-for="post in posts" :key="post.id">
      <ImageOne :objectUrl="getIndexImagePath(post)"/>
      <div class="goDetail row">
        <!--   クエリに投稿IDをセットし/detail?postId=XXXへルーティング     -->
        <NuxtLink :to="{path: '/detail', query: {postId: `${post.id}`}}">
          #{{ post.id }} {{ post.dog_name }}
        </NuxtLink>
      </div>
      <TextBox :title="'犬種'" v-model="post.breed"/>
      <Gender v-model="post.gender"/>
      <TextBox :title="'自己紹介'" v-model="post.introduction"/>
      <TextBox :title="'投稿日時'" v-model="post.created_at"/>
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
import Pagenation from "~/components/Pagenation";
import TextBox from "~/components/post/TextBox";
import ImageOne from "~/components/post/ImageOne";
import Gender from "~/components/post/Gender";
import {publishingList} from "~/consts/publishingList";

export default {
  components: {
    Pagenation,
    TextBox,
    ImageOne,

    Gender,
  },

  mounted() {
    this.getIndex(this.currentPage).then(this.generatePagination());
  },

  data() {
    return {
      objectUrlBase: "https://bbsapp-img.s3.us-east-2.amazonaws.com/",

      // TODO: 初期値を書いた方が良いか確認すること
      posts: [], // 投稿記事

      //ページネーション設定
      currentPage: 1, //現在のページ（初期は1）
      showPages: 5, //ページネーションを何ページ表示するか（奇数でないとずれる）
      perPage: 20, //1ページの表示件数
      totalCount: 0, //取得したレコードの総件数
      totalPages: 0, //算出後の総ページ数
    }
  },
  computed: {},
  methods: {
    /**
     * 1投稿のpositionが最小値(原則0)のobjectUrlを取得する
     * @param {object} post
     */
    getIndexImagePath(post) {
      let min = Math.min(...post.post_images.map(x => x.position));
      let objectKeyByMinPosition = post.post_images.filter(e => (e.position === min))[0].objectKey;
      return this.objectUrlBase + objectKeyByMinPosition
    },
    /**
     * 指定したページの投稿一覧を取得する
     * @param {int} pageNum
     */
    getIndex(page) {
      return new Promise((resolve, reject) => {
        this.$axios.get('/api/index',
          {
            params: {
              page: page,
              publishing: publishingList.public,
            }
          }
        )
          .then((response) => {
            if ((response.status !== 200)) {
              console.error(`Error:${response.statusText}, ${this.getIndex.name}`)
            } else {
              this.posts = response.data;
            }
          }).catch(err => console.error(err.response));
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
        this.$axios.get('/api/post_count',
          {
            params: {
              publishing: publishingList.public,
            }
          }
        )
          .then((response) => {
            if ((response.status !== 200)) {
              console.error(`Error:${response.statusText}, ${this.generatePagination.name}`)
            } else {
              this.totalCount = response.data.count; // 公開/非公開記事数 ex: 41
              this.totalPages = Math.ceil(this.totalCount / this.perPage); // 総ページ数 ex: 3
            }
          }).catch(err => console.error(err.response));
      })
    },
  },
}
</script>

<!-- cssは原則、assetsから自動で読み込む-->
<style scoped>
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
