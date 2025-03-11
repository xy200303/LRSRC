<template>
    <el-upload
        class="avatar-uploader"
        :headers="headers"
        :action="uploadFileApi"
        :on-error="handleError"
        :on-success="handleSuccess"
        :before-upload="beforeAvatarUpload"
        :show-file-list="false"
    >
    <img v-if="props.modelValue" :src="props.modelValue" class="avatar" :style="{ width: width, height: height }"/>
    <el-icon v-else class="avatar-uploader-icon" :style="{ width: width, height: height }">
      <Plus />
    </el-icon>
    </el-upload>
</template>

<script lang="ts" setup>
import { reactive, ref, watch } from 'vue'
import { ElMessage } from 'element-plus'
import { Plus } from '@element-plus/icons-vue'
import type { UploadFile, UploadProps } from 'element-plus'
import { getToken } from '@/utils';
// 定义props
const props = defineProps({
    modelValue: {
      type: String,
      default: ''
    },
    width:{
        type:String,
        default:'100px'
    },
    height:{
        type:String,
        default:'100px'
    },
  });
const uploadFileApi=import.meta.env.VITE_APP_API_URL+'/uploadFile'
const headers = reactive({
  Authorization: `Bearer ${getToken()}`
})
const imageUrl = ref(props.modelValue)

// 监听变化以同步到localUrl
watch(() => props.modelValue, (newValue) => {
  imageUrl.value = newValue;
  emit("update:modelValue",imageUrl.value)
});
// 定义emits
const emit = defineEmits(['update:modelValue',"error","success"]);

//上传成功
const handleSuccess: UploadProps['onSuccess'] = (response,uploadFile) => {
    const file_id=response.data.file_id
    const url=import.meta.env.VITE_APP_API_URL+'/downloadFile?id='+file_id
    imageUrl.value = url
    emit("update:modelValue",imageUrl.value)
    emit("success",response)
    ElMessage.success("上传成功")
}
//上传失败
const handleError= (error: Error, uploadFile: UploadFile, uploadFiles: UploadFile) => {
    emit("error",error)
    ElMessage.error("上传失败")
}
//上传前检查
const beforeAvatarUpload: UploadProps['beforeUpload'] = (rawFile) => {
  if (!/^image\/.+$/.test(rawFile.type)) {
    ElMessage.error('必须上传图像格式')
    return false
  } else if (rawFile.size / 1024 / 1024 > 10) {
    ElMessage.error('图像 大小不能超过10MB')
    return false
  }
  return true
}
</script>

<style scoped>
.avatar-uploader .avatar {
  display: block;
}
</style>

<style>
.avatar-uploader .el-upload {
  border: 1px dashed var(--el-border-color);
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
  transition: var(--el-transition-duration-fast);
}

.avatar-uploader .el-upload:hover {
  border-color: var(--el-color-primary);
}

.el-icon.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  text-align: center;
}
</style>