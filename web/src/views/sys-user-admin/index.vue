<script setup>
import Search from "@/components/search/index.vue";
import Table from "@/components/table/index.vue";
import {onMounted, ref} from "vue";
import {getSysAdminUsersList, updateSysAdminUsers} from "@/api/sys_admin_users.js";
import Status from "@/components/status/index.vue";
import {Message} from "@arco-design/web-vue";
import {formatTime} from "@/utils/time.js";
import Form from "./form.vue";
import {list,form,formRef,columns, formList, page, searchList} from "./index.js";


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
const showEdit = (row) => {
  form.value = row
  formRef.value.setVisible(true)
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
        {{ formatTime(record.created_at) }}
      </template>
      <template #optional="{ record }">
        <a-button size="mini" @click="showEdit(record)">编辑</a-button>
      </template>
    </Table>
  </div>
  <Form ref="formRef" :model="formList" :form="form"></Form>
</template>

<style scoped>

</style>