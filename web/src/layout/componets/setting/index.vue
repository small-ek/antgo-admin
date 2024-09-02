<script setup>
import {useLayout} from "@/stores/layout.js";
import {ref} from "vue";
import {useTheme} from "@/utils/theme.js";

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
</script>

<template>
  <a-drawer :width="250" :visible="useLayout().isSetting" title="布局设置" placement="right" :closable="false" :footer="false" :drawerStyle="{padding:'2%'}" @cancel="useLayout().setState('isSetting', false)">
    <a-divider orientation="center" class="setting-title" style="margin-top: 50px;margin-bottom: 30px">
      <icon-layout/>
      布局设置
    </a-divider>

    <a-space direction="vertical" :size="16" style="display: block;">
      <a-row :gutter="24" justify="space-between">
        <a-col :span="12">
          <a-tooltip content="纵向">
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

        <a-col :span="12">
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

        <a-col :span="12">
          <a-tooltip content="横向">
            <div :class="['layout-item shadow layout-transverse', { 'is-select': useLayout().layout === 'transverse' }]">
              <div class="layout-dark"></div>
              <div class="layout-content"></div>
              <icon-check-circle-fill v-if="useLayout().layout==='transverse'" size="20" class="icon-circle"/>
            </div>
          </a-tooltip>
        </a-col>

        <a-col :span="12">
          <a-tooltip content="分栏">
            <div :class="['layout-item shadow layout-columns', { 'is-select': useLayout().layout === 'transverse' }]">
              <div class="layout-dark"></div>
              <div class="layout-light"></div>
              <div class="layout-content"></div>
              <icon-check-circle-fill v-if="useLayout().layout==='transverse'" size="20" class="icon-circle"/>
            </div>
          </a-tooltip>
        </a-col>

        <a-col :span="24" class="mb-15">
          <a-row justify="space-between" align="center">
            <a-col :span="10">
              <a-tooltip content="侧边栏颜色变化深色模式">
                <span>侧边栏深色</span>
              </a-tooltip>
            </a-col>
            <a-col :span="5">
              <a-switch v-model="useLayout().asideInverted" checked-value="yes" unchecked-value="no"/>
            </a-col>
          </a-row>
        </a-col>

        <a-col :span="24" class="mb-15">
          <a-row justify="space-between" align="center">
            <a-col :span="10">
              <a-tooltip content="头部颜色变化深色模式">
                <span>头部反转色</span>
              </a-tooltip>
            </a-col>
            <a-col :span="5">
              <a-switch v-model="useLayout().headerInverted" checked-value="yes" unchecked-value="no"/>
            </a-col>
          </a-row>
        </a-col>
      </a-row>
    </a-space>

    <a-divider orientation="center" class="setting-title" style="margin-bottom: 30px">
      <icon-dashboard/>
      全局主题
    </a-divider>

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


</style>