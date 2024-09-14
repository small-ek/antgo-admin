<script setup>
import Breadcrumbs from '@/layout/componets/breadcrumbs/index.vue'
import {ref} from "vue";
import {useLayout} from '@/stores/layout.js'
import headerLogo from "@/layout/componets/headerLogo/index.vue";

const isFullscreen = ref(false)

const onCollapse = () => {
  if (useLayout().windowWidth < 768) {
    useLayout().setState("isCollapsed", false)
    useLayout().setState("showMobileMenu", !useLayout().showMobileMenu)
  } else {
    useLayout().setState("isCollapsed", !useLayout().isCollapsed)
  }
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
  <a-layout-header :class="[{'dark':useLayout().isDarkHeader===true}]">
    <a-row justify="space-between" align="center" :wrap="false">
      <a-col :flex="useLayout().sidebarWidth+'px'" v-if="useLayout().layout==='classic'">
        <headerLogo></headerLogo>
      </a-col>
      <a-col flex="auto">
        <a-tooltip content="收缩菜单">
          <a-button @click="onCollapse" class="btn-icon shadow first-icon" shape="circle">
            <icon-menu-unfold size="19" v-if="useLayout().isCollapse"/>
            <icon-menu-fold size="19" v-else/>
          </a-button>
        </a-tooltip>

        <a-tooltip content="刷新页面" v-if="useLayout().isRefresh">
          <a-button @click="onRefresh" class="btn-icon shadow" shape="circle">
            <icon-refresh size="19"/>
          </a-button>
        </a-tooltip>
        <Breadcrumbs v-if="useLayout().windowWidth>768"></Breadcrumbs>
      </a-col>

      <a-col :flex="useLayout().windowWidth>768?'320px':'250px'">
        <template v-if="useLayout().isFullScreen">
          <a-tooltip content="全屏开关" v-if="useLayout().windowWidth>768">
            <a-button @click="onFullscreen" class="btn-icon shadow" shape="circle">
              <icon-fullscreen size="19" v-if="!isFullscreen"/>
              <icon-fullscreen-exit size="19" v-else/>
            </a-button>
          </a-tooltip>
        </template>

        <a-tooltip content="搜索菜单内容" v-if="useLayout().isSearch">
          <a-button class="btn-icon shadow" shape="circle">
            <icon-search size="19"/>
          </a-button>
        </a-tooltip>
        <a-dropdown trigger="hover" v-if="useLayout().isLanguage">
          <a-button class="btn-icon shadow" shape="circle">
            <icon-language size="19"/>
          </a-button>
          <template #content>
            <a-doption>简体中文</a-doption>
            <a-doption>English</a-doption>
          </template>
        </a-dropdown>
        <template v-if="useLayout().setting">
          <a-tooltip content="设置" v-if="useLayout().windowWidth>768">
            <a-button class="btn-icon shadow" shape="circle" @click="useLayout().setState('showSetting',true)">
              <icon-settings size="19"/>
            </a-button>
          </a-tooltip>
        </template>


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
.first-icon {
  margin-left: 1vw;
}

.dark {
  background: black !important;
  color: white !important;

  .btn-icon {
    background: #313132;
    color: white;
  }
}

.btn-icon {
  background: var(--color-bg-4);
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