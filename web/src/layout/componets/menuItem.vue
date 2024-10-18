<script setup>
import {defineProps} from 'vue';

const props = defineProps({
  item: {
    type: Object,
    required: true
  }
});
</script>

<template>
  <a-sub-menu v-if="item.children&&item.children.length>0" :key="item.id">
    <template #icon>
      <font-awesome-icon v-if="item.icon!==''" :icon="item.icon"/>
      <font-awesome-icon v-if="item.icon===''" icon="fa-solid fa-table-list"/>
    </template>
    <template #title>
      <span>{{ item.title }}</span>
    </template>
    <template v-for="child in item.children" :key="child.id">
      <MenuItem :item="child"/>
    </template>
  </a-sub-menu>
  <a-menu-item v-else :key="item.id">
    <template #icon>
      <font-awesome-icon v-if="item.icon!==''" :icon="item.icon"/>
      <font-awesome-icon v-if="item.icon===''" icon="fa-solid fa-table-list"/>
    </template>
    {{ item.title }}
  </a-menu-item>
</template>