//キャメルケースに修正 genderList
// 関数的な扱い
// genderList
// 例えばvalueだけの配列を返すとか、functionを用意する時に、
const GENDER_LIST = [
  {
    value: 0,
    label: "女の子"
  },
  {
    value: 1,
    label: "男の子"
  },
  {
    value: 2,
    label: "その他/不明"
  }
]
/*
KW: 連想配列
const GENDERS = {
  0: "女の子",
  1: "男の子",
  2: "その他/不明",
}
 */

export default GENDER_LIST;
