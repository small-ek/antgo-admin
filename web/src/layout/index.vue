<script setup>
//纵向布局
import vertical from '@/layout/vertical/index.vue'
//经典布局
import classic from '@/layout/classic/index.vue'
//横向布局
import transverse from '@/layout/transverse/index.vue'
//分栏布局
import columns from '@/layout/columns/index.vue'
import {useLayout} from "@/stores/layout.js";
console.log(useLayout().layout)
//布局组件
const Layout = {
  vertical,
  classic,
  transverse,
  columns
}
</script>

<template>
  <component :is="Layout[useLayout().layout]"/>
</template>

<style scoped>

</style>