<script setup>
import FilterBar from "@/components/filter-bar/index.vue";
import Table from "@/components/table/index.vue";
import {onMounted, ref} from "vue";
import Status from "@/components/status/index.vue";
import {Message} from "@arco-design/web-vue";
import {formatTime} from "@/utils/time.js";
import EditForm from "@/components/edit-form/index.vue";
import {
  createSysAdminUsers,
  deleteSysAdminUsers,
  getSysAdminUsersList,
  updateSysAdminUsers
} from "@/api/sys_admin_users.js";
import {
  columns,
  formData,
  formList,
  formListIndexMap,
  formRef,
  formRules,
  formTitle,
  ids,
  list,
  page,
  searchList,
  tableRef
} from "./index.js";


const passwordFormRef = ref()
//表单列表
const passwordFormList = ref([
  {
    label: '密码',
    key: 'password',
    value: "",
    type: 'password',
    placeholder: '请输入新密码'
  }
]);


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

const sorterChange = (field,sort) => {
  console.log(field);
  console.log(sort);
  if (sort&&field) {
    console.log(sort === 'descend')
    fetchPageList({
      currentPage: page.value.current,
      pageSize: page.value.pageSize,
      filter_map: page.value.searchForm,
      order: [field],
      desc: [sort === 'descend']
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
  formList.value[formListIndexMap['password']].type = row ? "hidden" : "password";
  formTitle.value = row ? '编辑' : '添加';
  formData.value = {...row};
  formRef.value.setVisible(true);
};

const showPasswordEdit = (row) => {
  formData.value = {...row};
  passwordFormRef.value.setVisible(true);
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
  </FilterBar>

  <div class="container ant-card">
    <!-- 表格-->
    <Table ref="tableRef" :columns="columns" :data="list" :total="page.total" :current="page.current"
           :pageSize="page.pageSize" @changePage="changePage"
           @pageSizeChange="pageSizeChange"
           @sorterChange="sorterChange" @select="select" @deletes="deletesItem" @showEdit="showEdit">

      <template #status="{ record }">
        <Status :row="record" @onClick="updatedStatus"></Status>
      </template>
      <template #created_at="{ record }">
        {{ formatTime(record.created_at) }}
      </template>
      <template #optional="{ record }">
        <a-button size="mini" @click="showEdit(record)">编辑</a-button>

        <a-button size="mini" class="ml-10" @click="showPasswordEdit(record)">新密码</a-button>

        <a-popconfirm content="确认要删除当前项目吗?" okText="确认删除" cancelText="取消" @ok="deletesItem(record.id)">
          <a-button size="mini" type="outline" class="ml-10" status="danger">
            删除
          </a-button>
        </a-popconfirm>
      </template>
    </Table>
  </div>
  <!-- 编辑表单-->
  <EditForm :title="formTitle" ref="formRef" :model="formList" :form="formData" :rules="formRules" @submit="submit">
  </EditForm>

  <!-- 编辑密码-->
  <EditForm title="修改密码" ref="passwordFormRef" :model="passwordFormList" :form="formData" :rules="formRules"
            @submit="submit">
  </EditForm>
</template>

<style scoped>

</style>