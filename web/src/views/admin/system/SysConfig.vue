<template>
    <el-tabs v-model="activeName" class="demo-tabs" @tab-click="handleClick">
      <!-- 基本配置 -->
      <el-tab-pane label="基本配置" name="first">
        <el-card>
          <template #header>
          <div class="card-header">
              <span>登录/注册配置</span>
          </div>
          </template>
          <el-form label-position="left" label-width="auto">
              <el-form-item label="是否启用注册">
                  <el-switch v-model="sysConfigMap.sys_register_enable"></el-switch>
              </el-form-item>
          </el-form>
        </el-card>
        <!-- 添加保存按钮 -->
        <el-form-item style="margin-top: 20px;">
          <el-button type="primary" @click="saveSettings">保存设置</el-button>
        </el-form-item>
      </el-tab-pane>

      <!-- 文件上传配置 -->
      <el-tab-pane label="文件配置" name="second">
        <el-card>
            <template #header>
            <div class="card-header">
                <span>文件存储配置</span>
            </div>
            </template>
            <el-form label-position="left" label-width="auto">
            <el-form-item label="文件存储方式">
            <el-select v-model="sysConfigMap.sys_file_storage" placeholder="Select">
                <el-option
                v-for="item in sys_file_storage"
                :key="item.value"
                :label="item.label"
                :value="item.value"
                />
            </el-select>

            </el-form-item>
                <el-form-item label="文件上传限制大小(MB)">
                <el-input-number v-model="sysConfigMap.sys_file_max_size" :min="1"></el-input-number>
                
            </el-form-item>
        </el-form>
        </el-card>

        <!-- OSS配置 -->
        <el-card style="margin-top: 20px;" v-if="sysConfigMap.sys_file_storage!='local'">
            <template #header>
            <div class="card-header">
                <span>OSS对象存储配置</span>
            </div>
            </template>
           
            <el-form label-position="left" label-width="auto">
        <!-- 阿里云OSS配置 -->
         <div v-if="sysConfigMap.sys_file_storage=='aliyun'">
            <el-form-item label="阿里云OSS AccessKey">
            <el-input v-model="sysConfigMap['sys_aliyun_oss_ak']"></el-input>
            </el-form-item>
            <el-form-item label="阿里云OSS SecretKey">
            <el-input v-model="sysConfigMap['sys_aliyun_oss_sk']"></el-input>
            </el-form-item>
            <el-form-item label="阿里云OSS Bucket">
            <el-input v-model="sysConfigMap['sys_aliyun_oss_bucket']"></el-input>
            </el-form-item>
            <el-form-item label="阿里云OSS Region">
            <el-input v-model="sysConfigMap['sys_aliyun_oss_regin']"></el-input>
            </el-form-item>
         </div>

        <!-- 腾讯云COS配置 -->
         <div v-if="sysConfigMap.sys_file_storage=='tencent'">
            <el-form-item label="腾讯云COS AccessKey">
            <el-input v-model="sysConfigMap['sys_tencent_oss_ak']"></el-input>
            </el-form-item>
            <el-form-item label="腾讯云COS SecretKey">
            <el-input v-model="sysConfigMap['sys_tencent_oss_sk']"></el-input>
            </el-form-item>
            <el-form-item label="腾讯云COS Base URL">
            <el-input v-model="sysConfigMap['sys_tencent_oss_base_url']"></el-input>
            </el-form-item>
         </div>
       
        <!-- 华为云OBS配置 -->
         <div v-if="sysConfigMap.sys_file_storage=='huawei'">
            <el-form-item label="华为云OBS AccessKey">
            <el-input v-model="sysConfigMap['sys_huawei_oss_ak']"></el-input>
            </el-form-item>
            <el-form-item label="华为云OBS SecretKey">
            <el-input v-model="sysConfigMap['sys_huawei_oss_sk']"></el-input>
            </el-form-item>
            <el-form-item label="华为云OBS Bucket">
            <el-input v-model="sysConfigMap['sys_huawei_oss_bucket']"></el-input>
            </el-form-item>
            <el-form-item label="华为云OBS Endpoint">
            <el-input v-model="sysConfigMap['sys_huawei_oss_endpoint']"></el-input>
            </el-form-item>
         </div>
       

      </el-form>

        </el-card>
        <!-- 添加保存按钮 -->
        <el-form-item style="margin-top: 20px;">
            <el-button type="primary" @click="saveSettings">保存设置</el-button>
        </el-form-item>
        </el-tab-pane>

        <!-- AI大模型配置 -->
      <el-tab-pane label="AI大模型配置" name="ai">
        <el-card>
          <template #header>
          <div class="card-header">
              <span>AI大模型配置</span>
          </div>
          </template>
          <el-form label-position="left" label-width="auto">
            <el-form-item label="AI类型">
              <el-select v-model="sysConfigMap.sys_ai_type" placeholder="Select">
                <el-option
                v-for="item in sys_ai_type"
                :key="item.value"
                :label="item.label"
                :value="item.value"
                />
            </el-select>
            </el-form-item>
            <el-form-item label="BASE_URL">
              <el-input v-model="sysConfigMap.sys_ai_base_url"></el-input>
            </el-form-item>
            <el-form-item label="API_KEY">
              <el-input v-model="sysConfigMap.sys_ai_api_key"></el-input>
            </el-form-item>
            <el-form-item label="模型名称">
              <el-select v-model="sysConfigMap.sys_ai_model" placeholder="Select">
                <el-option
                v-for="item in sys_ai_model"
                :key="item.value"
                :label="item.label"
                :value="item.value"
                />
            </el-select>
            </el-form-item>
            <el-form-item label="单词对话最大Tokens">
              <el-input-number v-model="sysConfigMap.sys_ai_max_per_tokens"></el-input-number>
            </el-form-item>
          </el-form>
        </el-card>
        <!-- 添加保存按钮 -->
        <el-form-item style="margin-top: 20px;">
          <el-button type="primary" @click="saveSettings">保存设置</el-button>
        </el-form-item>
      </el-tab-pane>

      <!-- 邮箱配置 -->
      <el-tab-pane label="邮箱配置" name="thirty">
        <el-card>
            <template #header>
            <div class="card-header">
                <span>SMTP配置</span>
            </div>
            </template>
            <el-form label-position="left" label-width="auto" autocomplete="off">
                <el-form-item label="SMTP主机">
                <el-input v-model="sysConfigMap.sys_smtp_host"></el-input>
                </el-form-item>
                <el-form-item label="SMTP端口">
                <el-input-number v-model="sysConfigMap.sys_smtp_port" :min="1" :max="65535"></el-input-number>
                </el-form-item>
                <el-form-item label="SMTP用户名">
                <el-input v-model="sysConfigMap.sys_smtp_username" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item label="SMTP密码">
                <el-input v-model="sysConfigMap.sys_smtp_password" show-password autocomplete="new-password"></el-input>
                </el-form-item>
                <el-form-item label="发件人">
                <el-input v-model="sysConfigMap.sys_smtp_sender"></el-input>
                </el-form-item>
            </el-form>
        </el-card>
        <!-- 添加保存按钮 -->
        <el-form-item style="margin-top: 20px;">
            <el-button type="primary" @click="saveSettings">保存设置</el-button>
        </el-form-item>
        
       </el-tab-pane>
      
      <el-tab-pane label="邮箱模板" name="fourth">
        <el-card>
            <template #header>
            <div class="card-header">
                <span>邮箱模板配置</span>
            </div>
            </template>
            <el-form label-position="left" label-width="auto">
              <el-form-item label="验证码标题">
                  <el-input v-model="sysConfigMap.sys_smtp_captcha_title"></el-input>
                </el-form-item>
            </el-form>
        </el-card>
        <!-- 添加保存按钮 -->
        <el-form-item style="margin-top: 20px;">
            <el-button type="primary" @click="saveSettings">保存设置</el-button>
        </el-form-item>
        
      </el-tab-pane>
    </el-tabs>
  </template>
  <script lang="ts" setup>
  import { getCurrentInstance, onMounted, ref } from 'vue'
  import { ElMessage, type TabsPaneContext } from 'element-plus'
