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
  total: {
    type: Number,
    required: true
  },
  current_page: {
    type: Number,
    required: true
  },
  page_size: {
    type: Number,
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
const emit = defineEmits(['changePage', 'pageSizeChange'])
const onChangePage = (page) => {
  emit('changePage', page)
}
const onPageSizeChange = (size) => {
  console.log(size)
  emit('pageSizeChange', size)
}
</script>

<template>
  <a-table :columns="props.columns" :data="props.data" :row-selection="props.rowSelection"
           stripe :pagination="false"/>
  <div class="ant-page">
    <a-pagination :total="props.total" :current="props.current" :page-size="page_size" show-total show-jumper
                  show-page-size @change="onChangePage"
                  @pageSizeChange="onPageSizeChange"/>
  </div>
</template>

<style scoped>
.ant-page {
  margin: 0 auto;
  margin-top: 30px;
  width: 650px;
}

</style>