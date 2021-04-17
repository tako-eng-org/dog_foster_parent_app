<template>
  <div class="dog-name">
    <div class="row">
      <p>犬の名前 : {{ dogName }}</p>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    dogName: String,
  },
}
</script>
