<template>
  <div class="component">
    <div class="row">
      <!--   TODO readonly=falseで入力モード   -->
      {{ title }}:<input type="text" readonly class="form-control-plaintext" readonly="true" :value="detail">
    </div>
  </div>
</template>

<script>
export default {
  props: {
    title: String,
    detail: String,
    readonly: true,
  },
}
// TODO if readonlyの時は、inputの内容を入力する
</script>
