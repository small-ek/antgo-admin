<script setup>
//纵向布局
import Vertical from '@/layout/vertical/index.vue'
//经典布局
import Classic from '@/layout/classic/index.vue'
//横向布局
import Transverse from '@/layout/transverse/index.vue'

//布局组件
const Layout = {
  Vertical,
  Classic,
  Transverse
}
</script>

<template>
  <component :is="Layout['Vertical']"/>
</template>

<style scoped>

</style>