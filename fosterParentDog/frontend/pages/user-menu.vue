<template>
  <div class="user-menu">
    ユーザメニュー
    <button type="button" @click="logout">ログアウト</button>
  </div>
</template>

<script>
export default {
  data() {
    return {
      userName: '',
      password: '',
    }
  },
  methods: {
    logout() {
      this.$cookies.get("user_name", this.userName);
      this.$cookies.get("password", this.password);
      this.$axios.post('/api/logout')
        .then((response) => {
          if (response.status === 200) {
            console.log(`ログアウト完了:${response.status}:${response.statusText}`)
            this.$router.push('/')
          } else if (response.status === 400) {
            console.error(`ログアウトエラー:${response.status}:${response.statusText}`)
            this.$router.push('/login')
          } else {
            console.error(`予期しないログアウトエラー:${response.status}:${response.statusText}`)
          }
        })
        .catch((err) => {
          console.error(err.response)
        })
        .finally(() => {
        })
    },
  },
}
</script>

<!-- cssはassetsから自動で読み込む-->
<style scoped>
</style>
