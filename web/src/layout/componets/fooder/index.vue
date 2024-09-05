<script setup>

import {useLayout} from "@/stores/layout.js";
</script>

<template>
  <a-layout-footer>{{ useLayout().footerText }}</a-layout-footer>
</template>

<style scoped>

</style>