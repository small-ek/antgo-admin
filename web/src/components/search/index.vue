<script setup>
import {defineEmits, defineProps, reactive, ref} from "vue";
import formItem from "@/components/formItem/index.vue";

const props = defineProps({
  model: {
    type: Object,
    required: true
  }
});
const formRef = ref(null)
const isExpand = ref(false)
const onChangeIsExpand = () => {
  isExpand.value = !isExpand.value
}
const emit = defineEmits(['search'])
const form = reactive({})
const onSearch = () => {
  emit('search', form)

}
const onReset = () => {
  formRef.value.resetFields()
}
</script>

<template>
  <div class="container ant-card is-always-shadow search-list" v-if="props.model">
    <a-form ref="formRef" :model="form" class="form" layout="horizontal" auto-label-width size="small"
            label-align="right" :label-col-props="{span: 8, offset: 0}" :wrapper-col-props="{span: 16, offset: 0}">
      <a-row :gutter="15">
        <template v-for="(row,index) in props.model">
          <a-col :xs="24" :sm="24" :md="12" :lg="12" :xl="8" :xxl="8" v-if="index<3">
            <formItem :row="row" :form="form"></formItem>
          </a-col>
          <a-col :xs="24" :sm="24" :md="12" :lg="12" :xl="8" :xxl="8" v-else-if="isExpand===true">
            <formItem :row="row" :form="form"></formItem>
          </a-col>
        </template>
        <a-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" :xxl="24">
          <div style="float: right;margin-bottom: 20px">
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
            <span v-show="isExpand===false">
              展开
            <icon-caret-down/>
            </span>
              <span v-show="isExpand===true">
              收起
            <icon-caret-up/>
            </span>
            </a-button>
          </div>
        </a-col>
      </a-row>

    </a-form>
  </div>
</template>

<style scoped>
.search-list {

}

.form {
  width: 95%;
  margin: 0 auto;
  margin-top: 20px;
}

.ml-10 {
  margin-left: 10px;
}
</style>