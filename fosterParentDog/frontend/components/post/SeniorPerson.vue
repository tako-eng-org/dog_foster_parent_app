<template>
  <div class="senior-person">
    <div v-if="readonly">
      高齢者への譲渡: {{ seniorPersonValue(value) }}
    </div>
    <div v-else>
      <label>高齢者への譲渡
        <select v-model="inputValue" class="form-control">
          <option v-for="(seniorPerson, index) in seniorPersonList()" :value="index">
            {{ seniorPerson }}
          </option>
        </select>
      </label>
    </div>
  </div>
</template>

<script>
import {seniorPersonList, seniorPersonValue} from "~/consts/seniorPersonList";

export default {
  data() {
    return {
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
    seniorPersonList,
    seniorPersonValue,
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
