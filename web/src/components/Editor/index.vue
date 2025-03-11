<template>
    <umo-editor :templates="templates"  ref="editorRef" v-bind="options" @assistant="onAssistant" @file-upload="onFileUpload" @changed="onChanged" />
  </template>
  <script setup>
    import { onMounted, ref } from 'vue'
    import { UmoEditor } from '@umoteam/editor';
    import { uploadFile } from '@/api/file';
    import { ElMessage } from 'element-plus';
    import {OpenAI} from 'openai'
    import { getToken } from '@/utils';
    const emit = defineEmits(["handleUploadFileSuccess","handleUploadFileError","update:modelValue","update:textContent","handle"])
    const editorRef = ref(null);
    const templates = ref([
    {
      title: '工作任务',
      description: '工作任务模板',
      content: '<p>工作任务</p>',
    },
    {
      title: '工作周报',
      description: '工作周报模板',
      content: '<h2>工作任务</h2>',
    },
  ])

  function getBaseURL() {
    let protocol = window.location.protocol; // 获取协议，例如 "http:" 或 "https:"
    let hostname = window.location.hostname; // 获取主机名，例如 "example.com"
    let port = window.location.port; // 获取端口号，如果没有指定端口则为空字符串
    // 构建基础URL
    let baseURL = protocol + "//" + hostname;
    // 如果存在端口号，则添加到URL中
    if (port) {
        baseURL += ":" + port;
    }
    return baseURL;
}

  const onAssistant = async (payload, content) => {
        console.log(payload, content)
        const { command, lang, input, output } = payload
        const langs = {
            'en-US': '英文',
            'zh-CN': '中文',
        }
        const options = {
            stream: true,
            model: '...',
            messages: [
            {
                role: 'system',
                content: `你是一个文档助手，根据用户输入的文本或者HTML内容，以及对应操作指令，生成符合要求的文档内容。要求如下：1.如果指令不是要求翻译内容，请使用${langs[lang]}返回，否则按用户要求翻译的语言返回；2.返回${output === 'rich-text' ? '富文本（HTML）' : '纯文本（剔除内容中的HTML标记）'}格式；3.如果用户输入的指令你不能理解，在返回的内容前加上“[ERROR]: ”，4.除此之外不返回任何其他多余的内容。`,
            },
            {
                role: 'user',
                content: `对以下内容进行：【${command}】操作。\n${input}`,
            },
            ],
        }
        let baseURL=import.meta.env.VITE_APP_API_URL+"/ai"
        if (import.meta.env.VITE_APP_API_URL.indexOf("http") === -1) {
            baseURL=getBaseURL()+baseURL
        }
        const client=new OpenAI({
            baseURL:baseURL,
            defaultHeaders: {
                'Authorization': "Bearer "+getToken()
            },
            apiKey:"",
            dangerouslyAllowBrowser: true
        })
        const completion = await client.chat.completions.create(options)
        const stream = new ReadableStream({
            async start(controller) {
                for await (const chunk of completion) {
                    controller.enqueue(chunk.choices[0]?.delta?.content || '')
                }
                controller.close()
            },
        })
        return stream
    }

    const onFileUpload = async (file) => {
        let formData = new FormData();
        // 添加文件到 formData 对象中
        formData.append('file', file);
        try{
            const res=await uploadFile(formData)
            const url=import.meta.env.VITE_APP_API_URL+"/downloadFile?id="+res.data.file_id
            const r={
                id: res.data.file_id,
                url: url,
            }; 
            emit("handleUploadFileSuccess",r)
            return r;
        }catch(err){
            emit("handleUploadFileError",err)
            return false
        }
    }
    const props = defineProps({
    modelValue: {
      type: String,
      default: ""
    }
    });
    onMounted(()=>{
        editorRef.value.setContent(props.modelValue)
    })

    watch(() => props.modelValue,(newVal) => {
        emit("update:modelValue",newVal)
    },
    { deep: true }
    );
    
    const setContent=(content)=>{
        editorRef.value.setContent(content)
    }
    defineExpose({
        setContent
    });

    function onChanged(){
        emit("update:modelValue",editorRef.value.getContent("html"))
        emit("update:textContent",editorRef.value.getContent("text"))
    }
    
    const options = ref({
        file: {
            allowedMimeTypes: [],
            maxSize: 1024 * 1024 * 10, // 10M
        },
            assistant: {
            enabled: true,
            maxlength: 100,
            commands: [
                    {
                        label: { en_US: 'Continuation', zh_CN: '续写', ru_RU: 'Продолжение' },
                        value: { en_US: 'Continuation', zh_CN: '续写', ru_RU: 'Продолжение' },
                    },
                    {
                        label: { en_US: 'Rewrite', zh_CN: '重写', ru_RU: 'Переписать' },
                        value: { en_US: 'Rewrite', zh_CN: '重写', ru_RU: 'Переписать' },
                    },
                    {
                        label: { en_US: 'Abbreviation', zh_CN: '缩写', ru_RU: 'Аббревиатура' },
                        value: { en_US: 'Abbreviation', zh_CN: '缩写', ru_RU: 'Аббревиатура' },
                    },
                    {
                        label: { en_US: 'Expansion', zh_CN: '扩写', ru_RU: 'Расширение' },
                        value: { en_US: 'Expansion', zh_CN: '扩写', ru_RU: 'Расширение' },
                    },
                    {
                        label: { en_US: 'Polish', zh_CN: '润色', ru_RU: 'Польский' },
                        value: { en_US: 'Polish', zh_CN: '润色', ru_RU: 'Польский' },
                    },
                    {
                        label: { en_US: 'Proofread', zh_CN: '校阅', ru_RU: 'Корректура' },
                        value: { en_US: 'Proofread', zh_CN: '校阅', ru_RU: 'Корректура' },
                    },
                    {
                        label: { en_US: 'Translate', zh_CN: '翻译', ru_RU: 'Перевести' },
                        value: { en_US: 'Translate to chinese', zh_CN: '翻译成英文', ru_RU: 'Перевести на китайский' },
                        autoSend: false,
                    },
            ],
        },  
    });
  </script>