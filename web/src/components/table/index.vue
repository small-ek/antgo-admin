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
  current: {
    type: Number,
    required: true
  },
  pageSize: {
    type: Number,
    required: true
  },
  pageSizeOptions: {
    type: Array,
    required: false,
    default: () => [10, 20, 50, 100, 200, 500, 1000]
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
const emit = defineEmits(['changePage', 'pageSizeChange', 'onChangeTable'])
const onChangePage = (page) => {
  emit('changePage', page)
}
const onPageSizeChange = (size) => {
  emit('pageSizeChange', size)
}
const slot = props.columns.filter(item => item.slotName)

const onChangeTable = (data, extra, currentDataSource) => {
  emit('onChangeTable', data, extra, currentDataSource)
}
</script>

<template>
  <a-table :columns="props.columns" :data="props.data" :row-selection="props.rowSelection"
           stripe :pagination="false"  @change="onChangeTable">
    <!--自定义插槽-->
    <template #[row.slotName]="{record, rowIndex}" v-for="row in slot">
      <slot :name="row.slotName" :record="record" :rowIndex="rowIndex"></slot>
    </template>
  </a-table>
  <div class="ant-page">
    <a-pagination :total="props.total" :current="props.current" :page-size="props.pageSize" show-total show-jumper
                  show-page-size @change="onChangePage" :page-size-options="props.pageSizeOptions"
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