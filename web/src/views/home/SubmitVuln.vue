<template>
  <!-- 当 showSuccess 为 true 时，显示提交成功信息 -->
  <div v-if="showSuccess">
    <el-result
        icon="success"
        title="提交成功"
        sub-title="操作已经成功，请等待漏洞审核"
      >
      <!-- 结束 el-result 组件，该组件用于展示提交成功的信息 -->
      </el-result>
  </div>
  <!-- 当 showSuccess 为 false 时，显示漏洞内容编辑页面 -->
  <div v-else>
    <!-- 漏洞内容编辑页面 -->
    <el-card>
      <!-- 表单部分，绑定数据和验证规则 -->
      <el-form :model="vuln" :rules="formRules" ref="articleForm" label-position="left" label-width="auto">
        <!-- 漏洞标题输入项 -->
        <el-form-item label="漏洞标题" required prop="title">
          <el-input v-model="vuln.title" placeholder="单位名称+漏洞类型，如：某单位存在SQL注入漏洞" maxlength="100" show-word-limit></el-input>
        </el-form-item>
        <!-- 漏洞属性选择项 -->
        <el-form-item label="漏洞属性" required prop="attribute">
          <el-radio
            v-for="item in sys_vuln_attribute"
            :key="item.value"
            :label="item.value"
            v-model="vuln.attribute"
          >
            {{ item.label }}
          </el-radio>
        </el-form-item>
        <!-- 厂商名称输入项 -->
        <el-form-item label="厂商名称" required prop="muna_name">
          <el-autocomplete v-model="vuln.muna_name" placeholder="请输入厂商名称" :fetch-suggestions="querySearchAsync" maxlength="100" show-word-limit></el-autocomplete>
        </el-form-item>
        <!-- 厂商域名输入项 -->
        <el-form-item label="厂商域名" required prop="muna_domain">
          <el-input v-model="vuln.muna_domain" placeholder="请输入厂商域名" ></el-input>
        </el-form-item>
        <!-- 漏洞类型选择项 -->
        <el-form-item label="漏洞类型" required prop="type">
          <el-cascader
            v-model="vuln_type"
            :options="sys_vuln_type"
          />
        </el-form-item>
        <!-- 漏洞等级选择项 -->
        <el-form-item label="漏洞等级" required prop="level">
          <el-radio
            v-for="item in sys_vuln_level"
            :key="item.value"
            :label="item.value"
            v-model="vuln.level"
          >
            {{ item.label }}
          </el-radio>
        </el-form-item>
        <!-- 漏洞描述输入项 -->
        <el-form-item label="漏洞描述" required prop="desc">
          <el-input :rows="5" type="textarea"  v-model="vuln.desc" placeholder="简要描述漏洞概况以及影响，请勿在此填写漏洞URL等具体漏洞信息" maxlength="500" show-word-limit></el-input>
        </el-form-item>
        <!-- 漏洞URL输入项，根据 cate_type 决定是否显示 -->
        <el-form-item label="漏洞URL" required prop="url" v-if="formItemVerbose.url">
          <el-input v-model="vuln.url" placeholder="URL格式：以http://或https://开头" maxlength="100" show-word-limit></el-input>
        </el-form-item>
        <!-- 漏洞PoC包输入项，根据 cate_type 决定是否显示 -->
        <el-form-item label="PoC请求包" required prop="poc" v-if="formItemVerbose.poc">
          <el-input type="textarea" :rows="10" v-model="vuln.poc" placeholder="请输入漏洞PoC请求包,使用标准HTTP格式包" ></el-input>
        </el-form-item>
        <!-- 漏洞详情输入项 -->
        <el-form-item label="漏洞详情" required prop="detail">
          <simple-editor v-model="vuln.detail"></simple-editor>
        </el-form-item>
        <!-- 漏洞附件上传项 -->
        <el-form-item label="漏洞附件" required prop="detail">
          <x-upload-file
            :limit="1"
            v-model="file_obj_list"
            accept=".zip,.tar,.tag.gz"
          />
        </el-form-item>
        <!-- 修复建议输入项 -->
        <el-form-item label="修复建议" required prop="repair_suggestion">
          <simple-editor  height="200px" v-model="vuln.repair_suggestion"></simple-editor>
        </el-form-item>
        <!-- 所属地区选择项 -->
        <el-form-item label="所属地区" required prop="county">
          <data-select data="area" v-model="region"></data-select>
        </el-form-item>
        <!-- 所属行业选择项 -->
        <el-form-item label="所属行业" required prop="industry">
          <data-select data="cate" v-model="vuln.industry"></data-select>
        </el-form-item>
      </el-form>
      <!-- 提交按钮，点击触发 handleSubmitVuln 方法 -->
      <el-button type="primary" @click="handleSubmitVuln">{{"提交漏洞"}}</el-button>
    </el-card>
  </div>
</template>

<script lang="ts" setup>
import { computed, getCurrentInstance, onMounted, reactive, ref } from 'vue';
import { useRoute } from 'vue-router';
import { submitVuln, buildVulnTypeTree } from '@/api/admin/vuln';
import { ElMessage } from 'element-plus';
import { getVuln } from '@/api/admin/vuln';
import {VulnInfo} from "@/types/vuln";
import { SysDictItem } from '@/types/dict';

// 获取当前路由信息
const route = useRoute();
// 控制是否显示提交成功信息
const showSuccess = ref<boolean>(false);
// 获取当前实例
const { proxy } = getCurrentInstance();


