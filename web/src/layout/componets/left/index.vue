<script setup>
import Menu from "@/layout/componets/menu/index.vue";
import Logo from "@/layout/componets/logo/index.vue";
import {onMounted, onUnmounted} from "vue";
import {useLayout} from '@/stores/layout.js'


const handleResize = () => {
  useLayout().setState("windowWidth", window.innerWidth)
  useLayout().setState("windowHeight", window.innerHeight)
  if (useLayout().windowWidth < 768) {
    useLayout().setState("isCollapse", false)
  } else {
    useLayout().setState("mobileVisible", false)
  }
};

onMounted(() => {
  window.addEventListener('resize', handleResize);
});

onUnmounted(() => {
  window.removeEventListener('resize', handleResize);
});


const onCancel = () => {
  useLayout().setState("mobileVisible", false)
}
</script>

<template>
  <!--手机菜单-->
  <a-drawer v-if="useLayout().windowWidth<768" :width="250" :visible="useLayout().mobileVisible" placement="left" :closable="false" :header="false" :footer="false" :drawerStyle="{padding:'0px'}" @cancel="onCancel" class="ant-left">
    <a-layout-sider hide-trigger collapsible :collapsed="useLayout().isCollapse" :width="250" class="ant-menu">
      <!--logo-->
      <Logo></Logo>
      <!--菜单-->
      <Menu></Menu>
    </a-layout-sider>
  </a-drawer>
  <!--PC菜单-->
  <a-layout-sider v-if="useLayout().windowWidth>768" hide-trigger collapsible :collapsed="useLayout().isCollapse" :width="250" class="ant-menu">
    <!--logo-->
    <Logo></Logo>
    <!--菜单-->
    <Menu></Menu>
  </a-layout-sider>
</template>

<style scoped>
.ant-menu {
  box-shadow: 2px 0 8px 0 rgba(29, 35, 41, .05);
  background: #232324;
  height: 100vh;
}

</style>