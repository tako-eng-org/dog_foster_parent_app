<template>
  <div class="post-prefecture">
    <div v-if="readonly">
      譲渡可能都道府県:
      <div v-for="(prefecture, index) in value" :key="index">
        {{ prefectureValue(prefecture.post_prefecture_id) }}
      </div>
    </div>
    <div v-else>
      <div class="form-check">
        <div v-for="(prefecture, index) in prefectureList()" :key="index">
          <input type="checkbox" id="check" :value="Number(index)" v-model="inputValue">
          <label for="check">{{ prefecture }}</label>
        </div>
        <p>value: {{ value }}</p>
        <p>inputValue: {{ inputValue }}</p>
      </div>
    </div>
  </div>
</template>

<script>
import {prefectureList, prefectureValue} from "~/consts/prefectureList";

export default {
  data() {
    return {
      ioMode: { // read/write時の表示切り替え
        'form-control-plaintext': this.readonly,
        readonly: this.readonly,
      },

    }
  },

  model: {
    event: "change",
  },

  props: {
    value: { //子コンポーネントから親コンポーネントへバインディングする設定
      default: () => ([]), //default:[]にした場合、 [Vue warn]: Invalid default value for prop "value": Props with type Object/Array must use a factory function to return the default value.
      required: true
    },

    readonly: {
      type: Boolean,
      default: true,
      required: false,
    },
  },

  methods: {
    prefectureList,
    prefectureValue,
  },

  computed: {
    inputValue: { //子コンポーネントから親コンポーネントへバインディングする設定
      get() {
        return this.value;
      },
      set(newValue) {
        this.$emit('change', newValue);
      }
    },
  },
}
</script>
