<template>
  <div class="post-prefecture">
    <div class="row">
      <p>{{ title }}</p>
      <div v-for="postPrefecture in postPrefectureList" v-bind:key="postPrefecture.id">
        <p>{{ getLabel(itemMap, postPrefecture.post_prefecture_id) }} </p>
        <!--        <p>{{ $getLabel($PREFECTURE, postPrefecture.post_prefecture_id) }} </p>-->
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    title: String,
    postPrefectureList: Object,
    itemMap: Object,
  },

  methods: {
    // TODO getLabel関数を共通化する
    getLabel(mapName, i) { // 第1引数：mapの変数名, 第2引数:数値
      return mapName.find(map => map.value === i).label //指定したmapのvalueに紐づくlavelを返す
    },
  },
}
</script>
