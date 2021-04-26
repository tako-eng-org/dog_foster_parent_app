<template>
  <div class="image-list">
    <div class="container" v-for="imagePath in imagePathList" v-bind:key="imagePath.post_id">
      <ImageOne :imagePath="imagePath.image_path"/>
    </div>
  </div>
</template>

<script>
import ImageOne from "~/components/post/ImageOne";

export default {
  components: {
    ImageOne,
  },
  props: {
    imagePathList: Object,
  },
}
</script>
