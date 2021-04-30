<template>
  <!-- TODO 新規追加,更新の順に作成する -->
  <div class="post-edit">
    <div class="container">
      <!--   **********debug**********   -->
      <button type="submit" class="btn btn-primary" @click="testMethod">console</button>
      <!--   **********debug**********   -->

      <TextBox :title="'犬の名前'" v-model="post.postBase.dog_name" :readonly="false"/>
      <p class="text-primary">debug_犬の名前:{{ post.postBase.dog_name }}</p>

      <TextBox :title="'犬種'" v-model="post.postBase.breed" :readonly="false"/>
      <p class="text-primary">debug_犬種:{{ post.postBase.breed }}</p>

      <Gender v-model="post.postBase.gender" :readonly="false"/>
      <p class="text-primary">debug_性別:{{ post.postBase.gender }}</p>

      <Spay v-model="post.postBase.spay" :readonly="false"/>
      <p class="text-primary">debug_去勢手術:{{ post.postBase.spay }}</p>

      <TextBox :title="'年齢'" v-model="post.postBase.old" :readonly="false"/>
      <p class="text-primary">debug_年齢:{{ post.postBase.old }}</p>

      <SinglePerson v-model="post.postBase.single_person" :readonly="false"/>
      <p class="text-primary">debug_単身者への譲渡:{{ post.postBase.single_person }}</p>

      <SeniorPerson v-model="post.postBase.senior_person" :readonly="false"/>
      <p class="text-primary">debug_高齢者への譲渡:{{ post.postBase.senior_person }}</p>

      <TransferStatus v-model="post.postBase.transfer_status" :readonly="false"/>
      <p class="text-primary">debug_譲渡ステータス:{{ post.postBase.transfer_status }}</p>

      <TextBox :title="'自己紹介'" v-model="post.postBase.introduction" :readonly="false"/>
      <p class="text-primary">debug_自己紹介:{{ post.postBase.introduction }}</p>

      <TextBox :title="'アピールポイント'" v-model="post.postBase.appeal_point" :readonly="false"/>
      <p class="text-primary">debug_アピールポイント:{{ post.postBase.appeal_point }}</p>

      <!--   TODO: 譲渡可能都道府県のCreate実装   -->
      <PostPrefecture v-model="post.postPrefectureList.post_prefecture_id" :readonly="false"/>
      <p class="text-primary">debug_譲渡可能都道府県:{{ post.postPrefectureList.post_prefecture_id }}</p>

      <TextBox :title="'その他特記事項'" v-model="post.postBase.other_message" :readonly="false"/>
      <p class="text-primary">debug_その他特記事項:{{ post.postBase.other_message }}</p>

      <hr>
      <!--   TODO トップ画像をアップロードする   -->
      <p>トップ画像をアップロードしてください(サムネイルや目立つ箇所に表示されます)</p>
      <UploadOne v-model="topImageBase64"/>
      <p class="text-primary">debug_picture is: {{ topImageBase64 }}</p>
      <!--      <img :src="topImage"/>-->

      <!--   TODO サブ画像をアップロードする(任意,9枚まで)uploadOne使うか？   -->

      <!--  投稿ボタン  -->
      <button type="submit" class="btn btn-primary" @click="postDetail">投稿する</button>

      <!--  下書きに保存ボタン  -->
      <!--          <button type="submit" class="btn btn-primary" @click="postDetail">下書きに保存</button>-->
    </div>

  </div>
</template>

<script>
import TextBox from "~/components/post/TextBox";
import ImageOne from "~/components/post/ImageOne"
import ImageList from "~/components/post/ImageList";
import PostPrefecture from "~/components/post/PostPrefecture"
import UserProfile from "~/components/post/UserProfile";
import Gender from "~/components/post/Gender";
import SeniorPerson from "~/components/post/SeniorPerson";
import SinglePerson from "~/components/post/SinglePerson";
import Spay from "~/components/post/Spay";
import TransferStatus from "~/components/post/TransferStatus";
import UploadOne from "~/components/post/UploadOne";

import CustomInput from "~/components/post/CustomInput";

export default {
  components: {
    TextBox,
    ImageOne,
    ImageList,
    PostPrefecture,
    UserProfile,
    Gender,
    SeniorPerson,
    SinglePerson,
    Spay,
    TransferStatus,
    UploadOne,

    CustomInput,
  },
  mounted() {
  },
  data() {
    return {
      topImageBase64: null,

      post: {
        postBase: { //基礎投稿
          // id: 投稿時に自動付与するため、新規作成では無し
          // created_at:
          // updated_at:
          publishing: 1,
          dog_name: "test_dog_name",
          breed: "test_breed",
          gender: 0,
          spay: 0,
          old: "test_old",
          single_person: 0,
          senior_person: 0,
          transfer_status: 0,
          introduction: "test_introduction",
          appeal_point: "test_appeal_point",
          other_message: "test_other_message",
          top_image_path: "test_top_image_path",
        },

        imagePathList: { //画像パスリスト
          image_path: [],
        },

        postPrefectureList: { //譲渡可能都道府県リスト
          post_prefecture_id: [],
        },

        // user: { //ユーザー情報
        //   user_id: 1,//debug
        // },
      },
    }
  },
  computed: {},
  methods: {
    /**
     * debug用
     */
    testMethod() {
      console.log("-------testMethod");
      // console.log(this.postData);
    },

    /**
     * 投稿記事を追加する
     */
    postDetail() {
      let postJson = {
        postBase: { //投稿基礎情報
          publishing: this.post.postBase.publishing,
          dog_name: this.post.postBase.dog_name,
          breed: this.post.postBase.breed,
          gender: this.post.postBase.gender,
          spay: this.post.postBase.spay,
          old: this.post.postBase.old,
          single_person: this.post.postBase.single_person,
          senior_person: this.post.postBase.senior_person,
          transfer_status: this.post.postBase.introduction,
          introduction: this.post.postBase.introduction,
          appeal_point: this.post.postBase.appeal_point,
          other_message: this.post.postBase.other_message,
          user_id: null, //ログインユーザーのIDが入る
          top_image_path: this.post.postBase.top_image_path,
        },

        imagePathList: { //画像パスリスト
          image_path: [
            "12345",
            "6789",
          ],
        },

        postPrefectureList: { //譲渡可能都道府県
          post_prefecture_id: [
            1,
            2,
            3,
          ],
        },

        // user: { //投稿者情報
        //   user_id: 2,
        // },
      }

      console.log("--------postJSON");
      console.log(postJson);
      console.log("--------postDetail");
      // return new Promise((resolve, reject) => {
      // this.$axios.defaults.timeout = 100000;
      this.$axios.post('/api/post_test', postJson)
        .then((response) => {
          if ((response.status !== 201)) {
            console.error("error---------------------");
            console.error(`Error:${response.statusText}, ${this.postDetail.name}`)
          } else {
            console.log("aaaaaaaaaaaaaaaaaaaaaaaaa");
          }
        }).catch(err => console.error(err));
      // })
    },
  },
}
</script>

<!-- cssはassetsから自動で読み込む-->
<style scoped></style>
