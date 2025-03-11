<template>
  <div v-if="showSuccess">
    <el-result
        icon="success"
        :title="mode=='add'?'发表成功':'更新成功'"
        sub-title="操作已经成功"
      >
      </el-result>
  </div>
    <div v-else>
        <!-- 文章内容编辑页面 -->
        <el-card>
            <el-form :model="article" :rules="formRules" ref="articleForm" label-position="top">
                <el-form-item label="文章标题" required prop="title">
                    <el-input v-model="article.title" placeholder="请输入标题" maxlength="100" show-word-limit></el-input>
                </el-form-item>
                <el-form-item label="文章内容" prop="content">
                  <editor v-model="article.content" ref="editorRef"/>
                </el-form-item>
            </el-form>
        </el-card>
        <!-- 文章属性设置 -->
        <el-card style="margin-top: 20px;">
            <el-form :model="article" :rules="formRules" ref="articleForm1">
                <!-- 标签 -->
                <el-form-item label="文章标签"  prop="tag_list">
                    <tag-input v-model="article.tag_list"></tag-input>
                </el-form-item>
                <el-form-item label="添加封面" porp="poster_url">
                  <upload-image
                    v-model="article.poster_url"
                  />
                </el-form-item>
                <el-form-item label="添加专栏" prop="group_id">
                  <el-select v-model="article.group_id" placeholder="请选择文章专栏">
                      <el-option
                      v-for="item in articleGroup"
                      :key="item.id"
                      :label="item.name"
                      :value="item.id"
                      />
                  </el-select>
                </el-form-item>
                <el-form-item label="文章类型" prop="type">
                  <el-select v-model="article.type" placeholder="请选择文章类型">
                      <el-option
                      v-for="item in sys_article_type"
                      :key="item.value"
                      :label="item.label"
                      :value="item.value"
                      />
                  </el-select>
                </el-form-item>
                <el-form-item label="添加摘要" prop="desc" required>
                    <el-input :rows="3" type="textarea" maxlength="256" show-word-limit v-model="article.desc" placeholder="摘要:将会在首页列表进行展示，作为搜索的基本条件"></el-input>
                    <!-- 添加的提取摘要按钮 -->
                    <el-button 
                        type="primary" 
                        size="small" 
                        :loading="extractSummaryLoading"
                        @click="extractSummary" 
                        style="margin-top: 10px;">
                        一键提取摘要
                    </el-button>
                  </el-form-item>
                <el-form-item label="限制等级" prop="level_limit">
                  <el-input-number v-model="article.level_limit" :min="0" :max="6" />
                </el-form-item>
            </el-form>

        </el-card>

        <!-- 文章附件 -->
        <el-card style="margin-top: 20px;">
          <el-form :model="article" label-position="top">
              <!-- 标签 -->
              <el-form-item label="附件列表">
                <x-upload-file
                  v-model="article.file_obj_list"
                />
              </el-form-item>
            </el-form>
        </el-card>

        <!-- 发布文章 -->
      <div class="parent-container">
        <el-affix position="bottom" :offset="0" style="position:sticky;bottom:0;">
          <el-card class="toolbar">
            <el-row type="flex" justify="space-between">
              <el-col :span="8">
                <span>共:{{ wordCount }} 字</span>
              </el-col>
              <el-col :span="8" style="text-align: center;">
                <el-button type="primary" @click="publishArticle">{{mode=="add"?'发布文章':'更新文章'}}</el-button>
              </el-col>
            </el-row>
          </el-card>
        </el-affix>
      </div>
    </div>
  </template>

