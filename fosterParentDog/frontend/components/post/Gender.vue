<template>
  <div class="gender">
    <div v-if="readonly">
      <!--   TODO: indexなどで表示時は、数値1で受け取り文字列を表示する。（props:Number）-->
      <!--   が、入力時は、文字列で渡ってくる(props:String)のため、型チェックはしないほうが良いか？   -->
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
      inputValue: {{ inputValue }}
    </div>
  </div>

  <!--  FIXME: 表示モードでも、非アクティブでプルダウンが表示されてしまう  -->
  <!--    性別 :-->
  <!--      <select v-model="selected" :class="ioMode">-->
  <!--        <option v-for="gender in genderList()"-->
  <!--                :value="gender.value">-->
  <!--          {{ gender }}-->
  <!--        </option>-->
  <!--      </select>-->
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
      // type: String,
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
        this.$emit('input', value);
      }
    },
  },
}
</script>
