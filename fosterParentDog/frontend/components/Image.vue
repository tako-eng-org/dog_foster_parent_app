<template>
  <div class='col-xs-12'>
    <img class='float-left'
         style='padding:0;margin:0 15px 0 0;'
         v-bind:src="imagepath"
         width="180"
         height="180">
  </div>
</template>

<script>
export default {
  props: {
    imagepath: String, //画像URL
  },
}
</script>
