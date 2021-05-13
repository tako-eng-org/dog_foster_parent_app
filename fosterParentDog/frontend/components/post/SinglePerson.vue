<template>
  <div class="single-person">
    <div v-if="readonly">
      単身者への譲渡: {{ singlePersonValue(value) }}
    </div>
    <div v-else>
      <div class="form-group row">
        <label>単身者への譲渡
          <select v-model="inputValue" class="form-control">
            <option v-for="(singlePerson, index) in singlePersonList()" :value="index">
              {{ singlePerson }}
            </option>
          </select>
        </label>
      </div>
    </div>
  </div>
</template>

<script>
import {singlePersonList, singlePersonValue} from "~/consts/singlePersonList";

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
    singlePersonList,
    singlePersonValue,
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
