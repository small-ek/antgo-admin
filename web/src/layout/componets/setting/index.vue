<script setup>
import {useLayout} from "@/stores/layout.js";
import {ref} from "vue";
import {useTheme} from "@/utils/theme.js";
import {IconMoon, IconSun, IconTranslate} from '@arco-design/web-vue/es/icon';

const history = ref([useLayout().primary])
const addHistory = (visible, color) => {
  if (!visible) {
    const index = history.value.indexOf(color);
    if (index !== -1) {
      history.value.splice(index, 1);
    }
    history.value.unshift(color);
    useLayout().setState("primary", color)
    useTheme().changePrimary(color)
  }
}

const themeList = ref([{title: "浅色", value: "light", icon: IconSun}, {title: "深色", value: "dark", icon: IconMoon}, {title: "跟随系统", value: "auto", icon: IconTranslate}])

const onUpdateHeader = (value) => {
  if (value === 'static') {
    useLayout().setState("isFixedHeader", false)
  } else {
    useLayout().setState("isFixedHeader", true)
  }
}
</script>

<template>
  <a-drawer :width="350" :visible="useLayout().showSetting" title="布局设置" placement="right" :closable="true" :footer="false" :drawerStyle="{padding:'2%'}" @cancel="useLayout().setState('showSetting', false)">
    <div class="drawer-body">
      <a-alert type="warning" closable banner center>
        <template #title>
          提示
        </template>
        应用配置可实时预览效果，但只是在当前浏览器使用,清楚缓存会导致失效。
      </a-alert>
      <!--布局设置-->
      <a-divider orientation="center" class="setting-title" style="margin-top: 20px;">
        <icon-layout/>
        布局设置
      </a-divider>
      <a-col :span="24" class="mb-15 title">
        布局
      </a-col>
      <a-space direction="vertical" :size="16" style="display: block;">
        <a-row :gutter="24" justify="space-between">
          <a-col :span="12" @click="useLayout().setState('layout','vertical')">
            <a-tooltip content="垂直">
              <div :class="['layout-item shadow layout-vertical', { 'is-select': useLayout().layout === 'vertical' }]">
                <div class="layout-dark"></div>
                <div class="layout-container">
                  <div class="layout-light"></div>
                  <div class="layout-content"></div>
                </div>
                <icon-check-circle-fill v-if="useLayout().layout==='vertical'" size="20" class="icon-circle"/>
              </div>
            </a-tooltip>
          </a-col>
          <a-col :span="12" @click="useLayout().setState('layout','classic')">
            <a-tooltip content="经典">
              <div :class="['layout-item shadow layout-classic', { 'is-select': useLayout().layout === 'classic' }]">
                <div class="layout-dark"></div>
                <div class="layout-container">
                  <div class="layout-light"></div>
                  <div class="layout-content"></div>
                </div>
                <icon-check-circle-fill v-if="useLayout().layout==='classic'" size="20" class="icon-circle"/>
              </div>
            </a-tooltip>
          </a-col>
          <a-col :span="12" @click="useLayout().setState('layout','transverse')">
            <a-tooltip content="横向">
              <div :class="['layout-item shadow layout-transverse', { 'is-select': useLayout().layout === 'transverse' }]">
                <div class="layout-dark"></div>
                <div class="layout-content"></div>
                <icon-check-circle-fill v-if="useLayout().layout==='transverse'" size="20" class="icon-circle"/>
              </div>
            </a-tooltip>
          </a-col>
          <a-col :span="12" @click="useLayout().setState('layout','columns')">
            <a-tooltip content="分栏">
              <div :class="['layout-item shadow layout-columns', { 'is-select': useLayout().layout === 'columns' }]">
                <div class="layout-dark"></div>
                <div class="layout-light"></div>
                <div class="layout-content"></div>
                <icon-check-circle-fill v-if="useLayout().layout==='columns'" size="20" class="icon-circle"/>
              </div>
            </a-tooltip>
          </a-col>

          <a-col :span="24" class="mb-15 title">
            侧边栏
          </a-col>
          <a-col :span="24" class="mb-15">
            <a-row justify="space-between" align="center">
              <a-col :span="10">
                侧边栏深色
              </a-col>
              <a-col :span="5">
                <a-switch v-model="useLayout().isDarkSidebar" :checked-value="true" :unchecked-value="false"/>
              </a-col>
            </a-row>
          </a-col>
          <a-col :span="24" class="mb-15">
            <a-row justify="space-between" align="center">
              <a-col :span="10">
                手风琴模式
              </a-col>
              <a-col :span="5">
                <a-switch v-model="useLayout().isAccordion" :checked-value="true" :unchecked-value="false"/>
              </a-col>
            </a-row>
          </a-col>
          <a-col :span="24" class="mb-15">
            <a-row justify="space-between" align="center">
              <a-col :span="10">
                宽度
              </a-col>
              <a-col :span="14">
                <a-input-number v-model="useLayout().sidebarWidth" placeholder="请输入" :precision="0" :step="1" :min="200" :max="320"/>
              </a-col>
            </a-row>
          </a-col>

          <a-col :span="24" class="mb-15 title">
            顶部栏
          </a-col>
          <a-col :span="24" class="mb-15">
            <a-row justify="space-between" align="center">
              <a-col :span="10">
                <a-tooltip content="头部颜色变化深色模式">
                  <span>头部深色</span>
                </a-tooltip>
              </a-col>
              <a-col :span="5">
                <a-switch v-model="useLayout().isDarkHeader" :checked-value="true" :unchecked-value="false"/>
              </a-col>
            </a-row>
          </a-col>
          <a-col :span="24" class="mb-15">
            <a-row justify="space-between" align="center">
              <a-col :span="10">
                显示面包屑导航
              </a-col>
              <a-col :span="5">
                <a-switch v-model="useLayout().isBreadcrumb" :checked-value="true" :unchecked-value="false"/>
              </a-col>
            </a-row>
          </a-col>
          <a-col :span="24" class="mb-15">
            <a-row justify="space-between" align="center">
              <a-col :span="10">
                显示多语言
              </a-col>
              <a-col :span="5">
                <a-switch v-model="useLayout().isLanguage" :checked-value="true" :unchecked-value="false"/>
              </a-col>
            </a-row>
          </a-col>
          <a-col :span="24" class="mb-15">
            <a-row justify="space-between" align="center">
              <a-col :span="10">
                显示全屏
              </a-col>
              <a-col :span="5">
                <a-switch v-model="useLayout().isFullScreen" :checked-value="true" :unchecked-value="false"/>
              </a-col>
            </a-row>
          </a-col>
          <a-col :span="24" class="mb-15">
            <a-row justify="space-between" align="center">
              <a-col :span="10">
                显示刷新
              </a-col>
              <a-col :span="5">
                <a-switch v-model="useLayout().isRefresh" :checked-value="true" :unchecked-value="false"/>
              </a-col>
            </a-row>
          </a-col>
          <a-col :span="24" class="mb-15">
            <a-row justify="space-between" align="center">
              <a-col :span="10">
                显示搜索
              </a-col>
              <a-col :span="5">
                <a-switch v-model="useLayout().isSearch" :checked-value="true" :unchecked-value="false"/>
              </a-col>
            </a-row>
          </a-col>
          <a-col :span="24" class="mb-15">
            <a-row justify="space-between" align="center">
              <a-col :span="10">
                模式
              </a-col>
              <a-col :span="14">
                <a-select placeholder="请选择" v-model="useLayout().header" @change="onUpdateHeader">
                  <a-option value="static">静止</a-option>
                  <a-option value="fixed">固定</a-option>
                  <a-option value="adaptive">滚动隐藏和显示</a-option>
                </a-select>
              </a-col>
            </a-row>
          </a-col>
        </a-row>
      </a-space>
      <a-col :span="24" class="mb-15 title">
        标签栏
      </a-col>
      <a-col :span="24" class="mb-15">
        <a-row justify="space-between" align="center">
          <a-col :span="10">
            显示标签栏
          </a-col>
          <a-col :span="5">
            <a-switch v-model="useLayout().isTabs" :checked-value="true" :unchecked-value="false"/>
          </a-col>
        </a-row>
      </a-col>
      <a-col :span="24" class="mb-15">
        <a-row justify="space-between" align="center">
          <a-col :span="10">
            是否显示图标
          </a-col>
          <a-col :span="5">
            <a-switch v-model="useLayout().isTabsIcon" :checked-value="true" :unchecked-value="false"/>
          </a-col>
        </a-row>
      </a-col>
      <a-col :span="24" class="mb-15">
        <a-row justify="space-between" align="center">
          <a-col :span="10">
            启动拖拽排序
          </a-col>
          <a-col :span="5">
            <a-switch v-model="useLayout().isTabsDraggable" :checked-value="true" :unchecked-value="false"/>
          </a-col>
        </a-row>
      </a-col>
      <a-col :span="24" class="mb-15">
        <a-row justify="space-between" align="center">
          <a-col :span="10">
            选项卡的样式
          </a-col>
          <a-col :span="14">
            <a-select placeholder="请选择" v-model="useLayout().tabsType">
              <a-option value="line">线条</a-option>
              <a-option value="card">卡片</a-option>
              <a-option value="card-gutter">卡片间隙</a-option>
              <a-option value="text">文本</a-option>
              <a-option value="rounded">圆角</a-option>
              <a-option value="capsule">胶囊居中</a-option>
            </a-select>
          </a-col>
        </a-row>
      </a-col>
      <a-col :span="24" class="mb-15 title">
        底部栏
      </a-col>
      <a-col :span="24" class="mb-15">
        <a-row justify="space-between" align="center">
          <a-col :span="10">
            显示底部栏
          </a-col>
          <a-col :span="5">
            <a-switch v-model="useLayout().isFooter" :checked-value="true" :unchecked-value="false"/>
          </a-col>
        </a-row>
      </a-col>
      <a-col :span="24" class="mb-15">
        <a-row justify="space-between" align="center">
          <a-col :span="10">
            底部文案
          </a-col>
          <a-col :span="14">
            <a-input placeholder="请输入" v-model="useLayout().footerText"/>
          </a-col>
        </a-row>
      </a-col>

      <!--全局主题-->
      <a-divider orientation="center" class="setting-title">
        <icon-dashboard/>
        全局主题
      </a-divider>
      <a-col :span="24" class="mb-15">
        <a-row justify="center" align="center">
          <a-radio-group v-model="useLayout().theme" @change="useTheme().updateDark()">
            <template v-for="item in themeList" :key="item">
              <a-radio :value="item.value">
                <template #radio="{ checked }">
                  <a-tag size="large" :checked="checked" bordered checkable>
                    <template #icon>
                      <component :is="item.icon"></component>
                    </template>
                    {{ item.title }}
                  </a-tag>
                </template>
              </a-radio>
            </template>
          </a-radio-group>
        </a-row>
      </a-col>

      <a-col :span="24" class="mb-15">
        <a-row justify="space-between" align="center">
          <a-col :span="10">
            主题颜色
          </a-col>
          <a-col :span="5">
            <a-color-picker
                :defaultValue="useLayout().primary"
                :historyColors="history"
                showHistory
                showPreset
                format="rgb"
                @popup-visible-change="addHistory"
            />
          </a-col>
        </a-row>
      </a-col>
      <a-col :span="24" class="mb-15">
        <a-row justify="space-between" align="center">
          <a-col :span="10">
            色弱模式
          </a-col>
          <a-col :span="5">
            <a-switch v-model="useLayout().isColorBlind" @change="useTheme().changeGreyOrWeak('weak',useLayout().isColorBlind)" :checked-value="true" :unchecked-value="false"/>
          </a-col>
        </a-row>
      </a-col>
      <a-col :span="24" class="mb-15">
        <a-row justify="space-between" align="center">
          <a-col :span="10">
            灰色模式
          </a-col>
          <a-col :span="5">
            <a-switch v-model="useLayout().isGrey" @change="useTheme().changeGreyOrWeak('grey',useLayout().isGrey)" :checked-value="true" :unchecked-value="false"/>
          </a-col>
        </a-row>
      </a-col>
    </div>
  </a-drawer>
