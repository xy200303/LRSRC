<template>
    <el-upload
      v-model:file-list="fileList"
      class="upload-demo"
      action="http://127.0.0.1:8123/api/v1/uploadFile"
      multiple
      :on-preview="handlePreview"
      :on-remove="handleRemove"
      :before-remove="beforeRemove"
      :limit="3"
      :on-success="onSuccess"
      :on-exceed="handleExceed"
      :headers="headers"
    >
      <el-button type="primary">Click to upload</el-button>
      <template #tip>
        <div class="el-upload__tip">
          jpg/png files with a size less than 500KB.
        </div>
      </template>
    </el-upload>
  </template>
  <script lang="ts" setup>
  import { reactive, ref } from 'vue'
  
  import type { UploadFile, UploadFiles, UploadProps, UploadUserFile } from 'element-plus'
import { getToken } from '@/utils/auth'
  // 自定义请求头
const headers = reactive({
  Authorization: `Bearer ${getToken()}`
})

  const fileList = ref<UploadUserFile[]>([
  ])
  
  const onSuccess:UploadProps["onSuccess"]=(response: any, uploadFile: UploadFile, uploadFiles: UploadFiles) => {
    console.log(fileList)
  }

  const handleRemove: UploadProps['onRemove'] = (file, uploadFiles) => {
    console.log(file, uploadFiles)
  }
  
  const handlePreview: UploadProps['onPreview'] = (uploadFile) => {
    console.log(uploadFile)
  }
  
  const handleExceed: UploadProps['onExceed'] = (files, uploadFiles) => {
    ElMessage.warning(
      `The limit is 3, you selected ${files.length} files this time, add up to ${
        files.length + uploadFiles.length
      } totally`
    )
  }
  
  const beforeRemove: UploadProps['beforeRemove'] = (uploadFile, uploadFiles) => {
    return ElMessageBox.confirm(
      `Cancel the transfer of ${uploadFile.name} ?`
    ).then(
      () => true,
      () => false
    )
  }
  </script>