const SENIOR_PERSON_LIST = {
  0: "不可",
  1: "可",
  2: "応相談",
}
export const seniorPersonList = () => {
  return SENIOR_PERSON_LIST;
}

export const seniorPersonValue = (i) => {
  return SENIOR_PERSON_LIST[i]
}
