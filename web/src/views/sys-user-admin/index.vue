<script setup>
import FilterBar from "@/components/filterBar/index.vue";
import Table from "@/components/table/index.vue";
import {onMounted} from "vue";
import {getSysAdminUsersList, updateSysAdminUsers} from "@/api/sys_admin_users.js";
import Status from "@/components/status/index.vue";
import {Message} from "@arco-design/web-vue";
import {formatTime} from "@/utils/time.js";
import EditForm from "./form.vue";
import {columns, formData, formList, formRef, formRules, list, page, searchList} from "./index.js";


//搜索
const search = (row) => {
  page.value.searchForm = row
  page.value.current = 1
  getPageList({
    currentPage: page.value.current,
    pageSize: page.value.pageSize,
    updateTotal: true,
    filter_map: row
  })
}

/**
 * 获取分页列表
 * @param currentPage 当前页
 * @param pageSize  每页显示多少条
 * @param updateTotal   是否更新总数
 * @param filter_map   过滤条件
 * @param order  排序字段
 * @param desc  是否倒序
 * @returns {Promise<void>}
 */
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

/**
 * 分页变化
 * @param current
 */
const changePage = (current) => {
  getPageList({
    currentPage: current,
    pageSize: page.value.pageSize,
    filter_map: page.value.searchForm
  })
}

/**
 * 每页显示多少条
 * @param size
 */
const pageSizeChange = (size) => {
  getPageList({
    currentPage: page.value.pageSize,
    pageSize: size,
    filter_map: page.value.searchForm
  })
}

/**
 * 表格变化 修改排序
 * @param data
 * @param extra
 * @param currentDataSource
 */
const changeTable = (data, extra, currentDataSource) => {
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

/**
 * 更新状态
 * @param status
 * @param row
 */
const updatedStatus = (status, row) => {
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
/**
 * 显示编辑
 * @param row
 */
const showEdit = (row) => {
  formData.value = {...row}
  formRef.value.setVisible(true)
}

const submit = (row) => {
  if(row.id){
    updateSysAdminUsers(row).then((res) => {
      if (res.code === 0) {
        Message.success('操作成功')
        formRef.value.setVisible(false)
        getPageList({
          currentPage: page.value.current,
          pageSize: page.value.pageSize,
          filter_map: page.value.searchForm
        })
      }
    })
  }

}
</script>

<template>
  <FilterBar :model="searchList" @search="search">
  </FilterBar>

  <div class="container ant-card">
    <Table :columns="columns" :data="list" :total="page.total" :current="page.current"
           :pageSize="page.pageSize" @changePage="changePage" @pageSizeChange="pageSizeChange"
           @changeTable="changeTable">
      <template #status="{ record }">
        <Status :row="record" @onClick="updatedStatus"></Status>
      </template>
      <template #created_at="{ record }">
        {{ formatTime(record.created_at) }}
      </template>
      <template #optional="{ record }">
        <a-button size="mini" @click="showEdit(record)">编辑</a-button>
      </template>
    </Table>
  </div>
  <EditForm ref="formRef" :model="formList" :form="formData" :rules="formRules" @submit="submit"></EditForm>
</template>

<style scoped>

</style>