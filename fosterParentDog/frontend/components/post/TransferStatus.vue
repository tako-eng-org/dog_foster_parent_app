<template>
  <div class="transfer-status">
    <div class="row">
      <div v-if="readonly">
        <p>譲渡ステータス: {{ getLabel(TRANSFER_STATUS_LIST, itemValue) }}</p>
      </div>
      <div v-else>
        <select v-model="selected">
          <option v-for="transferStatus in TRANSFER_STATUS_LIST" v-bind:value="transferStatus.value">
            譲渡ステータス {{ transferStatus.label }}
          </option>
        </select>
      </div>
    </div>
  </div>
</template>

<script>
import TRANSFER_STATUS_LIST from "~/consts/transferStatusList";
import getLabel from "~/consts/getLabel";

export default {
  data() {
    return {
      TRANSFER_STATUS_LIST,
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
