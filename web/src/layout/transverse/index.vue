<template>
  <a-layout class="ant-layout">
    <!--头部-->

      <a-layout class="ant-container">

        <!--左侧-->
        <Left/>
        <a-layout>
          <a-scrollbar class="scrollbar" @scroll="useTheme().handleScroll">
          <!--头部标签-->
          <div :class="{'affix': useLayout().isFixedHeader, 'affix-hidden': !useLayout().isFixedHeader&&useLayout().header==='adaptive'}">
            <Header/>
            <Tags/>
          </div>

          <!--内容-->
          <a-layout class="content-layout">
            <a-layout-content>

              <router-view></router-view>

            </a-layout-content>
          </a-layout>

          <!--底部-->
          <Footer v-if="useLayout().isFooter"></Footer>
          </a-scrollbar>
        </a-layout>

      </a-layout>

    <!--布局设置-->
    <Setting></Setting>
  </a-layout>
</template>
<script setup>
import Left from "@/layout/componets/left/index.vue";
import Header from "@/layout/componets/header/index.vue";
import Footer from "@/layout/componets/fooder/index.vue";
import Setting from "@/layout/componets/setting/index.vue"
import Tags from "@/layout/componets/tabs/index.vue"
import {useLayout} from "@/stores/layout.js";
import {useTheme} from "@/utils/theme.js";
</script>
<style scoped>
.ant-container {
  margin-top: -1px;
}

.ant-layout :deep(.arco-layout-header) {
  height: 64px;
  line-height: 64px;
  position: relative;
  top: -1px;
  background: var(--color-bg-4);
  color: var(--color-text-1);
}

.ant-layout :deep(.arco-layout-footer) {
  height: 30px;
  font-weight: 400;
  font-size: 13px;
  line-height: 30px;
  color: var(--color-text-3);
}

.ant-layout :deep(.arco-layout-content) {
  font-weight: 400;
  font-size: 14px;
  background: var(--color-bg-4);
  color: var(--color-text-1);
}

.ant-layout :deep(.arco-layout-footer),
.ant-layout :deep(.arco-layout-content) {
  display: flex;
  flex-direction: column;
  justify-content: center;
  font-stretch: condensed;
  text-align: center;
}

.ant-container {
  height: 100vh;
}

</style>