<script setup>
import FilterBar from "@/components/filterBar/index.vue";
import Table from "@/components/table/index.vue";
import {onMounted} from "vue";
import Status from "@/components/status/index.vue";
import {Message} from "@arco-design/web-vue";
import {formatTime} from "@/utils/time.js";
import EditForm from "./form.vue";
import {
  createSysAdminUsers,
  deleteSysAdminUsers,
  getSysAdminUsersList,
  updateSysAdminUsers
} from "@/api/sys_admin_users.js";
import {columns, formData, formList, formRef, formRules, ids, list, page, searchList, tableRef} from "./index.js";

const fetchPageList = async (params) => {
  const res = await getSysAdminUsersList(params.currentPage, params.pageSize, params.filter_map, params.order, params.desc);
  list.value = res.data.items;
  page.value.current = params.currentPage;
  page.value.pageSize = params.pageSize;
  if (params.updateTotal) {
    page.value.total = res.data.total;
  }
};

const search = (row) => {
  page.value.searchForm = row;
  page.value.current = 1;
  fetchPageList({currentPage: 1, pageSize: page.value.pageSize, updateTotal: true, filter_map: row});
};

onMounted(() => {
  fetchPageList({currentPage: page.value.current, pageSize: page.value.pageSize, updateTotal: true});
});

const changePage = (current) => {
  fetchPageList({currentPage: current, pageSize: page.value.pageSize, filter_map: page.value.searchForm});
};

const pageSizeChange = (size) => {
  fetchPageList({currentPage: page.value.pageSize, pageSize: size, filter_map: page.value.searchForm});
};

const changeTable = (data, extra) => {
  if (extra.sorter) {
    fetchPageList({
      currentPage: page.value.current,
      pageSize: page.value.pageSize,
      filter_map: page.value.searchForm,
      order: [extra.sorter.field],
      desc: [extra.sorter.direction === 'descend']
    });
  }
};

const updatedStatus = (status, row) => {
  row.status = status;
  updateSysAdminUsers(row).then((res) => {
    if (res.code === 0) {
      Message.success('操作成功');
    } else {
      Message.error('操作失败');
      row.status = status === 2 ? 1 : 2;
    }
  });
};

const showEdit = (row) => {
  // 如果是编辑，密码默认隐藏
  formList.value[1].type = row ? "hidden" : "password";
  formData.value = {...row};
  formRef.value.setVisible(true);
};

const submit = (row) => {
  const action = row.id ? updateSysAdminUsers : createSysAdminUsers;
  action(row).then((res) => {
    if (res.code === 0) {
      Message.success('操作成功');
      formRef.value.setVisible(false);
      reload();
    }
  });
};

const select = (selectedRowKeys) => {
  ids.value = selectedRowKeys;
};

const deletesItem = (id) => {
  if (id > 0) {
    ids.value.push(id);
  }
  if (ids.value.length === 0) {
    Message.warning('请选择要删除的项目');
    return;
  }
  deleteSysAdminUsers(ids.value).then((res) => {
    if (res.code === 0) {
      Message.success('操作成功');
      ids.value = [];
      tableRef.value.clearSelected();
      reload();
    }
  });
};
// 重新加载
const reload = () => {
  fetchPageList({
    currentPage: page.value.current,
    pageSize: page.value.pageSize,
    filter_map: page.value.searchForm
  });
};
</script>

<template>
  <!--过滤栏-->
  <FilterBar :model="searchList" @search="search">
    <template #topBtn>
      <a-button type="outline" @click="showEdit()">
        <template #icon>
          <icon-plus/>
        </template>
        添加
      </a-button>
      <a-popconfirm content="确认要删除选中项目吗?" okText="确认删除" cancelText="取消" @ok="deletesItem">
        <a-button type="outline" class="ml-10" status="danger" v-show="ids.length>0">
          <template #icon>
            <icon-delete/>
          </template>
          批量删除
        </a-button>
      </a-popconfirm>

    </template>
  </FilterBar>

  <div class="container ant-card">
    <!-- 表格-->
    <Table ref="tableRef" :columns="columns" :data="list" :total="page.total" :current="page.current"
           :pageSize="page.pageSize" @changePage="changePage"
           @pageSizeChange="pageSizeChange"
           @changeTable="changeTable" @select="select">
      <template #status="{ record }">
        <Status :row="record" @onClick="updatedStatus"></Status>
      </template>
      <template #created_at="{ record }">
        {{ formatTime(record.created_at) }}
      </template>
      <template #optional="{ record }">
        <a-button size="mini" @click="showEdit(record)">编辑</a-button>

        <a-popconfirm content="确认要删除当前项目吗?" okText="确认删除" cancelText="取消" @ok="deletesItem(record.id)">
          <a-button size="mini" type="outline" class="ml-10" status="danger">
            删除
          </a-button>
        </a-popconfirm>
      </template>
    </Table>
  </div>
  <!-- 编辑表单-->
  <EditForm ref="formRef" :model="formList" :form="formData" :rules="formRules" @submit="submit">
  </EditForm>
</template>

<style scoped>

</style>