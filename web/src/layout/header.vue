<script setup>

import {IconCaretLeft, IconCaretRight} from "@arco-design/web-vue/es/icon/index.js";
import {defineProps, ref} from "vue";

const emit = defineEmits(['onCollapse'])
const collapsed = ref(false)
const props = defineProps({
  collapsed: {
    type: Boolean,
    default: false
  },
});
collapsed.value = props.collapsed

const onCollapse = () => {
  collapsed.value = !collapsed.value
  emit("onCollapse")
}
</script>

<template>
  <a-layout-header style="padding-left: 20px;">
    <a-row class="grid-demo">
      <a-col :xs="2" :sm="4" :md="6" :lg="8" :xl="10" :xxl="8">
        <a-avatar :style="{ backgroundColor: '#14a9f8' }">Arco</a-avatar>
      </a-col>
      <a-col :xs="20" :sm="16" :md="12" :lg="8" :xl="4" :xxl="8">
        <a-button shape="round" @click="onCollapse">
          <IconCaretRight v-if="collapsed"/>
          <IconCaretLeft v-else/>
        </a-button>
      </a-col>
      <a-col :xs="2" :sm="4" :md="6" :lg="8" :xl="10" :xxl="8">
        Col
      </a-col>
    </a-row>


  </a-layout-header>
</template>

<style scoped>
.logo {
  width: 100px;
  height: 32px;
  margin: 12px 8px;
  background: rgba(255, 255, 255, 0.2);
}
</style>