<template>
  <div class="user-profile">
    <div class="container">
      <p>投稿者情報</p>
      <p>ニックネーム : {{ value.nickname }}</p>
      <p>プロフィール : {{ value.profile }}</p>
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
