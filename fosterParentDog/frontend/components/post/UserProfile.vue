<template>
  <div class="user-profile">
    <div class="container" v-bind:key="user.id">
      <div class="row"><p>投稿者情報</p></div>
      <div class="row"><p>ニックネーム : {{ user.nickname }}</p></div>
      <div class="row"><p>プロフィール : {{ user.profile }}</p></div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    user: Object,
  },
}
</script>
