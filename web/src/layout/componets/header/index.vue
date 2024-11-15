<script setup>
import Breadcrumbs from '@/layout/componets/breadcrumbs/index.vue'
import {onMounted, onUnmounted, ref} from "vue";
import {useLayout} from '@/stores/layout.js'
import headerLogo from "@/layout/componets/headerLogo/index.vue";
import {useTheme} from "@/utils/theme.js";
import Menu from "@/layout/componets/menu/index.vue";
import {useMenu} from "@/stores/menu.js";
import {useRouter} from "vue-router";

const router = useRouter()
const isFullscreen = ref(false)
const menuRef = ref()
const isSearch = ref(false);
const searchList = ref(useMenu().subMenu)

onMounted(async () => {
  if (useLayout().layout === 'transverse') {
    menuRef.value.setList(useMenu().menuTree)
  }
})
//收缩菜单
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

const onShowSearch = () => {
  isSearch.value = true
}

const onInputSeach = (e) => {
  searchList.value = useMenu().subMenu.filter(item => {
    return item.title.indexOf(e) !== -1 || item.pinyin.indexOf(e) !== -1
  })
}
const jump = (row) => {
  router.push({path: "/" + row.path});
  isSearch.value = false
}
//控制屏幕宽度
onMounted(() => {
  window.addEventListener('resize', useTheme().handleResize);
});

onUnmounted(() => {
  window.removeEventListener('resize', useTheme().handleResize);
});

</script>

<template>
  <a-layout-header :class="[{'dark':useLayout().isDarkHeader===true}]">
    <a-row justify="space-between" align="center" :wrap="false">
      <a-col :flex="useLayout().sidebarWidth+'px'"
             v-if="useLayout().layout!=='columns'&&useLayout().layout!=='vertical'&&useLayout().windowWidth>992">
        <headerLogo></headerLogo>
      </a-col>
      <a-col :flex="useLayout().layout==='transverse'?'80px':'auto'">
        <a-tooltip content="收缩菜单" v-if="useLayout().layout==='vertical'||useLayout().layout==='classic'">
          <a-button @click="onCollapse" class="btn-icon shadow first-icon" shape="circle">
            <icon-menu-unfold size="19" v-if="useLayout().isCollapse"/>
            <icon-menu-fold size="19" v-else/>
          </a-button>
        </a-tooltip>
        <a-tooltip content="收缩菜单"
                   v-if="(useLayout().layout==='transverse'||useLayout().layout==='columns')&&useLayout().windowWidth<765">
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
        <template
            v-if="useLayout().layout==='vertical'||useLayout().layout==='classic'||useLayout().layout==='columns'">
          <Breadcrumbs v-if="useLayout().windowWidth>768"></Breadcrumbs>
        </template>
      </a-col>
      <a-col flex="auto" v-show="useLayout().layout==='transverse'&&useLayout().windowWidth>765">
        <div style="width: 100%">
          <Menu ref="menuRef" mode="horizontal"></Menu>
        </div>
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
          <a-button class="btn-icon shadow" shape="circle" @click="onShowSearch">
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
          <img alt="avatar"
               src="https://p1-arco.byteimg.com/tos-cn-i-uwbnlip3yd/3ee5f13fb09879ecb5185e440cef6eb9.png~tplv-uwbnlip3yd-webp.webp"/>
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
  <!--搜索-->
  <a-modal v-model:visible="isSearch" draggable :footer="false" ok-text="关闭"
           :width="useLayout().windowWidth>768?'500px':'95%'">
    <template #title>
      <h4>菜单搜索</h4>
    </template>
    <a-space direction="vertical" :size="16" style="display: block;">
      <a-row class="grid-demo">
        <a-col :span="24">
          <div>
            <a-input-search placeholder="请输入需要跳转的关键字或者拼音" @input="onInputSeach"/>
          </div>
        </a-col>

        <a-col :span="24" class="mt-10">
          <a-scrollbar style="height:400px;overflow-x: auto;">
            <a-typography-title :heading="6">
              跳转
            </a-typography-title>
            <div v-for="row in searchList">
              <div class="hand search-link-list m-auto" @click="jump(row)">
                <font-awesome-icon v-if="row.icon!==''" :icon="row.icon"/>
                <font-awesome-icon v-if="row.icon===''" icon="fa-solid fa-table-list"/>
                <span class="search-link-text">{{ row.title }}</span>
              </div>
            </div>
          </a-scrollbar>
        </a-col>
      </a-row>
    </a-space>
  </a-modal>
</template>

<style scoped>
.first-icon {
  margin-left: 1vw;
}

.dark {
  background: var(--color-menu-dark-bg) !important;
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

.search-link-list {
  width: 400px;
  height: 40px;
  line-height: 40px;
  font-weight: bold;
  border-radius: 10px;
  padding: 0 10px;
  font-size: 15px;
  margin-top: 5px;
}

.search-link-text {
  margin-left: 10px;
}

.search-link-list:hover {
  transform: scale(1.10);
  transition: transform 0.8s;
  background: #f4f4f4;
}
</style>