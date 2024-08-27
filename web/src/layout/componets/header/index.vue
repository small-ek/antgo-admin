<script setup>
import Breadcrumbs from '@/layout/componets/breadcrumbs/index.vue'
import {defineProps, ref} from "vue";
import {useLayout} from '@/stores/layout.js'

const emit = defineEmits(['onCollapse'])
const collapsed = ref(false)
const isFullscreen = ref(false)
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

//刷新
const onRefresh = () => {
  window.location.reload()
}

//全屏
const onFullscreen = () => {
  if (!isFullscreen.value) {
    const element = document.documentElement;
    if (element.requestFullscreen) {
      element.requestFullscreen();
    } else if (element.mozRequestFullScreen) { // Firefox
      element.mozRequestFullScreen();
    } else if (element.webkitRequestFullscreen) { // Chrome, Safari and Opera
      element.webkitRequestFullscreen();
    } else if (element.msRequestFullscreen) { // IE/Edge
      element.msRequestFullscreen();
    }
  } else {
    if (document.exitFullscreen) {
      document.exitFullscreen();
    } else if (document.mozCancelFullScreen) { // Firefox
      document.mozCancelFullScreen();
    } else if (document.webkitExitFullscreen) { // Chrome, Safari and Opera
      document.webkitExitFullscreen();
    } else if (document.msExitFullscreen) { // IE/Edge
      document.msExitFullscreen();
    }
  }
  isFullscreen.value = !isFullscreen.value
}
</script>

<template>
  <a-layout-header style="padding-left: 20px;">
    <a-row justify="space-between" align="center" :wrap="false">
      <a-col flex="320px">
        <a-tooltip content="收缩菜单">
          <a-button @click="onCollapse" class="btn-icon shadow" shape="circle">
            <icon-menu-unfold size="19" v-if="useLayout().isCollapse"/>
            <icon-menu-fold size="19" v-else/>
          </a-button>
        </a-tooltip>

        <a-tooltip content="刷新页面">
          <a-button @click="onRefresh" class="btn-icon shadow" shape="circle">
            <icon-refresh size="19"/>
          </a-button>
        </a-tooltip>
        <Breadcrumbs></Breadcrumbs>
      </a-col>

      <a-col flex="320px">
        <a-tooltip content="全屏开关">
          <a-button @click="onFullscreen" class="btn-icon shadow" shape="circle">
            <icon-fullscreen size="19" v-if="!isFullscreen"/>
            <icon-fullscreen-exit size="19" v-else/>
          </a-button>
        </a-tooltip>
        <a-tooltip content="搜索菜单内容">
          <a-button  class="btn-icon shadow" shape="circle">
            <icon-search size="19"/>
          </a-button>
        </a-tooltip>
        <a-dropdown trigger="hover">
          <a-button class="btn-icon shadow" shape="circle">
            <icon-language size="19"/>
          </a-button>
          <template #content>
            <a-doption>简体中文</a-doption>
            <a-doption>English</a-doption>
          </template>
        </a-dropdown>
        <a-tooltip content="设置">
          <a-button class="btn-icon shadow" shape="circle">
            <icon-settings size="19"/>
          </a-button>
        </a-tooltip>

        <span class="head-row">
          <a-avatar class="head-img">
          <img alt="avatar" src="https://p1-arco.byteimg.com/tos-cn-i-uwbnlip3yd/3ee5f13fb09879ecb5185e440cef6eb9.png~tplv-uwbnlip3yd-webp.webp"/>
        </a-avatar>
        <a-dropdown :popup-max-height="false" trigger="hover">
          <span class="head-text">Antgo Admin</span>
          <template #content>
            <a-doption>信息设置</a-doption>
            <a-doption>退出登录</a-doption>
          </template>
        </a-dropdown>
        </span>
      </a-col>
    </a-row>


  </a-layout-header>
</template>

<style scoped>
.btn-icon {
  background: white;
  margin-top: 16px;
  margin-left: 7px;
}

.head-img {
  width: 30px;
  height: 30px;
  margin-left: 10px;
}

.head-text {
  margin-left: 10px;
  cursor: pointer;
}

.head-row {
  position: relative;
  top: -5px
}
</style>