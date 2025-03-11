<template>
  <div :style="containerStyle">
    <Toolbar
      style="border-bottom: 1px solid #ccc"
      :editor="editorRef"
      :defaultConfig="toolbarConfig"
      :mode="mode"
    />
    <Editor
      :style="{  height: height }"
      v-model="content"
      :defaultConfig="editorConfig"
      :mode="mode"
      @onCreated="handleCreated"
      @paste="handlePaste"
      @onChange="handleContentChange"
    />
  </div>
</template>

<script setup lang="ts">
import '@wangeditor/editor/dist/css/style.css' // 引入 css

import { onBeforeUnmount, ref, shallowRef, onMounted, watch, computed } from 'vue'
import { Editor, Toolbar } from '@wangeditor/editor-for-vue'
import type { IEditorConfig, InsertFnType } from '@wangeditor/editor'
import { uploadFile } from '@/api/file'

  // 定义props
  const props = defineProps({
    modelValue: {
      type: String,
      default: ""
    },
    height:{
      type:String,
      default:"500px"
    },
    z_index:{
      type:Number,
      default:100
    }
  });

  
// 动态计算容器样式
const containerStyle = computed(() => ({
  border: '1px solid #ccc',
  zIndex: props.z_index
}));
const editorRef = shallowRef()
// 内容 HTML
const content=ref(props.modelValue)
watch(() => props.modelValue, (newValue) => {
  content.value=newValue
  emit('update:modelValue', newValue) // 向父组件上报最新内容
});
const mode = 'default'

const toolbarConfig = {}
const editorConfig: Partial<IEditorConfig> = {
  placeholder: '请输入内容...',
  MENU_CONF: {},
}

const emit = defineEmits(['handleUploadFileError', 'handleUploadFileSuccess',"update:modelValue","update:textContent"])

editorConfig.MENU_CONF['uploadImage'] = {
  // 自定义上传
  async customUpload(file: File, insertFn: InsertFnType) {
    handleFileUploadImage(file,insertFn)
  },
}

editorConfig.MENU_CONF['uploadVideo'] = {
  // 自定义插入视频
  async customUpload(file: File, insertFn: InsertFnType) {
    handleFileUploadVideo(file,insertFn)
  },
}

const handleContentChange = (editor) => {
  emit('update:modelValue', editor.getHtml()) // 向父组件上报最新内容
  emit('update:textContent', editor.getText()) // 向父组件上报最新内容
}


// 监听粘贴事件
const handlePaste = (editor, event) => {
  const items = (event.clipboardData || window.clipboardData).items
  let file = null

  for (let i = 0; i < items.length; i++) {
    if (items[i].type.indexOf('image') !== -1) {
      file = items[i].getAsFile()
      if (file) {
        // 调用自定义上传方法
        handleFileUploadImage(file, (url) => {
          editor.insertEmbed(editor.getSelection().index, 'image', url)
        })
      }
    } else if (items[i].type.indexOf('video') !== -1) {
      file = items[i].getAsFile()
      if (file) {
        // 调用自定义上传方法
        handleFileUploadVideo(file, (url) => {
          editor.insertEmbed(editor.getSelection().index, 'video', url)
        })
      }
    }
  }
}

// 组件销毁时，也及时销毁编辑器
onBeforeUnmount(() => {
  const editor = editorRef.value
  if (editor == null) return
  editor.destroy()
})

const handleCreated = (editor) => {
  editorRef.value = editor // 记录 editor 实例，重要！
}

//处理上传图片
function handleFileUploadImage(file, insertFn){
  const formData=new FormData()
  formData.append("file",file)
    uploadFile(formData).then((res)=>{
       const file_id=res.data.file_id
       const url=import.meta.env.VITE_APP_API_URL+"/downloadFile?id="+file_id
       emit('handleUploadFileSuccess', {
          id:file_id,
          url:url
        })
        insertFn(url, "", url)
    }).catch((err)=>{
      emit("handleUploadFileError",err)
    })
}
//处理上传视频
function handleFileUploadVideo(file, insertFn){
  const formData=new FormData()
  formData.append("file",file)
    uploadFile(formData).then((res)=>{
       const file_id=res.data.file_id
       const url=import.meta.env.VITE_APP_API_URL+"/downloadFile?id="+file_id
        emit('handleUploadFileSuccess', {
          id:file_id,
          url:url
        })
        insertFn(url,"", url)
    }).catch((err)=>{
      emit('handleUploadFileError', err)
    })
}
</script>