<script setup>

import {useLayout} from "@/stores/layout.js";
</script>

<template>
  <a-drawer :width="250" :visible="useLayout().isSetting" title="布局设置" placement="right" :closable="false"  :footer="false" :drawerStyle="{padding:'2%'}" @cancel="useLayout().setState('isSetting', false)">
    asaa
  </a-drawer>
</template>

<style scoped>

</style>