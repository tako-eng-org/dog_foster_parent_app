import Vue from 'vue'

Vue.prototype.$TRANSFER_STATUS = [
  {
    value: 0,
    label: "応募受付中"
  },
  {
    value: 1,
    label: "やりとり中"
  },
  {
    value: 2,
    label: "お試し中"
  },
  {
    value: 3,
    label: "応募終了"
  }
]