<template>
    <div>
        <!-- 显示数据页面 -->
         <el-card style="margin-bottom: 20px;display: flex;align-items: center;">
                <el-form :inline="true" :model="queryForm"  label-width="auto" >
                    <el-form-item label="标题">
                        <el-input v-model="queryForm.title" placeholder="请输入文章标题" clearable />
                    </el-form-item>
                    <el-form-item label="添加专栏" prop="group_id">
                        <el-select v-model="queryForm.group_id" placeholder="请选择文章专栏" style="width: 150px;">
                            <el-option
                            v-for="item in articleGroup"
                            :key="item.id"
                            :label="item.name"
                            :value="item.id"
                            />
                        </el-select>
                    </el-form-item>
                    <el-form-item label="文章类型" prop="type">
                        <el-select v-model="queryForm.type" placeholder="请选择文章类型" style="width: 150px;">
                            <el-option
                            v-for="item in sys_article_type"
                            :key="item.value"
                            :label="item.label"
                            :value="item.value"
                            />
                        </el-select>
                    </el-form-item>
                    <el-form-item>
                        <el-button type="primary" @click="listData">搜索</el-button>
                        <el-button type="info" @click="resetQueryForm">重置</el-button>
                    </el-form-item>
                </el-form>
         </el-card>
        <div style="font-weight: lighter; font-size: 13px;">
            <!-- 卡片列表 -->
            <el-card v-for="article in tableData" :key="article.id" style="margin-bottom: 10px;">
                    <div style="display: flex; align-items: center;">
                        <!-- 封面 -->
                        <img :src="article.poster_url" alt="文章封面" style="width: 100px; height: 100px; object-fit: cover; margin-right: 16px;">
                        <div>
                            <!-- 文章标题 -->
                            <router-link :to="'/articleDetail?id=' + article.id" class="link-type">
                                <h2 style="margin: 0; margin-bottom: 8px;">{{ article.title }}</h2>
                            </router-link>
                            <!-- 文章摘要 -->
                            <span style="margin: 0; margin-bottom: 8px;">{{ article.desc }}</span>
                            <!-- 文章浏览数、创建时间和作者放在一行 -->
                            <div style="display: flex; gap: 16px;margin-top: 7px;">
                                    <!-- 使用图标替换浏览数文本 -->
                                <span style="display: flex">
                                    <Icon icon="hugeicons:view" style="margin-right: 5px;"></Icon>{{ article.view_count }}
                                </span>
                                <span style="display: flex">
                                    <Icon icon="hugeicons:view" style="margin-right: 5px;"></Icon>{{ formatDate(article.created_at) }}
                                </span>
                                <span> {{ formatDate(article.created_at) }}</span>
                                <span>{{ article.created_by }}</span>
                            </div>
                        </div>
                    </div>
            </el-card>
            <el-pagination style="margin-top: 20px;margin-left: 20px;"
                :page-sizes="[5, 10, 15, 20]"
                layout="sizes, prev, pager, next"
                :total="page.total"
                v-model:currentPage="page.page"
                v-model:page-size="page.page_size"
                @size-change="listData"
                @current-change="listData"
            />
        </div>
    </div>
</template>
<script lang="ts" setup>
import { ref, onMounted, getCurrentInstance } from 'vue'
import { formatDate } from '@/utils/datetime'
const  {proxy}  = getCurrentInstance();
const { sys_article_type } = proxy.useDict("sys_article_type");
import { listHomeArticle} from '@/api/admin/article'
import { Icon } from '@iconify/vue/dist/iconify.js';
import { listHomeAllArticleGroup } from '@/api/admin/article';
import { clearValues } from '@/utils';

const tableData=ref([])
const queryForm=ref({
    "group_id":null,
    "title":null,
    "type":null
})
const page=ref({
    "page":1,
    "page_size":5,
    "total":1
})
const articleGroup=ref({
  "name":"",
  "id":""
})
//获取文章分组
listHomeAllArticleGroup().then((res:any)=>{
    articleGroup.value=res.data
})
onMounted(()=>{
    listData()
})

function resetQueryForm(){
    queryForm.value=clearValues(queryForm.value)
}

//查看数据
function listData(){
    listHomeArticle({
        ...queryForm.value,
        ...page.value
    }).then((res:any)=>{
        tableData.value=res.data
        page.value.page=res.page
        page.value.page_size=res.page_size
        page.value.total=res.total
    })
}
</script>
