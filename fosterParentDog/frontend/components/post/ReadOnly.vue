<template>
  <div class="component">
    <div class="row">
      <p v-if="isView">{{ title }} : {{ detail }}</p>
      <!--      elseルートなし-->
      <!--      <p v-else>{{ itemTitle }} <input v-model="itemName" placeholder=""></p>-->
      <!--      <p>投稿日時 : {{ createdAt }}</p>-->
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
