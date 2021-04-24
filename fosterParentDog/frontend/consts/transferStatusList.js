const TRANSFER_STATUS_LIST = {
  0: "応募受付中",
  1: "やりとり中",
  2: "お試し中",
  3: "応募終了",
}

export const transferStatusList = () => {
  return TRANSFER_STATUS_LIST;
}
export const transferStatusValue = (i) => {
  return TRANSFER_STATUS_LIST[i]
}
