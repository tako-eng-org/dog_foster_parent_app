export default {
  target: 'static',
  ssr: false, // サーバーサイドレンダリングを無効化

  // Global page headers: https://go.nuxtjs.dev/config-head
  head: {
    title: 'frontend',
    htmlAttrs: {
      lang: 'en'
    },
    meta: [
      {charset: 'utf-8'},
      {name: 'viewport', content: 'width=device-width, initial-scale=1'},
      {hid: 'description', name: 'description', content: ''}
    ],
    link: [
      {rel: 'icon', type: 'image/x-icon', href: '/favicon.ico'}
    ]
  },

  // Global CSS: https://go.nuxtjs.dev/config-css
  css: [
    '~/assets/css/style.css',
  ],

  // Plugins to run before rendering page: https://go.nuxtjs.dev/config-plugins
  plugins: [
    '@/plugins/utils',
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

  // Axios module configuration: https://go.nuxtjs.dev/config-axios
  axios: {
    proxy: true,
    prefix: '/fosterparent',
  },

  // Build Configuration: https://go.nuxtjs.dev/config-build
  build: {},

  proxy: {
    '/': {
      target: process.env.NODE_ENV === 'production'
        ? 'http://localhost:8000' //本番は本番用ドメインに修正すること
        : 'http://localhost:8000',
// : 'http://dog_app:8000',
      logLevel: 'debug'
    },
  },

  router: {
    // .env.local or prod にて環境変数NODE_ENVを設定し、
    // モードによってルートパスを切り替える。
    // server.go の router.StaticFSに繋がる
    base: process.env.NODE_ENV === 'production'
      ? '/fosterparent/'
      : '/',
    // : '/',
  },
}
