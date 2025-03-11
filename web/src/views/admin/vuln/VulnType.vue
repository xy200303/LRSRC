<template>
    <div>
        <!-- 显示数据页面 -->
        <el-card style="height: 100%;  font-weight: bold;" shadow="always">
        <div style="font-weight: lighter; font-size: 17px;">
            <!-- 搜索 -->
            <div >
                <div style="display: flex; width: 300px; gap: 1%;">
                    <el-select v-model="queryData.parent_id" placeholder="选择漏洞类别">
                        <el-option
                        v-for="item in sys_vuln_cate"
                        :key="item.value"
                        :label="item.label"
                        :value="item.value"
                        />
                    </el-select>
                    <el-button type="primary" :icon="Search" @click="handleSearch">搜索</el-button>
                </div>
            </div>
            <!-- 操作栏 -->
            <div style="margin-top: 20px;">
                <el-button type="primary" @click="handleAddClick" :icon="Plus">添加</el-button>
                <el-button type="danger" @click="handleMulDeleteClick" :icon="Delete" :disabled="multideleteVisible">批量删除</el-button>
            </div>
            <!-- 表格页面 -->
            <el-table :data="tableData" @selection-change="handleSelectionChange">
                <el-table-column type="selection" width="55" />
                <el-table-column prop="id" label="ID" sortable width="55"/>
                <el-table-column prop="cate_type" label="漏洞类别" width="130">
                    <template #default="scope">
                        <dict-tag :options="sys_vuln_cate" :value="scope.row.parent_id" />
                    </template>
                </el-table-column>
                <el-table-column prop="type_name" label="漏洞类型" width="100"/>
                <el-table-column prop="desc" label="描述"/>
                <el-table-column label="创建时间" width="100">
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
        </div>
        </el-card>
        <!-- 对话框显示 -->
        <el-dialog :title="formMode=='add'?'添加数据':'编辑数据'" v-model="dialogVisible" width="30%" @close="handleCloseForm">
            <el-form :model="formData" label-position="left" label-width="auto">
                <el-form-item required label="漏洞类别">
                    <el-select v-model="formData.parent_id" placeholder="选择漏洞类别">
                        <el-option
                        v-for="item in sys_vuln_cate"
                        :key="item.value"
                        :label="item.label"
                        :value="item.value"
                        />
                    </el-select>
                </el-form-item>
                <el-form-item required label="漏洞类型名称">
                    <el-input v-model="formData.type_name"></el-input>
                </el-form-item>
                <el-form-item label="漏洞类型描述">
                    <el-input type="textarea" v-model="formData.desc"></el-input>
                </el-form-item>
            </el-form>
            <span slot="footer" class="dialog-footer">
                <el-button @click="dialogVisible=false">取消</el-button>
                <el-button type="primary" @click="handleSaveForm">保存</el-button>
            </span>
        </el-dialog>
    </div>
</template>
<script lang="ts" setup>
import { ref, computed, onMounted, getCurrentInstance, reactive } from 'vue'
import { formatDate } from '@/utils/datetime'
import {clearValues} from '@/utils'
import { Delete, Edit, Search, Share, Upload,Plus } from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'
import { listVulnType,deleteVulnType,createVulnType,updateVulnType,buildVulnTypeTree } from '@/api/admin/vuln'
const  {proxy}  = getCurrentInstance();
const sys_vuln_cate = ref([
    {value:1,label:"Web漏洞"},
    {value:2,label:"App漏洞"},
    {value:3,label:"Iot漏洞"},
    {value:4,label:"工控漏洞"},
    {value:5,label:"操作系统漏洞"},
])
const dialogVisible=ref(false)
const multipleSelection = ref([])
const multideleteVisible=ref(true)
const tableData=ref([])
const formMode=ref("add")

const queryData=ref({
    "parent_id":1,
})
const formData=ref({
    "parent_id":1,
    "type_name":null,
    "desc":null,
})

onMounted(()=>{
    listData()
})

//搜索
function handleSearch(){
    listData()
}
//查看数据
function listData(){
    listVulnType({
        ...queryData.value
    }).then((res:any)=>{
        tableData.value=res.data
    })
}
//清空FORM数据
function handleCloseForm(){
    dialogVisible.value=false
    formData.value=clearValues(formData.value)
}
//保存数据
function handleSaveForm(){
    if(formMode.value=="add"){
        createVulnType(formData.value).then((res:any)=>{
            ElMessage.success("添加数据成功")
            dialogVisible.value=false
            listData()
        })
    }else{
        updateVulnType(formData.value).then((res:any)=>{
            ElMessage.success("编辑数据成功")
            dialogVisible.value=false
            listData()
        })
    }

}

//删除数据
function deleteData(ids:[]){
    deleteVulnType({
        ids:ids
    }).then((res:any)=>{
        ElMessage.success("删除成功")
        listData()
    })
}
//点击添加
function handleAddClick(){
    dialogVisible.value=true
    formData.value.parent_id=1
    formMode.value="add"
}
//点击删除
function handleDeleteClick(row:any){
    ElMessageBox.confirm('确定要删除该漏洞类型吗？', '提示', {
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
    dialogVisible.value=true
    formData.value=row
    formMode.value="update"
}
//点击批量删除
function handleMulDeleteClick(){
    ElMessageBox.confirm('确定要删除该漏洞类型吗？', '提示', {
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