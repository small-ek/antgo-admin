<script setup>
import Search from "@/components/search/index.vue";
import Table from "@/components/table/index.vue";
import {onMounted, ref} from "vue";
import {getSysAdminUsersList} from "@/api/sys_admin_users.js";

const searchList = ref([
  {
    label: '姓名',
    key: 'name',
    value: "",
    type: 'input',
    placeholder: '请输入姓名'
  },
  {
    label: '年龄',
    key: 'age',
    value: "",
    type: 'number',
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
    value: '',
    type: 'datePicker',
    placeholder: '请选择创建时间',
  }, {
    label: '三方sdk参数',
    key: 'updated_at',
    value: [],
    type: 'dateRange',
    placeholder: '请选择创建时间',
  }, {
    label: '三方sdk参数',
    key: 'updated_at2',
    value: [],
    type: 'dateRange',
    placeholder: '请选择创建时间',
  }, {
    label: '三方sdk参数3',
    key: 'updated_at3',
    value: [],
    type: 'slot',
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
const data = ref([]);

const page = ref({
  current_page: 1,
  page_size: 10,
  total: 0
})
const onSearch = (row) => {
  console.log(row)
}

const getPageList = async (current_page, page_size) => {
  const res = await getSysAdminUsersList(current_page, page_size)
  data.value = res.data.items
  page.value.total = res.data.total
}

onMounted(() => {

  getPageList(page.current_page, page.page_size)
  // getSysAdminUsersList(page.value.current_page, page.value.page_size).then(res => {
  //   data.value = res.data.items
  //   page.value.total = res.data.total
  // })
})
const onChangePage = (current_page) => {
  console.log(current_page)
  getPageList(current_page, page.page_size)
}
</script>

<template>
  <Search :model="searchList" @search="onSearch">
    <template #updated_at3="{form}">
      <a-input v-model="form.updated_at3" autocomplete="off" class="input" placeholder="请输入" allow-clear/>
    </template>
  </Search>

  <div class="container ant-card">
    <Table :columns="columns" :data="data" :total="page.total" :current_page="page.current_page"
           :page_size="page.page_size" @changePage="onChangePage"></Table>
  </div>


</template>

<style scoped>

</style>