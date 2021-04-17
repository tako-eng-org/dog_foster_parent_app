<template>
  <div class="old">
    <div class="row">
      <p>年齢 : {{ old }}</p>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    old: String,
  },
}
</script>
