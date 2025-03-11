<template>
  <div>
    <!-- 系统标题和 Logo -->
    <div style="height:60px;display: flex; justify-content: center; align-items: center; color: #fff;background-color: #393939;">
      <img src="@/assets/logo.png" style="height: 40px;">
      <span style="margin-left: 1vw;" v-if="!isCollapsed">零日信安</span>
    </div>
    <SidebarTree :menuData="sortedMenuData" :activeMenu="activeMenu" :isCollapsed="isCollapsed" @select="select" />
  </div>
</template>

<script lang="ts" setup>
import { computed } from 'vue';
import { useRoute } from 'vue-router';

const emit = defineEmits(['select']);
const props = defineProps({
  menuData: {
    type: Array,
    required: true,
  },
  isCollapsed: {
    type: Boolean,
    required: true,
  },
});

// 排序函数：根据 meta.sort 的值对菜单数据进行排序
const sortMenuData = (menuData: any[]) => {
  return menuData
    .map(item => {
      // 如果有子菜单，递归排序子菜单
      if (item.children && item.children.length > 0) {
        item.children = sortMenuData(item.children);
      }
      return item;
    })
    .sort((a, b) => {
      // 根据 meta.sort 的值进行升序排序
      const sortA = a.meta?.sort ?? Infinity; // 如果没有 meta.sort，默认排到最后
      const sortB = b.meta?.sort ?? Infinity;
      return sortA - sortB;
    });
};

// 计算属性：返回排序后的菜单数据
const sortedMenuData = computed(() => {
  return sortMenuData(props.menuData);
});
const route = useRoute();
// 获取当前激活的菜单项
const activeMenu = computed(() => {
  return route.path;
});
// 菜单项点击事件
const select = (path: string) => {
  emit("select", path);
};
</script>

<style scoped>
.el-menu-vertical-demo:not(.el-menu--collapse) {
  width: 200px;
}
</style>