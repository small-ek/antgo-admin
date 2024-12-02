<script setup>
import Search from "@/components/search/index.vue";
import Table from "@/components/table/index.vue";
import {onMounted, ref} from "vue";
import {getSysAdminUsersList} from "@/api/sys_admin_users.js";

const searchList = ref([
  {
    label: '用户名',
    key: 'username',
    value: "",
    type: 'input',
    placeholder: '请输入姓名'
  },
  {
    label: '手机号',
    key: 'phone',
    value: "",
    type: 'input',
    placeholder: '请输入年龄'
  },
  {
    label: '性别',
    key: 'sex',
    value: "",
    type: 'select',
    placeholder: '请选择性别',
    options: [
      {
        label: '男',
        value: '1'
      },
      {
        label: '女',
        value: '2'
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
    minWidth: 30
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
  }, {
    title: '创建时间',
    dataIndex: 'created_at',
    tooltip: true,
    ellipsis: true,
    minWidth: 100
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
  getPageList({currentPage: page.value.current, pageSize: page.value.pageSize, updateTotal: true, filter_map: row})
}

const getPageList = async ({currentPage, pageSize, updateTotal = false, filter_map}) => {
  const res = await getSysAdminUsersList(currentPage, pageSize, filter_map)
  list.value = res.data.items
  page.value.current = currentPage
  page.value.pageSize = pageSize
  if (updateTotal) {
    page.value.total = res.data.total;
  }
}

onMounted(() => {
  getPageList({currentPage: page.value.current, pageSize: page.value.pageSize, updateTotal: true})
})
//切换页码
const onChangePage = (current) => {
  getPageList({currentPage: current, pageSize: page.value.pageSize, filter_map: page.value.searchForm})
}
//设置每页显示多少条
const onPageSizeChange = (size) => {
  getPageList({currentPage: page.value.pageSize, pageSize: size, filter_map: page.value.searchForm})
}
</script>

<template>
  <Search :model="searchList" @search="onSearch">
    <template #updated_at3="{form}">
      <a-input v-model="form.updated_at3" autocomplete="off" class="input" placeholder="请输入" allow-clear/>
    </template>
  </Search>

  <div class="container ant-card">
    <Table :columns="columns" :data="list" :total="page.total" :current="page.current"
           :pageSize="page.pageSize" @changePage="onChangePage" @pageSizeChange="onPageSizeChange"></Table>
  </div>
</template>

<style scoped>

</style>