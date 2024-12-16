<script setup>
import {defineProps} from "vue";

const props = defineProps({
  row: {
    type: Object,
    required: true
  },
  form: {
    type: Object,
    required: true
  }
});
const type = props.row.value

if (props.row.type === 'number' && typeof type !== 'number') {
  props.row.value = Number(props.row.value);
}

props.form[props.row.key] = props.row.value

const emit = defineEmits(['search'])
const onSearch = () => {
  emit('search')
}
/**
 * 树菜单过滤
 * @param searchValue
 * @param nodeData
 * @returns {boolean}
 */
const filterTreeNode = (searchValue, nodeData) => {
  return nodeData.title.toLowerCase().indexOf(searchValue.toLowerCase()) > -1;
}

</script>

<template>
  <!--文本框-->
  <a-form-item v-if="props.row.type==='input'" :field="props.row.key" :label="props.row.label+'：'" feedback>
    <a-input v-model="props.form[props.row.key]" autocomplete="off" class="input"
             :placeholder="props.row.placeholder||'请输入'" :max-length="props.row.maxLength||60" show-word-limit
             allow-clear @press-enter="onSearch" @clear="onSearch"/>
  </a-form-item>

  <!--密码-->
  <a-form-item v-if="props.row.type==='password'" :field="props.row.key" :label="props.row.label+'：'" feedback>
    <a-input-password v-model="props.form[props.row.key]" autocomplete="off" :max-length="props.row.maxLength||60"
                      show-word-limit :placeholder="props.row.placeholder||'请输入密码'" show-password allow-clear>
    </a-input-password>
  </a-form-item>

  <!--开关-->
  <a-form-item v-if="props.row.type==='switch'&&props.row.options" :field="props.row.key" :label="props.row.label+'：'"
               feedback>
    <a-switch v-model="props.form[props.row.key]"
              :checked-value="props.row.options[1]?props.row.options[1]['value']:'1'"
              :unchecked-value="props.row.options[0]?props.row.options[0]['value']:'2'">
      <template #checked v-if="props.row.options[1]&&props.row.options[1]['label']">
        {{ props.row.options[1]['label'] }}
      </template>
      <template #unchecked v-if="props.row.options[0]&&props.row.options[0]['label']">
        {{ props.row.options[0]['label'] }}
      </template>
    </a-switch>
  </a-form-item>

  <!--单选框按钮-->
  <a-form-item v-if="props.row.type==='radio-group-btn'&&props.row.options" :field="props.row.key"
               :label="props.row.label+'：'" feedback>
    <a-radio-group type="button" v-model="props.form[props.row.key]">
      <a-radio :value="row.value" v-for="row in props.row.options">{{ row.label }}</a-radio>
    </a-radio-group>
  </a-form-item>

  <!--文本许-->
  <a-form-item v-if="props.row.type==='textarea'" :field="props.row.key" :label="props.row.label+'：'" feedback>
    <a-textarea :placeholder="props.row.placeholder||'请输入'" :max-length="props.row.maxLength||200" allow-clear
                show-word-limit/>
  </a-form-item>

  <!--数字框-->
  <a-form-item v-if="props.row.type==='number'" :field="props.row.key" :label="props.row.label+'：'" feedback>
    <a-input-number v-model="props.form[props.row.key]" class="input" :placeholder="props.row.placeholder||'请输入'"
                    allow-clear @clear="onSearch"/>
  </a-form-item>

  <!--下拉框-->
  <a-form-item v-if="props.row.type==='select'" :field="props.row.key" :label="props.row.label+'：'" feedback>
    <a-select v-model="props.form[props.row.key]" :placeholder="props.row.placeholder||'请选择'" allow-clear
              class="input" @clear="onSearch" allow-search>
      <a-option v-for="option in props.row.options" :key="option.value" :value="option.value">
        {{ option.label }}
      </a-option>
    </a-select>
  </a-form-item>
  <!--树菜单-->
  <a-form-item v-if="props.row.type==='tree-select'" :field="props.row.key" :label="props.row.label+'：'" feedback>
    <a-tree-select v-model="props.form[props.row.key]" :data="props.row.options"
                   :placeholder="props.row.placeholder||'请选择'" allow-clear
                   :filter-tree-node="filterTreeNode"
                   class="input" allow-search>
    </a-tree-select>
  </a-form-item>
  <!--日期时间选择框-->
  <a-form-item v-if="props.row.type==='datePicker'" :field="props.row.key" :label="props.row.label+'：'" feedback>
    <a-date-picker v-model="props.form[props.row.key]" format="YYYY-MM-DD HH:mm:ss" style="width: 100%" show-time
                   :time-picker-props="{ defaultValue: '00:00:00' }" type="datetime"
                   :placeholder="props.row.placeholder||'请选择'" allow-clear @clear="onSearch"/>
  </a-form-item>
  <!--日期范围选择-->
  <a-form-item v-if="props.row.type==='dateRange'" :field="props.row.key" :label="props.row.label+'：'" feedback>
    <a-range-picker v-model="props.form[props.row.key]" showTime
                    :time-picker-props="{defaultValue:['00:00:00','23:59:59']}" style="width: 100%" @clear="onSearch"/>
  </a-form-item>

</template>

<style scoped>

.input {
  width: 100%;
}
</style>