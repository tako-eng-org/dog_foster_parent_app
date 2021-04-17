<template>
  <div class="spay">
    <div class="row">
      <p>去勢/避妊手術 : {{ $getLabel($SPAY, spay) }}</p>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    spay: Number,
  },
}
</script>
