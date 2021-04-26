<template>
  <div class="post-prefecture">
    <div class="row">
      <p>譲渡可能都道府県</p>
      <div v-if="readonly">
        <div v-for="postPrefecture in postPrefectureList" v-bind:key="postPrefecture.id">
          <p>{{ prefectureValue(postPrefecture.post_prefecture_id) }} </p>
        </div>
      </div>
      <div v-else>
        <select v-model="selected">
          <option v-for="prefecture in prefectureList" v-bind:value="prefecture.value">
            {{ prefecture.label }}
          </option>
        </select>
      </div>

    </div>
  </div>
</template>

<script>
import {prefectureList, prefectureValue} from "~/consts/prefectureList";

export default {
  data() {
    return {
      prefectureList,
    }
  },
  props: {
    postPrefectureList: Object,
    selected: "",
    readonly: true,
  },
  methods: {
    prefectureValue,
  }
}
</script>
