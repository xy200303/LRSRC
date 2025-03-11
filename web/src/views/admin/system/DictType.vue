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
                <el-button type="primary" @click="handleAddClick" :icon="Plus">添加</el-button>
                <el-button type="danger" @click="handleMulDeleteClick" :icon="Delete" :disabled="multideleteVisible">批量删除</el-button>
            </div>
            <!-- 表格页面 -->
            <el-table :data="tableData" @selection-change="handleSelectionChange">
                <el-table-column type="selection" width="55" />
                <el-table-column prop="dict_type" label="字典类型" sortable >
                    <template #default="scope">
                        <router-link :to="'/admin/system/dictData?dict_type=' + scope.row.dict_type" class="link-type">
                            <span>{{ scope.row.dict_type }}</span>
                        </router-link>
                    </template>
                </el-table-column>
                <el-table-column prop="name" label="字典名称" />
                <el-table-column prop="type" label="类型" >
                    <template #default="scope">
                        <dict-tag :options="sys_dict_type" :value="scope.row.type" />
                    </template>
                </el-table-column>
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
        <!-- 对话框显示 -->
        <el-dialog :title="formMode=='add'?'添加数据':'编辑数据'" v-model="dialogVisible" width="30%" @close="handleCloseForm">
            <el-form :model="formData" label-position="left" label-width="auto">
                <el-form-item required label="字典类型">
                    <el-input v-model="formData.dict_type" ></el-input>
                </el-form-item>
                <el-form-item required label="字典名称">
                    <el-input v-model="formData.name"></el-input>
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
import { ref, computed, onMounted, getCurrentInstance } from 'vue'
import { formatDate } from '@/utils/datetime'
import {clearValues} from '@/utils'
import { Delete, Edit, Search, Share, Upload,Plus } from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'
import { deleteDictType, listDictType, updateDictType ,createDictTyp,listDictDatae} from '@/api/admin/dict'
import { createDictType } from '@/api/admin/dict'
const  {proxy}  = getCurrentInstance();
const { sys_dict_type } = proxy.useDict("sys_dict_type");
const dialogVisible=ref(false)
const multipleSelection = ref([])
const multideleteVisible=ref(true)
const search=ref("")
const tableData=ref([])
const formMode=ref("add")
const formData=ref({
    "dict_type":null,
    "name":null,
})
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
    listDictType(page.value).then((res:any)=>{
        tableData.value=res.data
        page.value.page=res.page
        page.value.page_size=res.page_size
        page.value.total=res.total
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
        createDictType(formData.value).then((res:any)=>{
            ElMessage.success("添加数据成功")
            dialogVisible.value=false
            listData()
        })
    }else{
        updateDictType(formData.value).then((res:any)=>{
            ElMessage.success("编辑数据成功")
            dialogVisible.value=false
            listData()
        })
    }

}

//删除数据
function deleteData(ids:[]){
    deleteDictType({
        ids:ids
    }).then((res:any)=>{
        ElMessage.success("删除成功")
        listData()
    })
}
//点击添加
function handleAddClick(){
    console.log(sys_dict_type.value)
    dialogVisible.value=true
    formMode.value="add"
}
//点击删除
function handleDeleteClick(row:any){
    ElMessageBox.confirm('确定要删除该字典类型吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
  })
    .then(() => {
      // 用户点击了“确定”按钮
      deleteData([row.dict_type])
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
    ElMessageBox.confirm('确定要删除该字典类型吗？', '提示', {
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
    multipleSelection.value =  val.map(row => row.dict_type);
    if (val.length > 0) {
        multideleteVisible.value = false
    } else {
        multideleteVisible.value = true
    }
}

</script>