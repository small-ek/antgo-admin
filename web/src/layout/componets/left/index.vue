<script setup>
import Menu from "@/layout/componets/menu/index.vue";
import Logo from "@/layout/componets/logo/index.vue";
import {useLayout} from '@/stores/layout.js';
import LogoImg from "@/layout/componets/logoImg/index.vue";

const onCancel = () => {
  useLayout().setState("showMobileMenu", false);
};
</script>

<template>
  <!-- Mobile Menu -->
  <a-drawer v-if="useLayout().windowWidth < 768" :width="useLayout().sidebarWidth" :visible="useLayout().showMobileMenu"
            placement="left" :closable="false" :header="false" :footer="false" :drawerStyle="{ padding: '0px' }"
            @cancel="onCancel" class="drawer-left">
    <a-layout-sider hide-trigger collapsible :collapsed="useLayout().isCollapsed" :width="useLayout().sidebarWidth"
                    :theme="useLayout().isDarkSidebar ? 'dark' : 'light'" class="menu">
      <Logo/>
      <Menu/>
    </a-layout-sider>
  </a-drawer>

  <!-- PC Menu -->
  <a-layout-sider
      v-if="useLayout().windowWidth > 768 && (useLayout().layout === 'vertical' || useLayout().layout === 'classic')"
      hide-trigger collapsible :collapsed="useLayout().isCollapsed" :width="useLayout().sidebarWidth"
      :theme="useLayout().isDarkSidebar ? 'dark' : 'light'" class="menu">
    <Logo v-if="useLayout().layout === 'vertical'"/>
    <Menu/>
  </a-layout-sider>

  <div class="columns-menu" v-if="useLayout().windowWidth > 768 && useLayout().layout === 'columns'">
    <a-row justify="center" align="center">
      <a-col :span="12">
        <div class="logo-container">
          <LogoImg/>
        </div>
      </a-col>
      <a-col :span="20" v-for="row in 30" :key="row">
        <div class="menu-item hand">
          <div class="icon">
            <font-awesome-icon icon="fa-solid fa-house"/>
          </div>
          <div>首页首页</div>
        </div>
      </a-col>
    </a-row>
  </div>

  <a-layout-sider v-if="useLayout().windowWidth > 768 && useLayout().layout === 'columns'" hide-trigger collapsible
                  :collapsed="useLayout().isCollapsed" :width="useLayout().sidebarWidth"
                  :theme="useLayout().isDarkSidebar ? 'dark' : 'light'" class="menu">
    <Logo v-if="useLayout().layout === 'vertical' || useLayout().layout === 'columns'"/>
    <Menu/>
  </a-layout-sider>
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
  font-size: 12px;
  padding: 20px 5px;
}

.menu-item:hover,
.menu-item-checked {
  background: rgb(var(--primary-6));
}

.icon {
  margin-bottom: 5px;
}

.logo-container {
  margin: 10px 0 25px;
}

.menu-item:hover {
  transform: scale(1.15);
  transition: transform 0.5s;
}
</style>