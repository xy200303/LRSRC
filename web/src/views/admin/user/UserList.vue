<template>
    <div class="main-div">
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
                <el-table :data="tableData" @selection-change="handleSelectionChange" style="width: 100%">
                    <el-table-column type="selection" width="55" />
                    <el-table-column prop="username" label="用户名称" sortable />
                    <el-table-column prop="nickname" label="用户昵称" />
                    <el-table-column prop="email" label="邮箱" />
                    <el-table-column prop="phone" label="手机号" />
                    <el-table-column prop="level" label="等级" sortable />
                    <el-table-column prop="integral" label="积分" sortable />
                    
                    <el-table-column prop="is_admin" label="管理员" >
                        <template #default="scope">
                            <el-switch :disabled="scope.row.username=='admin'" v-model="scope.row.is_admin" @change="handleSwitch(scope.row)"/>
                        </template>
                    </el-table-column>
                    <el-table-column prop="status" label="状态" >
                        <template #default="scope">
                            <dict-tag :options="sys_user_status" :value="scope.row.status" />
                        </template>
                    </el-table-column>
                    <el-table-column label="创建时间">
                        <template #default="{ row }">
                            {{ formatDate(row.created_at) }}
                        </template>
                    </el-table-column>
                    <el-table-column label="操作" fixed="right" min-width="100">
                        <template #default="scope">
                            <div style="display: flex;">
                                <el-button link  size="small" type="primary" @click="handleEditClick(scope.row)" :icon="Edit">编辑</el-button>
                                <el-dropdown>
                                    <el-button link size="small" type="primary" :icon="Operation">更多</el-button>
                                    <template #dropdown>
                                        <el-dropdown-menu>
                                            <el-dropdown-item>
                                                <el-button link  size="small" type="primary" @click="handleDeleteClick(scope.row)" :icon="Lock">重置密码</el-button>
                                            </el-dropdown-item>
                                            <el-dropdown-item>
                                                <el-button link  size="small" type="danger" @click="handleDeleteClick(scope.row)" :icon="Delete">删除</el-button>
                                            </el-dropdown-item>
                                        </el-dropdown-menu>
                                    </template>
                                </el-dropdown>
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
        <el-dialog :title="formMode === 'add' ? '添加数据' : '编辑数据'" v-model="dialogVisible" @close="handleCloseForm" width="600" style="padding: 20px;">
            <el-form :model="formData" label-width="auto" ref="formRef" :rules="rules">
            <el-row :gutter="20">
                <!-- 第一列 -->
                <el-col :span="12">
                <!-- 用户名 -->
                <el-form-item label="用户名" prop="username" placeholder="请输入用户名" required>
                    <el-input v-model="formData.username" :disabled="formMode!='add'"></el-input>
                </el-form-item>

                <!-- 昵称 -->
                <el-form-item label="昵称" prop="nickname" required>
                    <el-input v-model="formData.nickname" placeholder="请输入昵称"></el-input>
                </el-form-item>

                <!-- 邮箱 -->
                <el-form-item label="邮箱" prop="email" required>
                    <el-input v-model="formData.email" placeholder="请输入邮箱"></el-input>
                </el-form-item>


                <!-- 性别 -->
                <el-form-item label="性别" prop="gender">
                    <el-select v-model="formData.gender" placeholder="选择性别">
                        <el-option
                        v-for="item in sys_gender_type"
                        :key="item.value"
                        :label="item.label"
                        :value="item.value"
                        />
                    </el-select>
                </el-form-item>


                </el-col>

                <!-- 第二列 -->
                <el-col :span="12">
                <!-- 状态 -->
                <el-form-item label="状态" prop="status">
                    <el-select v-model="formData.status" placeholder="选择用户状态">
                        <el-option
                        v-for="item in sys_user_status"
                        :key="item.value"
                        :label="item.label"
                        :value="item.value"
                        />
                    </el-select>
                </el-form-item>

                <!-- 地址 -->
                <el-form-item label="地址" prop="address">
                    <el-input v-model="formData.address" placeholder="请输入地址"></el-input>
                </el-form-item>

                <!-- 手机号 -->
                <el-form-item label="手机号" prop="phone" required>
                    <el-input v-model="formData.phone" placeholder="请输入手机号"></el-input>
                </el-form-item>


                <!-- 出生日期 -->
                <el-form-item label="出生日期" prop="birthdate">
                    <el-date-picker
                    v-model="formData.birthdate"
                    type="date"
                    placeholder="选择出生日期"
                    ></el-date-picker>
                </el-form-item>
                </el-col>


            </el-row>
            <el-row>

                <!-- 身份证号码 -->
                <el-form-item label="身份证号码" prop="id_card">
                    <el-input v-model="formData.id_card" placeholder="请输入身份证号码"></el-input>
                </el-form-item>
            </el-row>


            <!-- 提交和取消按钮 -->
            <el-row>
                <el-col :span="24" style="text-align: right;">
                <el-button @click="dialogVisible = false">取消</el-button>
                <el-button type="primary" @click="handleSaveForm">保存</el-button>
                </el-col>
            </el-row>
            </el-form>
        </el-dialog>
    </div>
