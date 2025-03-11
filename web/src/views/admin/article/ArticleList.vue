<template>
    <div>
        <!-- 显示数据页面 -->
        <el-card style="height: 100%;  font-weight: bold;" shadow="always">
        <div style="font-weight: lighter; font-size: 17px;">
            <!-- 搜索 -->
            <div >
                <div style="display: flex; width: 80%; gap: 1%;">
                    <el-input v-model="search" placeholder="搜索" clearable style="width: 30%;" />
                     <el-button type="primary" :icon="Search">搜索</el-button>
                </div>
            </div>
            <!-- 操作栏 -->
            <div style="margin-top: 20px;">
                <el-button type="danger" @click="handleMulDeleteClick" :icon="Delete" :disabled="multideleteVisible">批量删除</el-button>
            </div>
            <!-- 表格页面 -->
            <el-table :data="tableData" @selection-change="handleSelectionChange">
                <el-table-column type="selection" width="55" />
                <el-table-column prop="id" label="文章编号" />
                <el-table-column prop="title" label="文章标题" >
                    <template #default="scope">
                        <router-link :to="'/articleDetail?id=' + scope.row.id" class="link-type">
                            <span>{{ scope.row.title }}</span>
                        </router-link>
                    </template>
                </el-table-column>
                
                <el-table-column prop="type" label="类型" >
                    <template #default="scope">
                        <dict-tag :options="sys_article_type" :value="scope.row.type" />
                    </template>
                </el-table-column>
                <el-table-column prop="like_count" label="点赞数" />
                <el-table-column prop="view_count" label="阅读数" />
                <el-table-column prop="created_by" label="创建者" />
                <el-table-column label="创建时间">
                    <template #default="{ row }">
                        {{ formatDate(row.created_at) }}
                    </template>
                </el-table-column>
                <el-table-column label="操作">
                    <template #default="scope">
                        <div style="display: flex; gap: 8px;">
                            <el-button link size="small" type="primary" @click="handleEditClick(scope.row)" :icon="Edit">编辑</el-button>
                            <el-button link size="small" type="danger" @click="handleDeleteClick(scope.row)" :icon="Delete">删除</el-button>
                        </div>
                    </template>
                </el-table-column>
            </el-table>
            <el-pagination style="margin-top: 20px;margin-left: 20px;"
                :page-sizes="[5, 10, 15,20]"
                layout="sizes, prev, pager, next"
                :total="page.total"
                v-model:currentPage="page.page"
                v-model:page-size="page.page_size"
                @size-change="listData"
                @current-change="listData"
            />
        </div>
        </el-card>
    </div>
</template>
<script lang="ts" setup>
import { ref, onMounted, getCurrentInstance, reactive } from 'vue'
import { formatDate } from '@/utils/datetime'
import { Delete, Edit, Search, Share, Upload,Plus } from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'
const  {proxy}  = getCurrentInstance();
const { sys_article_type } = proxy.useDict("sys_article_type");
import { listArticle,deleteArticle} from '@/api/admin/article'
import { useRouter } from 'vue-router';

const multipleSelection = ref([])
const multideleteVisible=ref(true)
const search=ref("")
const tableData=ref([])
const router = useRouter();
const page=ref({
    "page":1,
    "page_size":5,
    "total":1
})

onMounted(()=>{
    listData()
})


//查看数据
function listData(){
    listArticle(page.value).then((res:any)=>{
        tableData.value=res.data
        page.value.page=res.page
        page.value.page_size=res.page_size
        page.value.total=res.total
    })
}


//删除数据
function deleteData(ids:[]){
    deleteArticle({
        ids:ids
    }).then((res:any)=>{
        ElMessage.success("删除成功")
        listData()
    })
}

//点击删除
function handleDeleteClick(row:any){
    ElMessageBox.confirm('确定要删除该文章吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
  })
    .then(() => {
      // 用户点击了“确定”按钮
      deleteData([row.id])
    });
}
//点击编辑
function handleEditClick(row:any){
    router.push("/admin/article/publishArtcile?id="+row.id)
}
//点击批量删除
function handleMulDeleteClick(){
    ElMessageBox.confirm('确定要删除该文章吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
  })
    .then(() => {
      // 用户点击了“确定”按钮
      deleteData(multipleSelection.value)
    });
}
//选中数据
function handleSelectionChange(val:any){
    multipleSelection.value =  val.map(row => row.id);
    if (val.length > 0) {
        multideleteVisible.value = false
    } else {
        multideleteVisible.value = true
    }
}

</script>