<template>
  <div v-if="source === 'local'" :style="iconStyle">
    <svg :class="svgClass" aria-hidden="true" :fill="props.color" :style="iconStyle">
      <use :xlink:href="iconName" />
    </svg>
  </div>
  <Icon
    v-else
    :icon="icon"
    :style="iconStyle"
  ></Icon>
</template>

<script setup>
import { computed } from "vue";
import { Icon } from "@iconify/vue";

// 定义 props
const props = defineProps({
  // 图标名称
  icon: {
    type: String,
    required: true,
  },
  // 图标来源：'local' 或 'remote'
  source: {
    type: String,
    default: "local", // 默认为本地资源
    validator: (value) => ["local", "remote"].includes(value),
  },
  // 图标大小
  size: {
    type: [String, Number],
    default: "24px",
  },
  // 图标颜色
  color: {
    type: String,
    default: "#000000",
  },
});

// 动态计算图标的样式
const iconStyle = computed(() => ({
  width: typeof props.size === "number" ? `${props.size}px` : props.size,
  height: typeof props.size === "number" ? `${props.size}px` : props.size,
  color: props.color,
  fill:props.color
}));

// 计算本地图标的引用路径
const iconName = computed(() => `#icon-${props.icon}`);

// 动态生成 SVG 类名
const svgClass = computed(() => "svg-icon");

// 错误处理：检查本地图标是否存在
onMounted(() => {
  if (props.source === "local") {
    const iconExists = document.querySelector(iconName.value);
    if (!iconExists) {
      console.warn(`[IconComponent] 未找到本地图标: ${props.icon}`);
    }
  }
});
</script>

<style scoped>
.svg-icon {
  vertical-align: middle;
  display: inline-block;
}
</style>