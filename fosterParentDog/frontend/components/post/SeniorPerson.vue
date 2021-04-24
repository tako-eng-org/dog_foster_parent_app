<template>
  <div class="senior-person">
    <div class="row">
      <div v-if="readonly">
        <p>高齢者への譲渡: {{ getLabel(SENIOR_PERSON_LIST, itemValue) }}</p>
      </div>
      <div v-else>
        <select v-model="selected">
          <option v-for="seniorPerson in SENIOR_PERSON_LIST" v-bind:value="seniorPerson.value">
            高齢者への譲渡 {{ seniorPerson.label }}
          </option>
        </select>
      </div>
    </div>
  </div>
</template>

<script>
import SENIOR_PERSON_LIST from "~/consts/seniorPersonList";
import getLabel from "~/consts/getLabel";

export default {
  data() {
    return {
      SENIOR_PERSON_LIST,
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
