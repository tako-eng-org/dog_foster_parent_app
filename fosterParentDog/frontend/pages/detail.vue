<template>
  <div class="detail">
    <div class="container" v-bind:key="post.postBase.id">
      <ImageTop :imagePath="post.postBase.top_image_path"/>
      <div class="post_id">
        <div class="row">
          <p>投稿No : {{ post.postBase.id }}</p>
        </div>
      </div>
      <TextOrTextBox :title="CON.DOG_NAME_TITLE"
                     :detail="post.postBase.breed"
                     :isView="true"/>
      <TextOrTextBox :title="CON.BREED_TITLE"
                     :detail="post.postBase.breed"
                     :isView="true"/>
      <LabelOrDropdown :title="CON.GENDER_TITLE"
                       :detail="CON.GENDER"
                       :itemValue="post.postBase.gender"
                       :isView="true"/>
      <LabelOrDropdown :title="CON.SPAY_TITLE"
                       :mapName="CON.SPAY"
                       :itemValue="post.postBase.spay"
                       :isView="true"/>
      <TextOrTextBox :title="CON.OLD_TITLE"
                     :detail="post.postBase.old"
                     :isView="true"/>
      <LabelOrDropdown :title="CON.SINGLE_PERSON_TITLE"
                       :mapName="CON.SINGLE_PERSON"
                       :itemValue="post.postBase.single_person"
                       :isView="true"/>
      <LabelOrDropdown :title="CON.SENIOR_PERSON_TITLE"
                       :mapName="CON.SENIOR_PERSON"
                       :itemValue="post.postBase.senior_person"
                       :isView="true"/>
      <LabelOrDropdown :title="CON.TRANSFER_STATUS_TITLE"
                       :mapName="CON.TRANSFER_STATUS"
                       :itemValue="post.postBase.transfer_status"
                       :isView="true"/>
      <TextOrTextBox :title="CON.INTRODUCTION_TITLE"
                     :detail="post.postBase.introduction"
                     :isView="true"/>
      <TextOrTextBox :title="CON.CREATED_AT_TITLE"
                     :detail="post.postBase.created_at"
                     :isView="true"/>
      <TextOrTextBox :title="'アピールポイント'"
                     :detail="post.postBase.appeal_point"
                     :isView="true"/>

      <PostPrefecture :title="'譲渡可能都道府県'"
                      :postPrefectureList="post.postPrefectureList"
                      :itemMap="prefecture_map"/>

      <ImageList :imagePathList="post.postPrefectureList"
                 :isView="true"/>
      <TextOrTextBox :title="'その他特記事項'"
                     :detail="post.postBase.other_message"
                     :isView="true"/>
    </div>

    <!--  画像リスト表示  -->
    <ImageList :imagePathList="post.imagePathList"/>

    <!--  ユーザープロフィール表示  -->
    <UserProfile :user="post.user"/>
  </div>
</template>

<script>
import CON from "~/components/const/const"
import TextOrTextBox from "~/components/post/TextOrTextBox";
import LabelOrDropdown from "~/components/post/LabelOrDropdown";
import ImageTop from "~/components/post/Image"
import ImageList from "~/components/post/ImageList";
import PostPrefecture from "~/components/post/PostPrefecture"
import UserProfile from "~/components/post/UserProfile";

export default {
  components: {
    CON,

    TextOrTextBox,
    LabelOrDropdown,

    ImageTop,
    ImageList,
    PostPrefecture,
    UserProfile,
  },

  mounted: function () {
    this.getDetail(this.$route.query.postId);
  },

  data: function () {
    return {
      CON: CON.data(), //const.vue読み込み。 //TODO コンストファイルをいい感じにする

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
      this.$axios.get('api/post', {
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
        }).catch(err => alert(err));
    },

  },
}
</script>

<!-- cssはassetsから自動で読み込む-->
<style scoped></style>
