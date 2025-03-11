<template>
  <div>
   <!-- 仪表盘 -->
     <el-card style="margin-top: 1%;" :header="t('app.webui.systemstatus')" body-style="display: flex; justify-content: space-around;">
        <div>
          <span style="display: flex; justify-content: center;">{{ t('app.webui.cpuuse') }}</span>
          <el-progress style="margin-top: 10%; font-weight: bold;" type="dashboard" stroke-width="12" :percentage="sysdata.cpu" :color="colors"></el-progress>
        </div>
        <div>
          <span style="display: flex; justify-content: center;">{{ t('app.webui.ramuse') }}</span>
          <el-progress style="margin-top: 10%; font-weight: bold;" type="dashboard" stroke-width="12" :percentage="sysdata.mem" :color="colors"></el-progress>
        </div>
        <div>
          <span style="display: flex; justify-content: center;">{{ t('app.webui.diskuse') }}</span>
          <el-progress style="margin-top: 10%; font-weight: bold;" type="dashboard" stroke-width="12" :percentage="sysdata.disk" :color="colors"></el-progress>
        </div>
     </el-card>
    <!-- 漏洞统计 -->
    <el-card :header="t('app.webui.vulnstatistics')" style="margin-top: 2%; margin-bottom: 2%;">
      <el-row justify="start"> <!-- space-evenly -->
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
      </el-row>
    </el-card>
  </div>
  <!-- 用户信息 -->
  <el-card :header="t('app.webui.userstatistics')" style="margin-top: 2%; margin-bottom: 2%;">
    <el-row justify="space-evenly">
      <StatisticCard
          :total-value="1234"
          total-name="总用户数"
        />
        <StatisticCard
          :total-value="1234"
          total-name="总管理员数"
        />
        <StatisticCard
          :total-value="1"
          total-name="在线人数"
        />
    </el-row>
  </el-card>

</template>
<script lang="ts" setup>
  import { useI18n } from 'vue-i18n';
  import { ref, onMounted } from 'vue';
  import { useRouter } from 'vue-router'
  import { Icon } from "@iconify/vue";
import { getSysStatus } from '@/api/admin/system';
  const { t } = useI18n();
  const router = useRouter();
  const sysdata = ref({
    "cpu":0,
    "disk":0,
    "mem":0
  })
  const vulndata = ref({})
  const colors = [
    { color: '#20a53a', percentage: 30 },
    { color: 'orange', percentage: 60 },
    { color: 'red', percentage: 100 }
  ]

  // let intervalId = setInterval(getSystem, 3000);

  onMounted(() => {
    // mountedFunctions.forEach(fn => {
    //   fn();
    // });
	  getSysStatus().then((res:any)=>{
      sysdata.value=res.data
    })
  });
  

//   async function getSystem() {
//     const token = sessionStorage.getItem('token')
//     try {
//       const config = {
//         headers: {
//           'Authorization': `Bearer ${token}`  // 使用Bearer schema
//         }
//       };
//       const response = await api.get('/api/v1/getsystemstatus', config)
//       if (response.data.code != 1) {
//         clearInterval(intervalId);
//         // 返回登录页
//         sessionStorage.removeItem('token')
//         sessionStorage.removeItem('username')
//         sessionStorage.removeItem('avatar')
//         router.push('/login')
//       }
//       sysdata.value = response.data.data
//     } catch (error) {
//       // 处理请求错误
//       //ElMessage.error(t('app.webui.loginerr2'));
//     }
//   }
//   async function getVuln() {
//   try {
//     const response = await api.get('/api/v1/getvulnabs')
//     vulndata.value = response.data
//   } catch (error) {
//     // 处理请求错误
//     //ElMessage.error(t('app.webui.loginerr2'));
//   }
// }

</script>
<style scoped>
</style>