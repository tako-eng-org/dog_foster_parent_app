<template>
  <div class="component">
    <div class="row">
      <p v-if="isView">{{ title }} : {{ getLabel(mapName, itemValue) }}</p>
    </div>
  </div>
</template>

<script>
import CON from "~/components/const/const"

export default {
  component: {
    CON,
  },
  data() {
    return {
      isView: true,
    }
  },
  props: {
    title: String,
    mapName: String,
    itemValue: Number,
  },
  methods: {
    // TODO getLabel関数を共通化する
    getLabel(mapName, i) { // 第1引数：mapの変数名, 第2引数:数値
      return mapName.find(map => map.value === i).label //指定したmapのvalueに紐づくlavelを返す
    },
  },
}
</script>
