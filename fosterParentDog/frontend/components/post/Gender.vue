<template>
  <div class="gender">
    <div class="row">
      <div v-if="readonly">
        <!--    TODO getLabelいらなくなる（連想配列）    -->
        <p>性別 : {{ getLabel(GENDER_LIST, itemValue) }}</p>
      </div>
      <div v-else>
        <select v-model="selected">
          <option v-for="gender in GENDER_LIST" v-bind:value="gender.value">
            {{ gender.label }}
          </option>
        </select>
      </div>
    </div>
  </div>
</template>

<script>
import GENDER_LIST from "~/consts/genderList";
import getLabel from "~/consts/getLabel";

export default {
  data() {
    return {
      GENDER_LIST,
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
