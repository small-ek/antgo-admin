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
      <a-row :gutter="24" justify="space-between">
        <a-col :span="12">
          <div class="relative layout-list shadow check">
            <a-tooltip content="纵向">
              <a-space direction="vertical" :size="16">
                <a-row :gutter="2" justify="space-between">
                  <a-col flex="18px">
                    <div class="layout-menu"></div>
                  </a-col>
                  <a-col flex="auto">
                    <div class="layout-head"></div>
                    <div class="layout-body" style="margin-top: 3px"></div>
                  </a-col>
                </a-row>
              </a-space>
              <icon-check-circle-fill v-if="useLayout().layout==='vertical'" size="20" class="icon-circle"/>
            </a-tooltip>
          </div>
        </a-col>

        <a-col :span="12">
          <div class="relative layout-list shadow">
            <a-tooltip content="经典">
              <a-space direction="vertical" :size="16">
                <a-row :gutter="2" justify="space-between">
                  <a-col>
                    <div class="layout-head" style="height: 13px"></div>
                  </a-col>
                  <a-col style="margin-top: 3px">
                    <a-row :gutter="2" justify="center">
                      <a-col flex="20px">
                        <div class="layout-menu" style="height: 45px"></div>
                      </a-col>
                      <a-col flex="auto">
                        <div class="layout-body" style="height: 45px"></div>
                      </a-col>
                    </a-row>
                  </a-col>
                </a-row>
              </a-space>
              <icon-check-circle-fill v-if="useLayout().layout==='classic'" size="20" class="icon-circle"/>
            </a-tooltip>
          </div>
        </a-col>

        <a-col :span="12">
          <div class="relative layout-list shadow">
            <a-tooltip content="横向">
              <a-space direction="vertical" :size="16">
                <a-row :gutter="2" justify="space-between">
                  <a-col>
                    <div class="layout-head" style="height: 13px"></div>
                  </a-col>
                  <a-col style="margin-top: 3px">
                    <a-row :gutter="2" justify="center">
                     
                      <a-col flex="auto">
                        <div class="layout-body" style="height: 45px"></div>
                      </a-col>
                    </a-row>
                  </a-col>
                </a-row>
              </a-space>
              <icon-check-circle-fill v-if="useLayout().layout==='classic'" size="20" class="icon-circle"/>
            </a-tooltip>
          </div>
        </a-col>
      </a-row>
    </a-space>

  </a-drawer>
</template>

<style scoped>
.layout-list {
  width: 93px;
  height: 63px;
  cursor: pointer;
  padding: 7px;
  border-radius: 8px;
  position: relative;
}

.layout-menu {
  height: 63px;
  background-color: rgb(var(--primary-6));
  width: 100%;
  border-radius: 4px;
}

.layout-head {
  height: 15px;
  width: 100%;
  background-color: rgb(var(--primary-4));
  border-radius: 4px;
  border: 1px solid rgb(var(--primary-1));
}

.layout-body {
  height: 43px;
  width: 100%;
  background-color: rgb(var(--primary-2));
  border-radius: 4px;
  border: 1px dashed rgb(var(--primary-1));
}

.icon-circle {
  position: absolute;
  color: rgb(var(--primary-6));
  right: 15px;
  bottom: 15px;
}

.check {
  border: 1px dashed rgb(var(--primary-6));
}
</style>