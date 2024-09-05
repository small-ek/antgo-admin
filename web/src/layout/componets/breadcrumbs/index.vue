<script setup>

import {useLayout} from "@/stores/layout.js";
</script>

<template>
  <a-breadcrumb class="ant-breadcrumb" v-if="useLayout().isBreadcrumb">
    <a-breadcrumb-item>
      <span :class="['ml-8',{'dark':useLayout().isDarkHeader===true}]">首页</span></a-breadcrumb-item>
    <a-breadcrumb-item><span :class="['ml-8',{'dark':useLayout().isDarkHeader===true}]">首页首页</span></a-breadcrumb-item>
    <a-breadcrumb-item><span :class="['ml-8',{'dark':useLayout().isDarkHeader===true}]">首页首页</span></a-breadcrumb-item>
  </a-breadcrumb>
</template>

<style scoped>
.ant-breadcrumb {
  position: relative;
  top: -6px;
  color: white;

  .dark {
    color: white;
  }

  .ml-8 {
    margin-left: 8px;
  }
}


</style>