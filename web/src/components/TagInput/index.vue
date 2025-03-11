<template>
    <div class="tags-container">
      <el-tag
        v-for="tag in props.modelValue"
        :key="tag.name"
        closable
        :disable-transitions="false"
        @close="handleClose(tag)"
      >
        {{ tag.name }}
      </el-tag>
      <el-input
        v-if="inputVisible"
        ref="InputRef"
        v-model="inputValue"
        size="small"
        @keyup.enter="handleInputConfirm"
        @blur="handleInputConfirm"
        style="width: 100px; margin-left: 10px;"
      />
      <el-button v-else class="button-new-tag" size="small" @click="showInput" style="margin-left: 10px;">
        添加标签
      </el-button>
    </div>
  </template>
  
  <script setup>
  import { reactive, ref, nextTick, watch } from 'vue';
  // Props
  const props = defineProps({
    modelValue: {
      type: Array,
      default: () => []
    }
  });
  
  // Emits
  const emit = defineEmits(['update:modelValue']);
  
  // Reactive state
  const tag_list = reactive(props.modelValue);
  
  // Refs
  const inputVisible = ref(false);
  const inputValue = ref('');
  const InputRef = ref(null);
  watch(() => props.modelValue,(newVal) => {
    tag_list.splice(0, tag_list.length, ...newVal); // 更新内部状态
    emit("update:modelValue",tag_list)
  },
  { deep: true }
);
  
  // Methods
  const handleClose = (tag) => {
    const index = tag_list.indexOf(tag);
    if (index > -1) {
      tag_list.splice(index, 1);
      emit('update:modelValue', tag_list);
    }
  };
  
  const showInput = () => {
    inputVisible.value = true;
    nextTick(() => {
      InputRef.value.focus();
    });
  };
  
  const handleInputConfirm = () => {
    if (inputValue.value) {
      tag_list.push({name:inputValue.value});
      emit('update:modelValue', tag_list);
    }
    inputVisible.value = false;
    inputValue.value = '';
  };
  </script>
  
  <style scoped>
  .tags-container {
    display: flex;
    flex-wrap: wrap;
  }
  
  .el-tag + .el-tag {
    margin-left: 10px;
  }
  
  .button-new-tag {
    margin-left: 10px;
    height: 30px;
    line-height: 30px;
    padding-top: 0;
    padding-bottom: 0;
  }
  
  .input-new-tag {
    width: 90px;
    margin-left: 10px;
  }
  </style>