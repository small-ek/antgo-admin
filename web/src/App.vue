<template>
  <div>
    <a-config-provider :locale="zhCN">
      <router-view></router-view>
    </a-config-provider>
  </div>
</template>

<script setup>
import zhCN from '@arco-design/web-vue/es/locale/lang/zh-CN';
</script>
<style scoped>

</style>
