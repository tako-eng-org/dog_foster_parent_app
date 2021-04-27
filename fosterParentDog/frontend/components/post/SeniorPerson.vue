<template>
  <div class="senior-person">
    <div v-if="readonly">
      高齢者への譲渡: {{ seniorPersonValue(value) }}
    </div>
    <div v-else>
      <form class="form-inline">
        <label for="senior-person">高齢者への譲渡</label>
        <select id="senior-person" v-model="inputValue" class="form-control">
          <option v-for="(seniorPerson, index) in seniorPersonList()" :value="index">
            {{ seniorPerson }}
          </option>
        </select>
      </form>
    </div>
  </div>
</template>

<script>
import {seniorPersonList, seniorPersonValue} from "~/consts/seniorPersonList";

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
    seniorPersonList,
    seniorPersonValue,
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
