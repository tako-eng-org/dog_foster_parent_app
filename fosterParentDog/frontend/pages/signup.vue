<template>
  <div class="container">
    <h1>新規登録</h1>

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
      console.log("loooooooooooooooooogin")
    },
  }
}
</script>

<!-- cssはassetsから自動で読み込む-->
<style></style>
