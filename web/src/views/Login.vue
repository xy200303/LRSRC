<template>
    <div class="login-container">
        <el-card class="login-card" :header="t('app.webui.login')" shadow="always">
            <el-input
                v-model="uname"
                maxlength="50"
                :placeholder="t('app.webui.username')"
                size="large"
                :prefix-icon="User"
                clearable
            />
            <el-input
                v-model="passwd"
                maxlength="50"
                type="password"
                show-password
                :placeholder="t('app.webui.password')"
                size="large"
                :prefix-icon="Lock"
                clearable
                @keyup.enter.native="Login"
            />
            <div class="login-option">
                <el-link href="/forgotpwd" style="margin-right: 10px;">{{ t('app.webui.forgot') }}</el-link>
                <el-link href="/register" v-if="register_enable">立即注册</el-link>
            </div>
            <el-button
                type="success"
                size="large"
                style="width: 60%; margin: 2% 20%; margin-bottom: 20px; font-weight: bold; font-size: 16px;"
                @click="Login"
                auto-insert-space
                >{{t('app.webui.login')}}
            </el-button>
        </el-card>
    </div>
</template>

<script lang="ts" setup>
    import { ref } from 'vue'
    import { User, Lock } from '@element-plus/icons-vue'
    import { useRouter } from 'vue-router'
    import { onMounted } from 'vue';
    import { useI18n } from 'vue-i18n';
	import {useUserStore} from "@/stores/userStores"
	import { ElMessage } from 'element-plus'
    import { getBaseSysConfigMap } from '@/api/common';
    const { t } = useI18n();
	const userStore=useUserStore()

    const register_enable=ref(false)
		
    onMounted(()=>{
		// alert(userStore.token)
        getBaseSysConfigMap(null).then((res)=>{
            register_enable.value=res.data.register_enable
        }).catch((err)=>{
        })
	});
    const router = useRouter();
    const uname = ref('');
    const passwd = ref('');

    async function Login() {
        if (uname.value == '' || passwd.value == '') {
            ElMessage.error('账号或密码不能为空');
            return;
        }
        try {
			const response=await userStore.login({
				username: uname.value,
				password: passwd.value,
			})
            // 检查登录是否成功
            if (response.data.token) {
                ElNotification({
                    title: t('app.webui.loginsucc'),
                    message: userStore.userInfo.username + ', ' + t('app.webui.welcome'),
                    type: 'success',
                });
                router.push('/')
            }
        } catch (error) {
            // 处理请求错误
            console.error(error);
        }
    }
    
</script>

<style scoped>
  .login-container {
    display: flex;
    justify-content: center;
  }

  .login-card {
    width: 500px;
    margin-top: 10%;
    margin-bottom: 10%;
    font-size: 20px;
    font-weight: bold;
    /*background: #303030;*/
  }
  .el-input {
    width: 80%;
    margin: 0 10%;
    padding: 15px;
  }

  .login-option {
    display: flex;
    justify-content: right;
    margin: 2% 10%;
  }
</style>