<template>
  <div class="detail">
    <div class="container" v-bind:key="post.postBase.id">
      <ImageComponent :imagepath="post.postBase.top_image_path"/>
      <div class="post_id">
        <div class="row">
          <p>投稿No : {{ post.postBase.id }}</p>
        </div>
      </div>
      <DogNameComponent :dogName="post.postBase.dog_name"/>
      <BreedComponent :breed="post.postBase.breed"/>
      <GenderComponent :gender="post.postBase.gender"/>
      <SpayComponent :spay="post.postBase.spay"/>
      <OldComponent :old="post.postBase.old"/>
      <SinglePersonComponent :singlePerson="post.postBase.single_person"/>
      <SeniorPersonComponent :seniorPerson="post.postBase.senior_person"/>
      <TransferStatusComponent :transferStatus="post.postBase.transfer_status"/>
      <IntroductionComponent :introduction="post.postBase.introduction"/>
      <CreatedAtComponent :createdAt="post.postBase.created_at"/>
      <AppealPointComponent :appealPoint="post.postBase.appeal_point"/>
      <PostPrefectureComponent :postPrefectureList="post.postPrefectureList"/>
      <OtherMessageComponent :otherMessage="post.postBase.other_message"/>
    </div>

    <!--  画像リスト表示  -->
    <ImageListComponent :imagePathList="post.imagePathList"/>

    <!--  ユーザープロフィール表示  -->
    <UserProfileComponent :user="post.user"/>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  components: {},

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
      }
    }
  },
  computed: {},
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
          this.post = {
            postBase: response.data.post,
            imagePathList: response.data.post_images,
            postPrefectureList: response.data.post_prefectures,
            user: response.data.user,
          }
        }
      }).catch(err => alert(err));
    },
  },
}
</script>

<!-- cssはassetsから自動で読み込む-->
<style scoped></style>
