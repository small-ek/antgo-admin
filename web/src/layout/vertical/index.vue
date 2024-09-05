<template>
  <a-layout class="ant-layout">
    <!--左侧-->
    <Left></Left>
    <a-layout>
      <!--头部-->
      <a-affix :offsetTop="1">
        <Header></Header>
      </a-affix>
      <a-affix :offsetTop="65">
        <Tags></Tags>
      </a-affix>
      <a-layout style="padding: 0 0.6vw;">

        <a-layout-content>
          <router-view></router-view>
        </a-layout-content>
        <!--底部-->
        <Footer></Footer>
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
import Tags from "@/layout/componets/tabs/index.vue"
import Setting from "@/layout/componets/setting/index.vue"
</script>
<style scoped>
.ant-layout {
  height: 100vh;
  background: var(--color-neutral-1);
}

.ant-layout :deep(.arco-layout-header) {
  height: 64px;
  line-height: 64px;
  padding-left: 20px;
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

</style>