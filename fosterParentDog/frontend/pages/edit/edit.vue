<template>
  <div className="container">
    <h1>記事編集画面</h1>

    <button type="submit"
            className="btn btn-primary"
            v-on:click="doEditConsole">
      テストボタン
    </button>

  </div>
</template>

<script>
// process.env.DEBUG = 'nuxt:*' // nuxt.jsについてログ出力する

export default {
  methods: {
    doEditConsole() {
      console.log("test")
    },
  }
}
</script>

<!-- cssはassetsから自動で読み込む-->
<style></style>