</template>

<style scoped lang="less">
.icon-circle {
  position: absolute;
  color: rgb(var(--primary-6));
  right: 12px;
  bottom: 10px;
}

.layout-item {
  position: relative;
  box-sizing: border-box;
  width: 100px;
  height: 70px;
  padding: 6px;
  cursor: pointer;
  border-radius: 5px;
  box-shadow: 0 0 5px 1px var(--el-border-color-dark);
  transition: all 0.2s;
}

//纵向布局
.layout-vertical {
  display: flex;
  justify-content: space-between;

  .layout-dark {
    width: 20%;
  }

  .layout-container {
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    width: 72%;

    .layout-light {
      height: 20%;
    }

    .layout-content {
      height: 67%;
    }
  }
}

//经典布局
.layout-classic {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  margin-bottom: 20px;

  .layout-dark {
    height: 22%;
  }

  .layout-container {
    display: flex;
    justify-content: space-between;
    height: 70%;

    .layout-light {
      width: 20%;
    }

    .layout-content {
      width: 70%;
    }
  }
}

//横向布局
.layout-transverse {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  margin-bottom: 15px;

  .layout-dark {
    height: 20%;
  }

  .layout-content {
    height: 67%;
  }
}

//分栏布局
.layout-columns {
  display: flex;
  justify-content: space-between;
  margin-bottom: 15px;

  .layout-dark {
    width: 14%;
  }

  .layout-light {
    width: 17%;
  }

  .layout-content {
    width: 55%;
  }
}

