<script setup>
import { useLayout } from "@/stores/layout.js";
</script>

<template>
  <a-breadcrumb v-if="useLayout().isBreadcrumb" class="breadcrumb">
    <a-breadcrumb-item>
      <span :class="{ 'ml-8': true, 'dark': useLayout().isDarkHeader }">首页</span>
    </a-breadcrumb-item>
    <a-breadcrumb-item>
      <span :class="{ 'ml-8': true, 'dark': useLayout().isDarkHeader }">首页首页</span>
    </a-breadcrumb-item>
    <a-breadcrumb-item>
      <span :class="{ 'ml-8': true, 'dark': useLayout().isDarkHeader }">首页首页</span>
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
  margin-left: 8px;
}
</style>