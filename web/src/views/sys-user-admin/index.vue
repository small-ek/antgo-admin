<script setup>
import Search from "@/components/search/index.vue";
import Table from "@/components/table/index.vue";
import {onMounted, ref} from "vue";
import {getSysAdminUsersList, updateSysAdminUsers} from "@/api/sys_admin_users.js";
import Status from "@/components/status/index.vue";
import {Message} from "@arco-design/web-vue";
import {formatTime} from "@/utils/time.js";

const searchList = ref([
  {
    label: '用户名',
    key: 'username',
    value: "",
    type: 'input',
    placeholder: '请输入姓名'
  },
  {
    label: '昵称',
    key: 'nick_name',
    value: "",
    type: 'input',
    placeholder: '请输入昵称'
  },
  {
    label: '电子邮箱',
    key: 'email',
    value: "",
    type: 'input',
    placeholder: '请输入电子邮箱'
  },
  {
    label: '手机号',
    key: 'phone',
    value: "",
    type: 'input',
    placeholder: '请输入年龄'
  },
  {
    label: '状态',
    key: 'status',
    value: "",
    type: 'select',
    placeholder: '请选择状态',
    options: [
      {
        label: '已禁用',
        value: 1
      },
      {
        label: '已启用',
        value: 2
      }],
  },

  {
    label: '创建时间',
    key: 'created_at',
    value: [],
    type: 'dateRange',
    placeholder: '请选择创建时间',
  }]);

const columns = [
  {
    title: '标识',
    dataIndex: 'id',
    minWidth: 30,
    sortable: {
      sortDirections: ['ascend', 'descend']
    }
  },
  {
    title: '用户名',
    dataIndex: 'username',
    minWidth: 50,
    tooltip: true,
    ellipsis: true,
  },
  {
    title: '昵称',
    dataIndex: 'nick_name',
    minWidth: 100,
    tooltip: true,
    ellipsis: true,
  },
  {
    title: '手机号',
    dataIndex: 'phone',
    minWidth: 100,
    tooltip: true,
    ellipsis: true,
  },
  {
    title: '电子邮箱',
    dataIndex: 'email',
    minWidth: 100,
    tooltip: true,
    ellipsis: true,
  }, {
    title: '状态',
    dataIndex: 'status',
    minWidth: 100,
    tooltip: true,
    ellipsis: true,
    slotName: 'status'
  }, {
    title: '创建时间',
    dataIndex: 'created_at',
    tooltip: true,
    ellipsis: true,
    minWidth: 100,
    slotName: 'created_at'
  }, {
    title: '操作',
    slotName: 'optional'
  }
];
const list = ref([]);

const page = ref({
  current: 1,
  pageSize: 10,
  total: 0,
  searchForm: {}
})

//搜索
const onSearch = (row) => {
  page.value.searchForm = row
  page.value.current = 1
  getPageList({
    currentPage: page.value.current,
    pageSize: page.value.pageSize,
    updateTotal: true,
    filter_map: row
  })
}

const getPageList = async ({currentPage = 1, pageSize = 10, updateTotal = false, filter_map = {}, order, desc}) => {
  const res = await getSysAdminUsersList(currentPage, pageSize, filter_map, order, desc)
  list.value = res.data.items
  page.value.current = currentPage
  page.value.pageSize = pageSize
  if (updateTotal) {
    page.value.total = res.data.total;
  }
}

onMounted(() => {
  getPageList({
    currentPage: page.value.current,
    pageSize: page.value.pageSize,
    updateTotal: true
  })
})
//切换页码
const onChangePage = (current) => {
  getPageList({
    currentPage: current,
    pageSize: page.value.pageSize,
    filter_map: page.value.searchForm
  })
}
//设置每页显示多少条
const onPageSizeChange = (size) => {
  getPageList({
    currentPage: page.value.pageSize,
    pageSize: size,
    filter_map: page.value.searchForm
  })
}
//表格变化
const onChangeTable = (data, extra, currentDataSource) => {
  if (extra.sorter) {
    getPageList({
      currentPage: page.value.current,
      pageSize: page.value.pageSize,
      filter_map: page.value.searchForm,
      order: [extra.sorter.field],
      desc: [extra.sorter.direction === 'descend'],
    })
  }
}
//更新状态
const onUpdatedStatus = (status, row) => {
  row.status = status
  updateSysAdminUsers(row).then((res) => {
    if (res.code === 0) {
      Message.success('操作成功')
    } else {
      Message.error('操作失败')
      row.status = status === 2 ? 1 : 2
    }
  })
}
</script>

<template>
  <Search :model="searchList" @search="onSearch">
  </Search>

  <div class="container ant-card">
    <Table :columns="columns" :data="list" :total="page.total" :current="page.current"
           :pageSize="page.pageSize" @changePage="onChangePage" @pageSizeChange="onPageSizeChange"
           @onChangeTable="onChangeTable">
      <template #status="{ record }">
        <Status :row="record" @onClick="onUpdatedStatus"></Status>
      </template>
      <template #created_at="{ record }">
        {{formatTime(record.created_at)}}
      </template>
      <template #optional="{ record }">
        <a-button>编辑</a-button>
      </template>
    </Table>
  </div>
</template>

<style scoped>

</style>