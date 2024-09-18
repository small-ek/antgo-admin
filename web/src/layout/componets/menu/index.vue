<script setup>
import {Message} from "@arco-design/web-vue";
import {useLayout} from "@/stores/layout.js";
import {defineProps} from "vue";

const onClickMenuItem = (key) => {
  // Message.info({content: `You select ${key}`, showIcon: true});
}
const props = defineProps({
  mode: {
    type: String,
    default: 'vertical'
  }
});
const onCollapse = (val, type) => {
  const content = type === 'responsive' ? '触发响应式收缩' : '点击触发收缩';
  if (useLayout().layout === 'vertical' || useLayout().layout === 'classic') {
    useLayout().setState("isCollapsed", val)
  }
}
</script>

<template>
  <a-menu
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
    <a-menu-item key="0_1" disabled>
      <IconHome/>
      Menu 1
    </a-menu-item>
    <a-menu-item key="0_2">
      <IconCalendar/>
      Menu 2
    </a-menu-item>
    <a-menu-item key="0_3">
      <IconCalendar/>
      Menu 3
    </a-menu-item>
    <a-sub-menu key="1">
      <template #title>
        <span><IconCalendar/>Navigation 1</span>
      </template>
      <a-menu-item key="1_1">Menu 1</a-menu-item>
      <a-menu-item key="1_2">Menu 2</a-menu-item>
      <a-sub-menu key="2" title="Navigation 2">
        <a-menu-item key="2_1">Menu 1</a-menu-item>
        <a-menu-item key="2_2">Menu 2</a-menu-item>
      </a-sub-menu>
      <a-sub-menu key="3" title="Navigation 3">
        <a-menu-item key="3_1">Menu 1</a-menu-item>
        <a-menu-item key="3_2">Menu 2</a-menu-item>
        <a-menu-item key="3_3">Menu 3</a-menu-item>
      </a-sub-menu>
    </a-sub-menu>
    <a-sub-menu key="4">
      <template #title>
        <span><IconCalendar/>Navigation 4</span>
      </template>
      <a-menu-item key="4_1">Menu 1</a-menu-item>
      <a-menu-item key="4_2">Menu 2</a-menu-item>
      <a-menu-item key="4_3">Menu 3</a-menu-item>
    </a-sub-menu>
  </a-menu>
</template>

<style scoped>
.arco-menu-dark .arco-menu-item.arco-menu-selected{
  background: rgb(var(--primary-6));
  border-radius: 10px;

}
.arco-menu-light .arco-menu-item.arco-menu-selected{
  background: rgb(var(--primary-6));
  border-radius: 10px;
  color: white;
}


.arco-menu-item:hover{
  background: rgb(var(--primary-1));
  border-radius: 10px;
}
</style>