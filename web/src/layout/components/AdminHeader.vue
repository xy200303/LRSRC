<template>
  <div>
    <el-container style="width: 100%;">
      <!-- 头部 -->
      <el-header style="background-color: #393939; display: flex; align-items: center; padding: 10px;">
        <!-- 折叠按钮 -->
        <div style="cursor: pointer; margin-right: 20px;">
          <el-icon
            v-if="!isCollapse"
            @click="handleCollapse"
            size="25"
            color="#008B8B"
          >
            <Fold />
          </el-icon>
          <el-icon
            v-else
            @click="handleCollapse"
            size="25"
            color="#008B8B"
          >
            <Expand />
          </el-icon>
        </div>
        <!-- 用户信息下拉菜单 -->
        <div style="cursor: pointer; margin-left: auto;">
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
      </el-header>
    </el-container>
  </div>
</template>

<script lang="ts" setup>
import { ref,onMounted  } from 'vue';
import { useI18n } from 'vue-i18n';
import { Fold, Expand } from '@element-plus/icons-vue';
const { t } = useI18n();
import {useUserStore} from "@/stores/userStores"
// 定义 props 接收父组件传递的下拉框数据
const props = defineProps({
  items: {
    type: Array as () => { value: number; label: string }[],
    required: true,
  },
});
const isCollapse = ref(false); // 侧边栏折叠状态
const userStore=useUserStore()
// 定义 emit 事件
const emit = defineEmits(['collapse',"commond"]);
// 折叠侧边栏
function handleCollapse() {
  isCollapse.value = !isCollapse.value;
  // 触发自定义事件，并将 isCollapse 的值传递给父组件
  emit('collapse', isCollapse.value);
}
onMounted(() => {
  userStore.getMyProfile()
});

function commond(command: number) {
  emit("commond",command)
}

</script>

<style scoped>
.el-header {
  background-color: #393939;
  display: flex;
  align-items: center;
  padding: 10px;
}

.el-dropdown {
  color: #fff;
}

.el-avatar {
  margin-right: 10px;
}
</style>