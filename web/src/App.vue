<template>
  <div>
    <a-config-provider :locale="zhCN">
      <router-view></router-view>
    </a-config-provider>
  </div>
</template>

<script setup>
import zhCN from '@arco-design/web-vue/es/locale/lang/zh-CN';
import {useTheme} from "@/utils/theme.js";

useTheme().defaultTheme()
</script>
<style scoped>

</style>
