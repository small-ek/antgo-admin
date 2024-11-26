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
    key: 'age',
    value: '',
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
    label: '修改时间修改时间',
    key: 'updated_at',
    value: [],
    type: 'dateRange',
    placeholder: '请选择创建时间',
  }]);

</script>

<template>
  <Search :model="searchList"></Search>
</template>

<style scoped>

</style>