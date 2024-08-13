<template>
  <div id="login-bg">
    <div class="relative login" style="width: 100%;height: 100vh">
      <div class="center-vertical login-form">
        <div class="card">
          <div class="card-form">
            <a-row :gutter="10">
              <a-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" class="text-center">
                <span style="font-size: 23px;font-weight:bold">后台登录</span>
              </a-col>
              <a-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" style="margin-top: 35px">
                <a-form :model="form" :rules="rules" layout="vertical" feedback @submit-success="onSubmit">
                  <a-form-item field="username" hide-asterisk>
                    <a-input v-model="form.username" autocomplete="off" placeholder="请输入用户名" size="large" allow-clear>
                      <template #prefix>
                        <icon-user/>
                      </template>
                    </a-input>
                  </a-form-item>
                  <a-form-item field="password" hide-asterisk>
                    <a-input-password v-model="form.password" autocomplete="off" placeholder="请输入密码" show-password allow-clear size="large">
                      <template #prefix>
                        <icon-lock/>
                      </template>
                    </a-input-password>
                  </a-form-item>
                  <a-form-item field="code" hide-asterisk>
                    <a-input v-model="form.code" style="width: 63%" autocomplete="off" placeholder="请输入验证码" size="large" allow-clear>
                      <template #prefix>
                        <icon-safe/>
                      </template>
                    </a-input>
                    <div class="login-code">
                      <img :src="show.pic" @click="reloadCaptcha" class="login-code-img" alt="验证码"/>
                    </div>
                  </a-form-item>
                  <a-form-item>
                    <a-button round html-type="submit" long type="primary" size="large">
                      登 &nbsp 录
                    </a-button>
                  </a-form-item>
                </a-form>
              </a-col>
            </a-row>
          </div>

        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import {onMounted, reactive} from 'vue'
import {IconLock, IconSafe, IconUser} from '@arco-design/web-vue/es/icon';
import {getCaptcha} from "@/api/common.js";

const rules = reactive({
  username: [
    {required: true, message: '请输入用户名', trigger: 'blur'},
  ],
  password: [
    {required: true, message: '请输入密码', trigger: 'blur',},
  ],
  code: [
    {required: true, message: '请输入验证码', trigger: 'blur',},
  ],
})

const form = reactive({
  username: '',
  password: '',
  code: "",
  id: ""
})

let show = reactive({
  pic: ""
})

//验证码
const reloadCaptcha = () => {
  getCaptcha().then(res => {
    show.pic = res.result.pic
    form.id = res.result.id
  })
}
onMounted(() => {
  reloadCaptcha()
})

const onSubmit = () => {
  console.log(form)
}
</script>

<style lang="less" scoped>
@import "index.less";
</style>
