<script setup>
import Menu from "@/layout/componets/menu/index.vue";
import Logo from "@/layout/componets/logo/index.vue";
import {useLayout} from '@/stores/layout.js'
import logoImg from "@/layout/componets/logoImg/index.vue"

const onCancel = () => {
  useLayout().setState("showMobileMenu", false)
}
</script>

<template>
  <!--手机菜单-->
  <a-drawer v-if="useLayout().windowWidth<768" :width="useLayout().sidebarWidth" :visible="useLayout().showMobileMenu" placement="left" :closable="false" :header="false" :footer="false" :drawerStyle="{padding:'0px'}" @cancel="onCancel" class="ant-left">
    <a-layout-sider hide-trigger collapsible :collapsed="useLayout().isCollapsed" :width="useLayout().sidebarWidth" :theme="useLayout().isDarkSidebar===true?'dark':'light'" :class="['ant-menu']">
      <!--logo-->
      <Logo></Logo>
      <!--菜单-->
      <Menu></Menu>
    </a-layout-sider>
  </a-drawer>
  <!--PC菜单-->
  <a-layout-sider v-if="useLayout().windowWidth>768&&(useLayout().layout==='vertical'||useLayout().layout==='classic')" hide-trigger collapsible :collapsed="useLayout().isCollapsed" :width="useLayout().sidebarWidth" :theme="useLayout().isDarkSidebar===true?'dark':'light'" :class="['ant-menu',]">
    <!--logo-->
    <Logo v-if="useLayout().layout==='vertical'"></Logo>
    <!--菜单-->
    <Menu></Menu>
  </a-layout-sider>

  <div class="right-columns" v-if="useLayout().windowWidth>768&&(useLayout().layout==='columns')">
    <a-row justify="center" align="center">
      <a-col :span="12">
        <div style="margin-bottom: 15px">
          <logoImg></logoImg>
        </div>
      </a-col>
      <a-col :span="20">
        <div class="menu-columns-text m-auto hand">
          <div>
            <font-awesome-icon :icon="['fas', 'house']" />
          </div>
          <span>首页</span></div>
      </a-col>
      <a-col :span="20">
        <div class="menu-columns-text m-auto hand"><span>表单 Form</span></div>
      </a-col>
      <a-col :span="20">
        <div class="menu-columns-text m-auto hand"><span>首页首页首页首页首页首页</span></div>
      </a-col>
    </a-row>
  </div>
  <!--PC菜单-->
  <a-layout-sider v-if="useLayout().windowWidth>768&&(useLayout().layout==='columns')" hide-trigger collapsible :collapsed="useLayout().isCollapsed" :width="useLayout().sidebarWidth" :theme="useLayout().isDarkSidebar===true?'dark':'light'" :class="['ant-menu',]">

    <!--logo-->
    <Logo v-if="useLayout().layout==='vertical'||useLayout().layout==='columns'"></Logo>
    <!--菜单-->
    <Menu></Menu>
  </a-layout-sider>
</template>

<style scoped>
.ant-menu {
  box-shadow: 2px 0 8px 0 rgba(29, 35, 41, .05);
  height: 100%;
}

.ant-menu-dark {
  background: #232324;
}

.right-columns {
  background: var(--color-menu-dark-bg);
  width: 70px;
  padding-top: 0.5vw;
}

.menu-columns-text {
  color: #f5f5f5;
  text-align: center;
  height: 70px;
  width: 60px;
  border-radius: 10px;
  font-size: 12px;
  line-height: 18px;
}

.menu-columns-text:hover {
  background: rgb(var(--primary-6));
}

.menu-columns-check {
  background: rgb(var(--primary-6));
}
</style>