import { getSysConfigMap, updateSysConfigMap } from '@/api/admin/config';

  const  {proxy}  = getCurrentInstance();
  const { sys_file_storage,sys_ai_type,sys_ai_model } = proxy.useDict("sys_file_storage","sys_ai_type","sys_ai_model");
  const activeName = ref('first')
  const sysConfigMap=ref({
    sys_aliyun_oss_ak: "",
    sys_aliyun_oss_bucket: "",
    sys_aliyun_oss_regin: "",
    sys_aliyun_oss_sk: "",
    sys_file_max_size: 100,
    sys_file_storage: 'local',
    sys_huawei_oss_ak: "",
    sys_huawei_oss_bucket: "",
    sys_huawei_oss_endpoint: "",
    sys_huawei_oss_sk: "",
    sys_register_enable: true,
    sys_smtp_captcha_html: "",
    sys_smtp_captcha_title:"",
    sys_smtp_host: "",
    sys_smtp_password: "",
    sys_smtp_port: 25,
    sys_smtp_sender: "",
    sys_smtp_username: "",
    sys_tencent_oss_ak: "",
    sys_tencent_oss_base_url: "",
    sys_tencent_oss_sk: "",
    sys_ai_type:"",
    sys_ai_model:"",
    sys_ai_base_url:"",
    sys_ai_api_key:"",
    sys_ai_max_per_tokens:4096,
  })
  function getSysConfig(){
    getSysConfigMap({
      
    }).then((res:any)=>{
       sysConfigMap.value= res.data
    })
  }

  onMounted(()=>{
    getSysConfig()
  })

  function saveSettings(){
    updateSysConfigMap(sysConfigMap.value).then((res)=>{
        ElMessage.success("保存成功")
    })

  }


  </script>
  
  <style>

  </style>