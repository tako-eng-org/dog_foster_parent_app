<template>
  <div class="detail">
    <div class="container" v-bind:key="post.postBase.id">
      <ImageOne :imagePath="post.postBase.top_image_path"/>
      <div class="post_id">
        <div class="row">
          <p>投稿No : {{ post.postBase.id }}</p>
        </div>
      </div>
      <TextBox :title="'犬の名前'"
               :detail="post.postBase.breed"
               :readonly="true"/>
      <TextBox :title="'犬種'"
               :detail="post.postBase.breed"
               :readonly="true"/>
      <Gender :itemValue="post.postBase.gender"
              :readonly="true"/>
      <Spay :itemValue="post.postBase.spay"
            :readonly="true"/>
      <TextBox :title="'年齢'"
               :detail="post.postBase.old"
               :readonly="true"/>
      <SinglePerson :itemValue="post.postBase.single_person"
                    :readonly="true"/>
      <SeniorPerson :itemValue="post.postBase.senior_person"
                    :readonly="true"/>
      <TransferStatus :itemValue="post.postBase.transfer_status"
                      :readonly="true"/>
      <TextBox :title="'自己紹介'"
               :detail="post.postBase.introduction"
               :readonly="true"/>
      <TextBox :title="'投稿日時'"
               :detail="post.postBase.created_at"
               :readonly="true"/>
      <TextBox :title="'アピールポイント'"
               :detail="post.postBase.appeal_point"
               :readonly="true"/>
      <PostPrefecture :postPrefectureList="post.postPrefectureList"
                      :readonly="true"/>
      <TextBox :title="'その他特記事項'"
               :detail="post.postBase.other_message"
               :readonly="true"/>
    </div>

    <!--  画像リスト表示  -->
    <ImageList :imagePathList="post.imagePathList"/>

    <!--  ユーザープロフィール表示  -->
    <UserProfile :user="post.user"/>
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
  },

  mounted: function () {
    this.getDetail(this.$route.query.postId);
  },

  data: function () {
    return {
      post: {
        postBase: {}, //基礎投稿
        imagePathList: {}, //画像パスリスト
        postPrefectureList: {}, //譲渡可能都道府県リスト
        user: {}, //ユーザー情報
      },
    }
  },
  computed: {},
  methods: {
    /**
     * 投稿IDに紐づいた投稿記事を取得する
     * @param {int} postID
     */
    getDetail(currentPostId) {
      this.$axios.get('/api/post', {
          params: {
            postId: currentPostId,
          }
        }
      )
        .then((response) => {
          if ((response.status !== 200)) {
            console.error(`Error:${response.statusText}, ${this.getDetail.name}`)
          } else {
            this.post = {
              postBase: response.data.post,
              imagePathList: response.data.post_images,
              postPrefectureList: response.data.post_prefectures,
              user: response.data.user,
            }
          }
        }).catch(err => console.error(err));
    },
  },
}
</script>

<!-- cssはassetsから自動で読み込む-->
<style scoped></style>
