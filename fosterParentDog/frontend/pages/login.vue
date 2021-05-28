<template>
  <div class="login">
    <p>ユーザー名 :<input type="text" name="user-name" placeholder="" v-model="userName"/></p>
    <p>パスワード :<input type="text" name="password" placeholder="" v-model="password"/></p>
    <p>
      <button type="button" @click="login">ログイン</button>
    </p>
    <p>
      <nuxt-link to="./re-password">パスワードを忘れた方はこちら</nuxt-link>
    </p>
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
    login() {
      const formData = new FormData()
      formData.append('user-name', this.userName)
      formData.append('password', this.password)
      console.log(...formData.entries()) //debug用/
      this.$axios.post('/api/login', formData)
        .then((response) => {
          if (response.status === 200) {
            console.log(`ログイン完了:${response.status}:${response.statusText}`)
            this.$router.push('/userMenu')
          } else if (response.status === 400) {
            console.error(`ログインエラー:${response.status}:${response.statusText}`)
            this.$router.push('/login')
          } else {
            console.error(`予期しないログインエラー:${response.status}:${response.statusText}`)
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
<style scoped></style>
