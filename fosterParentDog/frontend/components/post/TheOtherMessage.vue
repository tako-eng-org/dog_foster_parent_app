<template>
  <div class="other-message">
    <div class="row">
      <p>健康状態や譲渡条件などの特記事項 : {{ otherMessage }}</p>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    otherMessage: String,
  },
}
</script>
