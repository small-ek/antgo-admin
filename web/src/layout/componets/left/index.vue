<script setup>
import {onMounted, ref} from 'vue';
import Menu from "@/layout/componets/menu/index.vue";
import Logo from "@/layout/componets/logo/index.vue";
import {useLayout} from '@/stores/layout.js';
import LogoImg from "@/layout/componets/logoImg/index.vue";
import {getMenu} from "@/api/menu.js";
import {useTree} from "@/utils/tree.js";
import {useUserLoginStore} from "@/stores/userLogin.js";

const menuRef = ref();
const mobileMenuRef = ref();
const rightMenu = ref([]);
const parentMenu = ref([]);
const onCancel = () => {
  useLayout().setState("showMobileMenu", false);
};

onMounted(async () => {
  getMenu().then((res) => {
    parentMenu.value = useTree().buildTree(res.data.items);
    useUserLoginStore().setMenu(parentMenu.value)
    mobileMenuRef.value.setList(parentMenu.value)
    menuRef.value.setList(parentMenu.value)
    console.log(parentMenu.value)
    if (useLayout().layout !== 'columns') {

      rightMenu.value = parentMenu.value
    }
  });
})

const onParentMenu = (row) => {
  console.log('onParentMenu');
  menuRef.value.setList(row.children)
  rightMenu.value = row.children
  console.log(rightMenu.value)
  console.log(rightMenu.value.length)
};
</script>

<template>
  <!-- Mobile Menu -->
  <a-drawer v-show="useLayout().windowWidth < 768" :width="useLayout().sidebarWidth" :visible="useLayout().showMobileMenu"
            placement="left" :closable="false" :header="false" :footer="false" :drawerStyle="{ padding: '0px' }"
            @cancel="onCancel" class="drawer-left">
    <a-layout-sider hide-trigger collapsible :collapsed="useLayout().isCollapsed" :width="useLayout().sidebarWidth"
                    :theme="useLayout().isDarkSidebar ? 'dark' : 'light'" class="menu">
      <Logo/>
      <Menu ref="mobileMenuRef"/>
    </a-layout-sider>
  </a-drawer>

  <!-- PC Menu -->
  <a-layout-sider
      v-show="useLayout().windowWidth > 768 && (useLayout().layout === 'vertical' || useLayout().layout === 'classic')"
      hide-trigger collapsible :collapsed="useLayout().isCollapsed" :width="useLayout().sidebarWidth"
      :theme="useLayout().isDarkSidebar ? 'dark' : 'light'" class="menu">
    <Logo v-if="useLayout().layout === 'vertical'"/>
    <Menu ref="menuRef"/>
  </a-layout-sider>

  <!--分栏左侧菜单-->
  <a-layout-sider :width="80" v-if="useLayout().windowWidth > 768 && useLayout().layout === 'columns'">
    <div class="columns-menu">
      <a-row justify="center" align="center">
        <a-col :span="12">
          <div class="logo-container">
            <LogoImg/>
          </div>
        </a-col>
        <a-col :span="20" v-for="row in parentMenu" :key="row">
          <div class="menu-item hand" @click="onParentMenu(row)">
            <div class="icon" v-show="row.icon!==''">
              <font-awesome-icon :icon="row.icon" size="xl"/>
            </div>
            <div class="icon" v-show="row.icon===''">
              <font-awesome-icon icon="fa-solid fa-table-list"/>
            </div>
            <div>{{ row.title }}</div>
          </div>
        </a-col>
      </a-row>
    </div>
  </a-layout-sider>
  <!--分栏右侧菜单-->
  <div v-show="rightMenu.length > 0">
    <a-layout-sider v-show="useLayout().windowWidth > 768 && useLayout().layout === 'columns'" hide-trigger collapsible
                    :collapsed="useLayout().isCollapsed" :width="useLayout().sidebarWidth"
                    :theme="useLayout().isDarkSidebar ? 'dark' : 'light'" class="menu">
      <Logo v-if="useLayout().layout === 'vertical' || useLayout().layout === 'columns'"/>
      <Menu ref="menuRef"/>
    </a-layout-sider>
  </div>

</template>

<style scoped>
.menu {
  box-shadow: 2px 0 8px 0 rgba(29, 35, 41, 0.05);
  height: 100%;
}

.menu-dark {
  background: #232324;
}

.columns-menu {
  background: var(--color-menu-dark-bg);
  width: 80px;
  padding-top: 0.5vw;
  height: 99%;
  overflow: auto;
}

.columns-menu::-webkit-scrollbar {
  display: none;
}

.menu-item {
  color: #f5f5f5;
  text-align: center;
  border-radius: 10px;
  font-size: 11px;
  padding: 20px 5px;
}

.menu-item:hover,
.menu-item-checked {
  background: rgb(var(--primary-6));
}

.icon {
  margin-bottom: 6px;
}

.logo-container {
  margin: 10px 0 25px;
}

.menu-item:hover {
  transform: scale(1.15);
  transition: transform 0.5s;
}
</style>