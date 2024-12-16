<script setup>
import { useLayout } from "@/stores/layout.js";
import { useNavigation } from "@/utils/base.js";
import LogoImg from "@/layout/componets/logo-img/index.vue";
import LogoText from "@/layout/componets/logo-text/index.vue";
</script>

<template>
  <a-row justify="center" class="logo" align="center" @click="useNavigation().jump('user')">
    <a-col :span="5">
      <LogoImg />
    </a-col>
    <a-col :span="12">
      <span :class="{ 'logo-text': true, 'white': useLayout().isDarkHeader }">
        <LogoText />
      </span>
    </a-col>
  </a-row>
</template>

<style scoped>
.logo {
  height: 64px;
  cursor: pointer;
}

.logo img {
  position: relative;
  top: 10px;
  width: 35px;
  height: 35px;
}

.logo-text {
  font-size: 16px;
  font-weight: bold;
}

.white {
  color: white;
}
</style>