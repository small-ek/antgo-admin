<script setup>
import {defineProps, reactive, ref} from "vue";
import formItem from "@/components/formItem/index.vue";

const props = defineProps({
  model: {
    type: Object,
    required: true
  }
});
console.log(props.model)
const isExpand = ref(false)
const onChangeIsExpand = () => {
  isExpand.value = !isExpand.value
}

const form = reactive({})
const onSearch = () => {
  console.log(props.model)
}
</script>

<template>
  <div class="container ant-card is-always-shadow search-list" v-if="props.model">
    <a-form :model="form" class="form" layout="horizontal">
      <a-row :gutter="15">
        <template v-for="(row,index) in props.model">
          <a-col :xs="24" :sm="24" :md="12" :lg="12" :xl="8" :xxl="8" v-if="index<3">
            <formItem :row="row"></formItem>
          </a-col>
          <a-col :xs="24" :sm="24" :md="12" :lg="12" :xl="8" :xxl="8" v-else-if="isExpand===true">
            <formItem :row="row"></formItem>
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
            <a-button class="ml-10">
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
  width: 98%
}

.ml-10 {
  margin-left: 10px;
}
</style>