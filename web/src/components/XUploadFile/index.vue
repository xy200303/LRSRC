<template>
  <el-upload
    :drag="isDrag"
    v-model:file-list="FileList"
    :headers="headers"
    :action="uploadFileApi"
    :multiple="isMultiple"
    :limit="limit"
    :accept="acceptTypes" 
    :before-upload="beforeUpload" 
    @success="handleUploadSuccess"
    @error="handleUploadError"
    @remove="handleRemove"
  >
    <el-icon class="el-icon--upload" v-if="isDrag"><upload-filled /></el-icon>
    <div class="el-upload__text">
      拖拽文件至此或者 <em>点击上传</em>
    </div>
  </el-upload>
</template>

<script setup lang="ts">
import { ref, defineProps, defineEmits, computed } from 'vue';
import { getToken } from '@/utils';
import { UploadFilled } from '@element-plus/icons-vue';
import { ElMessage, UploadFile, UploadFiles } from 'element-plus';

// 定义props
const props = defineProps({
  isDrag: {
    type: Boolean,
    default: true // 默认启用拖拽
  },
  isMultiple: {
    type: Boolean,
    default: true
  },
  modelValue: {
    type: Array,
    default: () => [] // 默认为空数组
  },
  limit: {
    type: Number,
    default: 10
  },
  acceptTypes: {
    type: String,
    default: '.jpg,.jpeg,.png,.pdf,.zip,.tar,.tag.gz' // 默认允许的文件类型
  },
  maxSize: {
    type: Number,
    default: 10 * 1024 * 1024 // 默认最大文件大小为5MB（单位：字节）
  }
});

// 上传接口地址
const uploadFileApi = computed(() => `${import.meta.env.VITE_APP_API_URL}/uploadFile`);

// 请求头
const headers = computed(() => ({
  Authorization: `Bearer ${getToken()}`
}));

// 文件列表双向绑定
const FileList = computed({
  get: () => {
    return props.modelValue
  },
  set: (val) => emit('update:modelValue', val)
});

// 定义事件
const emit = defineEmits(['update:modelValue',  'remove', 'success', 'error']);

// 文件上传前校验
const beforeUpload = (file: UploadFile) => {
  const allowedTypes = props.acceptTypes.split(','); // 解析允许的文件类型
  // 校验文件大小
  const isLtMaxSize = file.size <= props.maxSize;
  if (!isLtMaxSize) {
    ElMessage.error(`文件大小不能超过 ${props.maxSize / (1024 * 1024)} MB`);
    return false;
  }
  // 校验文件类型
  const fileExtension = `.${file.name.split('.').pop()?.toLowerCase()}`; // 获取文件扩展名
  const isFileTypeAllowed = allowedTypes.includes(fileExtension);
  if (!isFileTypeAllowed) {
    ElMessage.error(`不支持的文件类型：${fileExtension}`);
    return false;
  }
  return true;
};

// 文件上传成功处理
const handleUploadSuccess = (response, file, fileList) => {
  const newFileId = response.data.file_id;
  file.file_id = newFileId;
  ElMessage.success("上传成功");
  emit('update:modelValue', fileList);
  emit("success", file);
};

// 文件上传失败处理
const handleUploadError = (error: Error, uploadFile: UploadFile, uploadFiles: UploadFiles) => {
  ElMessage.error("上传失败");
  emit("error", error);
};

// 文件移除处理
const handleRemove = (file, fileList) => {
  const fileIdToRemove = file.file_id;
  emit('update:modelValue', fileList);
  emit("remove", file);
};
</script>