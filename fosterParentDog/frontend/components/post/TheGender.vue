<template>
  <div class="gender">
    <div class="row">
      <p>性別 : {{ $getLabel($GENDER, gender) }}</p>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    gender: Number,
  },
}
</script>
