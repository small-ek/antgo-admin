<script setup>
import {defineProps, ref} from "vue";
import {ids} from "@/views/sys-user-admin/index.js";
import {useCsv} from "@/utils/csv.js";
import {onScreenFull} from "@/utils/screenfull.js";

const props = defineProps({
  name: {type: String, required: true},
  //列
  columns: {type: Object, required: true},
  //数据
  data: {type: Array, required: true},
  //总数
  total: {type: Number, required: true},
  //当前页
  current: {type: Number, required: true},
  //每页显示数量
  pageSize: {type: Number, required: true},
  //展开行
  expandable: {type: Object, required: false},
  //隐藏删除按钮
  hideDel: {type: Boolean, required: false, default: false},
  //隐藏添加按钮
  hideAdd: {type: Boolean, required: false, default: false},
  //自定义导出
  customExport: {type: Boolean, required: false, default: false},
  //隐藏导出按钮
  hideExport: {type: Boolean, required: false, default: false},
  //行key
  rowKey: {type: String, required: false, default: 'id'},
  //每页显示数量选项
  pageSizeOptions: {
    type: Array,
    required: false,
    default: () => [10, 20, 50, 100, 200, 500, 1000]
  },
  //行选择
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
});
//响应式数据
const selectedKeys = ref([]);
const columns = ref(props.columns.filter(column => column.visible));
const defaultColumn = ref(props.columns.filter(item => item.title && item.visible).map(item => item.title));
const slot = props.columns.filter(item => item.slotName)

//定义事件
const emit = defineEmits(['changePage', 'pageSizeChange', 'sorterChange', 'select', 'deletes', 'showEdit', 'exportData'])
/**
 * 分页
 * @param page
 */
const onChangePage = (page) => emit('changePage', page);
/**
 * 每页显示数量
 * @param size
 */
const pageSizeChange = (size) => emit('pageSizeChange', size)

/**
 * 切换表格
 * @param field
 * @param sort
 */
const sorterChange = (field, sort) => emit('sorterChange', field, sort)
/**
 * 选中
 * @param selectedRowKeys
 * @param selectedRows
 */
const select = (selectedRowKeys, selectedRows) => emit('select', selectedRowKeys)

/**
 * 删除
 */
const deletes = () => emit('deletes');

/**
 * 显示编辑
 */
const showEdit = () => emit('showEdit');

/**
 * 清空选中
 */
const clearSelected = () => {
  selectedKeys.value = []
}
/**
 * 全选
 * @param selected
 */
const selectAll = (selected) => {
  if (selected) {
    emit('select', props.data.map(item => item[props.rowKey]))
  } else {
    emit('select', [])
  }
}


/**
 * 切换列
 * @param value
 */
const changeColumn = (value) => {
  const visibleTitles = new Set(value);
  const updatedColumns = [];

  props.columns.forEach(column => {
    column.visible = visibleTitles.has(column.title);
    if (column.visible) {
      updatedColumns.push(column.title);
    }
  });
  columns.value = props.columns.filter(column => column.visible);
  defaultColumn.value = updatedColumns;
};

/**
 * 导出数据
 */
const exportData = () => {
  if (props.customExport) {
    emit('exportData')
  } else {
    useCsv().exportCsv([props.columns.filter(column => column.visible && column.title !== '操作').map(column => column.title), ...props.data.map(row => props.columns.filter(column => column.visible).map(column => row[column.dataIndex]))], "导出数据")
  }
}

defineExpose({
  clearSelected
})
</script>

<template>
  <div id="ant-table">
    <a-row :gutter="15" style="margin-bottom: 5px;margin-top: 5px">
      <a-col :xs="14" :sm="12" :md="12" :lg="12" :xl="12" :xxl="12">
        <a-button type="outline" @click="showEdit()" v-if="!hideAdd">
          <template #icon>
            <icon-plus/>
          </template>
          添加
        </a-button>
        <a-popconfirm content="确认要删除选中项目吗?" okText="确认删除" cancelText="取消" @ok="deletes">
          <a-button v-if="!hideDel" type="outline" class="ml-10" status="danger" v-show="ids.length>0">
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
                  <a-checkbox-group :default-value="defaultColumn" direction="vertical" @change="changeColumn">
                    <a-checkbox :value="row.title" v-for="row in props.columns">{{ row.title }}</a-checkbox>
                  </a-checkbox-group>
                </div>
              </template>
            </a-dropdown>
          </a-tooltip>

          <a-tooltip content="导出" v-if="!hideExport">
            <a-dropdown>
              <a-button type="outline" class="ml-10" size="mini" @click="exportData">
                <template #icon>
                  <icon-download/>
                </template>
              </a-button>
            </a-dropdown>
          </a-tooltip>

          <a-tooltip content="全屏">
            <a-dropdown>
              <a-button type="outline" class="ml-10" size="mini" @click="onScreenFull('ant-table')">
                <template #icon>
                  <icon-fullscreen/>
                </template>
              </a-button>
            </a-dropdown>
          </a-tooltip>
        </div>
      </a-col>
    </a-row>
    <a-table :columns="columns" :data="props.data" :row-selection="props.rowSelection" :pagination="false"
             @sorterChange="sorterChange" :expandable="props.expandable" :row-key="props.rowKey" :bordered="false"
             v-model:selectedKeys="selectedKeys" @select="select" @select-all="selectAll" stripe column-resizable>
      <!--自定义插槽-->
      <template #[row.slotName]="{record, rowIndex}" v-for="row in slot">
        <slot :name="row.slotName" :record="record" :rowIndex="rowIndex"></slot>
      </template>
    </a-table>

    <!--分页-->
    <div class="ant-page">
      <a-pagination :total="props.total" :current="props.current" :page-size="props.pageSize" show-total show-jumper
                    show-page-size @change="onChangePage" :page-size-options="props.pageSizeOptions"
                    @pageSizeChange="pageSizeChange"/>
    </div>
  </div>
</template>

<style scoped>
.ant-page {
  margin: 0 auto;
  margin-top: 30px;
  width: fit-content;
}

</style>