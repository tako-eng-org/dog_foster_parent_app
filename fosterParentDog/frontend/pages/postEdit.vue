<template>
  <!-- TODO 新規追加,更新の順に作成する -->
  <div class="post-edit">
    <div class="container">
      <!--   **********debug**********   -->
      <button type="submit" class="btn btn-primary" @click="testMethod">console</button>
      <!--   **********debug**********   -->

      <TextBox :title="'犬の名前'" v-model="post.dogName" :readonly="false"/>
      <p class="text-primary">debug_犬の名前:{{ post.dogName }}</p>

      <TextBox :title="'犬種'" v-model="post.breed" :readonly="false"/>
      <p class="text-primary">debug_犬種:{{ post.breed }}</p>

      <Gender v-model="post.gender" :readonly="false"/>
      <p class="text-primary">debug_性別:{{ post.gender }}</p>

      <Spay v-model="post.spay" :readonly="false"/>
      <p class="text-primary">debug_去勢手術:{{ post.spay }}</p>

      <TextBox :title="'年齢'" v-model="post.old" :readonly="false"/>
      <p class="text-primary">debug_年齢:{{ post.old }}</p>

      <SinglePerson v-model="post.singlePerson" :readonly="false"/>
      <p class="text-primary">debug_単身者への譲渡:{{ post.singlePerson }}</p>

      <SeniorPerson v-model="post.seniorPerson" :readonly="false"/>
      <p class="text-primary">debug_高齢者への譲渡:{{ post.seniorPerson }}</p>

      <TransferStatus v-model="post.transferStatus" :readonly="false"/>
      <p class="text-primary">debug_譲渡ステータス:{{ post.transferStatus }}</p>

      <TextBox :title="'自己紹介'" v-model="post.introduction" :readonly="false"/>
      <p class="text-primary">debug_自己紹介:{{ post.introduction }}</p>

      <TextBox :title="'アピールポイント'" v-model="post.appealPoint" :readonly="false"/>
      <p class="text-primary">debug_アピールポイント:{{ post.appealPoint }}</p>

      <!--   TODO: 譲渡可能都道府県のCreate実装   -->
      <!--      <PostPrefecture v-model="post.postPrefectureList.post_prefecture_id" :readonly="false"/>-->
      <!--      <p class="text-primary">debug_譲渡可能都道府県:{{ post.postPrefectureList.post_prefecture_id }}</p>-->
      <PostPrefecture v-model="post.postPrefectureIdList" :readonly="false"/>
      <p class="text-primary">debug_譲渡可能都道府県:{{ post.postPrefectureIdList }}</p>

      <TextBox :title="'その他特記事項'" v-model="post.otherMessage" :readonly="false"/>
      <p class="text-primary">debug_その他特記事項:{{ post.otherMessage }}</p>

      <hr>
      <!--   TODO トップ画像をアップロードする   -->
      <p>トップ画像をアップロードしてください(サムネイルや目立つ箇所に表示されます)</p>
      <UploadOne v-model="post.topImagePath"/>
      <p class="text-primary">debug_picture is: {{ post.topImagePath }}</p>

      <!--   TODO サブ画像をアップロードする(任意,9枚まで)uploadOne使うか？   -->

      <!--  投稿ボタン  -->
      <button type="submit" class="btn btn-primary" @click="postCreate">投稿する</button>

      <!--  debug_画像ボタン  -->
      <button type="submit" class="btn btn-primary" @click="imageCreate">画像投稿デバッグ用</button>

      <!--  下書きに保存ボタン  -->
      <!--          <button type="submit" class="btn btn-primary" @click="postCreate">下書きに保存</button>-->
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

  },
  mounted() {
  },
  data() {
    return {
      post: {
        //基礎投稿
        // id: 投稿時に自動付与するため、新規作成では無し
        // created_at:
        // updated_at:
        publishing: 1,
        dogName: "test_dog_name",
        breed: "test_breed",
        gender: 0,
        spay: 0,
        old: "test_old",
        singlePerson: 0,
        seniorPerson: 0,
        transferStatus: 0,
        introduction: "test_introduction",
        appealPoint: "test_appeal_point",
        otherMessage: "test_other_message",
        topImagePath: "test_top_image_path",

        imagePathList: [], //画像パスリスト

        postPrefectureIdList: [], //譲渡可能都道府県リスト

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
    postCreate() {
      let postJson = {
        publishing: this.post.publishing,
        dog_name: this.post.dogName,
        breed: this.post.breed,
        gender: this.post.gender,
        spay: this.post.spay,
        old: this.post.old,
        single_person: this.post.singlePerson,
        senior_person: this.post.seniorPerson,
        transfer_status: this.post.transferStatus,
        introduction: this.post.introduction,
        appeal_point: this.post.appealPoint,
        other_message: this.post.otherMessage,
        user_id: 1, // TODO: ログインユーザーのIDが入る(入れないとindex表示時にnullエラーになりdog_app再起動が必要になるためdebug用として1を入れている)
        top_image_path: this.post.topImagePath,
        image_path_list: [
          "12345",
          "6789",
        ],
        post_prefecture_id_list: this.post.postPrefectureIdList
      }
      // console.log("--------postJSON");
      // console.log(postJson);
      // console.log("--------postCreate");
      this.$axios.post('/api/post_create', postJson)
        .then((response) => {
          if ((response.status !== 201)) {
            console.error("error-201以外");
            console.error(`投稿エラー:${response.status}:${response.statusText}, ${this.postCreate.name}`)
          } else {
            console.log(`投稿完了:${response.status}:${response.statusText}, ${this.postCreate.name}`);
          }
        }).catch(err => console.error(err.response));
    },

    /**
     * top画像を投稿する（debug用）
     */
    imageCreate() {
      let postJson = {
        top_image_path: this.post.topImagePath,
      }
      console.log("--------postJSON");
      console.log(postJson);
      this.$axios.post('/api/image_create', postJson)
        .then((response) => {
          if ((response.status !== 201)) {
            console.error("error-201以外");
            console.error(`画像エラー:${response.status}:${response.statusText}, ${this.imageCreate.name}`)
          } else {
            console.log(`画像完了:${response.status}:${response.statusText}, ${this.imageCreate.name}`);
          }
        }).catch(err => console.error(err.response));
    },

  },
}
</script>

<!-- cssはassetsから自動で読み込む-->
<style scoped></style>
