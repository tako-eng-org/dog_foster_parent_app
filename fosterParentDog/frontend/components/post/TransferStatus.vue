<template>
  <div class="transfer-status">
    <div v-if="readonly">
      譲渡ステータス: {{ transferStatusValue(value) }}
    </div>
    <div v-else>
      <div class="form-group row">
        <label>譲渡ステータス
          <select v-model="inputValue" class="form-control">
            <option v-for="(transferStatus, index) in transferStatusList()" :value="index">
              {{ transferStatus }}
            </option>
          </select>
        </label>
      </div>
    </div>
  </div>
</template>

<script>
import {transferStatusList, transferStatusValue} from "~/consts/transferStatusList";

export default {
  data() {
    return {}
  },

  props: {
    value: { //子コンポーネントから親コンポーネントへバインディングする設定
      required: true
    },

    readonly: {
      type: Boolean,
      default: true,
      required: false,
    },
  },

  methods: {
    transferStatusList,
    transferStatusValue,
  },

  computed: {
    inputValue: { //子コンポーネントから親コンポーネントへバインディングする設定
      get() {
        return this.value;
      },
      set(value) {
        this.$emit('input', Number(value));
      }
    },
  },
}
</script>
