<template>
  <div class="detail">
    <div class="container" v-bind:key="post.postBase.id">
      <ImageTop :imagePath="post.postBase.top_image_path"/>
      <div class="post_id">
        <div class="row">
          <p>投稿No : {{ post.postBase.id }}</p>
        </div>
      </div>
      <ViewOrTextBox :title="'犬の名前'"
                     :detail="post.postBase.breed"
                     :isView="true"/>
      <ViewOrTextBox :title="'犬種'"
                     :detail="post.postBase.breed"
                     :isView="true"/>
      <LabelOrDropdown :title="'性別'"
                       :detail="gender_map"
                       :itemValue="post.postBase.gender"
                       :isView="true"/>
      <LabelOrDropdown :title="'去勢/避妊手術'"
                       :mapName="spay_map"
                       :itemValue="post.postBase.spay"
                       :isView="true"/>
      <ViewOrTextBox :title="'年齢'"
                     :detail="post.postBase.old"
                     :isView="true"/>
      <LabelOrDropdown :title="'単身者への譲渡'"
                       :mapName="single_person_map"
                       :itemValue="post.postBase.single_person"
                       :isView="true"/>
      <LabelOrDropdown :title="'高齢者への譲渡'"
                       :mapName="senior_person_map"
                       :itemValue="post.postBase.senior_person"
                       :isView="true"/>
      <LabelOrDropdown :title="'譲渡ステータス'"
                       :mapName="transfer_status_map"
                       :itemValue="post.postBase.transfer_status"
                       :isView="true"/>
      <ViewOrTextBox :title="'自己紹介'"
                     :detail="post.postBase.introduction"
                     :isView="true"/>
      <ViewOnly :title="'投稿日時'"
                :detail="post.postBase.created_at"
                :isView="true"/>
      <ViewOrTextBox :title="'アピールポイント'"
                     :detail="post.postBase.appeal_point"
                     :isView="true"/>

      <PostPrefecture :title="'譲渡可能都道府県'"
                      :postPrefectureList="post.postPrefectureList"
                      :itemMap="prefecture_map"/>

      <ImageList :imagePathList="post.postPrefectureList"
                 :isView="true"/>
      <ViewOrTextBox :title="'その他特記事項'"
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
import ViewOrTextBox from "~/components/post/TextOrTextBox";
import ImageTop from "~/components/post/Image"
import ImageList from "~/components/post/ImageList";
import UserProfile from "~/components/post/UserProfile";
import LabelOrDropdown from "~/components/post/LabelOrDropdown";
import ViewOnly from "~/components/post/ReadOnly";
import PostPrefecture from "~/components/post/PostPrefecture"
import CON from "~/components/const/const"

export default {
  components: {
    ViewOrTextBox,
    ImageList,
    UserProfile,
    ImageTop,
    LabelOrDropdown,
    ViewOnly,
    PostPrefecture,
    CON,
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

      // components/const/const.vueから定数読み出し
      gender_map: CON.data().GENDER,
      prefecture_map: CON.data().PREFECTURE,
      senior_person_map: CON.data().SENIOR_PERSON,
      single_person_map: CON.data().SINGLE_PERSON,
      spay_map: CON.data().SPAY,
      transfer_status_map: CON.data().TRANSFER_STATUS,

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
