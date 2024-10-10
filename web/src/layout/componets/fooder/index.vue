<script setup>

import {useLayout} from "@/stores/layout.js";
</script>

<template>
  <a-layout-footer style="background: white">{{ useLayout().footerText }}</a-layout-footer>
</template>

<style scoped>

</style>