<script setup>
import {reactive} from "vue";

const form = reactive({
  value1: '',
  value2: '',
  value3: '',
  value4: '',
  value5: '',
})
</script>

<template>
  <div class="container ant-card search-list">
    <a-form :model="form">
      <a-row :gutter="16">
        <a-col :span="8">
          <a-form-item field="value1" label="Value 1" label-col-flex="100px">
            <a-input v-model="form.value1" placeholder="please enter..." />
          </a-form-item>
        </a-col>
        <a-col :span="8">
          <a-form-item field="value2" label="Value 2" label-col-flex="80px">
            <a-input v-model="form.value2" placeholder="please enter..." />
          </a-form-item>
        </a-col>
        <a-col :span="8">
          <a-form-item field="value3" label="Value 3" label-col-flex="80px">
            <a-input v-model="form.value3" placeholder="please enter..." />
          </a-form-item>
        </a-col>
      </a-row>
      <a-row :gutter="16">
        <a-col :span="16">
          <a-form-item field="value4" label="Value 4" label-col-flex="100px">
            <a-input v-model="form.value4" placeholder="please enter..." />
          </a-form-item>
        </a-col>
        <a-col :span="8">
          <a-form-item field="value5" label="Value 5" label-col-flex="80px">
            <a-input v-model="form.value5" placeholder="please enter..." />
          </a-form-item>
        </a-col>
      </a-row>
    </a-form>
  </div>
</template>

<style scoped>
.search-list{
  width: 70vw;
  min-height: 10vw;
}
</style>