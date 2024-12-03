<template>
  <div class="ant-layout">
    <!--头部-->
    <Header/>
    <a-layout class="ant-container">

      <!--左侧-->
      <Left/>
      <a-layout>
        <a-spin :loading="loading" tip="加载中...">
          <a-scrollbar class="scrollbar" style="height: calc(100vh - 69px);overflow: auto;"
                       @scroll="useTheme().handleScroll">
            <!--头部标签-->
            <div
                :class="{'affix': useLayout().isFixedHeader, 'affix-hidden': !useLayout().isFixedHeader&&useLayout().header==='adaptive'}">
              <Tags/>
            </div>
            <!--内容-->
            <div class="content-layout">
              <a-layout-content>
                <router-view></router-view>
              </a-layout-content>
            </div>

            <!--底部-->
            <Footer v-if="useLayout().isFooter"></Footer>
            <a-back-top target-container=".scrollbar" :style="{position:'fixed'}"/>
          </a-scrollbar>
        </a-spin>
      </a-layout>

    </a-layout>
    <!--布局设置-->
    <Setting></Setting>
  </div>
</template>
<script setup>
import Left from "@/layout/componets/left/index.vue";
import Header from "@/layout/componets/header/index.vue";
import Footer from "@/layout/componets/fooder/index.vue";
import Setting from "@/layout/componets/setting/index.vue"
import Tags from "@/layout/componets/tabs/index.vue"
import {useLayout} from "@/stores/layout.js";
import {useTheme} from "@/utils/theme.js";
import {loading} from "@/utils/loading.js";
</script>
<style scoped lang="less">
@import "../layout.less";

.ant-container {
  margin-top: -1px;
}

.ant-layout :deep(.arco-layout-header) {
  top: -1px;
}

.ant-container {
  height: calc(100vh - 64px);
}

</style>