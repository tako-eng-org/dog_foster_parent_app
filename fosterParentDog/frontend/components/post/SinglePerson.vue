<template>
  <div class="single-person">
    <div class="row">
      <div v-if="readonly">
        <p>単身者への譲渡: {{ getLabel(SINGLE_PERSON_LIST, itemValue) }}</p>
      </div>
      <div v-else>
        <select v-model="selected">
          <option v-for="singlePerson in SINGLE_PERSON_LIST" v-bind:value="singlePerson.value">
            単身者への譲渡 {{ singlePerson.label }}
          </option>
        </select>
      </div>
    </div>
  </div>
</template>

<script>
import SINGLE_PERSON_LIST from "~/consts/singlePersonList";
import getLabel from "~/consts/getLabel";

export default {
  data() {
    return {
      SINGLE_PERSON_LIST,
      getLabel,
    }
  },
  props: {
    itemValue: Number,
    selected: "",
    readonly: true,
  },
}
</script>
