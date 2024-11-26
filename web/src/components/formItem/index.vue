<script setup>
import {defineProps} from "vue";

const props = defineProps({
  row: {
    type: Object,
    required: true
  }
});
const type = props.row.value

if (props.row.type === 'number' && typeof type !== 'number') {
  props.row.value = 0
}

if (props.row.type === 'dateRange' && typeof type !== 'object') {
  props.row.value = 0
}
</script>

<template>
  <!--文本框-->
  <a-form-item v-if="props.row.type==='input'" :field="props.row.key" :label="props.row.label+'：'" auto-label-width
               feedback label-col-flex="130px">
    <a-input v-model="props.row.value" autocomplete="off" placeholder="请输入" allow-clear/>
  </a-form-item>
  <!--数字框-->
  <a-form-item v-if="props.row.type==='number'" :field="props.row.key" :label="props.row.label+'：'" auto-label-width
               feedback label-col-flex="130px">
    <a-input-number :default-value="props.row.value" :style="{width:'100%'}" placeholder="请输入" allow-clear/>
  </a-form-item>
  <!--下拉框-->
  <a-form-item v-if="props.row.type==='select'" :field="props.row.key" :label="props.row.label+'：'" auto-label-width
               feedback label-col-flex="130px">
    <a-select v-model="props.row.value" placeholder="请选择" allow-clear>
      <a-option v-for="option in props.row.options" :key="option.value" :value="option.value">
        {{ option.label }}
      </a-option>
    </a-select>
  </a-form-item>
  <!--日期时间选择框-->
  <a-form-item v-if="props.row.type==='datePicker'" :field="props.row.key" :label="props.row.label+'：'" auto-label-width
               feedback label-col-flex="130px">
    <a-date-picker v-model="props.row.value" format="YYYY-MM-DD HH:mm:ss" show-time
                   :time-picker-props="{ defaultValue: '00:00:00' }" type="datetime" placeholder="请选择" allow-clear/>
  </a-form-item>
  <!--日期范围选择-->
  <a-form-item v-if="props.row.type==='dateRange'" :field="props.row.key" :label="props.row.label+'：'" auto-label-width
               feedback label-col-flex="130px">
    <a-range-picker
        v-model="props.row.value"

        style="width: 100%"
    />
  </a-form-item>
</template>

<style scoped>
.search-list {

}

.form {
  width: 98%
}

.ml-10 {
  margin-left: 10px;
}
</style>