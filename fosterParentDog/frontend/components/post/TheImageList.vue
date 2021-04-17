<template>
  <div class="image-list">
    <div class="container" v-for="imagePath in imagePathList" v-bind:key="imagePath.post_id">
      <ImageComponent :imagepath="imagePath.image_path"/>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    imagePathList: Object,
  },
}
</script>
