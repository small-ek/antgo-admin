<script setup>
import Search from "@/components/search/index.vue";
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
  }]);

const onSearch = (form) => {
  console.log(form)
  console.log(searchList.value)
}
</script>

<template>
  <Search :model="searchList" @search="onSearch"></Search>
</template>

<style scoped>

</style>