<template>
  <div class="post-prefecture">
    <div class="row">
      <p>譲渡可能都道府県</p>
      <div v-if="readonly">
        <div v-for="postPrefecture in postPrefectureList" v-bind:key="postPrefecture.id">
          <!--        <p>{{postPrefecture.label}}</p>-->
          <p>{{ getLabel(PREFECTURE_LIST, postPrefecture.post_prefecture_id) }} </p>
          <!--        <p>{{ $getLabel($PREFECTURE, postPrefecture.post_prefecture_id) }} </p>-->
        </div>
      </div>
      <div v-else>
        <select v-model="selected">
          <option v-for="prefecture in PREFECTURE_LIST" v-bind:value="prefecture.value">
            {{ prefecture.label }}
          </option>
        </select>
      </div>

    </div>
  </div>
</template>

<script>
import PREFECTURE_LIST from "~/consts/prefectureList";
import getLabel from "~/consts/getLabel";

export default {
  data() {
    return {
      PREFECTURE_LIST,
      getLabel,
    }
  },
  props: {
    postPrefectureList: Object,
    selected: "",
    readonly: true,
  },
}
</script>
