<template>
  <div class="spay">
    <div v-if="readonly">
      去勢/避妊手術: {{ spayValue(value) }}
    </div>
    <div v-else>
      <form class="form-inline">
        <label>去勢/避妊手術
          <select v-model="inputValue" class="form-control">
            <option v-for="(spay, index) in spayList()" :value="index">
              {{ spay }}
            </option>
          </select>
        </label>
      </form>
      inputValue: {{ inputValue }}<br>
      type: {{ typeof (inputValue) }}
    </div>
  </div>
</template>

<script>
import {spayList, spayValue} from "~/consts/spayList";

export default {
  data() {
    return {
      ioMode: { // read/write時の表示切り替え
        'form-control-plaintext': this.readonly,
        readonly: this.readonly,
      }
    }
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
    spayList,
    spayValue,
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
