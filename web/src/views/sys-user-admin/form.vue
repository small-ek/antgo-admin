<script setup>
import {modalWidth} from "@/utils/helper.js";
import {defineExpose, defineProps, reactive, ref} from "vue";
import formItem from "@/components/formItem/index.vue";

const visible = ref(false)
const props = defineProps({
  model: {
    type: Array,
    required: true
  },
  form:{
    type: Object,
    required: true
  }
});
// const form = reactive({})
const formRef = ref(null)

const setVisible = (value) => {
  visible.value = value
}
defineExpose({
  setVisible
})

const onSubmit = () => {

}
const onReset = () => {
  formRef.value.resetFields()
}
</script>

<template>
  <a-modal v-model:visible="visible" draggable title-align="start" :mask-closable="false" :footer="false"
           :width="modalWidth('800px')">
    <template #title>
      <a-typography-title :heading="5">
        编辑
      </a-typography-title>
    </template>
    <a-form ref="formRef" :model="props.form" class="form" layout="horizontal" auto-label-width size="small"
            label-align="right" :label-col-props="{span: 8, offset: 0}" :wrapper-col-props="{span: 16, offset: 0}"
            @submit-success="onSubmit">
      <a-row :gutter="15">
        <template v-for="(row,index) in props.model">
          <a-col :xs="24" :sm="24" :md="24" :lg="24" :xl="12" :xxl="12">
            <formItem v-if="row.type!=='slot'" :row="row" :form="props.form"></formItem>
            <!--自定义插槽-->
            <a-form-item v-else :field="row.key" :label="row.label+'：'" feedback>
              <slot :name="row.key" :form="props.form"></slot>
            </a-form-item>
          </a-col>
        </template>
        <a-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" :xxl="24">
          <div style="float: right;margin-bottom: 20px">
            <a-button @click="onReset">
              <template #icon>
                <icon-refresh/>
              </template>
              重置
            </a-button>

            <a-button class="ml-10" type="primary" @click="onSubmit">
              <template #icon>
                <icon-search/>
              </template>
              提交
            </a-button>
          </div>
        </a-col>
      </a-row>

    </a-form>
  </a-modal>
</template>

<style scoped>
.form {
  width: 95%;
  margin: 0 auto;
  margin-top: 20px;
}

.ml-10 {
  margin-left: 10px;
}
</style>