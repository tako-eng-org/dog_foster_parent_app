<template>
  <div class="container">
    <h1>login</h1>

    <nuxt-link to="./signup">新規登録はこちら</nuxt-link>

    <button type="submit"
            class="btn btn-primary"
            v-on:click="doLoginConsole">
      テストボタン
    </button>
  </div>
</template>

<script>
process.env.DEBUG = 'nuxt:*' // nuxt.jsについてログ出力する

export default {
  methods:{
    doLoginConsole() {
      console.log("******************login******************")
    },
  }
}
</script>

<!-- cssはassetsから自動で読み込む-->
<style></style>
