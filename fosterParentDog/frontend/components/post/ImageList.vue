<template>
  <div class="image-list">
    <div class="container" v-for="imagePath in imagePathList" v-bind:key="imagePath.post_id">
      <ImageOne :imagePath="imagePath.image_path"/>
    </div>
  </div>
</template>

<script>
import ImageOne from "~/components/post/Image";

export default {
  data() {
    return {
      isView: true,
    }
  },
  props: {
    imagePathList: Object,
  },
  components: {
    ImageOne,
  },
}
</script>
