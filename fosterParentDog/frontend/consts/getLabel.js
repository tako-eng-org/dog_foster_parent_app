let getLabel = (mapName, i) => {
  return mapName.find(map => map.value === i).label
}

export default getLabel;
// export default ({}, inject) => {
//   inject('getLabel', getLabel);
// }
