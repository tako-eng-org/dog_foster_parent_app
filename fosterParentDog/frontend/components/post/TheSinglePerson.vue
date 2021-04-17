<template>
  <div class="single-person">
    <div class="row">
      <p>単身者への譲渡 : {{ $getLabel($SINGLE_PERSON, singlePerson) }}</p>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    singlePerson: Number,
  },
}
</script>
