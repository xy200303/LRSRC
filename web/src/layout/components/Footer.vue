<template>
  <el-footer class="footer">
    <div class="footer-content">
      <span v-if="icpNumber">备案号：{{ icpNumber }}</span>
      <span>{{copyRight}}</span>
    </div>
  </el-footer>
</template>

<script lang="ts" setup>
defineProps({
  icpNumber: {
    type: String,
    default: import.meta.env.VITE_APP_ICP, // 如果没有备案号，可以不传
  },
  copyRight:{
    type:String,
    default:import.meta.env.VITE_APP_COPYRIGHT
  }
});
</script>

<style scoped>
.footer {
  padding: 10px;
  background-color: #f5f5f5; /* 背景色 */
  color: #666; /* 文字颜色 */
  text-align: center; /* 文字居中 */
  border-top: 1px solid #e4e4e4; /* 顶部边框 */
  height: 60px; /* 脚部高度 */
  background-color: #fff; /* 脚部背景色 */
  text-align: center; /* 文字居中 */
  box-shadow: 0 -1px 4px rgba(0, 21, 41, 0.08); /* 脚部阴影 */
}

.footer-content {
  display: flex;
  flex-direction: column; /* 垂直排列 */
  gap: 4px; /* 内容间距 */
}

@media (min-width: 768px) {
  .footer-content {
    flex-direction: column; /* 大屏幕下水平排列 */
    justify-content: center; /* 水平居中 */
    gap: 4px; /* 增加间距 */
  }
}
</style>