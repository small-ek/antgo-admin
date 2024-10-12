<script setup>
import { defineProps } from 'vue';
import { IconApps, IconBug, IconBulb } from '@arco-design/web-vue/es/icon';

const props = defineProps({
  item: {
    type: Object,
    required: true
  }
});

const iconMap = {
  apps: IconApps,
  bug: IconBug,
  bulb: IconBulb
};
</script>

<template>
  <a-sub-menu v-if="item.children" :key="item.id">
    <template #icon>
      <component :is="iconMap[item.icon]" />
    </template>
    <template #title>
      <span>{{ item.title }}</span>
    </template>
    <template v-for="child in item.children" :key="child.id">
      <MenuItem :item="child" />
    </template>
  </a-sub-menu>
  <a-menu-item v-else>{{ item.title }}</a-menu-item>
</template>