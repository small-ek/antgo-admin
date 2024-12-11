<script setup>
import {modalWidth} from "@/utils/helper.js";
import {defineExpose, defineProps, ref} from "vue";
import FormItem from "@/components/formItem/index.vue";

const isVisible = ref(false)
const props = defineProps({
  model: {
    type: Array,
    required: true
  },
  form: {
    type: Object,
    required: true
  },
  rules: {
    type: Object,
    required: false
  }
});
const formRef = ref(null)
const emit = defineEmits(['submit'])
const setVisible = (value) => {
  isVisible.value = value
}

const submit = () => {
  emit('submit', props.form)
}
const reset = () => {
  formRef.value.resetFields()
}

defineExpose({
  setVisible
})
</script>

<template>
  <a-modal v-model:visible="isVisible" draggable title-align="start" :mask-closable="false" :footer="false"
           :width="modalWidth('800px')">
    <template #title>
      <a-typography-title :heading="5">
        编辑
      </a-typography-title>
    </template>
    <a-form ref="formRef" :model="props.form" :rules="props.rules" class="form" layout="horizontal" auto-label-width
            size="small"
            label-align="right" :label-col-props="{span: 8, offset: 0}" :wrapper-col-props="{span: 16, offset: 0}"
            @submit-success="submit">
      <a-row :gutter="15">
        <template v-for="(row,index) in props.model">
          <a-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" :xxl="24">
            <FormItem v-if="row.type!=='slot'" :row="row" :form="props.form"></FormItem>
            <!--自定义插槽-->
            <a-form-item v-else :field="row.key" :label="row.label+'：'" feedback>
              <slot :name="row.key"></slot>
            </a-form-item>
          </a-col>
        </template>
        <a-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" :xxl="24">
          <div style="float: right;margin-bottom: 20px">
            <a-button @click="reset">
              <template #icon>
                <icon-refresh/>
              </template>
              重置
            </a-button>

            <a-button class="ml-10" type="primary" html-type="submit">
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

</style>