<template>
  <div class="login">
    <form>
      <p>ユーザー名 :<input type="text" name="user-name" placeholder="" v-model="userName"/></p>
      <p>パスワード :<input type="text" name="password" placeholder="" v-model="password"/></p>
      <p><input type="submit" @click="login" value="ログイン"/></p>
    </form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      userName: "",
      password: "",
    }
  },
  methods: {
    login() {
      let formData = new FormData();
      formData.append('user_name', this.userName);
      formData.append('password', this.password);
      console.log(...formData.entries()); //debug用
      let startMsec = new Date();
      while (new Date() - startMsec < 10000) ;// 指定ミリ秒間だけループさせる（CPUは常にビジー状態）
      this.$axios.post('/api/login', formData)
        .then((response) => {
          if ((response.status !== 201)) {
            console.error("error-201以外");
            console.error(`ログインエラー:${response.status}:${response.statusText}, ${this.postCreate.name}`)
          } else {
            console.log(`ログイン完了:${response.status}:${response.statusText}, ${this.postCreate.name}`);
            // TODO: ログイン後の処理実装
          }
        })
        .catch(err => console.error(err.response));
    },
  },
}
</script>

<!-- cssはassetsから自動で読み込む-->
<style scoped></style>
