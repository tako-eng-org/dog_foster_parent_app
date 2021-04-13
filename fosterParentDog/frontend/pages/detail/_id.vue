<template>
  <div class="detail">
    <div class="container" v-for="post in posts" v-bind:key="post.id">
      <div class='col-xs-12'>
        <img class='float-left'
             style='padding:0;margin:0 15px 0 0;'
             v-bind:src="post.top_image_path"
             width="180"
             height="180">
      </div>
      <div class="row"><p>投稿No : {{ post.id }}</p></div>
      <div class="row"><p>犬の名前 : {{ post.dog_name }}</p></div>
      <div class="row"><p>犬種 : {{ post.breed }}</p></div>
      <div class="row"><p>性別 : {{ $getLabel(genderMap, post.gender) }}</p></div>
      <div class="row"><p>去勢/避妊手術 : {{ $getLabel(spayMap, post.spay) }}</p></div>
      <div class="row"><p>年齢 : {{ post.old }}</p></div>
      <div class="row"><p>単身者への譲渡 : {{ $getLabel(singlePersonMap, post.single_person) }}</p></div>
      <div class="row"><p>高齢者への譲渡 : {{ $getLabel(seniorPersonMap, post.senior_person) }}</p></div>
      <div class="row"><p>譲渡ステータス : {{ $getLabel(transferStatusMap, post.transter_status) }}</p></div>
      <div class="row"><p>自己紹介 : {{ post.introduction }}</p></div>
      <div class="row"><p>投稿日時 : {{ post.created_at }}</p></div>
      <div class="row"><p>性格アピールポイント : {{ post.appeal_point }}</p></div>
      <div class="row"><p>譲渡可能都道府県 : </p>
        <div class="transferable" v-for="postPrefecture in postPrefectures" v-bind:key="postPrefecture.post_id">
          <p>{{ $getLabel(prefectureMap, postPrefecture.post_prefecture_id) }} </p>
        </div>
      </div>
      <div class="row"><p>健康状態や譲渡条件などの特記事項 : {{ post.other_message }}</p></div>
    </div>

    <div class="container" v-for="imagePath in imagePaths" v-bind:key="imagePath.post_id">
      <div class='col-xs-12'>
        <img class='float-left'
             style='padding:0;margin:0 15px 0 0;'
             v-bind:src="imagePath.image_path"
             width="180"
             height="180">
      </div>
    </div>

    <div class="container" v-bind:key="user.id">
      <div class="row"><p>投稿者情報</p></div>
      <div class="row"><p>ニックネーム : {{ user.nickname }}</p></div>
      <div class="row"><p>プロフィール : {{ user.profile }}</p></div>
    </div>
  </div>
</template>

<script>
const axios = require('axios');
process.env.DEBUG = 'nuxt:*' // nuxt.jsについてログ出力する

import genderMap from '@/assets/json/gender.json'
import prefectureMap from '@/assets/json/prefecture.json'
import seniorPersonMap from '@/assets/json/senior_person.json'
import singlePersonMap from '@/assets/json/single_person.json'
import spayMap from '@/assets/json/spay.json'
import transferStatusMap from '@/assets/json/transfer_status.json'


export default {
  data: function () {
    return {
      //storeの共通資材
      genderMap: genderMap,
      prefectureMap: prefectureMap,
      seniorPersonMap: seniorPersonMap,
      singlePersonMap: singlePersonMap,
      spayMap: spayMap,
      transferStatusMap: transferStatusMap,

      posts: [], // 投稿記事
      imagePaths: [], //画像URL
      postPrefectures: [], //譲渡可能都道府県
      user: {}, //ユーザー情報
    }
  },
  components: {},
  computed: {},
  mounted: function () {
    this.getDetail(this.$route.params['id']);
    this.getPostImagePaths(this.$route.params['id']);
    this.getPostPrefecture(this.$route.params['id']);
    this.getUser(this.$route.params['id']);
  },

  methods: {
    /**
     * 投稿IDに紐づいた基礎投稿情報を取得する
     * @param {int} postID
     */
    getDetail(currentPostId) {
      axios.get('api/post', {
        params: {postId: currentPostId,}
      }).then((response) => {
        if ((response.status !== 200)) {
          throw new Error('レスポンスエラー')
        } else {
          this.posts = response.data;
        }
      })
    },

    /**
     * 投稿IDに紐づいた画像パスを取得する
     * @param {int} postID
     */
    getPostImagePaths(currentPostId) {
      axios.get('api/images', {
        params: {postId: currentPostId,}
      }).then((response) => {
        if ((response.status !== 200)) {
          throw new Error('レスポンスエラー')
        } else {
          this.imagePaths = response.data;
        }
      })
    },

    /**
     * 投稿IDに紐づいた譲渡可能都道府県を取得する
     * @param {int} postID
     */
    getPostPrefecture(currentPostId) {
      axios.get('api/post_prefecture', {
        params: {postId: currentPostId,}
      }).then((response) => {
        if ((response.status !== 200)) {
          throw new Error('レスポンスエラー')
        } else {
          this.postPrefectures = response.data;
        }
      })
    },

    /**
     * 投稿IDのユーザーIDに紐づいたユーザー情報を取得する
     * @param {int} postID
     */
    getUser(currentPostId) {
      axios.get('api/user', {
        params: {postId: currentPostId,}
      }).then((response) => {
        if ((response.status !== 200)) {
          throw new Error('レスポンスエラー')
        } else {
          this.user = response.data;
        }
      })
    }

  },
}
</script>

<!-- cssはassetsから自動で読み込む-->
<style></style>

