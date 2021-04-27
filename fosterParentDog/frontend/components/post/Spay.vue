<template>
  <div class="spay">
    <div v-if="readonly">
      去勢/避妊手術: {{ spayValue(value) }}
    </div>
    <div v-else>
      <form class="form-inline">
        <label for="spay">去勢/避妊手術</label>
        <select id="spay" v-model="inputValue" class="form-control">
          <option v-for="(spay, index) in spayList()" :value="index">
            {{ spay }}
          </option>
        </select>
      </form>
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
    spayList,
    spayValue,
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
