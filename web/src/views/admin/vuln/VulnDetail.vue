<template>
    <div v-if="vuln.id">
        <el-card>
            <template #header>
              <div style="display: flex;justify-content: space-between;align-items: center;">
                <!-- 漏洞基本信息 -->
                 <div>
                    <div style="display: flex;align-items: center;">
                      <span style="margin-right: 20px;font-size: 1.3rem;">{{ vuln.title }}</span><dict-tag :options="sys_vuln_level" :value="vuln.level" />
                    </div>
                    <span style="font-size: 0.9rem;">时间:{{ formatDate(vuln.created_at) }}| 提交者:{{vuln.created_by}}</span>
                 </div>
                 <!-- 漏洞审核按钮 -->
                  <div style="display: flex;justify-content: center;align-items: center;">
                     <!-- 显示漏洞状态 -->
                      <div>
                        <Icon :icon="vuln.status" source="local" size="60px"/>
                      </div>
                  </div>
              </div>
            </template>
            <div class="vuln-details">
              <el-form :model="vuln">
                <el-form-item label="关联厂商:">
                  <span>{{ vuln.muna_name }}</span>
                </el-form-item>
                <el-form-item label="漏洞编号:">
                  <span>{{ vuln.id }}</span>
                </el-form-item>
                <el-form-item label="漏洞属性:">
                  <dict-tag :options="sys_vuln_attribute" :value="vuln.attribute" />
                </el-form-item>
                <el-form-item label="漏洞类型:">
                  <span>{{vuln.cate_name_obj.type_name}}/{{ vuln.type_obj.type_name }}</span>
                </el-form-item>
              </el-form>
              <!-- 漏洞描述 -->
              <el-form
                  label-position="top"
                  label-width="auto"
                  :model="vuln"
                >
                <el-form-item label="漏洞描述:">
                  <span>{{ vuln.desc }}</span>
                </el-form-item>
                <el-form-item label="漏洞详情:" style="content-html">
                  <span v-html="vuln.detail"></span>
                </el-form-item>
                <el-form-item label="漏洞附件:" style="content-html">
                  <span style="margin-right: 20px;">{{ vuln.attachment_name }}</span>
                  <el-button type="primary" @click="downloadFile(vuln.attachment_id)">下载</el-button>
                </el-form-item>
                <el-form-item label="修复方案:" style="content-html">
                  <span v-html="vuln.repair_suggestion"></span>
                </el-form-item>
                <div v-if="vuln.status!='under_review'">
                  <el-form-item label="审核意见:">
                  <span v-html="vuln.audit_opinion"></span>
                </el-form-item>
                <el-form-item label="审核员:">
                  <span v-html="vuln.auditor"></span>
                </el-form-item>
                </div>
              </el-form>
            </div>
            <!-- 审核按钮 -->
            <template #footer>
              <!-- 显示审核操做 -->
              <div>
                <el-button type="primary" @click="handleAuditClick">审核</el-button>
              </div>
            </template>
        </el-card>
        <!-- 对话框显示 -->
        <el-dialog title="审核漏洞" v-model="dialogVisible" @close="handleCloseForm" width="600" style="padding: 20px;">
            <el-form :model="formData" label-width="auto"  label-position="top" ref="formRef" :rules="rules">
              <el-form-item label="漏洞编号">
                  <span v-html="formData.vuln_id"></span>
              </el-form-item>
              <el-form-item label="漏洞审核状态">
                <el-select v-model="formData.status" placeholder="选择审核状态">
                    <el-option
                      v-for="item in filteredStatus"
                      :key="item.value"
                      :label="item.label"
                      :value="item.value"
                    />
                </el-select>
              </el-form-item>
              <el-form-item label="漏洞等级">
                <el-select v-model="formData.level" placeholder="选择漏洞等级">
                    <el-option
                      v-for="item in sys_vuln_level"
                      :key="item.value"
                      :label="item.label"
                      :value="item.value"
                    />
                </el-select>
              </el-form-item>
              <el-form-item label="漏洞评分">
                <el-input-number placeholder="请输入漏洞评分" :min="1" :max="10" v-model="formData.score"></el-input-number>
              </el-form-item>
              <el-form-item label="漏洞审核意见">
                <el-input placeholder="请输入漏洞审核意见" type="textarea" v-model="formData.audit_opinion"></el-input>
              </el-form-item>
              <!-- 提交和取消按钮 -->
              <el-row>
                  <el-col :span="24" style="text-align: right;">
                    <el-button @click="dialogVisible = false">取消</el-button>
                    <el-button type="primary" @click="handleSubmitForm">提交</el-button>
                  </el-col>
              </el-row>
            </el-form>
        </el-dialog>
    </div>
    <el-empty :image-size="200" v-else/>
