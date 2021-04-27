<template>
  <div class="transfer-status">
    <div v-if="readonly">
      譲渡ステータス: {{ transferStatusValue(value) }}
    </div>
    <div v-else>
      <form class="form-inline">
        <label for="transfer-status">譲渡ステータス</label>
        <select id="transfer-status" v-model="inputValue" class="form-control">
          <option v-for="(transferStatus, index) in transferStatusList()" :value="index">
            {{ transferStatus }}
          </option>
        </select>
      </form>
    </div>
  </div>
</template>

<script>
import {transferStatusList, transferStatusValue} from "~/consts/transferStatusList";

export default {
  data() {
    return {
      ioMode: { // read/write時の表示切り替え
        'form-control-plaintext': this.readonly,
        readonly: this.readonly,
      },
    }
  },

  props: {
    value: { //子コンポーネントから親コンポーネントへバインディングする設定
      type: String,
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
        this.$emit('input', value);
      }
    },
  },
}
</script>
