<template>
  <div class="top">
    <h1>里親募集 掲示板</h1>
    <!--  投稿記事 一覧表示  -->
    <div class="container" v-for="post in posts" v-bind:key="post.id">
      <ImageTop :imagePath="post.top_image_path"/>
      <div class="goDetail">
        <div class="row">
          <!--   クエリに投稿IDをセットし/detail?postId=XXXへルーティング     -->
          <NuxtLink :to="{path: '/detail', query: {postId: `${post.id}`}}">
            #{{ post.id }} {{ post.dog_name }}
          </NuxtLink>
        </div>
      </div>
      <ViewOrTextBox :title="'犬種'"
                     :name="post.breed"
                     :isView="true"/>
      <LabelOrDropdown :title="'性別'"
                       :mapName="gender_map"
                       :itemValue="post.gender"
                       :isView="true"/>
      <ViewOrTextBox :title="'自己紹介'"
                     :name="post.introduction"
                     :isView="true"/>
      <ViewOnly :title="'投稿日時'"
                :name="post.created_at"
                :isView="true"/>
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
import Pagenation from "../components/Pagenation";
import ViewOrTextBox from "~/components/post/TextOrTextBox";
import LabelOrDropdown from "~/components/post/LabelOrDropdown";
import ViewOnly from "~/components/post/ReadOnly";
import ImageTop from "~/components/post/Image";
import CON from "~/components/const/const";

export default {
  components: {
    Pagenation,
    ViewOrTextBox,
    LabelOrDropdown,
    ViewOnly,
    ImageTop,
    CON,
  },


  mounted() {
    this.getIndex(this.currentPage).then(this.generatePagination());
    console.log("--------this.gender_map");
    console.log(this.gender_map);
  },

  data() {
    return {
      gender_map: CON.data().GENDER,

      posts: [], // 投稿記事

      //ページネーション設定
      currentPage: 1, //現在のページ（初期は1）
      showPages: 5, //ページネーションを何ページ表示するか（奇数でないとずれる）
      perPage: 20, //1ページの表示件数
      totalCount: Number, //取得したレコードの総件数
      totalPages: Number, //算出後の総ページ数

      publishing: {
        private: 0,
        public: 1,
      }
    }
  },
  computed: {},
  methods: {
    /**
     * 指定したページの投稿一覧を取得する
     * @param {int} pageNum
     */
    getIndex(page) {
      return new Promise((resolve, reject) => {
        this.$axios.get('api/index',
          {
            params: {
              page: page,
              publishing: this.publishing.public,
            }
          }
        )
          //
          // {params: {page: page, publishing: 0}})
          .then((response) => {
            if ((response.status !== 200)) {
              console.error(`Error:${response.statusText}, ${this.getIndex.name}`)
            } else {
              this.posts = response.data;
              console.log("-------this.posts");
              console.log(this.posts);
            }
          }).catch(err => alert(err));
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
        this.$axios.get('api/post_count',
          {
            params: {
              publishing: this.publishing.public,
            }
          }
        )
          .then((response) => {
            if ((response.status !== 200)) {
              console.error(`Error:${response.statusText}, ${this.generatePagination.name}`)
            } else {
              this.totalCount = response.data.count; // 公開済記事数 ex: 41
              this.totalPages = Math.ceil(this.totalCount / this.perPage); // 総ページ数 ex: 3
            }
          }).catch(err => alert(err));
      })
    },
  },
}
</script>

<!-- cssはassetsから自動で読み込む-->
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
