<template>
  <a-space direction="vertical" v-if="useLayout().isTabs" :size="16" class="ant-tabs-space">
    <a-row justify="space-between" align="center" :wrap="false">
      <a-col :xs="21" :sm="21" :md="21" :lg="22" :xl="23" :xxl="23">
        <VueDraggable v-model="list" :animation="150" target=".arco-tabs-nav-tab-list" direction="horizontal" :disabled="!useLayout().isTabsDraggable">
          <a-tabs :type="useLayout().tabsType" class="ant-tabs" size="large" :editable="true" @delete="onDelete" @tab-click="onClick" auto-switch lazy-load :active-key="activeKey">
            <a-tab-pane v-for="(item, index) in list" :key="item.key" :closable="index!==0" :class="[{'tab0':index===0}]">
              <template #title>
                <icon-apps v-if="useLayout().isTabsIcon" class="cursor-move"/>
                <span class="no-select">{{ item.title }}</span>
              </template>
            </a-tab-pane>
          </a-tabs>
        </VueDraggable>
      </a-col>
      <a-col flex="44px">
        <div class="tabs-switch">
          <a-dropdown :popup-max-height="false">
            <a-button class="tabs-btn" size="mini">
              <icon-down/>
            </a-button>
            <template #content>
              <a-doption @click="onClose(item.key)" v-for="item in closeText">
                <component :is="item.icon"></component>
                {{ item.title }}
              </a-doption>
            </template>
          </a-dropdown>
        </div>
      </a-col>
    </a-row>
  </a-space>
</template>

<script setup>
import {ref} from 'vue';
import {useLayout} from "@/stores/layout.js";
import {VueDraggable} from 'vue-draggable-plus'
import {IconClose, IconCloseCircle, IconDoubleLeft, IconDoubleRight, IconRefresh} from '@arco-design/web-vue/es/icon'

const activeKey = ref('1');
//右侧关闭文字
const closeText = ref([
  {key: 0, title: '关闭左侧', icon: IconDoubleLeft},
  {key: 1, title: '关闭右侧', icon: IconDoubleRight},
  {key: 2, title: '关闭其他', icon: IconClose},
  {key: 3, title: '关闭所有', icon: IconCloseCircle},
  {key: 4, title: '刷新', icon: IconRefresh}
])

const list = ref([
  {
    key: '0',
    title: 'Tab 0',
    content: 'Content of Tab Panel 1'
  },
  {
    key: '1',
    title: 'Tab 1',
    content: 'Content of Tab Panel 1'
  },
  {
    key: '2',
    title: 'Tab 2',
    content: 'Content of Tab Panel 2'
  },
  {
    key: '3',
    title: 'Tab 3',
    content: 'Content of Tab Panel 3'
  },
  {
    key: '4',
    title: 'Tab 4',
    content: 'Content of Tab Panel 4'
  },
  {
    key: '5',
    title: 'Tab 5',
    content: 'Content of Tab Panel 5'
  },
  {
    key: '6',
    title: 'Tab 6',
    content: 'Content of Tab Panel 6'
  },
  {
    key: '7',
    title: 'Tab 7',
    content: 'Content of Tab Panel 7'
  },{
    key: '8',
    title: 'Tab 8',
    content: 'Content of Tab Panel 8'
  },{
    key: '9',
    title: 'Tab 9',
    content: 'Content of Tab Panel 9'
  },{
    key: '10',
    title: 'Tab 10',
    content: 'Content of Tab Panel 10'
  },
]);

const onClick = (key) => {
  activeKey.value = key
}

const onDelete = (key) => {
  list.value = list.value.filter(item => item.key !== key)
};

//onClose(0) 关闭左侧
//onClose(1) 关闭右侧
//onClose(0) 关闭其他
//onClose(0) 关闭所有
const onClose = (key) => {
  if (key === 0) {
    for (let i = (activeKey.value - 1); i > 0; i--) {
      onDelete(i.toString())
    }
  }
  if (key === 1) {
    for (let i = list.value.length - 1; i > Number(activeKey.value); i--) {
      onDelete(i.toString())
    }
  }
  if (key === 2) {
    list.value = list.value.filter(item => item.key === activeKey.value || item.key === "0")
  }
  if (key === 3) {
    list.value = [list.value[0]]
  }
  if (key === 4) {
    window.location.reload()
  }
}
</script>

<style scoped>
.cursor-move {
  cursor: move;
}

.affix {
  position: sticky;
  top: 44px;
  z-index: 999;
}

.ant-tabs-space {
  display: block;
  background: rgb(var(--gray-1));
  position: relative;
  top: -1px
}

.ant-tabs {
  width: 96%;
  height: 33px;
  margin: 10px auto;

}

.tags-text {
  margin-left: 5px;
}

.tabs-switch {
  margin-top: 0;
}

.tabs-btn {
  background: var(--color-bg-5);
  width: 32px;
  height: 32px;
  color: var(--color-text-2);

}

:deep(.arco-tabs-nav-button-right) {
  color: var(--color-text-1);
}

:deep(.arco-tabs-nav-button-disabled) {
  color: var(--color-text-4) !important;
}

:deep(.arco-tabs-nav-button-left) {
  color: var(--color-text-1);
}
</style>