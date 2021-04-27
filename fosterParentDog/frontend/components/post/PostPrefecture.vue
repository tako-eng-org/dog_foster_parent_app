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
          <input type="checkbox" id="check" :value="index" v-model="testList">
          <label for="check">{{ prefecture }}</label>
        </div>
        <p>testList: {{ testList }}</p>
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
      testList: [],

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
    prefectureList,
    prefectureValue,
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
