<script setup>

import {reactive, ref} from "vue";

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
</script>

<template>
  <div class="container ant-card user-list">
    <a-tabs type="rounded" size="large">
      <a-tab-pane key="1" title="基础信息">
        <a-row :gutter="10">
          <a-col :xs="22" :sm="18" :md="15" :lg="13" :xl="10" class="m-auto" style="margin-top: 35px">
            <a-form :model="form" :rules="rules" layout="horizontal" auto-label-width feedback
                    @submit-success="onSubmit">
              <a-form-item field="nick_name" label="昵称" hide-asterisk>
                <a-input v-model="form.nick_name" autocomplete="off" placeholder="请输入昵称" allow-clear>
                </a-input>
              </a-form-item>
              <a-form-item field="phone" label="手机" hide-asterisk>
                <a-input v-model="form.phone" autocomplete="off" placeholder="请输入手机" allow-clear>
                </a-input>
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
        Content of Tab Panel 2
      </a-tab-pane>
    </a-tabs>
  </div>
</template>

<style scoped>
.user-list {
  height: 600px;
}
</style>