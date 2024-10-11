<script setup>
import {useLayout} from "@/stores/layout.js";
import {useNavigation} from "@/utils/base.js";
import logoText from "@/layout/componets/logoText/index.vue"
import logoImg from "@/layout/componets/logoImg/index.vue"
</script>

<template>
  <a-row justify="center" class="logo-container" align="center" v-if="!useLayout().isCollapsed" @click="useNavigation().jump('user')">
    <a-col :span="5" v-if="useLayout().layout!=='columns'">
      <logoImg></logoImg>
    </a-col>
    <a-col :span="12">
      <span :class="['logo-text',{'white':useLayout().isDarkSidebar===true}]">
        <logoText></logoText>
      </span>
    </a-col>
  </a-row>
  <a-row justify="center" class="logo-container" align="center" v-if="useLayout().isCollapsed">
    <a-col :span="20">
      <logoImg></logoImg>
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