<template>
  <div class="component">
    <div class="row">
      <div v-if="readonly">
        <p>{{ title }}: {{ detail }}</p>
      </div>
      <!--   TODO Readonly plain text を使うと、v-ifがいらなくなる   -->
      <!--   <input type="text" readonly class="form-control-plaintext" id="staticEmail" value="email@example.com">   -->
      <!--   class="form-control-plaintext" 忘れずに   -->
      <div v-else>
        <p>{{ title }} <input readonly="true" v-model="name" placeholder=""></p>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    title: String,
    detail: String,
    readonly: true,
  },
}
// TODO if readonluの時は、inputの内容を入力する
</script>

<!--cssの調整で、inputの表示を変えることができるかも -->
