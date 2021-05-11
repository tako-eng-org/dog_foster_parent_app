<template>
  <div class="detail">
    <div class="container">
      <!--   FIXME: 画像表示について動作しません。後日改修予定   -->
      <!--   トップ画像表示   -->
      <!--            <ImageOne :objectUrl="getTopObjectUrl(post.post_images)"/>-->
      <div class="post_id row">
        <p>投稿No : {{ post.id }}</p>
      </div>
      <TextBox :title="'犬の名前'" :value="post.dog_name"/>
      <TextBox :title="'犬種'" :value="post.breed"/>
      <Gender :value="post.gender"/>
      <Spay :value="post.spay"/>
      <TextBox :title="'年齢'" :value="post.old"/>
      <SinglePerson :value="post.single_person"/>
      <SeniorPerson :value="post.senior_person"/>
      <TransferStatus :value="post.transfer_status"/>
      <TextBox :title="'自己紹介'" :value="post.introduction"/>
      <TextBox :title="'投稿日時'" :value="post.created_at"/>
      <TextBox :title="'アピールポイント'" :value="post.appeal_point"/>
      <PostPrefecture :value="post.post_prefectures"/>
      <TextBox :title="'その他特記事項'" :value="post.other_message"/>
    </div>

    <!--  トップ画像以外の画像を複数表示  -->
    <!--    <ImageList :objectUrlList="post.postimages"/>-->

    <hr>
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

  data() {
    return {
      post: {},
    }
  },
  computed: {},
  methods: {
    /**
     * トップ画像（1投稿のpositionが最小値(原則0)のobjectUrl）を取得する
     * @param {object} post
     */
    // FIXME: この関数は動作しません。map取れない Cannot read property 'map' of undefined
    // getTopObjectUrl(postImageObj) {
    // let min = Math.min(postImageObj.map(x => x.position));
    // let ImagePathByMinPosition = (postObj.post_images || []).filter(e => (e.position === min))[0].objectKey
    // return ImagePathByMinPosition
    // },

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
            this.post = response.data;
          }
        }).catch(err => console.error(err.response));
    },
  },
}
</script>

<!-- cssはassetsから自動で読み込む-->
<style scoped></style>
