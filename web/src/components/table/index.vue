<script setup>
import {defineProps} from "vue";

const props = defineProps({
  columns: {
    type: Object,
    required: true
  },
  data: {
    type: Array,
    required: true
  }
});
</script>

<template>
  <a-table :columns="props.columns" :data="props.data" />
</template>

<style scoped>

</style>