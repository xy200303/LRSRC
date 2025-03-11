<template>
    <div>
        <el-tabs v-model="activeName" class="demo-tabs">
            <el-tab-pane label="个人信息" name="profile"></el-tab-pane>
            <el-tab-pane label="修改密码" name="updatePassword"></el-tab-pane>
        </el-tabs>
        <!-- 个人信息 -->
        <el-card style="padding: 30px;" v-if="activeName === 'profile'">
            <el-form ref="userInfoFormRef" :model="userInfoFormData" :rules="userInfoRules" label-width="auto">
                <el-form-item label="用户头像">
                    <upload-image
                        v-model="userInfoFormData.avatar"
                    />
                </el-form-item>
                <!-- 昵称 -->
                <el-form-item label="昵称" prop="nickname" required>
                    <el-input v-model="userInfoFormData.nickname" placeholder="请输入昵称"></el-input>
                </el-form-item>
                <!-- 邮箱 -->
                <el-form-item label="邮箱" prop="email" required>
                    <el-input v-model="userInfoFormData.email" placeholder="请输入邮箱"></el-input>
                </el-form-item>
                <!-- 手机号 -->
                <el-form-item label="手机号" prop="phone" required>
                    <el-input v-model="userInfoFormData.phone" placeholder="请输入手机号"></el-input>
                </el-form-item>
                <!-- 性别 -->
                <el-form-item label="性别" prop="gender">
                    <el-select v-model="userInfoFormData.gender" placeholder="选择性别">
                        <el-option
                            v-for="item in sys_gender_type"
                            :key="item.value"
                            :label="item.label"
                            :value="item.value"
                        />
                    </el-select>
                </el-form-item>
                <!-- 地址 -->
                <el-form-item label="地址" prop="address">
                    <el-input v-model="userInfoFormData.address" placeholder="请输入地址"></el-input>
                </el-form-item>
                <!-- 出生日期 -->
                <el-form-item label="出生日期" prop="birthdate">
                    <el-date-picker
                        v-model="userInfoFormData.birthdate"
                        type="date"
                        placeholder="选择出生日期"
                    ></el-date-picker>
                </el-form-item>
            </el-form>
            <el-button type="primary" @click="handleSaveUserInfoClick">保存</el-button>
        </el-card>
        <!-- 修改密码 -->
        <el-card v-if="activeName === 'updatePassword'">
            <el-form ref="changePasswordFormRef" :model="changePasswordFormData" :rules="changePasswordRules" label-width="auto">
                <el-form-item label="原始密码" prop="old_password" required>
                    <el-input v-model="changePasswordFormData.old_password" placeholder="请输入旧密码"></el-input>
                </el-form-item>
                <!-- 新密码 -->
                <el-form-item label="新密码" prop="new_password" required>
                    <el-input v-model="changePasswordFormData.new_password" placeholder="请输入新密码"></el-input>
                </el-form-item>
                <el-form-item label="确认新密码" prop="cfm_new_password" required>
                    <el-input v-model="changePasswordFormData.cfm_new_password" placeholder="请确认新密码"></el-input>
                </el-form-item>
            </el-form>
            <el-button type="primary" @click="handleChangePasswordClick">修改密码</el-button>
        </el-card>
    </div>
</template>
<script lang="ts" setup>
import { ref, onMounted, getCurrentInstance } from 'vue'
import { useI18n } from 'vue-i18n';
import { useUserStore } from '@/stores/userStores';
import { changePassword, updateProfile } from '../../api/login';
import { ElMessage } from 'element-plus';

const { t } = useI18n();
const { proxy } = getCurrentInstance();
const { sys_user_status, sys_gender_type } = proxy.useDict("sys_user_status", "sys_gender_type");
const userStore = useUserStore()
const userInfo = userStore.userInfo
const activeName = ref('profile')
const userInfoFormData = ref({
    avatar: userInfo.avatar,
    nickname: userInfo.nickname,
    email: userInfo.email,
    phone: userInfo.phone,
    gender: userInfo.gender,
    birthdate: userInfo.birthdate,
    address: userInfo.address
})
// 修改密码表单
const changePasswordFormData = ref({
    old_password: '',
    new_password: '',
    cfm_new_password: ''
})
// 个人信息表单验证规则
const userInfoRules = ref({
    nickname: [
        { required: true, message: '请输入昵称', trigger: 'blur' }
    ],
    email: [
        { required: true, message: '请输入邮箱', trigger: 'blur' },
        { type: 'email', message: '请输入正确的邮箱地址', trigger: ['blur', 'change'] }
    ],
    phone: [
        { required: true, message: '请输入手机号', trigger: 'blur' },
    ]
})
// 修改密码表单验证规则
const changePasswordRules = ref({
    old_password: [
        { required: true, message: '请输入原始密码', trigger: 'blur' }
    ],
    new_password: [
        { required: true, message: '请输入新密码', trigger: 'blur' },
        { min: 6, message: '密码长度不能少于6位', trigger: 'blur' }
    ],
    cfm_new_password: [
        { required: true, message: '请确认新密码', trigger: 'blur' },
        {
            validator: (rule, value, callback) => {
                if (value !== changePasswordFormData.value.new_password) {
                    callback(new Error('两次输入的密码不一致'));
                } else {
                    callback();
                }
            },
            trigger: 'blur'
        }
    ]
})

// 定义表单引用
const userInfoFormRef = ref(null)
const changePasswordFormRef = ref(null)

// 更新用户信息
function handleSaveUserInfoClick() {
    userInfoFormRef.value.validate((valid) => {
        if (valid) {
            updateProfile({
                ...userInfoFormData.value
            }).then((res) => {
                ElMessage.success("更新成功")
                userStore.getMyProfile().then((res) => {

                }).catch((err) => {

                })
            }).catch((err) => {

            })
        } else {
            ElMessage.error('表单验证失败，请检查输入内容')
            return false
        }
    })
}

// 更新密码
function handleChangePasswordClick() {
    changePasswordFormRef.value.validate((valid) => {
        if (valid) {
            changePassword({
                ...changePasswordFormData.value
            }).then((res) => {
                ElMessage.success("修改密码成功")
                userStore.removeUserInfo()
            }).catch((err) => {
            })
        } else {
            ElMessage.error('表单验证失败，请检查输入内容')
            return false
        }
    })
}
</script>