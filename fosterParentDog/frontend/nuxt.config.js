require('dotenv').config()
let axiosTarget = '';
let pathRewrite = {};
if (process.env.NODE_ENV === 'production') {
  // 本番用
  axiosTarget = 'https://tako-eng.com/'//本番用に後で書き換える
} else if (process.env.NODE_ENV === 'staging') {
  //Docker用
  axiosTarget = 'http://localhost:8000/fosterparent'
} else {
  //devサーバー用
  // axiosTarget = 'http://localhost:8000/fosterparent/api'
  axiosTarget = 'http://[::1]:8000/fosterparent/api'
  pathRewrite = {"^/api": ""}
}

export default {
  target: 'server',
  ssr: false, // サーバーサイドレンダリングを無効化

  // Global page headers: https://go.nuxtjs.dev/config-head
  head: {
    title: 'frontend',
    htmlAttrs: {
      lang: 'ja'
    },
    meta: [
      {charset: 'utf-8'},
      {name: 'viewport', content: 'width=device-width, initial-scale=1'},
      {hid: 'description', name: 'description', content: ''}
    ],
    link: [
      // {rel: 'icon', type: 'image/x-icon', href: '/favicon.ico'}
    ]
  },

  // Global CSS: https://go.nuxtjs.dev/config-css
  css: [
    '~/assets/css/style.css',
  ],

  // Plugins to run before rendering page: https://go.nuxtjs.dev/config-plugins
  plugins: [
    // '@/plugins/utils',
  ],

  // Auto import components: https://go.nuxtjs.dev/config-components
  components: true,

  // Modules for dev and build (recommended): https://go.nuxtjs.dev/config-modules
  buildModules: [],

  // Modules: https://go.nuxtjs.dev/config-modules
  modules: [
    // https://go.nuxtjs.dev/axios
    '@nuxtjs/axios',

    'bootstrap-vue/nuxt',
  ],

  // Build Configuration: https://go.nuxtjs.dev/config-build
  build: {},

  // Axios module configuration: https://go.nuxtjs.dev/config-axios
  axios: {
    debug: true,
    proxy: true,
    // baseURL: 'http://localhost:8000/fosterparent/',
    host: process.env.API_HOST,
    port: process.env.API_PORT,
    prefix: '/fosterparent/',
  },

  proxy: {
    'api/': {// /にすると、/が全部プロキシを挟んでしまうので、api/を挟んだ場合のみにする。
      target: axiosTarget,
      pathRewrite: pathRewrite,
    },
  },

  router: {
    // .env.local or prod にて環境変数NODE_ENVを設定し、
    // モードによってルートパスを切り替える。
    // server.go の router.StaticFSに繋がる
    base: process.env.NODE_ENV === 'development'
      ? '/fosterparent/'
      : '/fosterparent/',
  },

  vue: {
    config: {
      productionTip: true,
      devtools: true
    }
  }
}
