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


</script>

<template>
  <!--文本框-->
  <a-form-item v-if="props.row.type==='input'" :field="props.row.key" :label="props.row.label+'：'" feedback>
    <a-input v-model="props.form[props.row.key]" autocomplete="off" class="input" :placeholder="props.row.placeholder||'请输入'" allow-clear/>
  </a-form-item>
  <!--数字框-->
  <a-form-item v-if="props.row.type==='number'" :field="props.row.key" :label="props.row.label+'：'" feedback>
    <a-input-number v-model="props.form[props.row.key]" class="input" :placeholder="props.row.placeholder||'请输入'" allow-clear/>
  </a-form-item>
  <!--下拉框-->
  <a-form-item v-if="props.row.type==='select'" :field="props.row.key" :label="props.row.label+'：'" feedback>
    <a-select v-model="props.form[props.row.key]" :placeholder="props.row.placeholder||'请选择'" allow-clear class="input">
      <a-option v-for="option in props.row.options" :key="option.value" :value="option.value">
        {{ option.label }}
      </a-option>
    </a-select>
  </a-form-item>
  <!--日期时间选择框-->
  <a-form-item v-if="props.row.type==='datePicker'" :field="props.row.key" :label="props.row.label+'：'" feedback>
    <a-date-picker v-model="props.form[props.row.key]" format="YYYY-MM-DD HH:mm:ss" style="width: 100%" show-time
                   :time-picker-props="{ defaultValue: '00:00:00' }" type="datetime" :placeholder="props.row.placeholder||'请选择'" allow-clear/>
  </a-form-item>
  <!--日期范围选择-->
  <a-form-item v-if="props.row.type==='dateRange'" :field="props.row.key" :label="props.row.label+'：'" feedback>
    <a-range-picker v-model="props.form[props.row.key]" style="width: 100%"/>
  </a-form-item>

</template>

<style scoped>

.input {
  width: 100%;
}
</style>