.layout-dark {
  background-color: rgb(var(--primary-6));
  border-radius: 3px;
}

.layout-light {
  background-color: rgb(var(--primary-5));
  border-radius: 3px;
}

.layout-content {
  background-color: rgb(var(--primary-3));
  border: 1px dashed rgb(var(--primary-1));
  border-radius: 3px;
}

//选中
.is-select {
  border: 1px dashed rgb(var(--primary-6));
}

.title {
  font-weight: bold;
  letter-spacing: -.025em;
  font-size: 1.17em;
}

.setting-title {
  margin-bottom: 30px;
  margin-top: 30px;
}

.custom-radio-card {
  border: 1px solid var(--color-border-2);
  border-radius: 4px;
  width: 80px;
  box-sizing: border-box;
}


.custom-radio-card-title {
  color: var(--color-text-1);
  font-size: 14px;
  font-weight: bold;
  margin-bottom: 8px;
  text-align: center;
  margin-top: 10px;
}

.custom-radio-card:hover,
.custom-radio-card-checked,
.custom-radio-card:hover .custom-radio-card-mask,
.custom-radio-card-checked .custom-radio-card-mask {
  border-color: rgb(var(--primary-6));
}


.custom-radio-card:hover .custom-radio-card-title,
.custom-radio-card-checked .custom-radio-card-title {
  color: rgb(var(--primary-6));
}

.custom-radio-card-checked .custom-radio-card-mask-dot {
  background-color: rgb(var(--primary-6));
}
</style>