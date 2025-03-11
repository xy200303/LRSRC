<template>
  <div class="layout-container">
    <!-- 侧边栏 -->
    <Sidebar @select="select" :isCollapsed="isCollapsed" :menu-data="menuData"/>
    <!-- 右侧区域（头部 + 主内容区域 + 脚部） -->
    <div class="right-container">
      <!-- 头部 -->
      <Header :items="dropdownItems" @collapse="handleCollapse" @commond="commond"/>
      <!-- 主内容区域 -->
      <AppMain />
      <!-- 脚部 -->
      <Footer/>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { onMounted, ref } from 'vue';
import { useRouter } from 'vue-router';
import Header from './components/AdminHeader.vue';
import Sidebar from './components/Sidebar.vue';
import AppMain from './components/AppMain.vue';
import Footer from './components/Footer.vue';
import {useUserStore} from "@/stores/userStores"
import usePermissionStore from '../stores/modules/permission';
const router = useRouter();
const isCollapsed=ref(false)
const userStore=useUserStore()

// 定义下拉框数据
const dropdownItems = ref([
    { value: 1, label: '首页' },
    { value: 2, label: '我的信息' },
    { value: 3, label: '退出登录' },
    ]);
    //定义路由数据
const menuData=usePermissionStore().getAdminRoutes()

// 处理菜单点击事件
function select(index: string) {
  router.push(index); // 跳转到对应路由
}

//处理折叠点击事件
function handleCollapse(v:boolean){
	isCollapsed.value=v
}

async function commond(command:number){
	if (command == 1) {
		router.push("/");
	}
	if (command == 2) {
		router.push("/admin/profile");
	}
	if (command == 3) {
	try{
		const res=await userStore.logOut()
		ElMessage.success('退出登录成功');
		router.push("/login")
	  }catch(err){
		ElMessage.error('退出登录失败');
	  }
	}
}

</script>

<style scoped>
.layout-container {
  display: flex;
  min-height: 100vh; /* 确保布局占满整个视口高度 */
  max-width: 100vw;
}


/* 右侧区域样式 */
.right-container {
  flex: 1; /* 占据剩余空间 */
  display: flex;
  flex-direction: column;
  min-width: 0; 
  /* 防止 flex 容器溢出 */
}


</style>