<template>
    <el-menu class="el-menu" mode="horizontal" :ellipsis="false" :router="true">
      <div style="margin-right: 100px;">
        <img src="@/assets/home_logo.png" style="height: 40px;">
      </div>
      <el-menu-item index="home" >首页</el-menu-item>
      <el-menu-item index="article" >技术培训</el-menu-item>
      <!-- <el-menu-item index="rank" >排行榜</el-menu-item>
      <el-menu-item index="search" >搜索</el-menu-item> -->
       <!-- 用户信息下拉菜单 -->
       <div style="cursor: pointer; margin-left: auto;display: flex;flex-direction: row;align-items: center;">
          <!-- 提交漏洞 -->
          <div style="margin-right: 40px;">
            <el-button @click="handleOnClick" type="primary">提交漏洞</el-button>
          </div>
          <!-- 用户个人信息 -->
          <el-dropdown style="color: #fff;" trigger="click" @command="commond">
            <span style="display: flex; align-items: center;">
              <el-avatar style="margin-right: 10px;" :src="userStore.userInfo.avatar" /> {{ userStore.userInfo.username }}
            </span>
            <template #dropdown>
              <el-dropdown-menu>
                <!-- 动态渲染下拉框内容 -->
                <el-dropdown-item
                    v-for="item in props.items"
                    :key="item.value"
                    :command="item.value"
                  >
                    {{ item.label }}
                  </el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </div>
      
    </el-menu>
  </template>
  
<script lang="ts" setup>
import { ref,onMounted  } from 'vue';
import { useI18n } from 'vue-i18n';

const { t } = useI18n();
import {useUserStore} from "@/stores/userStores"
import { useRouter } from 'vue-router';
const router = useRouter();
const userStore=useUserStore()
// 定义 props 接收父组件传递的下拉框数据
const props = defineProps({
  items: {
    type: Array as () => { value: number; label: string }[],
    required: true,
  },
});
// 定义 emit 事件
const emit = defineEmits(["commond"]);

onMounted(() => {
  userStore.getMyProfile()
});

function handleOnClick(){
  router.push("submitVuln");
}

function commond(command: number) {
  emit("commond",command)
}
</script>
  
<style scoped>
.el-menu {
  align-items: center;
  width: 100%;
  --el-menu-bg-color: #fff;
  --el-menu-text-color: #000;
}
.el-menu--horizontal > .el-menu-item:nth-child(1) {
  margin-right: auto;
}
</style>
