<template>
  <a-layout class="ant-layout">
    <Left />
    <a-layout class="ant-container">
      <a-scrollbar class="scrollbar" @scroll="useTheme().handleScroll">
        <div :class="{'affix': useLayout().isFixedHeader, 'affix-hidden': !useLayout().isFixedHeader && useLayout().header === 'adaptive'}">
          <Header />
          <Tags />
        </div>
        <a-layout class="content-layout">
          <a-layout-content>
            <router-view />
          </a-layout-content>
          <Footer v-if="useLayout().isFooter" />
          <a-back-top target-container=".scrollbar" :style="{position:'fixed'}" />
        </a-layout>

      </a-scrollbar>
    </a-layout>
    <Setting />
  </a-layout>
</template>

<script setup>
import Left from '@/layout/componets/left/index.vue';
import Header from '@/layout/componets/header/index.vue';
import Footer from '@/layout/componets/fooder/index.vue';
import Tags from '@/layout/componets/tabs/index.vue';
import Setting from '@/layout/componets/setting/index.vue';
import { useLayout } from '@/stores/layout.js';
import { useTheme } from '@/utils/theme.js';
</script>

<style scoped lang="less">
@import "../layout.less";
//.container {
//  height: 100vh;
//  overflow: auto;
//}
</style>