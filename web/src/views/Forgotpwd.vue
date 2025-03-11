<template>
  <div class="forgot-container">
      <el-card class="forgot-card" :header="t('app.webui.forgot')" shadow="always">
          <div style="display: flex; align-items: center;">
              <el-input
                  style="width: 75%; margin-left: 10%; padding: 15px;"
                  v-model="formData.email"
                  maxlength="50"
                  :placeholder="t('app.webui.email')"
                  size="large"
                  :prefix-icon="Message"
                  clearable
              />
              <el-button :loading="captcha_loading" type="primary" size="small" @click="handleCaptchaClick">
                {{ captcha_loading?remaintime+'s':'获取验证码' }}
              </el-button>
              <!-- <el-countdown format="ss" :value="remaintime" @finish="captcha_loading=false" value-style="font-size: 14px;" /> -->
          </div>
          <el-input
              style="width: 75%; margin-left: 10%; padding: 15px;"
              v-model="formData.code"
              maxlength="50"
              :placeholder="t('app.webui.captcha')"
              size="large"
              :prefix-icon="Camera"
              clearable
          />
          <el-input
              style="width: 75%; margin-left: 10%; padding: 15px;"
              v-model="formData.password"
              maxlength="50"
              type="password"
              show-password
              :placeholder="t('app.webui.newpassword')"
              size="large"
              :prefix-icon="Unlock"
              clearable
          />
          <el-input
              style="width: 75%; margin-left: 10%; padding: 15px;"
              v-model="formData.cfmpassword"
              maxlength="50"
              type="password"
              show-password
              :placeholder="t('app.webui.confirmpassword')"
              size="large"
              :prefix-icon="Lock"
              clearable
          />
          <div class="forgot-option">
              <el-link href="/login">{{ t('app.webui.returnlogin') }}</el-link>
          </div>
          <el-button
            type="success"
            size="large"
            style="width: 60%; margin: 2% 20%; margin-bottom: 20px; font-weight: bold; font-size: 16px;"
            auto-insert-space
            @click="handleSubmit"
            >{{ t('app.webui.submit') }}
          </el-button>
      </el-card>
  </div>
</template>

<script lang="ts" setup>
  import { ref } from 'vue'
  import { User, Lock, Unlock, Message, Camera } from '@element-plus/icons-vue'
  import { useI18n } from 'vue-i18n';
  import { getCaptcha } from '@/api/login';
  import { ElMessage } from 'element-plus';
import { forgetPassword } from '@/api/login';

  const { t } = useI18n()
  const formData=ref({
    email:"",
    code:"",
    password:"",
    cfmpassword:""
  })
  const captcha_loading=ref(false)
  const remaintime = ref(0);

  //获取验证码
  function handleCaptchaClick() {
    const emailRegex = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;
    if (!emailRegex.test(formData.value.email)) {
      ElMessage.error("请输入有效的邮箱地址");
      return false;
    }
    if(!captcha_loading.value){
      getCaptcha({
        type:"email",
        data:formData.value.email,
        service_type:"forget"
      }).then((res)=>{
        ElMessage.success("验证码发送成功")
        captcha_loading.value = true;
        remaintime.value = 30;
        const intervalId = setInterval(() => {
            if (remaintime.value > 0) {
                remaintime.value--;
            } else {
                captcha_loading.value=false
                clearInterval(intervalId);
            }
        }, 1000);
      }).catch((err)=>{
        
      })
    }
  }

// 数据验证函数
function validateForm() {
  const { email, code, password, cfmpassword } = formData.value;
  // 验证邮箱
  const emailRegex = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;
  if (!emailRegex.test(email)) {
    ElMessage.error("请输入有效的邮箱地址");
    return false;
  }

  // 验证验证码
  if (code.length === 0) {
    ElMessage.error("请输入验证码");
    return false;
  }

  // 验证新密码
  if (password.length < 6) {
    ElMessage.error("新密码长度不能少于6位");
    return false;
  }

  // 验证确认密码
  if (password !== cfmpassword) {
    ElMessage.error("两次输入的密码不一致");
    return false;
  }

  return true;
}

// 处理提交事件
function handleSubmit() {
  if (validateForm()) {
    // 这里可以添加提交表单的逻辑
    console.log("表单验证通过，提交表单", formData.value);
    forgetPassword({
      email: formData.value.email,
      password: formData.value.password,
      code: formData.value.code
    }).then((res)=>{
      ElMessage.success("重置密码成功，请登录")
    }).catch((err)=>{
    })
  }
}

</script>

<style scoped>
.forgot-container {
  display: flex;
  justify-content: center;
}

.forgot-card {
  width: 500px;
  margin-top: 10%;
  margin-bottom: 10%;
  font-size: 20px;
  font-weight: bold;
  /*background: #303030;*/
}

.forgot-option {
  display: flex;
  justify-content: right;
  margin: 2% 10%;
}
</style>