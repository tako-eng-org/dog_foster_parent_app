<template>
  <div class="single-person">
      <div v-if="readonly">
        単身者への譲渡: {{ singlePersonValue(value) }}
      </div>
    <div v-else>
      <form class="form-inline">
        <label for="select">単身者への譲渡</label>
        <select id="select" v-model="inputValue" class="form-control">
          <option v-for="(singlePerson, index) in singlePersonList()" :value="index">
            {{ singlePerson }}
          </option>
        </select>
      </form>
    </div>
  </div>
</template>

<script>
import {singlePersonList, singlePersonValue} from "~/consts/singlePersonList";

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
    singlePersonList,
    singlePersonValue,
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
