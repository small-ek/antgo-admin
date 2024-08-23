<script setup>
import Menu from "@/layout/componets/menu/index.vue";
import Logo from "@/layout/componets/logo/index.vue";
import {defineProps, ref} from "vue";
import {useLayout} from '@/stores/layout.js'

const collapsed = ref(false)
const props = defineProps({
  collapsed: {
    type: Boolean,
    default: false
  },
});
collapsed.value = props.collapsed


</script>

<template>
  <a-layout-sider hide-trigger collapsible :collapsed="useLayout().isCollapse" :width="240" class="ant-left">
    <!--logo-->
    <Logo></Logo>
    <!--菜单-->
    <Menu></Menu>
  </a-layout-sider>
</template>

<style scoped>
.ant-left{
  box-shadow: 2px 0 8px 0 rgba(29,35,41,.05)
}
</style>