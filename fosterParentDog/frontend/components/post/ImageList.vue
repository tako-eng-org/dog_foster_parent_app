<template>
  <div class="image-list">
    <div class="container" v-for="objectUrl in objectUrlList" :key="objectUrl.post_id">
      <ImageOne :objectUrl="objectUrl.object_key"/>
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
    objectUrlList: Object,
  },
}
</script>
