<template>
  <div class="login">
    <!--    <button @click="setCookie">Cookieセット</button>-->
    <!--    <button @click="getCookie">Cookie表示</button>-->
    {{ getCookie("mysession") }}
    <form>
      <p>ユーザー名 :<input type="text" name="user_name" placeholder="" v-model="userName"/></p>
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

      cookieValue: "",
    }
  },
  methods: {
    login() {
      // TODO: デバッグ用。後で削除する
      // let formData = new FormData();
      // formData.append('user_name', this.userName);
      // formData.append('password', this.password);
      // console.log(...formData.entries()); //debug用
      // this.$axios.post('/api/login', formData)
      this.$cookies.set("user_name", this.userName)
      this.$cookies.set("password", this.password)
      // FIXME: どのルートにも入ってない
      this.$axios.post('/api/login')
        .then((response) => {
          if ((response.status !== 200)) {
            console.error("error-200以外");
            console.error(`ログインエラー:${response.status}:${response.statusText}, ${this.postCreate.name}`)
            this.sleep(10000);
          } else {
            console.log(`ログイン完了:${response.status}:${response.statusText}, ${this.postCreate.name}`);
            this.sleep(10000);
            this.$router.push('/index')
          }
        })
        .catch((err) => {
          console.error(err.response)
          this.sleep(10000);
        });
    },
    getCookie(cookieName) {
      this.cookieValue = this.$cookies.get(cookieName); // Cookieから値を取得
    },
    // setCookie(){
    //   this.$cookies.set("test", "cookie-test"); // Cookieに値をセット
    // },
    sleep(msec) {
      let startMsec = new Date();
      while (new Date() - startMsec < msec) ;// 10秒=10000 指定ミリ秒間だけループさせる（CPUは常にビジー状態）
    }
  },
}
</script>

<!-- cssはassetsから自動で読み込む-->
<style scoped></style>
