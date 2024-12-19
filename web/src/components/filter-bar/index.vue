<script setup>
import {defineEmits, defineProps, reactive, ref} from "vue";
import formItem from "@/components/form-item/index.vue";
import {useLayout} from "@/stores/layout.js";

const props = defineProps({
  model: {type: Object, required: true}
});
const formRef = ref(null)
const isExpand = ref(true)
const emit = defineEmits(['search'])
const form = reactive({})

const onSearch = () => emit('search', form)

const onChangeIsExpand = () => {
  isExpand.value = !isExpand.value
}

const onReset = () => {
  formRef.value.resetFields()
}
</script>

<template>

  <div class="container ant-card" v-if="props.model">
    <a-form ref="formRef" :model="form" class="form" :layout="useLayout().windowWidth>768?'horizontal':'vertical'"
            size="small" label-align="right" :label-col-props="{span: 8, offset: 0}"
            :wrapper-col-props="{span: 16, offset: 0}"
            @submit-success="onSearch">
      <a-grid :cols="{ xs: 1, sm: 1, md: 2, lg: 2, xl: 3, xxl: 3 }" :colGap="1" :rowGap="5" :collapsed="isExpand">
        <template v-for="(row,index) in props.model">
          <a-grid-item class="demo-item">
            <formItem v-if="row.type!=='slot'" :row="row" :form="form" @search="onSearch"></formItem>
            <a-form-item v-else :field="row.key" :label="row.label+'：'" feedback>
              <slot :name="row.key" :form="form"></slot>
            </a-form-item>
          </a-grid-item>
        </template>
        <a-grid-item suffix #="{ overflow }">
          <div style="float: right;margin-bottom: 10px">
            <a-button type="primary" @click="onSearch">
              <template #icon>
                <icon-search/>
              </template>
              筛选
            </a-button>
            <a-button class="ml-10" @click="onReset">
              <template #icon>
                <icon-refresh/>
              </template>
              重置
            </a-button>
            <a-button type="text" class="ml-10" @click="onChangeIsExpand">
            <span v-show="isExpand===true">
              展开
            <icon-caret-down/>
            </span>
              <span v-show="isExpand===false">
              收起
            <icon-caret-up/>
            </span>
            </a-button>
          </div>
        </a-grid-item>
      </a-grid>
    </a-form>
  </div>
</template>

<style scoped>
.form {
  width: 95%;
  margin: 0 auto;
  margin-top: 20px;
}

</style>