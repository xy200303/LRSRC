<template>
    <div id="app">
      <el-cascader
        :options="options"
        v-model="localSelectedOptions"
        @change="handleChange"
      />
    </div>
  </template>
  
  <script setup>
  import { ref, watch, computed } from 'vue'
  import { categoryData } from 'element-china-category-data'
  import { regionData} from 'element-china-area-data'
  // 定义 props
  const props = defineProps({
    modelValue: {
      type: Array,
      default: () => []
    },
    data:{
      type:String,
      default:"area"
    }
  })
  
  // 定义 emit
  const emit = defineEmits(['update:modelValue',"change"])
  
  const options=ref([])
  // 数据选项
  if (props.data=="area"){
    options.value = regionData
  }else{
    options.value =categoryData
  }
  
  // 本地状态用于双向绑定
  const localSelectedOptions = computed({
    get() {
      return props.modelValue
    },
    set(value) {
      emit('update:modelValue', value)
    }
  })
  
  
  // 处理选择变化
  const handleChange = (value) => {
    emit('update:modelValue', value)
    emit("change",value)
  }
  
  // 监听外部传入的 modelValue 变化
  watch(() => props.modelValue, (newVal) => {
      localSelectedOptions.value = newVal
  }, { immediate: true })
  </script>
  
  <style scoped>
  /* 添加你的样式 */
  </style>