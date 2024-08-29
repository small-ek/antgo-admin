<script setup>

import {useLayout} from "@/stores/layout.js";
</script>

<template>
  <a-drawer :width="250" :visible="useLayout().isSetting" title="布局设置" placement="right" :closable="false" :footer="false" :drawerStyle="{padding:'2%'}" @cancel="useLayout().setState('isSetting', false)">
    <a-divider orientation="center" style="margin-top: 50px">
      <icon-layout/>
      布局样式
    </a-divider>

    <a-space direction="vertical" :size="16" style="display: block;">
      <a-row :gutter="24">
        <a-col :span="12">
          <a-tooltip content="纵向">
            <div class="layout-list shadow">
              <a-row :gutter="5" justify="center" align="center" style="padding-top: 8px">
                <a-col flex="17px">
                  <div class="layout-menu"></div>
                </a-col>
                <a-col flex="64px" class="layout-container">
                  <div class="layout-head"></div>
                  <div class="layout-body"></div>
                </a-col>
              </a-row>
              <icon-check-circle-fill v-if="useLayout().layout==='vertical'" size="20" class="icon-circle"/>
            </div>
          </a-tooltip>
        </a-col>

        <a-col :span="12">
          <a-tooltip content="经典">
            <div class="layout-list shadow">
              <a-row :gutter="5" justify="center" align="center" style="padding-top: 8px">
                <a-col flex="17px">
                  <div class="layout-head"></div>
                </a-col>
                <a-col flex="17px">
                  <div class="layout-menu"></div>
                </a-col>
                <a-col flex="64px">
                  <div class="layout-body"></div>
                </a-col>
              </a-row>
              <icon-check-circle-fill v-if="useLayout().layout==='classic'" size="20" class="icon-circle"/>
            </div>
          </a-tooltip>
        </a-col>
      </a-row>
    </a-space>

  </a-drawer>
</template>

<style scoped>
.layout-list {
  width: 100px;
  height: 70px;
  cursor: pointer;
  border-radius: 2px;
  position: relative;
}

.layout-menu {
  height: 58px;
  background-color: rgb(var(--primary-6));
  width: 100%;
  border-radius: 4px;
}

.layout-head {
  height: 20%;
  width: 100%;
  background-color: rgb(var(--primary-4));
  border-radius: 4px;
  border: 1px solid rgb(var(--primary-1));
}

.layout-body {
  height: 67%;
  width: 100%;
  background-color: rgb(var(--primary-2));
  border-radius: 4px;
  margin-top: 3px;
  border: 1px dashed rgb(var(--primary-1));
}

.icon-circle {
  position: absolute;
  color: rgb(var(--primary-6));
  right: 15px;
  bottom: 10px;
}
</style>