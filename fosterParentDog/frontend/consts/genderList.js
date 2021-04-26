const GENDER_LIST = {
  0: "女の子",
  1: "男の子",
  2: "その他/不明",
}

export const genderList = () => {
  return GENDER_LIST;
}

export const genderValue = (i) => {
  return GENDER_LIST[i]
}
