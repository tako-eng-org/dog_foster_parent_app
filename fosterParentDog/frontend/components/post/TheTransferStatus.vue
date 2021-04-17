<template>
  <div class="transfer-status">
    <div class="row">
      <p>譲渡ステータス : {{ $getLabel($TRANSFER_STATUS, transferStatus) }}</p>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    transferStatus: Number,
  },
}
</script>
