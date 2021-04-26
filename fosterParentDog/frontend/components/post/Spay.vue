<template>
  <div class="spay">
    <div class="row">
      <div v-if="readonly">
        <p>去勢/避妊手術: {{ spayValue(itemValue) }}</p>
      </div>
      <div v-else>
        <select v-model="selected">
          <option v-for="spay in spayList" v-bind:value="spay.value">
            去勢/避妊手術 {{ spay.label }}
          </option>
        </select>
      </div>
    </div>
  </div>
</template>

<script>
import {spayList, spayValue} from "~/consts/spayList";

export default {
  data() {
    return {
      spayList,
    }
  },
  props: {
    itemValue: Number,
    selected: "",
    readonly: true,
  },
  methods: {
    spayValue,
  }
}
</script>
