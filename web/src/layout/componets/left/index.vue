<script setup>
import Menu from "@/layout/componets/menu/index.vue";
import Logo from "@/layout/componets/logo/index.vue";
import {computed, defineProps, onMounted, onUnmounted, ref} from "vue";
import {useLayout} from '@/stores/layout.js'

const collapsed = ref(false)
const mobileVisible = ref(false)
const props = defineProps({
  collapsed: {
    type: Boolean,
    default: false
  },
});
collapsed.value = props.collapsed

const windowWidth = ref(window.innerWidth);

const handleResize = () => {
  windowWidth.value = window.innerWidth;
  if (windowWidth.value < 768) {
    useLayout().setState("isCollapse", false)
    mobileVisible.value = true
  } else {
    mobileVisible.value = false
  }
};

onMounted(() => {
  window.addEventListener('resize', handleResize);
});

onUnmounted(() => {
  window.removeEventListener('resize', handleResize);
});


const onCancel = () => {
  mobileVisible.value = false;
}
</script>

<template>
  <a-drawer v-if="windowWidth<768" :width="250" :visible="mobileVisible" placement="left" :closable="false" :header="false" :footer="false" :drawerStyle="{padding:'0px'}" @cancel="onCancel">
    <a-layout-sider hide-trigger collapsible :collapsed="useLayout().isCollapse" :width="250" class="ant-left">
      <!--logo-->
      <Logo></Logo>
      <!--菜单-->
      <Menu></Menu>
    </a-layout-sider>
  </a-drawer>

  <a-layout-sider v-if="windowWidth>768" hide-trigger collapsible :collapsed="useLayout().isCollapse" :width="250" class="ant-left">
    <!--logo-->
    <Logo></Logo>
    <!--菜单-->
    <Menu></Menu>
  </a-layout-sider>
</template>

<style scoped>
.ant-left {
  box-shadow: 2px 0 8px 0 rgba(29, 35, 41, .05);
  background: #232324;
  height: 100vh;
}
</style>