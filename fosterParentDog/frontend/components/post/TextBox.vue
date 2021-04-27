<template>
  <div class="form-inline">
    <label for="text-box">{{ title }}:</label>
    <input id="text-box"
           type="text"
           class="form-control"
           v-model="inputValue"
           :class="ioMode"
    >
  </div>
</template>

<script>
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
      default: "debug_TextBoxのvalue初期値",
      required: true
    },

    readonly: {
      type: Boolean,
      default: true,
      required: false,
    },

    title: {
      type: String,
      default: "debug_TextBoxのtitle初期値",
      required: true,
    },

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
