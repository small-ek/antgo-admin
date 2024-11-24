<script setup>

import {reactive, ref} from "vue";
import {updateUserinfo} from "@/api/auth.js";
import {Message} from "@arco-design/web-vue";
import {useUserLoginStore} from "@/stores/userLogin.js";

const rules = reactive({
  nick_name: [
    {required: true, message: '请输入昵称', trigger: 'blur'},
  ],
  email: [
    {
      type: 'email',
      required: false,
    }
  ],
  phone: [
    {
      validator: (value, cb) => {
        console.log(isNaN(value))
        if (value && (value.toString().length !== 11 || isNaN(value))) {
          cb('请输入正确的手机号')
        } else {
          cb()
        }
      }
    }
  ],
})

const passwordRules = reactive({
  password: [
    {required: true, message: '请输入密码', trigger: 'blur'},
  ],
  new_password: [
    {required: true, message: '请输入新密码', trigger: 'blur'},
  ],
  confirm_password: [
    {required: true, message: '请输入确认密码', trigger: 'blur'},
    {
      validator: (value, cb) => {
        if (value !== passwordForm.value.new_password) {
          cb('两次密码输入不一致')
        } else {
          cb()
        }
      }
    }
  ]
})
const form = ref({
  nick_name: useUserLoginStore().userInfo.nick_name,
  phone: useUserLoginStore().userInfo.phone,
  email: useUserLoginStore().userInfo.email,
  status: useUserLoginStore().userInfo.status,
  username: useUserLoginStore().userInfo.username,
})

const passwordForm = ref({
  password: '',
  new_password: '',
  confirm_password: ''
})

const onSubmit = () => {
  updateUserinfo(form.value).then(res => {
    if (res.code === 0) {
      useUserLoginStore().setUserInfo(form.value)
      Message.success('保存成功')
    }
  })
}

const onSubmitPassword = () => {
  updateUserinfo(passwordForm.value).then(res => {
    if (res.code === 0) {
      useUserLoginStore().setUserInfo(passwordForm.value)
      Message.success('保存成功')
    }
  })
}
</script>

<template>
  <div class="container ant-card user-list">
    <a-tabs type="rounded" size="large">
      <a-tab-pane key="1" title="基础信息">
        <a-row :gutter="10">
          <a-col :xs="22" :sm="18" :md="15" :lg="13" :xl="10" class="m-auto" style="margin-top: 35px">
            <a-form :model="form" :rules="rules" layout="horizontal" auto-label-width feedback
                    @submit-success="onSubmit">
              <a-form-item field="nick_name" label="昵称">
                <a-input v-model="form.nick_name" autocomplete="off" placeholder="请输入昵称" allow-clear>
                </a-input>
              </a-form-item>
              <a-form-item field="phone" label="手机" hide-asterisk>
                <a-input v-model="form.phone" autocomplete="off" placeholder="请输入手机" allow-clear/>
              </a-form-item>
              <a-form-item field="email" label="电子邮件" hide-asterisk>
                <a-input v-model="form.email" autocomplete="off" placeholder="请输入电子邮件" allow-clear>
                </a-input>
              </a-form-item>
              <a-form-item>
                <a-button round html-type="submit" type="primary">
                  保存
                </a-button>
              </a-form-item>
            </a-form>
          </a-col>
        </a-row>
      </a-tab-pane>
      <a-tab-pane key="2" title="安全设置">
        <a-row :gutter="10">
          <a-col :xs="22" :sm="18" :md="15" :lg="13" :xl="10" class="m-auto" style="margin-top: 35px">
            <a-form :model="passwordForm" :rules="passwordRules" layout="horizontal" auto-label-width feedback
                    @submit-success="onSubmitPassword">
              <a-form-item field="password" label="原密码">
                <a-input v-model="passwordForm.password" autocomplete="off" placeholder="请输入原密码" allow-clear>
                </a-input>
              </a-form-item>
              <a-form-item field="new_password" label="新密码">
                <a-input v-model="passwordForm.new_password" autocomplete="off" placeholder="请输入新密码" allow-clear>
                </a-input>
              </a-form-item>
              <a-form-item field="confirm_password" label="确认新密码">
                <a-input v-model="passwordForm.confirm_password" autocomplete="off" placeholder="请输入确认新密码"
                         allow-clear>
                </a-input>
              </a-form-item>
              <a-form-item>
                <a-button round html-type="submit" type="primary">
                  保存
                </a-button>
              </a-form-item>
            </a-form>
          </a-col>
        </a-row>
      </a-tab-pane>
    </a-tabs>
  </div>
</template>

<style scoped>
.user-list {
  height: 600px;
}
</style>