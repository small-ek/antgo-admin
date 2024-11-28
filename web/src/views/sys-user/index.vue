<script setup>
import Search from "@/components/search/index.vue";
import Table from "@/components/table/index.vue";
import {ref} from "vue";

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
    title: 'Name',
    dataIndex: 'name',
  },
  {
    title: 'Salary',
    dataIndex: 'salary',
  },
  {
    title: 'Address',
    dataIndex: 'address',
  },
  {
    title: 'Email',
    dataIndex: 'email',
  },
];
const data = ref([{
  key: '1',
  name: 'Jane Doe',
  salary: 23000,
  address: '32 Park Road, London',
  email: 'jane.doe@example.com'
}, {
  key: '2',
  name: 'Alisa Ross',
  salary: 25000,
  address: '35 Park Road, London',
  email: 'alisa.ross@example.com'
}, {
  key: '3',
  name: 'Kevin Sandra',
  salary: 22000,
  address: '31 Park Road, London',
  email: 'kevin.sandra@example.com'
}, {
  key: '4',
  name: 'Ed Hellen',
  salary: 17000,
  address: '42 Park Road, London',
  email: 'ed.hellen@example.com'
}, {
  key: '5',
  name: 'William Smith',
  salary: 27000,
  address: '62 Park Road, London',
  email: 'william.smith@example.com'
}]);
const onSearch = (row) => {
  console.log(row)
}

</script>

<template>
  <Search :model="searchList" @search="onSearch">
    <template #updated_at3="{form}">
      <a-input v-model="form.updated_at3" autocomplete="off" class="input" placeholder="请输入" allow-clear/>
    </template>
  </Search>

  <div class="container ant-card">
    <Table :columns="columns" :data="data"></Table>
  </div>



</template>

<style scoped>

</style>