<template>
    <div class="layout-container">
        <HomeHeader :items="dropdownItems" @commond="commond" ></HomeHeader>
        <AppMain></AppMain>
        <Footer></Footer>
    </div>
  </template>
  
<script lang="ts" setup>
import HomeHeader from "./components/HomeHeader.vue"
import AppMain from "./components/AppMain.vue"
import Footer from "./components/Footer.vue"
import { useRouter } from 'vue-router';
import {useUserStore} from "@/stores/userStores"
import { ref } from "vue";

const router = useRouter();
const userStore=useUserStore()
// 定义下拉框数据
// 定义下拉框数据
const dropdownItems = ref([
    { value: 1, label: '个人中心' },
    { value: 3, label: '退出登录' },
]);

// 如果是管理员，则插入“后台管理”选项
if (userStore.userInfo.is_admin) {
    dropdownItems.value.splice(1, 0, { value: 2, label: '后台管理' });
}

async function commond(command:number){
	if (command == 1) {
		router.push("/user/profile");
	}
	if (command == 2) {
		router.push("/admin");
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
    if(command==4){
        router.push("/admin");
    }
}
</script>

<style scoped>
.layout-container {
  display: flex;
  flex-direction: column;
  min-height: 100vh; /* 确保布局占满整个视口高度 */
  max-width: 100vw;
}

</style>