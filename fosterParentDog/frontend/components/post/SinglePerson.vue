<template>
  <div class="single-person">
    <div class="row">
      <div v-if="readonly">
        <p>単身者への譲渡: {{ singlePersonValue(itemValue) }}</p>
      </div>
      <div v-else>
        <select v-model="selected">
          <option v-for="singlePerson in singlePersonList" v-bind:value="singlePerson.value">
            単身者への譲渡 {{ singlePerson.label }}
          </option>
        </select>
      </div>
    </div>
  </div>
</template>

<script>
import {singlePersonList, singlePersonValue} from "~/consts/singlePersonList";

export default {
  data() {
    return {
      singlePersonList,
    }
  },
  props: {
    itemValue: Number,
    selected: "",
    readonly: true,
  },
  methods: {
    singlePersonValue,
  }
}
</script>
