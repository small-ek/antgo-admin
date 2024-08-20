<template>
  <a-layout class="ant-layout">
    <!--顶部-->
    <Header :collapsed="collapsed" @onCollapse="onCollapse"></Header>
    <a-layout>
      <!--左侧-->
      <Left :collapsed="collapsed"></Left>
      <a-layout style="padding: 0 24px;">
        <!--面包屑-->
        <Breadcrumbs></Breadcrumbs>
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
import {Message} from '@arco-design/web-vue';
import Left from "@/layout/left.vue";
import Header from "@/layout/header.vue";
import Breadcrumbs from "@/layout/breadcrumbs.vue";
import Footer from "@/layout/footer.vue";

const collapsed = ref(true);
const onCollapse = () => {
  collapsed.value = !collapsed.value;
};
const onClickMenuItem = (key) => {
  Message.info({content: `You select ${key}`, showIcon: true});
}
</script>
<style scoped>
.ant-layout {
  height: 100vh;
  background: var(--color-fill-2);
  border: 1px solid var(--color-border);
}

.ant-layout :deep(.arco-layout-sider-light) .logo {
  background: var(--color-fill-2);
}

.ant-layout :deep(.arco-layout-header) {
  height: 64px;
  line-height: 64px;
  background: var(--color-bg-3);
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