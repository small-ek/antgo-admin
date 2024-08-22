<script setup>
import logo from '@/assets/img/logo.png'
import {useLayout} from "@/stores/layout.js";
</script>

<template>
  <a-row justify="center" class="logo-container" align="center" v-if="!useLayout().isCollapse">
    <a-col :span="6">
      <img :src="logo" alt=""/>
    </a-col>
    <a-col :span="15">
      <span class="logo-text">Antgo admin</span>
    </a-col>
  </a-row>
  <a-row justify="center" class="logo-container" align="center" v-if="useLayout().isCollapse">
    <a-col :span="20">
      <img :src="logo" alt=""/>
    </a-col>

  </a-row>
</template>

<style scoped>
.logo-container{
  height: 64px;
  img{
    width: 35px;
    height: 35px;
  }
  .logo-text{
    font-size: 16px;
    font-weight: bold;
  }
}
</style>