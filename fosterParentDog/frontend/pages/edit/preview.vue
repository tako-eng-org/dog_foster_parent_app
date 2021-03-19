<template>
  <div class="container">
    <h1>preview</h1>

    <button type="submit"
            class="btn btn-primary"
            v-on:click="doEditConsole">
      テストボタン
    </button>
  </div>
</template>

<script>
// process.env.DEBUG = 'nuxt:*' // nuxt.jsについてログ出力する

export default {
  methods:{
    doEditConsole() {
      console.log("****************preview********************")
    },
  }
}
</script>

<!-- cssはassetsから自動で読み込む-->
<style></style>
