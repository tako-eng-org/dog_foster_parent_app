const SINGLE_PERSON_LIST = {
  0: "不可",
  1: "可",
  2: "応相談",
}

export const singlePersonList = () => {
  return SINGLE_PERSON_LIST;
}

export const singlePersonValue = (i) => {
  return SINGLE_PERSON_LIST[i]
}
