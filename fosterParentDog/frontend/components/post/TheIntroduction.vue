<template>
  <div class="introduction">
    <div class="row">
      <p>自己紹介 : {{ introduction }}</p>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    introduction: String,
  },
}
</script>
