<script setup>
import {defineProps, ref} from 'vue';
import {useLayout} from "@/stores/layout.js";
import MenuItem from "@/layout/componets/menuItem/index.vue";

const list=ref([])

//点击菜单跳转
const onClickMenuItem = (key) => {
  // Message.info({content: `You select ${key}`, showIcon: true});

  const item = list.value.find(item => item.id === key);
  if (item && item.path) {
    router.push({ path: item.path });
  } else {
    console.error(`No route found for key: ${key}`);
  }

}

const props = defineProps({
  mode: {
    type: String,
    default: 'vertical'
  }
});

const setList = (val) => {
  list.value = val
  console.log(list.value)
}
const onCollapse = (val, type) => {
  const content = type === 'responsive' ? '触发响应式收缩' : '点击触发收缩';
  if (useLayout().layout === 'vertical' || useLayout().layout === 'classic') {
    useLayout().setState("isCollapsed", val)
  }
}

defineExpose({
  setList
})
</script>

<template>
  <a-menu
      v-if="list.length>0"
      :theme="useLayout().isDarkSidebar?'dark':'light'"
      :accordion="useLayout().isAccordion"
      :defaultOpenKeys="['1']"
      :defaultSelectedKeys="['0_3']"
      :style="{ width: '100%' }"
      breakpoint="lg"
      :collapsed="useLayout().isCollapsed"
      @menuItemClick="onClickMenuItem"
      @collapse="onCollapse"
      :mode="props.mode"
  >
    <template v-for="item in list" :key="item.id">
      <MenuItem :item="item" />
    </template>
  </a-menu>
</template>

<style>
.arco-menu-dark .arco-menu-item.arco-menu-selected{
  background: rgb(var(--primary-6));
  border-radius: 10px;
}

.arco-menu-light .arco-menu-item.arco-menu-selected{
  background: rgb(var(--primary-6));
  border-radius: 15px;
  color: white;
}

.arco-menu-item:hover{
  background: rgb(var(--primary-1));
  transform: scale(1.01);
  transition: transform 0.8s;
}

.arco-menu-dark .arco-menu-item .arco-icon, .arco-menu-dark .arco-menu-group-title .arco-icon, .arco-menu-dark .arco-menu-pop-header .arco-icon, .arco-menu-dark .arco-menu-inline-header .arco-icon, .arco-menu-dark .arco-menu-item .arco-menu-icon, .arco-menu-dark .arco-menu-group-title .arco-menu-icon, .arco-menu-dark .arco-menu-pop-header .arco-menu-icon, .arco-menu-dark .arco-menu-inline-header .arco-menu-icon:hover{
  color: white;
}

.arco-menu-light .arco-menu-item.arco-menu-selected .arco-menu-icon{
  color: white !important;
}
</style>