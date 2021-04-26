<template>
  <div class="gender">
    <div class="row">
      <!--   TODO v-ifでなくreadonlyによる表示/入力(水平展開)   -->
      <div v-if="readonly">
        <p>性別 : {{ genderValue(itemValue) }}</p>
      </div>
      <div v-else>
        <select v-model="selected">
          <option v-for="gender in genderList" v-bind:value="gender.value">
            {{ gender.label }}
          </option>
        </select>
      </div>
    </div>
  </div>
</template>

<script>
import {genderList, genderValue} from "~/consts/genderList";

export default {
  data() {
    return {
      genderList,
    }
  },
  props: {
    itemValue: Number,
    selected: "",
    readonly: true,
  },
  methods: {
    genderValue,
  }
}
</script>
