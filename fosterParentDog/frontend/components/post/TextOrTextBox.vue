<template>
  <div class="component">
    <div class="row">
      <p v-if="isView">{{ title }} : {{ detail }}</p>
      <p v-else>{{ title }} <input v-model="name" placeholder=""></p>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      isView: true,
    }
  },
  props: {
    title: String,
    detail: String,
  },
}
</script>
