<template>
  <a-layout class="ant-layout">
    <!--左侧-->
    <Left/>
    <a-layout class="ant-container">
      <a-scrollbar class="scrollbar" @scroll="useTheme().handleScroll">
        <!--头部-->
        <div :class="{'affix': useLayout().isFixedHeader, 'affix-hidden': !useLayout().isFixedHeader&&useLayout().header==='adaptive'}">
          <Header/>
          <Tags/>
        </div>
        <!--内容-->
        <a-layout class="content-layout">
          <a-layout-content>
            <router-view></router-view>
          </a-layout-content>

          <!--底部-->
          <Footer v-if="useLayout().isFooter"></Footer>
        </a-layout>
      </a-scrollbar>
    </a-layout>

    <!--布局设置-->
    <Setting></Setting>
  </a-layout>
</template>
<script setup>
import Left from "@/layout/componets/left/index.vue";
import Header from "@/layout/componets/header/index.vue";
import Footer from "@/layout/componets/fooder/index.vue";
import Tags from "@/layout/componets/tabs/index.vue"
import Setting from "@/layout/componets/setting/index.vue"
import {useLayout} from "@/stores/layout.js";
import {useTheme} from "@/utils/theme.js";
</script>
<style scoped lang="less">
@import "../layout.less";
</style>