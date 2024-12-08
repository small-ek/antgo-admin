<script setup>
import {defineProps, ref} from "vue";

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
  expandable: { //展开行
    type: Array,
    required: false
  },
  rowKey: {
    type: String,
    required: false,
    default: 'id'
  },
  rowSelection: {
    type: Object,
    required: false,
    default: () => {
      return {
        type: 'checkbox',
        showCheckedAll: true,
      }
    }
  }
});
const selectedKeys = ref([])
const emit = defineEmits(['changePage', 'pageSizeChange', 'changeTable', 'select'])
const onChangePage = (page) => {
  emit('changePage', page)
}
const pageSizeChange = (size) => {
  emit('pageSizeChange', size)
}

const slot = props.columns.filter(item => item.slotName)

const changeTable = (data, extra, currentDataSource) => {
  emit('changeTable', data, extra, currentDataSource)
}
const select = (selectedRowKeys, selectedRows) => {
  emit('select', selectedRowKeys)
}

const clearSelected = () => {
  selectedKeys.value = []
}

const selectAll = (selected) => {
  if (selected) {
    emit('select', props.data.map(item => item[props.rowKey]))
  } else {
    emit('select', [])
  }
}
defineExpose({
  clearSelected
})
</script>

<template>
  <a-table :columns="props.columns" :data="props.data" :row-selection="props.rowSelection" stripe :pagination="false"
           @change="changeTable" :expandable="props.expandable" :row-key="props.rowKey"
           v-model:selectedKeys="selectedKeys" @select="select" @select-all="selectAll">
    <!--自定义插槽-->
    <template #[row.slotName]="{record, rowIndex}" v-for="row in slot">
      <slot :name="row.slotName" :record="record" :rowIndex="rowIndex"></slot>
    </template>
  </a-table>
  <div class="ant-page">
    <a-pagination :total="props.total" :current="props.current" :page-size="props.pageSize" show-total show-jumper
                  show-page-size @change="onChangePage" :page-size-options="props.pageSizeOptions"
                  @pageSizeChange="pageSizeChange"/>
  </div>
</template>

<style scoped>
.ant-page {
  margin: 0 auto;
  margin-top: 30px;
  width: 650px;
}

</style>