<script lang="ts" setup>
import { computed, getCurrentInstance, onMounted, reactive, ref } from 'vue';
import { listAllArticleGroup } from '@/api/admin/article';
import {extractTextFromHtml} from "@/utils"
import { summaryContent } from '@/api/ai';
import {createArticle} from "@/api/admin/article"
import { ElMessage, FormInstance } from 'element-plus';
import { useRoute } from 'vue-router';
import { getArticle } from '@/api/admin/article';
import { watch } from 'vue';
import { updateArticle } from '@/api/admin/article';
const route=useRoute()
const showSuccess=ref(false)
const  {proxy}  = getCurrentInstance();
const { sys_article_type } = proxy.useDict("sys_article_type");
const articleGroup=ref({
  "name":"",
  "id":""
})
const mode=ref("add")
const extractSummaryLoading=ref(false)
const wordCount = ref(0)
const editorRef=ref()
const article=ref({
    title:"",
    tag_list:[],
    desc:"",
    poster_url:'',
    level_limit:0,
    group_id:null,
    type:'1',
    content:"",
    file_list:null,
    file_obj_list:[]//{name,file_id}
})

// 动态计算 file_list
watch(() => article.value.file_obj_list,(newFileObjList) => {
    // 更新 file_list
    article.value.file_list = newFileObjList.map(fileObj => fileObj.file_id);
  },
  { deep: true }
);

// const tag_list=ref([{name:111}])

//初始化数据
onMounted(()=>{
  const id=route.query.id;
  if(id){
    mode.value="update"
    getArticle({
      id:id
    }).then((res)=>{
      article.value=res.data
      //更新文章内容
      editorRef.value.setContent(article.value.content)
    }).catch(err=>{
      mode.value="add"
    })
  }
})

// 新增验证规则
const formRules = reactive({
  title: [
    { required: true, message: '请输入文章标题', trigger: 'blur' },
    { max: 100, message: '标题长度不能超过100个字符', trigger: 'blur' }
  ],
  content: [
    { 
      required: true,
      validator: (rule: any, value: any, callback: any) => {
        if (wordCount.value==0) {
          callback(new Error('请输入文章内容'))
        } else {
          callback()
        }
      },
      trigger: 'blur'
    }
  ],
  tag_list: [
    { 
      type: 'array',
      required: true,
      validator: (rule: any, value: any, callback: any) => {
        if (article.value.tag_list.length === 0) {
          callback(new Error('请至少添加一个标签'))
        } else {
          callback()
        }
      },
      trigger: 'change'
    }
  ],
  desc: [
    { required: true, message: '请输入文章摘要', trigger: 'blur' },
    { max: 256, message: '摘要长度不能超过256个字符', trigger: 'blur' }
  ]
})
//发布文章
function publishArticle(){
  Promise.all([
    (proxy.$refs.articleForm as FormInstance).validate(),
    (proxy.$refs.articleForm1 as FormInstance).validate()
  ]).then(() => {
    console.log(article);
    if(mode.value=="add"){
        // // 提交逻辑...
        console.log(article.value.file_list)
        createArticle(article.value).then((res)=>{
          ElMessage.success("发表文章成功")
          showSuccess.value=true
        }).catch((err)=>{
        })
    }else{
      updateArticle(article.value).then((res)=>{
          ElMessage.success("更新文章成功")
          showSuccess.value=true
        }).catch((err)=>{
        })
    }

  }).catch(() => {
    return false
  })
}

//获取文章分组
listAllArticleGroup().then((res:any)=>{
    articleGroup.value=res.data
})

//提取摘要
function extractSummary(){
  extractSummaryLoading.value=true
  const c=extractTextFromHtml(article.value.content)
  if(c==""){
    ElMessage.error("没有正文内容，无法提取")
    extractSummaryLoading.value=false
    return
  }
  summaryContent({
    content:c
  }).then(res=>{
    article.value.desc=res.data
    extractSummaryLoading.value=false
  }).catch(err=>{
    extractSummaryLoading.value=false
  })
}


watch(() => article.value.content, (newVal) => {
  const c=extractTextFromHtml(article.value.content)
  wordCount.value=c.replace(/\s+/g, '').length
}, { deep: true });


</script>
<style scoped>
.toolbar {
  /* margin-top: 20px; */
  height: 60px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
}
.parent-container {
  width: 100%;
}

.el-affix {
  width: 100% !important; /* 强制继承父宽度 */
}
</style>
