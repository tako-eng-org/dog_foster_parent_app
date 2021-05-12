<template>
  <div class="user-profile">
    <div class="container">
      <div class="row"><p>投稿者情報</p></div>
      <div class="row"><p>ニックネーム : {{ value.nickname }}</p></div>
      <div class="row"><p>プロフィール : {{ value.profile }}</p></div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    value: { //子コンポーネントから親コンポーネントへバインディングする設定
      default: () => ({}),
      required: true
    },
  },
}
</script>
