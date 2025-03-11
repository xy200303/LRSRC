<template>
  <div class="content">
    <!-- 公司介绍 -->
    <el-card style="width: 100%;margin-bottom: 20px;">
        <div class="company-intro" @mousemove="handleMouseMove">
          <div class="company-intro__text" style="margin-bottom: 40px;">
            <h2 class="company-intro__title">零日信安-一企一案网络安全解决方案</h2>
            <p class="company-intro__description">诚邀广大安全专家共捍亿万互联网用户安全</p>
          </div>
          <el-button type="primary" class="company-intro__button" @click="handleSubmitClick">在线提交漏洞</el-button>
        </div>
    </el-card>
    <el-card>
      <!--  -->
      <div style="display: flex;flex-direction: row;justify-content: center;align-items: center;">
          <StatisticCard
          icon="total_vuln"
          :total-value="1234"
          :increase-value="56"
          total-name="总漏洞数"
          increase-name="本周新增"
          style="margin: 40px;"
        />
        <!--  -->
        <StatisticCard
          icon="total_poc"
          :total-value="1234"
          :increase-value="56"
          total-name="总POC数"
          increase-name="本周新增"
          style="margin: 40px;"
        />
        <!--  -->
        <StatisticCard
          icon="total_exp"
          :total-value="1234"
          :increase-value="56"
          total-name="总EXP数"
          increase-name="本周新增"
          style="margin: 40px;"
        />
        <!--  -->
        <StatisticCard
          icon="total_product"
          :total-value="1234"
          :increase-value="56"
          total-name="总影响产品数"
          increase-name="本周新增"
          style="margin: 40px;"
        />
      </div>
    </el-card>

    </div>
</template>

<script lang="ts" setup>
import { useI18n } from 'vue-i18n';
import { ref, onMounted, getCurrentInstance, reactive } from 'vue'
import { useRouter } from 'vue-router';

const  {proxy}  = getCurrentInstance();
const { sys_vuln_attribute,sys_vuln_level,sys_vuln_status } = proxy.useDict("sys_vuln_attribute","sys_vuln_level","sys_vuln_status");
const { t } = useI18n();
const data = ref({})
const tableData=ref([])
const page=ref({
    "page":1,
    "page_size":5,
    "total":1
})
const router=useRouter()
function handleSubmitClick(){
  router.push("/submitVuln")
}

onMounted(()=>{
    // listData()
})

const handleMouseMove = (e) => {
  const { offsetX, offsetY } = e;
  const target = e.currentTarget;
  const { offsetWidth, offsetHeight } = target;
  const x = (offsetX / offsetWidth - 0.5) * 20;
  const y = (offsetY / offsetHeight - 0.5) * 20;
  target.style.transform = `perspective(1000px) rotateX(${y}deg) rotateY(${x}deg)`;
};
</script>

<style scoped>
.content {
  width: 100%;
  display: flex;
  margin: auto;
  flex-direction: column;
  justify-content: center;
  align-content: center;
}

.company-intro {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  background: url('https://img.shetu66.com/2023/07/14/1689302106613476.png') no-repeat center center fixed;
  background-size: cover;
  color: white;
  padding: 20px;
  border-radius: 8px;
  height: 500px;
  width: 100%;
  text-align: center;
  transition: transform 0.3s ease; /* 添加过渡效果 */
  transform-style: preserve-3d; /* 保留 3D 效果 */
}

.company-intro__text {
  text-align: center;
}

.company-intro__title {
  font-size: 28px;
  font-weight: bold;
  color: white;
  margin-bottom: 10px;
}

.company-intro__description {
  font-size: 16px;
  color: white;
  line-height: 1.6;
}

.company-intro__button {
  padding: 12px 24px;
  font-size: 16px;
  border-radius: 4px;
}
</style>