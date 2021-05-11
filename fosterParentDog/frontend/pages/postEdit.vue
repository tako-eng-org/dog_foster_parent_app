<template>
  <!-- TODO 新規追加,更新の順に作成する -->
  <div class="post-edit">
    <div class="container">
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

      <PostPrefecture v-model="post.postPrefectureIdList" :readonly="false"/>
      <p class="text-primary">debug_譲渡可能都道府県:{{ post.postPrefectureIdList }}</p>

      <TextBox :title="'その他特記事項'" v-model="post.otherMessage" :readonly="false"/>
      <p class="text-primary">debug_その他特記事項:{{ post.otherMessage }}</p>

      <hr>
      <!-- 画像アップロード -->
      <!--   TODO: 画像の投稿機能は後日再度実装予定。ファイルをアップロードし、ファイル名で表示する。post_idを直指定。プレビュー機能はなし-->

      <!--      <div class="upload-top-image">-->
      <!--        <input type="file" id="top-image" ref="input">-->
      <!--        <button type="submit" @click="ImageUpload(0)">トップ画像をアップロードする</button>-->
      <!--        <ImageOne :objectUrl="getTopObjectUrl()"/>-->
      <!--      </div>-->

      <!--  投稿ボタン  -->
      <button type="submit" class="btn btn-primary" @click="postCreate(1)">投稿する</button>
      <!--  下書きに保存ボタン  -->
      <button type="submit" class="btn btn-primary" @click="postCreate(0)">下書きに保存する</button>
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
      picture: null,

      post: {
        //基礎投稿
        // id: 投稿時に自動付与するため、新規作成では無し。update実装時に更新予定
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
        user_id: 1, // TODO: ログインユーザーのIDが入る(入れないとindex表示時にnullエラーになりdog_app再起動が必要になるためdebug用として1を入れている)
        postImages: [ //画像パスリスト
          {
            position: 0,
            objectKey: "images/1_767121c4-4b8b-4874-9690-008e0fcd3fae.png",
            objectUrl: "https://bbsapp-img.s3.us-east-2.amazonaws.com/images/000002.png",
          },
        ],
        postPrefectureIdList: [], //譲渡可能都道府県リスト
      },

      topObjectUrl: "", //https://bbsapp-img.s3.us-east-2.amazonaws.com/images/1_8a217b9f-b0de-441f-b68d-861a7cd5f4be.png
      topObjectKey: "", //images/1_8a217b9f-b0de-441f-b68d-861a7cd5f4be.png
    }
  },
  computed: {},
  methods: {
    /**
     * 1投稿のpositionが最小値(原則0)のobjectKeyを取得する
     * @param {object} post
     */
    getTopObjectUrl() {
      let min = Math.min(...this.post.postImages.map(x => x.position));
      let objectKeyByMinPosition = this.post.postImages.filter(e => (e.position === min))[0].object_key;
      console.log("objectKeyByMinPosition-------");
      console.log(objectKeyByMinPosition);
      return objectKeyByMinPosition
    },

    /**
     * 投稿記事を追加する
     */
    postCreate(publishing) {
      let postJson = {
        publishing: publishing,
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
        user_id: this.user_id,
        top_object_key: this.topObjectKey,
        object_key_list: [
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
        })
        .catch(err => console.error(err.response));
    },

    /**
     * debug用
     */
    ImageUpload(position) {
      let file = this.$refs.input;
      let formData = new FormData();
      formData.append('image', file.files[0]);
      formData.append('user_id', this.post.user_id);
      formData.append('position', position);
      console.log(...formData.entries());
      return new Promise((resolve, reject) => {
        this.$axios.post('/api/image_upload', formData, {
            headers: {
              'Content-Type': 'multipart/form-data',
            },
          }
        )
          .then((response) => {
            if ((response.status !== 201)) {
              console.error(`画像アップロードエラー:${response.status} ${response.statusText}`)
            } else {
              console.log(`画像アップロード完了:${response.status} ${response.statusText}`);
              // this.topObjectUrl = response.data.object_url;
              // this.topObjectKey = response.data.object_key;
              // console.log(this.topObjectUrl);
              // console.log(this.topObjectKey);
              let imageArr = {
                position: response.data.position,
                objectKey: response.data.object_key,
                objectUrl: response.data.object_url,
              };
              console.log("imageArr is --------");
              console.log(imageArr);
              this.post.postImages.push(imageArr);
              console.log("this.post.postImages are --------");
              console.log(this.post.postImages);
            }
          })
          .catch(function (error) {
            if (error.response) {
              // 要求がなされたとサーバがステータスコードで応答した// 2XXの範囲外
              console.log(error.response.data);
              console.log(error.response.status);
              console.log(error.response.headers);
            } else if (error.request) {
              // 要求がなされたが、応答が受信されなかった
              console.log(error.request);
            } else {
              // トリガーしたリクエストの設定に何かしらのエラーがある
              console.log('Error', error.message);
            }
            console.log(error.config);
          })
      });
    },

  },
}
</script>

<!-- cssはassetsから自動で読み込む-->
<style scoped>
</style>