</template>

<script setup lang="ts">
import { computed, getCurrentInstance, nextTick, onMounted, reactive, ref, watch } from 'vue';
import { useRoute } from 'vue-router';
import { formatDate } from '@/utils/datetime';
import { getVuln } from '@/api/admin/vuln';
import { auditVuln } from '@/api/admin/vuln';
import { ElMessage } from 'element-plus';

const dialogVisible=ref(false)
const route = useRoute();
const { proxy } = getCurrentInstance();
const { sys_vuln_attribute, sys_vuln_level, sys_vuln_status } = proxy.useDict("sys_vuln_attribute", "sys_vuln_level", "sys_vuln_status");

const vuln = ref({
    id: "",
    muna_name: "",
    title: "",
    desc: "",
    created_by: '',
    muna_domain: "",
    cate_type: 0,
    type: 0, // 子节点
    attribute: "eve",
    detail: "",
    repair_suggestion: "",
    // web 漏洞
    url: "",
    poc: "",
    // 附属信息
    province: "",
    city: "",
    county: "",
    industry: [],
    // 附件
    attachment_id: '',
    attachment_name: '',
    // 漏洞等级
    level: "low",
    created_at: "",
    //审核信息
    audit_opinion:'',
    auditor:'',
    status:'',
    score:0
});

const filteredStatus = computed(() =>
  sys_vuln_status.value.filter((item) => item.value !== "under_review")
);

//审核提交数据
const formData = ref({
  vuln_id: vuln.value.id,
  status: vuln.value.status,
  audit_opinion: vuln.value.audit_opinion, // 默认为空字符串
  level: vuln.value.level,
  score:vuln.value.score
});

//审核漏洞
function handleAuditClick(){
  dialogVisible.value=true
}

//关闭
function handleCloseForm(){
  dialogVisible.value=false
}
//提交漏洞审核
function handleSubmitForm(){
  auditVuln({
    ...formData.value
  }).then((res)=>{
    ElMessage.success("审核提交成功")
    //刷新数据
    location.reload()
  }).catch((err)=>{
  })
}

// 初始化数据
onMounted(() => {
    const id = route.query.id;
    if (id) {
        getVuln({
            id: id
        }).then((res) => {
            vuln.value = res.data;
            formData.value.vuln_id=vuln.value.id
            formData.value.audit_opinion=vuln.value.audit_opinion
            formData.value.status=vuln.value.status=='under_review'?'accepted':vuln.value.status
            formData.value.level=vuln.value.level
            formData.value.score=vuln.value.score
            console.log(vuln.value)
        });
    }
});

function downloadFile(id: string) {
    window.location.href = import.meta.env.VITE_APP_API_URL + "/downloadFile?id=" + id;
}


</script>

<style scoped>
  .content-html {
  
  }
  
  .content-html :deep(img) {
    display: block;
    max-width: 80%;
    height: auto;
    margin: 20px auto;
    border-radius: 4px;
    box-shadow: 0 2px 8px rgba(0,0,0,0.1);
  }
</style>