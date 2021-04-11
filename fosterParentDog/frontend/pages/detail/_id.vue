<template>
  <div className="container">
    <h1>Detail</h1>
    <!--    <div v-bind:post_id_>-->
    <h1>{{ $route.params.id }}</h1>
    <!--    </div>-->
    <div class="container" v-for="record in records" v-bind:key="record.id">
      <div class='col-xs-12'>
        <img class='float-left'
             style='padding:0;margin:0 15px 0 0;'
             v-bind:src="record.top_image_path"
             width="180"
             height="180">
      </div>
      <div class="row"><p>投稿No : {{ record.id }}</p></div>
      <div class="row"><p>犬の名前 : {{ record.dog_name }}</p></div>
      <div class="row"><p>犬種 : {{ record.breed }}</p></div>
      <div class="row"><p>性別 : {{ $getLabel(gender_map, record.gender) }}</p></div>
      <div class="row"><p>去勢/避妊手術 : {{ $getLabel(spay_map, record.spay) }}</p></div>
      <div class="row"><p>年齢 : {{ record.old }}</p></div>
      <div class="row"><p>単身者への譲渡 : {{ $getLabel(single_person_map, record.single_person) }}</p></div>
      <div class="row"><p>高齢者への譲渡 : {{ $getLabel(senior_person_map, record.senior_person) }}</p></div>
      <div class="row"><p>譲渡ステータス : {{ $getLabel(transfer_status_map, record.transter_status) }}</p></div>
      <div class="row"><p>自己紹介 : {{ record.introduction }}</p></div>
      <div class="row"><p>投稿日時 : {{ record.created_at }}</p></div>
      <div class="row"><p>性格アピールポイント : {{ record.appeal_point }}</p></div>
      <div class="row"><p>譲渡可能都道府県 : </p>
        <div class="transferable" v-for="record in transferable_prefectures" v-bind:key="record.post_id">
          <p>{{ $getLabel(prefecture_map, record.post_transferable_prefecture_id) }} </p>
        </div>
      </div>
      <div class="row"><p>健康状態や譲渡条件などの特記事項 : {{ record.other_message }}</p></div>
    </div>

    <div class="container" v-for="record in image_paths" v-bind:key="record.post_id">
      <div class='col-xs-12'>
        <img class='float-left'
             style='padding:0;margin:0 15px 0 0;'
             v-bind:src="record.image_path"
             width="180"
             height="180">
      </div>
    </div>

    <div class="row"><p>投稿者情報</p></div>
    <div class="container" v-for="record in user_profile" v-bind:key="record.id">
      <div class="row"><p>ニックネーム : {{ record.nickname }}</p></div>
      <div class="row"><p>プロフィール : {{ record.profile }}</p></div>
    </div>
  </div>
</template>

<script>
const axios = require('axios');
process.env.DEBUG = 'nuxt:*' // nuxt.jsについてログ出力する

import gender_map from '@/assets/json/gender.json'
import prefecture_map from '@/assets/json/prefecture.json'
import senior_person_map from '@/assets/json/senior_person.json'
import single_person_map from '@/assets/json/single_person.json'
import spay_map from '@/assets/json/spay.json'
import transfer_status_map from '@/assets/json/transfer_status.json'


export default {
  data: function () {
    return {
      //storeの共通資材
      gender_map: gender_map,
      prefecture_map: prefecture_map,
      senior_person_map: senior_person_map,
      single_person_map: single_person_map,
      spay_map: spay_map,
      transfer_status_map: transfer_status_map,

      records: [], // 投稿記事
      image_paths: [], //画像URL
      user_profile: [], //ユーザー情報
      transferable_prefectures: [], //譲渡可能都道府県
    }
  },
  components: {},
  computed: {
    // 表示対象の情報を返却する
    computedRecords() {

      if (this.$route.params.id !== undefined) {
        console.log("idがありません_computed")
      }

      this.doFetchOneRecords(this.$route.params['id']);
      this.doFetchPostImagePaths(this.$route.params['id']);
      this.doFetchTransferablePrefecture(this.$route.params['id']);
    },
  },

  mounted: function () {
    if (this.$route.params.id !== undefined) {
      console.log("idがありません_mounted")
    }
    this.doFetchOneRecords(this.$route.params['id']);
    this.doFetchPostImagePaths(this.$route.params['id']);
    this.doFetchTransferablePrefecture(this.$route.params['id']);
  },

  methods: {
    getLabel(mapName, i) {
      return mapName.find(map => map.value === i).label
    },

    // 1ページに表示する分、レコードを取得する
    doFetchOneRecords(currentPostId) {
      axios.get('/fosterparent/api/post', {
        params: {
          postId: currentPostId,
        }
      }).then((response) => {
        if ((response.status !== 200)) {
          throw new Error('レスポンスエラー')
        } else {
          let responseData = response.data
          // console.log(responseData)
          console.log("111111111111")
          console.log(response.data["transter_status"])
          console.log("111111111111")
          this.records = responseData;
        }
      })
    },

    // 1ページに表示する分、レコードを取得する
    doFetchPostImagePaths(currentPostId) {
      axios.get('/fosterparent/api/images', {
        params: {
          postId: currentPostId,
        }
      }).then((response) => {
        if ((response.status !== 200)) {
          throw new Error('レスポンスエラー')
        } else {
          this.image_paths = response.data;
        }
      })
    },

    // 投稿IDに紐づく譲渡可能都道府県を取得する
    doFetchTransferablePrefecture(currentPostId) {
      axios.get('/fosterparent/api/transferable_prefecture', {
        params: {
          postId: currentPostId,
        }
      }).then((response) => {
        if ((response.status !== 200)) {
          throw new Error('レスポンスエラー')
        } else {
          this.transferable_prefectures = response.data;
        }
      })
    },

    // 1ページに表示する分、ユーザー情報を取得する
    doFetchUserProfile(currentPostId) {
      axios.get('/fosterparent/api/user_profile', {
        params: {
          postId: currentPostId,
        }
      }).then((response) => {
        if ((response.status !== 200)) {
          throw new Error('レスポンスエラー')
        } else {
          this.user_profile = response.data;
        }
      })
    }

  },
}
</script>

<!-- cssはassetsから自動で読み込む-->
<style></style>

