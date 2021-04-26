<template>
  <div class="transfer-status">
    <div class="row">
      <div v-if="readonly">
        <p>譲渡ステータス: {{ transferStatusValue(itemValue) }}</p>
      </div>
      <div v-else>
        <select v-model="selected">
          <option v-for="transferStatus in transferStatusList" v-bind:value="transferStatus.value">
            譲渡ステータス {{ transferStatus.label }}
          </option>
        </select>
      </div>
    </div>
  </div>
</template>

<script>
import {transferStatusList, transferStatusValue} from "~/consts/transferStatusList";

export default {
  data() {
    return {
      transferStatusList,
    }
  },
  props: {
    itemValue: Number,
    selected: "",
    readonly: true,
  },
  methods: {
    transferStatusValue,
  }
}
</script>
