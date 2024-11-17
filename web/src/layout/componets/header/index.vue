<script setup>
import Breadcrumbs from '@/layout/componets/breadcrumbs/index.vue'
import {onMounted, onUnmounted, ref} from "vue";
import {useLayout} from '@/stores/layout.js'
import headerLogo from "@/layout/componets/headerLogo/index.vue";
import {useTheme} from "@/utils/theme.js";
import Menu from "@/layout/componets/menu/index.vue";
import {useMenu} from "@/stores/menu.js";
import {useRouter} from "vue-router";
import {useNavigation} from "@/utils/base.js";
import {useUserLoginStore} from "@/stores/userLogin.js";
import {Message} from "@arco-design/web-vue";
import Mousetrap from 'mousetrap';
import {modalWidth} from "@/utils/helper.js";

const router = useRouter()
const navigation = useNavigation(router)
const isFullscreen = ref(false)
const menuRef = ref()
const isSearch = ref(false);
const searchList = ref(useMenu().subMenu)
const isShortcutKey = ref(false)


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
  isSearch.value = !isSearch.value
}

const logout = () => {
  useUserLoginStore().logout()
  Message.info('退出登录成功')
  setTimeout(() => {
    window.location.reload()
  }, 500)
}

// 搜索快捷键
Mousetrap.bind('alt+s', onShowSearch);
Mousetrap.bind('command+s', onShowSearch);
Mousetrap.bind('alt+q', logout);
Mousetrap.bind('command+q', logout);

onUnmounted(() => {
  Mousetrap.unbind('alt+s');
  Mousetrap.unbind('command+s');
  Mousetrap.unbind('alt+q', logout);
  Mousetrap.unbind('command+q', logout);
});
const onInputSeach = (e) => {
  searchList.value = useMenu().subMenu.filter(item => {
    return item.title.indexOf(e) !== -1 || item.pinyin.indexOf(e) !== -1
  })
}
const jump = (row) => {
  isSearch.value = false
  navigation.jumpTab("/" + row.path)
}


const onDropdown = (e) => {
  if (e === '退出登录') {
    logout()
  }
  if (e === '快捷键介绍') {
    isShortcutKey.value = true
  }
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
          <a-tooltip content="偏好设置" v-if="useLayout().windowWidth>768">
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
        <a-dropdown :popup-max-height="false" trigger="hover" @select="onDropdown">
          <span class="head-text">Antgo Admin</span>
          <template #content>
            <a-doption>
                <template #icon>
                  <icon-user/>
                </template>
                <template #default>个人设置</template>
            </a-doption>
            <a-doption>
                <template #icon>
                  <icon-compass/>
                </template>
                <template #default>快捷键介绍</template>
            </a-doption>
            <a-doption>
                <template #icon>
                  <icon-poweroff/>
                </template>
                <template #default>退出登录</template>
            </a-doption>
          </template>
        </a-dropdown>
        </span>
      </a-col>
    </a-row>


  </a-layout-header>
  <!--搜索-->
  <a-modal v-model:visible="isSearch" draggable title-align="start" :footer="false" :width="modalWidth('500px')">
    <template #title>
      <a-typography-title :heading="5">
        菜单搜索
      </a-typography-title>
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

  <!--快捷键介绍-->
  <a-modal v-model:visible="isShortcutKey" title-align="start" draggable :footer="false" :width="modalWidth('500px')">
    <template #title>
      <a-typography-title :heading="5">
        快捷键介绍
      </a-typography-title>
    </template>
    <a-space direction="vertical" :size="16" style="display: block;">
      <a-row class="grid-demo">
        <a-col :span="24">
          <a-typography-title :heading="5">
            全局
          </a-typography-title>
        </a-col>
        <a-col :span="24">
          <a-space>
            <a-button>Alt</a-button>
            <a-button>S</a-button>
            <h4>唤起导航搜索</h4>
          </a-space>
        </a-col>
        <a-col :span="24">
          <a-space>
            <a-button>Alt</a-button>
            <a-button>Q</a-button>
            <h4>退出登录</h4>
          </a-space>
        </a-col>
        <a-col :span="24">
          <a-space>
            <a-button>ESC</a-button>
            <h4>关闭弹框</h4>
          </a-space>
        </a-col>
        <a-col :span="24">
          <a-typography-title :heading="5">
            标签栏
          </a-typography-title>
        </a-col>
        <a-col :span="24">
          <a-space>
            <a-button>Alt</a-button>
            <a-button>←</a-button>
            <h4>切换到上一个标签页</h4>
          </a-space>
        </a-col>
        <a-col :span="24">
          <a-space>
            <a-button>Alt</a-button>
            <a-button>→</a-button>
            <h4>切换到下一个标签页</h4>
          </a-space>
        </a-col>
        <a-col :span="24">
          <a-space>
            <a-button>Alt</a-button>
            <a-button>W</a-button>
            <h4>关闭当前标签页</h4>
          </a-space>
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
  background: rgb(var(--primary-6));
  color: white;
}
</style>