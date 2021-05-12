<template>
  <div class="gender">
    <div v-if="readonly">
      <!--   TODO: indexなどで表示時は、数値1で受け取り文字列を表示する（props:Number）。入力時は文字列(props:String)のため、型チェックはしないほうが良いか   -->
      性別: {{ genderValue(value) }}
    </div>
    <div v-else>
      <form class="form-inline">
        <label>性別
          <select v-model="inputValue" class="form-control">
            <option v-for="(gender, index) in genderList()" :value="index">
              {{ gender }}
            </option>
          </select>
        </label>
      </form>
    </div>
  </div>
</template>

<script>
import {genderList, genderValue} from "~/consts/genderList";

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
      required: true
    },

    readonly: {
      type: Boolean,
      default: true,
      required: false,
    },
  },

  methods: {
    genderList,
    genderValue,
  },

  computed: {
    inputValue: { //子コンポーネントから親コンポーネントへバインディングする設定
      get() {
        return this.value;
      },
      set(value) {
        // 連想配列のkeyとなる。Numberに変換しないとString型で返してしまいレコード登録できない（アプリ側のstruct型と一致せずJSON変換できない）。
        this.$emit('input', Number(value));
      }
    },
  },
}
</script>
