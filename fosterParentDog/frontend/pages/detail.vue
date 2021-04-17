<template>
  <div class="detail">
    <div class="container" v-bind:key="post.id">
      <ImageObj :imagepath="post.top_image_path"/>
      <div class="row"><p>投稿No : {{ post.id }}</p></div>
      <div class="row"><p>犬の名前 : {{ post.dog_name }}</p></div>
      <div class="row"><p>犬種 : {{ post.breed }}</p></div>
      <div class="row"><p>性別 : {{ $getLabel($GENDER, post.gender) }}</p></div>
      <div class="row"><p>去勢/避妊手術 : {{ $getLabel($SPAY, post.spay) }}</p></div>
      <div class="row"><p>年齢 : {{ post.old }}</p></div>
      <div class="row"><p>単身者への譲渡 : {{ $getLabel($SINGLE_PERSON, post.single_person) }}</p></div>
      <div class="row"><p>高齢者への譲渡 : {{ $getLabel($SENIOR_PERSON, post.senior_person) }}</p></div>
      <div class="row"><p>譲渡ステータス : {{ $getLabel($TRANSFER_STATUS, post.transter_status) }}</p></div>
      <div class="row"><p>自己紹介 : {{ post.introduction }}</p></div>
      <div class="row"><p>投稿日時 : {{ post.created_at }}</p></div>
      <div class="row"><p>性格アピールポイント : {{ post.appeal_point }}</p></div>
      <div class="row"><p>譲渡可能都道府県 : </p>
        <div class="transferable" v-for="postPrefecture in postPrefectures" v-bind:key="postPrefecture.post_id">
          <p>{{ $getLabel($PREFECTURE, postPrefecture.post_prefecture_id) }} </p>
        </div>
      </div>
      <div class="row"><p>健康状態や譲渡条件などの特記事項 : {{ post.other_message }}</p></div>
    </div>

    <div class="container" v-for="imagePath in imagePaths" v-bind:key="imagePath.post_id">
      <ImageObj :imagepath="imagePath.image_path"/>
    </div>

    <div class="container" v-bind:key="user.id">
      <div class="row"><p>投稿者情報</p></div>
      <div class="row"><p>ニックネーム : {{ user.nickname }}</p></div>
      <div class="row"><p>プロフィール : {{ user.profile }}</p></div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
import ImageObj from "~/components/Image.vue";

export default {
  data: function () {
    return {
      post: {}, //基礎投稿
      imagePaths: {}, //画像パスリスト
      postPrefectures: {}, //譲渡可能都道府県リスト
      user: {}, //ユーザー情報
    }
  },
  components: {
    ImageObj,
  },
  computed: {},
  mounted: function () {
    this.getDetail(this.$route.params['id']);
  },

  methods: {
    /**
     * 投稿IDに紐づいた投稿記事を取得する
     * @param {int} postID
     */
    getDetail(currentPostId) {
      axios.get('api/post', {
        params: {postId: currentPostId,}
      }).then((response) => {
        if ((response.status !== 200)) {
          console.error(`Error:${response.statusText}, ${this.getDetail.name}`)
        } else {
          this.post = response.data.post;
          this.imagePaths = response.data.post_images;
          this.postPrefectures = response.data.post_prefectures;
          this.user = response.data.user;
        }
      }).catch(err => alert(err));
    },
  },
}
</script>

<!-- cssはassetsから自動で読み込む-->
<style scoped></style>

