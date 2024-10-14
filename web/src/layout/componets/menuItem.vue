<script setup>
import { defineProps } from 'vue';
import { IconApps, IconBug, IconBulb } from '@arco-design/web-vue/es/icon';

const props = defineProps({
  item: {
    type: Object,
    required: true
  }
});
console.log(props.item)
const iconMap = {
  apps: IconApps,
  bug: IconBug,
  bulb: IconBulb
};
</script>

<template>
  <a-sub-menu v-if="item.children&&item.children.length>0" :key="item.id">
    <template #icon>
      <font-awesome-icon v-if="item.icon!==''" :icon="item.icon" size="xl"/>
      <font-awesome-icon v-if="item.icon===''" icon="fa-solid fa-table-list" />
    </template>
    <template #title>
      <span>{{ item.title }}</span>
    </template>
    <template v-for="child in item.children" :key="child.id">
      <MenuItem :item="child" />
    </template>
  </a-sub-menu>
  <a-menu-item v-else :key="item.id">
    <template #icon>
      <font-awesome-icon v-if="item.icon!==''" :icon="item.icon" size="xl"/>
      <font-awesome-icon v-if="item.icon===''" icon="fa-solid fa-table-list" />
    </template>
    {{ item.title }}</a-menu-item>
</template>