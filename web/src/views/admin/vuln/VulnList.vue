<template>
    <div>
        <!-- 显示数据页面 -->
        <el-card style="height: 100%;  font-weight: bold;" shadow="always">
        <div style="font-weight: lighter; font-size: 17px;">
            <!-- 搜索 -->
            <div >
                <div style="display: flex; width: 80%; gap: 1%;">
                    <el-input v-model="queryForm.muna_name" placeholder="厂商关键词" clearable style="width: 30%;" />
                    <el-select v-model="queryForm.status" placeholder="请选择漏洞审核状态">
                      <el-option
                      v-for="item in sys_vuln_status"
                      :key="item.value"
                      :label="item.label"
                      :value="item.value"
                      />
                    </el-select>
                    <el-button @click="listData" type="primary" :icon="Search">搜索</el-button>
                </div>
            </div>
            <!-- 操作栏 -->
            <div style="margin-top: 20px;">
                <el-button type="danger" @click="handleMulDeleteClick" :icon="Delete" :disabled="multideleteVisible">批量删除</el-button>
            </div>
            <!-- 表格页面 -->
            <el-table :data="tableData" @selection-change="handleSelectionChange">
                <el-table-column type="selection" width="55" />
                <el-table-column prop="id" label="漏洞编号" >
                    <template #default="scope">
                        <router-link :to="'/admin/vuln/vulnDetail?id=' + scope.row.id" class="link-type">
                            <span>{{ scope.row.id }}</span>
                        </router-link>
                    </template>
                </el-table-column>
                <el-table-column prop="title" label="漏洞标题" />
                <el-table-column prop="muna_name" label="影响企业" />
                <el-table-column prop="type" label="属性" >
                    <template #default="scope">
                        <dict-tag :options="sys_vuln_attribute" :value="scope.row.attribute" />
                    </template>
                </el-table-column>
                <el-table-column prop="type" label="漏洞等级" >
                    <template #default="scope">
                        <dict-tag :options="sys_vuln_level" :value="scope.row.level" />
                    </template>
                </el-table-column>
                <el-table-column label="漏洞类型">
                    <template #default="{ row }">
                        {{ row.cate_name_obj.type_name}}/{{ row.type_obj.type_name }}
                    </template>
                </el-table-column>
                <el-table-column label="创建时间">
                    <template #default="{ row }">
                        {{ formatDate(row.created_at) }}
                    </template>
                </el-table-column>
                <el-table-column prop="created_by" label="提交者" />
                <el-table-column label="审核状态">
                    <template #default="{ row }">
                        <dict-tag :options="sys_vuln_status" :value="row.status" />
                    </template>
                </el-table-column>
                <el-table-column label="操作" width="150">
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
const { sys_vuln_attribute,sys_vuln_level,sys_vuln_status } = proxy.useDict("sys_vuln_attribute","sys_vuln_level","sys_vuln_status");
import { listVuln,deleteVuln} from '@/api/admin/vuln'
import { useRouter } from 'vue-router';

const multipleSelection = ref([])
const multideleteVisible=ref(true)
const tableData=ref([])
const router = useRouter();
const page=ref({
    "page":1,
    "page_size":5,
    "total":1
})

const queryForm=ref({
    "muna_name":null,
    "type":null,
    "status":null,
    "muna_domain":null
})

onMounted(()=>{
    listData()
})


//查看数据
function listData(){
    listVuln({
        ...page.value,
        ...queryForm.value
    }).then((res:any)=>{
        tableData.value=res.data
        page.value.page=res.page
        page.value.page_size=res.page_size
        page.value.total=res.total
    })
}


//删除数据
function deleteData(ids:[]){
    deleteVuln({
        ids:ids
    }).then((res:any)=>{
        ElMessage.success("删除成功")
        listData()
    })
}

//点击删除
function handleDeleteClick(row:any){
    ElMessageBox.confirm('确定要删除该漏洞吗？', '提示', {
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
    router.push("/admin/vuln/UpdateVuln?id="+row.id)
}
//点击批量删除
function handleMulDeleteClick(){
    ElMessageBox.confirm('确定要删除这些漏洞吗？', '提示', {
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