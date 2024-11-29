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
  },
  rowSelection: {
    type: Object,
    required: false,
    default: () => {
      return {
        type: 'checkbox',
        showCheckedAll: true
      }
    }
  }
});
const pageChange = (page) => {
  console.log(page)
}
</script>

<template>
  <a-table :columns="props.columns" :data="props.data" :row-selection="props.rowSelection"
           stripe :pagination="false"/>
  <div class="ant-page">
    <a-pagination :total="1000" show-total show-jumper show-page-size />
  </div>
</template>

<style scoped>
.ant-page{
  margin:0 auto;
  margin-top: 30px;
  width: 650px;
}

</style>