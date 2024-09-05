<script setup>
import logo from '@/assets/img/logo.png'
import {useLayout} from "@/stores/layout.js";
import {useNavigation} from "@/utils/base.js";
</script>

<template>
  <a-row justify="center" class="logo-container" align="center" v-if="!useLayout().isCollapsed" @click="useNavigation().jump('user')">
    <a-col :span="5">
      <img :src="logo" alt=""/>
    </a-col>
    <a-col :span="12">
      <span :class="['logo-text',{'white':useLayout().isDarkSidebar===true}]">Antgo admin</span>
    </a-col>
  </a-row>
  <a-row justify="center" class="logo-container" align="center" v-if="useLayout().isCollapsed">
    <a-col :span="20">
      <img :src="logo" alt=""/>
    </a-col>

  </a-row>
</template>

<style scoped>
.logo-container {
  height: 64px;
  cursor: pointer;

  img {
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
}
</style>