</template>
<script lang="ts" setup>
import { ref, computed, onMounted, getCurrentInstance, reactive } from 'vue'
import { formatDate } from '@/utils/datetime'
import {listUser,createUser,updateUser,deleteUser,setUserAdmin} from '@/api/admin/user'
import {clearValues} from '@/utils'
import { Delete, Edit, Search, Share, Upload,Plus,Operation,Lock } from '@element-plus/icons-vue'
import { ElMessage, FormInstance, FormRules } from 'element-plus'
const  {proxy}  = getCurrentInstance();
const { sys_user_status,sys_gender_type } = proxy.useDict("sys_user_status","sys_gender_type");

const dialogVisible=ref(false)
const multipleSelection = ref([])
const multideleteVisible=ref(true)
const search=ref("")
const tableData=ref([])
const formMode=ref("add")
// 表单引用
const formRef = ref<FormInstance>()
// 表单验证规则
const rules = reactive<FormRules>({
    username: [
        { required: true, message: '请输入用户名', trigger: 'blur' },
    ],
    nickname: [
        { required: true, message: '请输入昵称', trigger: 'blur' },
    ],
    email: [
        { required: true, message: '请输入邮箱', trigger: 'blur' },
        { type: 'email', message: '请输入有效的邮箱地址', trigger: 'blur' },
    ],
    phone: [
        { required: true, message: '请输入手机号码', trigger: 'blur' },
        { pattern: /^1[3-9]\d{9}$/, message: '请输入有效的手机号码', trigger: 'blur' },
    ],
    gender: [
        { required: false, message: '请选择性别', trigger: 'change' },
    ],
    status: [
        { required: true, message: '请选择用户状态', trigger: 'change' },
    ],
    address: [
        { required: false }, // 如果不需要必填，则可以省略或设置为false
    ],
    id_card: [
        { pattern: /^[1-9]\d{5}(18|19|20)\d{2}(0[1-9]|1[0-2])(0[1-9]|[12]\d|3[01])\d{3}[0-9Xx]$/, message: '请输入有效的身份证号码', trigger: 'blur' },
    ],
})

const formData=ref({
    "username":null,
    "nickname":null,
    "email":null,
    "phone":null,
    "gender":null,
    "birthdate":null,
    "status":"1",
    "address":null,
    "id_card":null,
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
    listUser(page.value).then((res:any)=>{
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
    // 重置表单
    if (formRef.value) {
        formRef.value.resetFields()
    }
}
//保存数据
function handleSaveForm(){
    if (!formRef.value) return
    formRef.value.validate((valid) => {
        if (valid) {
            if(formMode.value=="add"){
                createUser(formData.value).then((res:any)=>{
                    ElMessage.success("添加数据成功")
                    dialogVisible.value=false
                    listData()
                })
            }else{
                updateUser(formData.value).then((res:any)=>{
                    ElMessage.success("编辑数据成功")
                    dialogVisible.value=false
                    listData()
                })
            }
        } else {
            console.log('表单验证失败')
            return false
        }
    })

}

//设置管理员
function handleSwitch(row:any){
    ElMessageBox.confirm('确定要为这个用户'+row.username+(row.is_admin?'设置':'取消')+'管理员权限吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
    })
        .then(() => {
            setUserAdmin({
                is_admin:row.is_admin,
                username:row.username
            }).then((res)=>{
                ElMessage.success(row.is_admin?'设置管理员成功':'取消管理员成功')
                //刷新数据
                listData()
            }).catch((err)=>{
                row.is_admin=!row.is_admin
            });
        }).catch(()=>{
            row.is_admin=!row.is_admin
        });
}

//删除数据
function deleteDictType(ids:[]){
    deleteUser({
        ids:ids
    }).then((res:any)=>{
        ElMessage.success("删除成功")
        listData()
    })
}
//点击添加
function handleAddClick(){
    dialogVisible.value=true
    formMode.value="add"
}
//点击删除
function handleDeleteClick(row:any){
    ElMessageBox.confirm('确定要删除该用户吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
  })
    .then(() => {
      // 用户点击了“确定”按钮
      deleteDictType([row.username]);
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
    ElMessageBox.confirm('确定要删除所选用户吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
  })
    .then(() => {
      // 用户点击了“确定”按钮
      deleteDictType(multipleSelection.value);
    });
}
//选中数据
function handleSelectionChange(val:any){
    multipleSelection.value =  val.map(row => row.username);
    if (val.length > 0) {
        multideleteVisible.value = false
    } else {
        multideleteVisible.value = true
    }
}

</script>

<style scoped>

</style>