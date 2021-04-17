<template>
  <div class="senior-person">
    <div class="row">
      <p>高齢者への譲渡 : {{ $getLabel($SENIOR_PERSON, seniorPerson) }}</p>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    seniorPerson: Number,
  },
}
</script>
