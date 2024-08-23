<template>
  <a-space direction="vertical" :size="16" style="display: block;">
    <a-row justify="space-between" align="center" :wrap="false">
      <a-col :span="20">
        <a-tabs type="card-gutter" class="ant-tabs" size="large" :editable="true" @delete="onDelete" @tab-click="onClick" auto-switch lazy-load :active-key="activeKey">
          <a-tab-pane v-for="(item, index) of data" :key="item.key" :closable="index!==0">
            <template #title>
              <icon-apps/>
              <text class="tags-text">{{ item.title }}</text>
            </template>
          </a-tab-pane>
        </a-tabs>
      </a-col>
      <a-col flex="45px">
        <div class="tabs-switch">
          <a-dropdown :popup-max-height="false">
            <a-button class="tabs-btn" size="mini">
              <icon-down/>
            </a-button>
            <template #content>
              <a-doption @click="onClose(0)">
                <icon-double-left/>
                关闭左侧
              </a-doption>
              <a-doption @click="onClose(1)">
                <icon-double-right/>
                关闭右侧
              </a-doption>
              <a-doption @click="onClose(2)">
                <icon-close/>
                关闭其他
              </a-doption>
              <a-doption @click="onClose(3)">
                <icon-close-circle/>
                全部关闭
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

const activeKey = ref('1');
const data = ref([
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
  }
]);

const onClick = (key) => {
  activeKey.value = key
}
const onDelete = (key) => {
  data.value = data.value.filter(item => item.key !== key)
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
    for (let i = data.value.length - 1; i > Number(activeKey.value); i--) {
      onDelete(i.toString())
    }
  }
  if (key === 2) {
    data.value = data.value.filter(item => item.key === activeKey.value || item.key === "0")
  }
  if (key === 3) {
    data.value = [data.value[0]]
  }
}
</script>
<style scoped>
.ant-tabs {
  width: 96%;
  height: 33px;
  margin: 10px auto;

}

.tags-text {
  margin-left: 5px;
}

.tabs-switch {
  margin-top: 0px;
}

.tabs-btn {
  background: white;
  width: 32px;
  height: 32px;
  color: black;
}
</style>