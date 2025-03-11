<template>
    <div class="register-container">
        <el-card class="register-card" header="注册" shadow="always">
            <el-input
                style="width: 75%; margin-left: 10%; padding: 15px;"
                v-model="formData.username"
                maxlength="50"
                placeholder="用户名"
                size="large"
                :prefix-icon="User"
                clearable
            />
            <el-input
                style="width: 75%; margin-left: 10%; padding: 15px;"
                v-model="formData.nickname"
                maxlength="50"
                placeholder="昵称"
                size="large"
                :prefix-icon="User"
                clearable
            />
            <el-input
                style="width: 75%; margin-left: 10%; padding: 15px;"
                v-model="formData.phone"
                maxlength="50"
                placeholder="手机号"
                size="large"
                :prefix-icon="Phone"
                clearable
            />
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
                placeholder="验证码"
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
                placeholder="密码"
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
                placeholder="重复密码"
                size="large"
                :prefix-icon="Lock"
                clearable
            />
            <div class="register-option">
                <el-link href="/login">{{ t('app.webui.returnlogin') }}</el-link>
            </div>
            <el-button
                type="success"
                size="large"
                style="width: 60%; margin: 2% 20%; margin-bottom: 20px; font-weight: bold; font-size: 16px;"
                auto-insert-space
                @click="handleRegister"
                >立即注册
            </el-button>
        </el-card>
    </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import { User, Lock,Phone,Camera,Message,Unlock } from '@element-plus/icons-vue'
import { useI18n } from 'vue-i18n';
import { ElMessage } from 'element-plus'
import { getCaptcha,register } from '@/api/login';
const { t } = useI18n();
const formData=ref({
    username:'',
    nickname:'',
    email:'',
    phone:'',
    password:'',
    cfmpassword:'',
    code:''
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
        service_type:"register"
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

function handleRegister() {
    // 验证用户名是否只包含英文和数字
    const usernameRegex = /^[a-zA-Z0-9]+$/;
    if (!usernameRegex.test(formData.value.username)) {
        ElMessage.error("用户名只能包含英文和数字");
        return;
    }

    // 验证密码和重复密码是否一致
    if (formData.value.password !== formData.value.cfmpassword) {
        ElMessage.error("两次输入的密码不一致");
        return;
    }

    // 验证邮箱是否有效
    const emailRegex = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;
    if (!emailRegex.test(formData.value.email)) {
        ElMessage.error("请输入有效的邮箱地址");
        return;
    }

    // 验证手机号是否为空
    if (formData.value.phone === '') {
        ElMessage.error("请输入手机号");
        return;
    }

    // 验证验证码是否为空
    if (formData.value.code === '') {
        ElMessage.error("请输入验证码");
        return;
    }

    // 验证昵称是否为空
    if (formData.value.nickname === '') {
        ElMessage.error("请输入昵称");
        return;
    }

    // 这里可以添加提交注册数据的逻辑
    console.log('提交注册数据:', formData.value);
    register({
        username:formData.value.username,
        nickname:formData.value.nickname,
        email:formData.value.email,
        phone:formData.value.phone,
        password:formData.value.password,
        code:formData.value.code
    }).then((res)=>{
        ElMessage.success("注册成功，请返回登录")
    }).catch((err)=>{
    })
}
</script>

<style scoped>
  .register-container {
    display: flex;
    justify-content: center;
  }

  .register-card {
    width: 500px;
    margin-top: 10%;
    margin-bottom: 10%;
    font-size: 20px;
    font-weight: bold;
    /*background: #303030;*/
  }
  
  .register-option {
    display: flex;
    justify-content: right;
    margin: 2% 10%;
  }
</style>