<template>
  <div>
    <template v-for="(item, index) in res">
      <template v-if="item.value==value">
        <span
          v-if="(item.el_tag_type == 'default' || item.el_tag_type == '' || item.el_tag_effect == null) && (item.el_tag_effect == '' || item.el_tag_effect == null)"
          :key="item.value"
          :index="index"
        >{{ item.label + " " }}</span>
        <el-tag
          v-else
          :disable-transitions="true"
          :key="item.value + ''"
          :index="index"
          :type="item.el_tag_type === 'primary' ? '' : item.el_tag_type"
          :effect="item.el_tag_effect"
        >{{ item.label + " " }}</el-tag>
      </template>
    </template>
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue';


const props = defineProps({
  // 数据
  options: {
    type: Array,
    default: null,
  },
  // 当前的值
  value: [Number, String, Array],
  // 当未找到匹配的数据时，显示value
  showValue: {
    type: Boolean,
    default: true,
  }
});
  const setOptions=(r)=>{
      res.value=r
    }
  defineExpose({
    setOptions
});
const res=ref(props.options)
// 监听 options 的变化
watch(() => props.options, (newVal, oldVal) => {
  res.value=newVal
}, { deep: true });
</script>

<style scoped>
.el-tag + .el-tag {
  margin-left: 10px;
}
</style>
