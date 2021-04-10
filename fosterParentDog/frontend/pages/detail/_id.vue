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
      <div class="row"><p>性別 : {{ getLabel(map_gender, record.gender) }}</p></div>
      <div class="row"><p>去勢/避妊手術 : {{ getLabel(map_spay, record.spay) }}</p></div>
      <div class="row"><p>年齢 : {{ record.old }}</p></div>
      <div class="row"><p>単身者への譲渡 : {{ getLabel(map_singlePerson, record.single_person) }}</p></div>
      <div class="row"><p>高齢者への譲渡 : {{ getLabel(map_seniorPerson, record.senior_person) }}</p></div>
      <div class="row"><p>譲渡ステータス : {{ getLabel(map_transferStatus, record.transter_status) }}</p></div>
      <div class="row"><p>自己紹介 : {{ record.introduction }}</p></div>
      <div class="row"><p>投稿日時 : {{ record.created_at }}</p></div>
      <div class="row"><p>性格アピールポイント : {{ record.appeal_point }}</p></div>
      <div class="row"><p>譲渡可能都道府県 : </p>
        <div class="transferable" v-for="record in transferable_prefectures" v-bind:key="record.post_id">
          <p>{{ getLabel(map_transferablePrefecture, record.post_transferable_prefecture_id) }} </p>
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

    <div class="container">
      <div class="row"><p>投稿者 : ＊ニックネーム＊</p></div>
    </div>
  </div>
</template>

<script>
const axios = require('axios');
process.env.DEBUG = 'nuxt:*' // nuxt.jsについてログ出力する

export default {
  data: function () {
    return {
      // postIDparam : params.id,
      records: [], // 投稿記事
      image_paths: [], //画像URL(複数可)
      user_profile: [],
      transferable_prefectures: [],

      map_gender: [
        {value: 0, label: "女の子"},
        {value: 1, label: "男の子"},
        {value: 2, label: "その他/不明"},
      ],

      map_spay: [
        {value: 0, label: "手術済"},
        {value: 1, label: "手術未"},
        {value: 2, label: "手術予定"},
      ],

      map_singlePerson: [
        {value: 0, label: "不可"},
        {value: 1, label: "可"},
        {value: 2, label: "応相談"},
      ],

      map_seniorPerson: [
        {value: 0, label: "不可"},
        {value: 1, label: "可"},
        {value: 2, label: "応相談"},
      ],

      map_transferStatus: [
        {value: 0, label: "応募受付中"},
        {value: 1, label: "やりとり中"},
        {value: 2, label: "お試し中"},
        {value: 3, label: "応募終了"},
      ],

      map_transferablePrefecture: [
        {value: 1, label: "北海道"},
        {value: 2, label: "青森県"},
        {value: 3, label: "岩手県"},
        {value: 4, label: "宮城県"},
        {value: 5, label: "秋田県"},
        {value: 6, label: "山形県"},
        {value: 7, label: "福島県"},
        {value: 8, label: "茨城県"},
        {value: 9, label: "栃木県"},
        {value: 10, label: "群馬県"},
        {value: 11, label: "埼玉県"},
        {value: 12, label: "千葉県"},
        {value: 13, label: "東京都"},
        {value: 14, label: "神奈川県"},
        {value: 15, label: "新潟県"},
        {value: 16, label: "富山県"},
        {value: 17, label: "石川県"},
        {value: 18, label: "福井県"},
        {value: 19, label: "山梨県"},
        {value: 20, label: "長野県"},
        {value: 21, label: "岐阜県"},
        {value: 22, label: "静岡県"},
        {value: 23, label: "愛知県"},
        {value: 24, label: "三重県"},
        {value: 25, label: "滋賀県"},
        {value: 26, label: "京都府"},
        {value: 27, label: "大阪府"},
        {value: 28, label: "兵庫県"},
        {value: 29, label: "奈良県"},
        {value: 30, label: "和歌山県"},
        {value: 31, label: "鳥取県"},
        {value: 32, label: "島根県"},
        {value: 33, label: "岡山県"},
        {value: 34, label: "広島県"},
        {value: 35, label: "山口県"},
        {value: 36, label: "徳島県"},
        {value: 37, label: "香川県"},
        {value: 38, label: "愛媛県"},
        {value: 39, label: "高知県"},
        {value: 40, label: "福岡県"},
        {value: 41, label: "佐賀県"},
        {value: 42, label: "長崎県"},
        {value: 43, label: "熊本県"},
        {value: 44, label: "大分県"},
        {value: 45, label: "宮崎県"},
        {value: 46, label: "鹿児島県"},
        {value: 47, label: "沖縄県"},
      ],

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

