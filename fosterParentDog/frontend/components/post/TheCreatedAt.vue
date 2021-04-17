<template>
  <div class="created-at">
    <div class="row">
      <p>投稿日時 : {{ createdAt }}</p>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    createdAt: String,
  },
}
</script>
