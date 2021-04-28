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

    </div>

    <!--  投稿ボタン  -->
    <button type="submit"
            class="btn btn-primary"
            @click="postDetail">
      投稿する
    </button>

    <!--  下書きに保存ボタン  -->
    <button type="submit"
            class="btn btn-primary"
            @click="postDetail">
      下書きに保存
    </button>

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

  mounted: function () {
    // this.getDetail(this.$route.query.postId);
  },

  data: function () {
    return {
      topImageBase64: null,

      // TODO 変数名をキャメルケースへ変更すること
      postData: {
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
        user_id: 1,
        top_image_path: "test_top_image_path",
        post_image_id: 1,
      },

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
          // post_id: null,
          // post_image_id: null,
          image_path: [],
        },

        postPrefectureList: { //譲渡可能都道府県リスト
          post_prefecture_id: [],
        },

        user: { //ユーザー情報
          user_id: 1,//debug
        },
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
      console.log(this.postData);
    },

    /**
     * 投稿記事を追加する
     */
    postDetail() {
      let params = new URLSearchParams();
      // id
      // created_at
      // updated_at
      params.append('publishing', this.post.postBase.publishing);
      params.append('dog_name', this.post.postBase.dog_name);
      params.append('breed', this.post.postBase.breed);
      params.append('gender', this.post.postBase.gender);
      params.append('spay', this.post.postBase.spay);
      params.append('old', this.post.postBase.old);
      params.append('single_person', this.post.postBase.single_person);
      params.append('senior_person', this.post.postBase.senior_person);
      params.append('transfer_status', this.post.postBase.transfer_status);
      params.append('introduction', this.post.postBase.introduction);
      params.append('appeal_point', this.post.postBase.appeal_point);
      params.append('other_message', this.post.postBase.other_message);
      params.append('top_image_path', this.post.postBase.top_image_path);

      // imagePathList

      // postPrefectureList
      params.append('post_prefecture_id', this.post.postPrefectureList.post_prefecture_id);

      // user
      // params.append('user_id', this.user.user_id);

      this.$axios.post('/api/posttest', params)
        .then((response) => {
          if ((response.status !== 200)) {
            console.error(`Error:${response.statusText}, ${this.postDetail.name}`)
          } else {
            console.log("create----------debug")
          }
        }).catch(err => console.error(err));
    },
  },
}
</script>

<!-- cssはassetsから自動で読み込む-->
<style scoped></style>
