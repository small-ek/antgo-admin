<script setup>
import {defineProps, ref} from "vue";
import {columns, ids} from "@/views/sys-user-admin/index.js";

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
    type: Object,
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
  },
  isDelete: {
    type: Boolean,
    required: false,
    default: true
  },
  isEdit: {
    type: Boolean,
    required: false,
    default: true
  }
});
const selectedKeys = ref([])
const emit = defineEmits(['changePage', 'pageSizeChange', 'changeTable', 'select', 'deletesItem', 'showEdit'])
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

const deletes = () => {
  emit('deletes')
}

const showEdit = () => {
  emit('showEdit')
}

defineExpose({
  clearSelected
})
</script>

<template>
  <a-row :gutter="15" style="margin-bottom: 10px">
    <a-col :xs="14" :sm="12" :md="12" :lg="12" :xl="12" :xxl="12">
      <a-button type="outline" @click="showEdit()" v-if="isEdit">
        <template #icon>
          <icon-plus/>
        </template>
        添加
      </a-button>
      <a-popconfirm content="确认要删除选中项目吗?" okText="确认删除" cancelText="取消" @ok="deletes">
        <a-button v-if="isDelete" type="outline" class="ml-10" status="danger" v-show="ids.length>0">
          <template #icon>
            <icon-delete/>
          </template>
          删除
        </a-button>
      </a-popconfirm>
    </a-col>
    <a-col :xs="10" :sm="12" :md="12" :lg="12" :xl="12" :xxl="12">
      <div style="float: right">
        <a-tooltip content="列设置">
          <a-dropdown>
            <a-button type="outline" size="mini">
              <template #icon>
                <icon-list/>
              </template>
            </a-button>
            <template #content>
              <div style="padding: 0 1vw">
                <a-checkbox-group :default-value="[]" direction="vertical">
                  <a-checkbox :value="row.dataIndex" v-for="row in columns">{{ row.title }}</a-checkbox>
                </a-checkbox-group>
              </div>
            </template>
          </a-dropdown>
        </a-tooltip>

        <a-tooltip content="导出">
          <a-dropdown>
            <a-button type="outline" class="ml-10" size="mini">
              <template #icon>
                <icon-download/>
              </template>
            </a-button>
          </a-dropdown>
        </a-tooltip>

        <a-tooltip content="打印">
          <a-dropdown>
            <a-button type="outline" class="ml-10" size="mini">
              <template #icon>
                <icon-printer/>
              </template>
            </a-button>
          </a-dropdown>
        </a-tooltip>

        <a-tooltip content="全屏">
          <a-dropdown>
            <a-button type="outline" class="ml-10" size="mini">
              <template #icon>
                <icon-fullscreen/>
              </template>
            </a-button>
          </a-dropdown>
        </a-tooltip>
      </div>
    </a-col>
  </a-row>
  <a-table :columns="props.columns" :data="props.data" :row-selection="props.rowSelection" :pagination="false"
           @change="changeTable" :expandable="props.expandable" :row-key="props.rowKey"
           v-model:selectedKeys="selectedKeys" @select="select" @select-all="selectAll" stripe column-resizable
           :bordered="{cell:true}">
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
  width: fit-content;
}

</style>