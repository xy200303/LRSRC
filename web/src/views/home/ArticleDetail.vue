<template>
  <div v-if="article.id">
    <el-card>
      <!-- 文章主体 -->
      <div>
          <!-- 标题 -->
          <h1 class="article-title">{{ article.title }}</h1>
          <!-- 元信息 -->
          <div class="article-meta">
            <span>作者：{{ article.created_by || '未知' }}</span>
            <el-divider direction="vertical" />
            <span>
              时间:{{ formatDate(article.created_at) }}
            </span>
            <el-divider direction="vertical" />
            <span>阅读量：{{ article.view_count || 0 }}</span>
          </div>
              <!-- 专栏和标签 -->
            <div class="article-meta">
              <span>专栏：{{ articleGroupInfo.name || '未分类' }}</span>
              <el-divider direction="vertical" />
                <span>文章标签:</span>
                  <el-tag
                  v-for="(tag, index) in article.tag_list"
                  :key="index"
                  type="primary"
                  style="margin-right: 5px;"
                >
                {{ tag.name }}
              </el-tag>
            </div>
          <!-- 文章内容 -->
          <div 
            class="umo-editor-content show-bookmark content-html" 
            v-html="article.content"
          ></div>
      </div>
    </el-card>
    <!-- 文章附件 -->
    <el-card style="margin-top: 20px;" v-if="article.file_obj_list!=null&&article.file_obj_list?.length!=0">
      <el-table
        :data="article.file_obj_list"
        style="width: 100%"
      >
      <el-table-column prop="name" label="文件名称" />
      <el-table-column prop="file_id" label="校验值" />
      <el-table-column prop="created_at" label="时间" >
        <template #default="{ row }">
          {{ formatDate(row.created_at) }}
        </template>
      </el-table-column>
      <el-table-column label="操作">
        <template #default="scope">
            <div style="display: flex; gap: 8px;">
                <el-button link size="small" type="primary" @click="downloadFile(scope.row.file_id)" :icon="Download">下载</el-button>
            </div>
        </template>
      </el-table-column>
      </el-table>
    </el-card>
    <!-- 评论展示页 -->
    <el-card style="margin-top: 20px;">
      <Comments v-model="comments" @tabChange="handleTabChange" @sendComment="handleSendArticleComment" @deleteComment="handleDeleteComment"></Comments>
    </el-card>

    <!-- 文章属性栏目 -->
    <!-- <div class="parent-container">
        <el-affix position="bottom" :offset="0" style="position:sticky;bottom:0;">
          
          <el-card class="toolbar">
            <div class="action-wrapper">
                <div class="action-item" @click="handleLike">
                    <Icon icon="mdi:like" width="23" :color="article.isLiked ? '#409EFF' : '#606266'" />
                    <span :class="{ 'active': article.isLiked }">{{ article.like_count }}</span>
                </div>
                <div class="action-item" @click="handleUnlike">
                    <Icon icon="subway:unlike" width="23" :color="article.isUnliked ? '#F56C6C' : '#606266'" />
                    <span :class="{ 'active-unlike': article.isUnliked }">{{ article.unlike_count }}</span>
                </div>
            </div>
          </el-card>
        </el-affix>
    </div> -->
  </div>
  <el-empty :image-size="200" v-else/>
  </template>
  
<script setup lang="ts">
import { getCurrentInstance, nextTick, onMounted, reactive, ref, watch } from 'vue';
import { useRoute } from 'vue-router';
import { getArticle } from '@/api/admin/article';
import { formatDate } from '@/utils/datetime'
import { Download } from '@element-plus/icons-vue';
import { getArticleGroup,listArticleComment,sendArticleComment } from '@/api/admin/article';
import { deleteMyArticleComment } from '@/api/admin/article';
import { ElMessage } from 'element-plus';
import { Icon } from "@iconify/vue";
const route=useRoute()

