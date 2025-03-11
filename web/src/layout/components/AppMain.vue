<template>
  <el-main>
    <router-view v-slot="{ Component }">
      <transition name="fade">
        <keep-alive>
          <div>
            <component :is="Component" />
          </div>
        </keep-alive>
      </transition>
    </router-view>
  </el-main>
</template>
<script lang="ts" setup>
</script>
