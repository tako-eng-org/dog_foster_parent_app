<template>
  <div class="post-prefecture">
    <div class="row">
      <p>譲渡可能都道府県 : </p>
      <div v-for="postPrefecture in postPrefectureList" v-bind:key="postPrefecture.post_id">
        <p>{{ $getLabel($PREFECTURE, postPrefecture.post_prefecture_id) }} </p>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    postPrefectureList: Object,
  },
}
</script>
