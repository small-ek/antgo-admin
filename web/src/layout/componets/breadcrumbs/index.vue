<script setup>
import {useRoute} from 'vue-router';
import {useMenu} from '@/stores/menu.js';
import {computed} from 'vue';
import {useLayout} from "@/stores/layout.js";

const route = useRoute();
const menuTree = useMenu().menuTree;
const layout = useLayout();

const matchedTitles = computed(() => {
  const titles = [];
  const findTitles = (nodes, path) => {
    for (const node of nodes) {
      if ("/" + node.path === path) {
        titles.push(node.title);
        return true;
      }
      if (node.children && findTitles(node.children, path)) {
        titles.push(node.title);
        return true;
      }
    }
    return false;
  };
  findTitles(menuTree, route.path);

  titles.length === 0 ? titles.push(route.meta.title) : "";
  return titles.reverse();
});
</script>

<template>
  <a-breadcrumb v-if="layout.isBreadcrumb" class="breadcrumb">
    <a-breadcrumb-item v-for="row in matchedTitles" :key="row">
      <span :class="{ 'ml-8': true, 'dark': layout.isDarkHeader }">{{ row }}</span>
    </a-breadcrumb-item>
  </a-breadcrumb>
</template>

<style scoped>
.breadcrumb {
  position: relative;
  top: -6px;
  color: white;
}

.dark {
  color: white;
}

.ml-8 {
  margin-left: 5px;
}
</style>