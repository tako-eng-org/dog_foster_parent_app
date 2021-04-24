<template>
  <div class="spay">
    <div class="row">
      <div v-if="readonly">
        <p>去勢/避妊手術: {{ getLabel(SPAY_LIST, itemValue) }}</p>
      </div>
      <div v-else>
        <select v-model="selected">
          <option v-for="spay in SPAY_LIST" v-bind:value="spay.value">
            去勢/避妊手術 {{ spay.label }}
          </option>
        </select>
      </div>
    </div>
  </div>
</template>

<script>
import SPAY_LIST from "~/consts/spayList";
import getLabel from "~/consts/getLabel";

export default {
  data() {
    return {
      SPAY_LIST,
      getLabel,
    }
  },
  props: {
    itemValue: Number,
    selected: "",
    readonly: true,
  },
}
</script>
