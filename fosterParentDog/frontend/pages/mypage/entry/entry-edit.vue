<template>
  <div class="container">
    <h1>entry-edit</h1>

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
      console.log("************************************")
    },
  }
}
</script>

<!-- cssはassetsから自動で読み込む-->
<style scoped></style>
