<template>
  <!-- 如果是顶层节点且需要隐藏 -->
  <template v-if="item.meta.isTopLevel">
    <!-- 不渲染顶层节点，只递归渲染子节点 -->
    <template v-for="child in item.children" :key="getFullPath(child)">
      <SidebarItem :item="child" :basePath="item.path" />
    </template>
  </template>

  <!-- 如果有子菜单，渲染 el-sub-menu -->
  <el-sub-menu v-else-if="item.children && item.children.length > 0" :index="getFullPath(item)">
    <template #title>
      <el-icon v-if="item.meta?.icon">
        <Icon :icon="item.meta.icon" style="font-size: 25px" />
      </el-icon>
      <span>{{ item.meta.title }}</span>
    </template>
    <!-- 递归渲染子菜单 -->
    <template v-for="child in item.children" :key="getFullPath(child)">
      <SidebarItem :item="child" :basePath="getFullPath(item)" />
    </template>
  </el-sub-menu>

  <!-- 如果没有子菜单，渲染 el-menu-item -->
  <el-menu-item v-else-if="!item.meta?.hidden" :index="getFullPath(item)">
    <el-icon v-if="item.meta?.icon">
      <Icon :icon="item.meta.icon" style="font-size: 25px" />
    </el-icon>
    <span>{{ item.meta.title }}</span>
  </el-menu-item>
</template>

<script lang="ts" setup>
import { Icon } from "@iconify/vue";

// 定义 props
const props = defineProps({
  item: {
    type: Object,
    required: true,
  },
  basePath: {
    type: String,
    default: "", // 父级路径，默认为空
  },
});

// 计算完整路径
const getFullPath = (item: any) => {
  return props.basePath ? `${props.basePath}/${item.path}` : item.path;
};
</script>