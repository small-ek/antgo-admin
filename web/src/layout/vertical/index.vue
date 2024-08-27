<template>
  <a-layout class="ant-layout">
    <!--左侧-->
    <Left :collapsed="collapsed"></Left>
    <a-layout>
      <!--头部-->
      <Header :collapsed="collapsed" @onCollapse="onCollapse"></Header>
      <Tags></Tags>
      <a-layout style="padding: 0 24px;">

        <a-layout-content>
          <router-view></router-view>
        </a-layout-content>
        <!--底部-->
        <Footer></Footer>
      </a-layout>
    </a-layout>
  </a-layout>
</template>
<script setup>
import {ref} from 'vue';
import Left from "@/layout/componets/left/index.vue";
import Header from "@/layout/componets/header/index.vue";
import Footer from "@/layout/componets/fooder/index.vue";
import Tags from "@/layout/componets/tabs/index.vue"

const collapsed = ref(true);
const onCollapse = () => {
  collapsed.value = !collapsed.value;
};

</script>
<style scoped>
.ant-layout {
  height: 100vh;
  background: var(--color-fill-2);
}

.ant-layout :deep(.arco-layout-header) {
  height: 64px;
  line-height: 64px;
  background: var(--color-bg-3);
  padding-left: 5px !important;
}

.ant-layout :deep(.arco-layout-footer) {
  height: 48px;
  color: var(--color-text-2);
  font-weight: 400;
  font-size: 14px;
  line-height: 48px;
}

.ant-layout :deep(.arco-layout-content) {
  color: var(--color-text-2);
  font-weight: 400;
  font-size: 14px;
  background: var(--color-bg-3);
}

.ant-layout :deep(.arco-layout-footer),
.ant-layout :deep(.arco-layout-content) {
  display: flex;
  flex-direction: column;
  justify-content: center;
  color: black;
  font-stretch: condensed;
  text-align: center;
}

</style>