<template>
  <div class="form-group row">
    <label>{{ title }}:
      <input type="text"
             v-model="inputValue"
             :class="className"
             :readonly="readonly"
      >
    </label>
  </div>
</template>

<script>
export default {
  data() {
    return {
      className: this.readonly ? 'form-control-plaintext' : 'form-control'
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
