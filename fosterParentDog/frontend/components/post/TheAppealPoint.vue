<template>
  <div class="appeal-point">
    <div class="row">
      <p>性格アピールポイント : {{ appealPoint }}</p>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    appealPoint: String,
  },
}
</script>
