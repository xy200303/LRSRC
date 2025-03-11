<template>
    <el-card shadow="always" style="padding: 20px;">
      <template #header>
        <el-row justify="space-between" style="font-weight: bold; font-size: 18px;">
          <span>{{ totalName }}</span>
          <el-icon size="25" style="margin-left: 20px;">
            <Icon :icon="icon" source="local" size="30px"></Icon>
          </el-icon>
        </el-row>
      </template>
      <el-col>
        <el-statistic :value="totalValue"></el-statistic>
        <div style="display: flex; font-size: 13px; margin-top: 6%; gap: 10px;">
          <span>{{ increaseName }}</span>
          <span>{{ increaseValue }}</span>
        </div>
      </el-col>
    </el-card>
  </template>
  
  <script>
  export default {
    name: "StatisticCard",
    props: {
      icon: {
        type: Object,
        required: true,
        description: "SVG 图标组件"
      },
      totalValue: {
        type: Number,
        required: true,
        description: "总数量"
      },
      increaseValue: {
        type: Number,
        required: true,
        description: "增加数量"
      },
      totalName: {
        type: String,
        required: true,
        description: "总数量名称"
      },
      increaseName: {
        type: String,
        required: true,
        description: "增加数量名称"
      }
    }
  };
  </script>
  
  <style scoped>
  /* 样式可以根据需要调整 */
  </style>