<template>
  <div id="login-bg">
    <div class="relative login" style="width: 100%;height: 100vh">
      <div class="center-vertical login-form">
        <div class="card">
          <div class="card-form">
            <a-row :gutter="10">
              <a-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" class="text-center">
                <span class="login-title">后台登录</span>
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
                      <img :src="captcha" @click="reloadCaptcha" class="login-code-img" alt="验证码"/>
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
import {onMounted, reactive, ref} from 'vue'
import {useRoute, useRouter} from "vue-router";
import {IconLock, IconSafe, IconUser} from '@arco-design/web-vue/es/icon';
import {getCaptcha, login} from "@/api/auth.js";
import {useI18n} from "vue-i18n"
import {Message} from "@arco-design/web-vue";
import {useUserLoginStore} from "@/stores/userLogin.js";
import {useNavigation} from "@/utils/base.js";
import {setMenu} from "@/utils/permission.js";

const router = useRouter();
const {jump} = useNavigation(router);
const route = useRoute()
const {t} = useI18n()
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

const form = ref({
  username: '',
  password: '',
  code: "",
  id: ""
})

let captcha = ref("")

//验证码
const reloadCaptcha = () => {
  getCaptcha().then(res => {
    captcha.value = res.data.pic
    form.value.id = res.data.id
  })
}
onMounted(() => {
  reloadCaptcha()
})

//登录
const onSubmit = () => {
  login(form.value).then(res => {
    if (res.code === 0) {
      useUserLoginStore().login(res.data)
      Message.info(t('tip.' + res.message))

      setMenu().then(() => {
        jump(route.query.redirect || '/')
      })
    } else {
      reloadCaptcha()
      Message.error(t('tip.' + res.message))
    }
  })
}
</script>

<style lang="less" scoped>
@import "index.less";
</style>
