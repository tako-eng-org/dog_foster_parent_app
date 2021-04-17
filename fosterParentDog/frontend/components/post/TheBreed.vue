<template>
  <div class="breed">
    <div class="row">
      <p>犬種 : {{ breed }}</p>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    breed: String,
  },
}
</script>
