<template>
    <el-menu
        :collapse="isCollapsed"
        class="el-menu-vertical-demo"
        @select="select"
        :default-active="activeMenu"
    >
      <!-- 遍历顶级菜单项 -->
      <template v-for="item in props.menuData" :key="item.path">
        <SidebarItem :item="item" />
      </template>
    </el-menu>
  </template>
  
  <script lang="ts" setup>
  import SidebarItem from './SidebarItem.vue'; // 引入 SidebarItem 组件
  
  // 定义 props
  const props = defineProps({
  menuData: {
    type: Array,
    required: true,
  },
  isCollapsed: {
    type: Boolean,
    required: true, // 假设这是必须的
  },
  activeMenu: {
    type: String,
    required: true,
  },
});
  // 定义 emit
    const emit = defineEmits(['select']);
    function select(path: any) {
        emit('select', path);
    }
  
  </script>
  
  <style scoped>
    .el-menu-vertical-demo:not(.el-menu--collapse) {
    width: 200px;
    }
  </style>