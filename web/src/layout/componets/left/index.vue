<script setup>
import Menu from "@/layout/componets/menu/index.vue";
import Logo from "@/layout/componets/logo/index.vue";
import {defineProps, ref} from "vue";
import {useLayout} from '@/stores/layout.js'

const store = useLayout();
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
  <a-layout-sider hide-trigger collapsible :collapsed="store.isCollapse">
    <!--logo-->
    <Logo></Logo>
    <!--菜单-->
    <Menu></Menu>
  </a-layout-sider>
</template>

<style scoped>

</style>