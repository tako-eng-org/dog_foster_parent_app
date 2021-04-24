const SPAY_LIST = {
  0: "手術済",
  1: "手術未",
  2: "手術予定",
}

export const spayList = () => {
  return SPAY_LIST;
}

export const spayValue = (i) => {
  return SPAY_LIST[i]
}
