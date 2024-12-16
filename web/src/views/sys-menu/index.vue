<script setup>
import FilterBar from "@/components/filter-bar/index.vue";
import Table from "@/components/table/index.vue";
import {onMounted} from "vue";
import Status from "@/components/status/index.vue";
import {Message} from "@arco-design/web-vue";
import {formatTime} from "@/utils/time.js";
import EditForm from "@/components/edit-form/index.vue";
import {createSysMenu, deleteSysMenu, getSysMenuList, updateSysMenu} from "@/api/sys_menu.js";
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

import {useTree} from "@/utils/tree.js";
import iconAll from "@/utils/icon.js";

const fetchPageList = async (params) => {
  const res = await getSysMenuList(params.currentPage, params.pageSize, params.filter_map, params.order, params.desc);
  list.value = useTree().buildTreeTable(res.data.items);
  console.log(list.value)
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
  updateSysMenu(row).then((res) => {
    if (res.code === 0) {
      Message.success('操作成功');
    } else {
      Message.error('操作失败');
      row.status = status === 2 ? 1 : 2;
    }
  });
};

const showEdit = (row, parent_id) => {
  setParentTree();
  formList.value[formListIndexMap['icon']].options = iconAll;
  if (parent_id) {
    row = {};
    row.parent_id = parent_id;
    formData.value = row;
  }
  formData.value = {...row};
  formTitle.value = row ? '编辑' : '添加';

  formRef.value.setVisible(true);
};

const setParentTree = () => {
  getSysMenuList(1, 1000).then((res) => {
    const rootOption = {key: 0, title: '根目录', parent_id: 0};
    formList.value[formListIndexMap['parent_id']].options = [rootOption, ...useTree().buildTreeSelect(res.data.items)];
  });
};

const submit = (row) => {
  const action = row.id ? updateSysMenu : createSysMenu;
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
  deleteSysMenu(ids.value).then((res) => {
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
    <Table :rowSelection="null" ref="tableRef" :columns="columns" :data="list" :total="page.total"
           :current="page.current"
           :pageSize="page.pageSize" @changePage="changePage"
           @pageSizeChange="pageSizeChange"
           @changeTable="changeTable" @select="select">
      <template #title="{ record }">
        <span class="icon" v-if="record.icon">
          <font-awesome-icon :icon="record.icon" size="1x"/>
        </span>
        <span class="icon" v-else>
          <font-awesome-icon icon="fa-table-list"/>
        </span>
        <span class="ml-10">{{ record.title }}</span>
      </template>
      <template #status="{ record }">
        <Status :row="record" @onClick="updatedStatus"></Status>
      </template>
      <template #created_at="{ record }">
        {{ formatTime(record.created_at) }}
      </template>
      <template #optional="{ record }">
        <a-button size="mini" @click="showEdit(record)">编辑</a-button>
        <a-button size="mini" class="ml-10" @click="showEdit(record,record.id)">添加子菜单</a-button>

        <a-popconfirm content="确认要删除当前项目吗?" okText="确认删除" cancelText="取消" @ok="deletesItem(record.id)">
          <a-button size="mini" type="outline" class="ml-10" status="danger">
            删除
          </a-button>
        </a-popconfirm>
      </template>
    </Table>
  </div>
  <!-- 编辑表单-->
  <EditForm :title="formTitle" :xxl="24" ref="formRef" :model="formList" :form="formData" :rules="formRules"
            @submit="submit">
    <template #icon>
      <a-button style="height: 27px" v-if="formData['icon']">
        <template #icon>
          <font-awesome-icon :icon="formData['icon']"/>
        </template>
      </a-button>
      <a-select v-model="formData['icon']" :placeholder="'请选择'" allow-clear
                class="input" allow-search>
        <a-option v-for="icon in iconAll" :key="icon" :value="icon">
          <font-awesome-icon :icon="icon"/>
          <span class="ml-10">{{ icon }}</span>
        </a-option>
      </a-select>
    </template>
  </EditForm>

</template>

<style scoped>

</style>