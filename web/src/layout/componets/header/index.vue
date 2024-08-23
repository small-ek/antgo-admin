<script setup>
import Breadcrumbs from '@/layout/componets/breadcrumbs/index.vue'
import {defineProps, ref} from "vue";
import {useLayout} from '@/stores/layout.js'

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
  useLayout().setState("isCollapse", !useLayout().isCollapse)
}
const onRefresh = () => {
  window.location.reload()
}
</script>

<template>
  <a-layout-header style="padding-left: 20px;">
    <a-row justify="space-between" align="center">
      <a-col flex="310px">
        <a-button @click="onCollapse" class="btn-icon">
          <icon-menu-unfold size="21" v-if="useLayout().isCollapse"/>
          <icon-menu-fold size="21" v-else/>
        </a-button>
        <a-button @click="onRefresh" class="btn-icon">
          <icon-refresh size="21"/>
        </a-button>
        <Breadcrumbs></Breadcrumbs>
      </a-col>
      <a-col flex="200px">

      </a-col>
      <a-col flex="200px">
        <Breadcrumbs></Breadcrumbs>
      </a-col>
<!--      <a-col :xs="2" :sm="4" :md="6" :lg="8" :xl="10" :xxl="8">-->
<!--        Col-->
<!--      </a-col>-->
    </a-row>


  </a-layout-header>
</template>

<style scoped>
.btn-icon {
  background: white;
  margin-top: 16px;
}
</style>