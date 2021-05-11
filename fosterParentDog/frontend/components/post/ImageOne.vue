<template>
  <div class='image-one col-xs-12'>
    <img class='float-left'
         style='padding:0;margin:0 15px 0 0;'
         :src="imagePath"
         width="180"
         height="180">
  </div>
</template>

<script>
export default {
  props: {
    imagePath: {
      required: true
    },
  },
}
</script>