// 获取系统字典数据
const { sys_vuln_attribute, sys_vuln_level }: { sys_vuln_attribute: SysDictItem[]; sys_vuln_level: SysDictItem[] } = proxy.useDict("sys_vuln_attribute", "sys_vuln_level");

// 存储漏洞信息的对象
const vuln = ref<VulnInfo>({
  muna_name: "",
  title: "",
  desc: "",
  created_by: '',
  muna_domain: "",
  cate_type: 0,
  type: 0,
  attribute: "eve",
  detail: "",
  repair_suggestion: "",
  url: "",
  poc: "",
  province: "",
  city: "",
  county: "",
  industry: [],
  attachment_id: '',
  attachment_name: '',
  level: "low"
});

// 存储漏洞类型数据
const sys_vuln_type = ref<SysDictItem[]>([]);

// 地区数据的计算属性
const region = computed<[string, string, string]>({
  get: () => [vuln.value.province, vuln.value.city, vuln.value.county],
  set: (val: [string, string, string]) => {
    vuln.value.province = val[0];
    vuln.value.city = val[1];
    vuln.value.county = val[2];
  }
});

// 漏洞类型数据的计算属性
const vuln_type = computed<[number, number]>({
  get: () => [vuln.value.cate_type, vuln.value.type],
  set: (val: [number, number]) => {
    vuln.value.cate_type = val[0];
    vuln.value.type = val[1];
  }
});

// 附件类型数据的计算属性
const file_obj_list = computed<{ name: string; file_id: string }[] | undefined>({
  get: () => {
    if (vuln.value.attachment_id) {
      return [{
        name: vuln.value.attachment_name,
        file_id: vuln.value.attachment_id
      }];
    }
    return undefined;
  },
  set: (val: { name: string; file_id: string }[] | undefined) => {
    if (val) {
      vuln.value.attachment_id = val[0]?.file_id;
      vuln.value.attachment_name = val[0]?.name;
    }
  }
});

// 参数控制器，根据 cate_type 决定是否显示某些表单项
const formItemVerbose = computed<{ url: boolean; poc: boolean }>(() => ({
  url: vuln.value.cate_type === 1,
  poc: vuln.value.cate_type === 1
}));

// 搜索公司的异步方法
const querySearchAsync = (queryString: string, cb: (arg: SysDictItem[]) => void) => {
  const data = ref<SysDictItem[]>([
    { value: "湖北经济学院" },
    { value: "武汉大学" }
  ]);
  const results = queryString
    ? data.value.filter(createFilter(queryString))
    : data.value;
  cb(results);
};

// 创建过滤函数
const createFilter = (queryString: string) => {
  return (restaurant: SysDictItem) => {
    return (
      restaurant.value.toLowerCase().indexOf(queryString.toLowerCase()) === 0
    );
  };
};

// 挂载后加载漏洞类型数据
onMounted(() => {
  ListVulnTypeData();
});

// 加载漏洞类型数据的函数
async function ListVulnTypeData() {
  try {
    const res = await buildVulnTypeTree();
    sys_vuln_type.value = res.data;
  } catch (err) {
    console.error(err);
  }
}

// 定义表单验证规则的类型
interface FormRules {
  [key: string]: {
    required?: boolean;
    message: string;
    trigger: string;
    max?: number;
    pattern?: RegExp;
  }[];
}

// 新增验证规则
const formRules = reactive<FormRules>({
  title: [
    { required: true, message: '请输入漏洞标题', trigger: 'blur' },
    { max: 100, message: '标题长度不能超过100个字符', trigger: 'blur' }
  ],
  attribute: [
    { required: true, message: '请选择漏洞属性', trigger: 'change' }
  ],
  muna_name: [
    { required: true, message: '请输入厂商名称', trigger: 'blur' },
    { max: 100, message: '厂商名称长度不能超过100个字符', trigger: 'blur' }
  ],
  muna_domain: [
    { required: true, message: '请输入厂商域名', trigger: 'blur' },
    { pattern: /^(?:[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,}$/, message: '请输入正确的域名格式，如：example.com或sub.example.com', trigger: 'blur' }
  ],
  type: [
    { required: true, message: '请选择漏洞类型', trigger: 'change' }
  ],
  url: [
    { required: true, message: '请输入漏洞URL', trigger: 'blur' },
    { pattern: /^https?:\/\/.+/, message: '请输入正确的URL格式', trigger: 'blur' }
  ],
  poc: [
    { required: true, message: '请输入漏洞PoC请求包', trigger: 'blur' }
  ],
  detail: [
    { required: true, message: '请输入漏洞详情', trigger: 'blur' }
  ],
  repair_suggestion: [
    { required: true, message: '请输入修复建议', trigger: 'blur' }
  ],
  county: [
    { required: true, message: '请选择所属地区', trigger: 'change' }
  ],
  industry: [
    { required: true, message: '请选择所属行业', trigger: 'change' }
  ]
});

// 挂载后初始化数据
onMounted(async () => {
  const id = route.query.id as string | undefined;
  if (id) {
    try {
      const res = await getVuln({ id });
      vuln.value = res.data;
    } catch (err) {
      console.error(err);
    }
  }
});

// 提交漏洞的方法
const handleSubmitVuln = async () => {
  const valid = await (proxy.$refs.articleForm as any).validate();
  if (valid) {
    try {
      await submitVuln({ ...vuln.value });
      ElMessage.success("提交成功");
      showSuccess.value = true;
    } catch (err) {
      ElMessage.error("提交失败");
    }
  } else {
    ElMessage.error("请填写完整所有必填项");
  }
};
</script>

<style scoped>
</style>