const articleGroupInfo=ref({
  name:"未分类"
})
const  {proxy}  = getCurrentInstance();
const { sys_article_type } = proxy.useDict("sys_article_type");
const article=ref({
    id:0,
    title:"",
    tag_list:[],
    desc:"",
    poster_url:'',
    level_limit:0,
    group_id:0,
    type:'1',
    content:"",
    created_at:"",
    created_by:"",
    view_count:0,
    like_count:0,
    unlike_count:0,
    file_list:[],
    file_obj_list:[],//{name,file_id}
    isLiked:false,
    isUnliked:false
})

const page=ref({
  page:1,
  page_size:20
})

const comments = ref([

]);

const query_type=ref("new")



const handleLike=()=>{
  article.value.isLiked=true
}
const handleUnlike=()=>{
  article.value.isUnliked=true
  

}

//删除评论数据
const handleDeleteComment = (comment_id) => {
    ElMessageBox.confirm(
        '确定要删除这条评论吗？',
        '提示',
        {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning',
        }
    )
    .then(() => {
        deleteMyArticleComment({
          ids:[comment_id]
        }).then((res)=>{
          ElMessage.success("删除成功")
          listArticleCommentData()
        }).catch((err)=>{
        })
    })
    .catch(() => {
    });
};

//切换评论显示按钮
function handleTabChange(tag_name){
  query_type.value=tag_name
  listArticleCommentData()
}

//显示评论数据
function listArticleCommentData(){
  listArticleComment({
    page:page.value.page,
    page_size:page.value.page_size,
    article_id:article.value.id,
    created_by:"",
    query_type:query_type.value
  }).then((res)=>{
    console.log(res.data)
    comments.value=res.data
  }).catch((err)=>{
  })
}

//发表评论
function handleSendArticleComment(val:string){
  sendArticleComment({
    content:val,
    article_id:article.value.id
  }).then((res)=>{
    ElMessage.success("评论成功")
    listArticleCommentData()
  }).catch((err)=>{
  })
}


//设置视频属性
const addControlsToVideos = (container) => {
  const videos = container.querySelectorAll('video[vnode="true"]');
  videos.forEach(video => {
    if (!video.hasAttribute('controls')) {
      video.setAttribute('controls', '');
    }
  });
};

//初始化数据
onMounted(()=>{
  const id=route.query.id;
  if(id){
    getArticle({
      id:id
    }).then((res)=>{
      article.value=res.data
      // 更新文章内容
      nextTick(() => {
        addControlsToVideos(document.querySelector('.umo-editor-content'));
      });
      //获取文章分组信息
      getArticleGroup({
        id:article.value.group_id
      }).then((res)=>{
        articleGroupInfo.value=res.data
      }).catch(err=>{
      })
      //获取评论信息
      listArticleCommentData()
  })
}
})

function downloadFile(id:string) {
      window.location.href =import.meta.env.VITE_APP_API_URL+"/downloadFile?id="+id
}
  </script>
  
  <style scoped>
  .article-title {
    text-align: center;
    font-size: 2.2rem;
    font-weight: 700;
    color: #333;
    margin-bottom: 20px;
    line-height: 1.3;
  }
  
  .article-meta {
    display: flex;
    justify-content: center;
    align-items: center;
    gap: 15px;
    color: #666;
    font-size: 0.9rem;
    margin-bottom: 20px;
  }

  
  /* 处理内容区域的样式 */
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
  

/* 评论 */
  .comment-card {
  margin-bottom: 10px;
}
.comment-content {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
}
.author {
  font-weight: bold;
}
.time {
  font-size: 12px;
  color: #999;
}

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

.action-wrapper {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 30px;
}

.action-item {
    display: flex;
    align-items: center;
    gap: 8px;
    cursor: pointer;
    transition: all 0.3s;
}

.action-item span {
    font-size: 16px;
    color: #606266;
    transition: color 0.3s;
}

.action-item:hover {
    transform: scale(1.1);
}

.active {
    color: #409EFF !important;
}

.active-unlike {
    color: #F56C6C !important;
}

</style>