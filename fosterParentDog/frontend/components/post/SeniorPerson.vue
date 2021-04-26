<template>
  <div class="senior-person">
    <div class="row">
      <div v-if="readonly">
        <p>高齢者への譲渡: {{ seniorPersonValue(itemValue) }}</p>
      </div>
      <div v-else>
        <select v-model="selected">
          <option v-for="seniorPerson in seniorPersonList" v-bind:value="seniorPerson.value">
            高齢者への譲渡 {{ seniorPerson.label }}
          </option>
        </select>
      </div>
    </div>
  </div>
</template>

<script>
import {seniorPersonList, seniorPersonValue} from "~/consts/seniorPersonList";

export default {
  data() {
    return {
      seniorPersonList,
    }
  },
  props: {
    itemValue: Number,
    selected: "",
    readonly: true,
  },
  methods: {
    seniorPersonValue,
  }
}
</script>
