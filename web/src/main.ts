import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import router from './router'
import ElementPlus from 'element-plus'
import { i18n } from './i18n'
import { createPinia } from 'pinia';
import { useDict } from '@/utils/dict'
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate';
import { useUmoEditor } from '@umoteam/editor';
// 字典标签组件
import DictTag from "@/components/DictTag"
import SidebarTree from "@/components/SidebarTree"
import TagInput from "@/components/TagInput"
import "@wangeditor/editor/dist/css/style.css"
import "@wangeditor/code-highlight/dist/css/style.css"

// 读取 localStorage 中的语言设置
const savedLanguage = localStorage.getItem('selectedLanguage');
if (savedLanguage) {
  i18n.global.locale = savedLanguage;
}

import 'virtual:svg-icons-register'

// 创建 Pinia 实例
const pinia = createPinia();
pinia.use(piniaPluginPersistedstate);
const app = createApp(App);
app.config.globalProperties.useDict = useDict
app.use(ElementPlus);
app.use(i18n);
app.use(router);
app.use(pinia)
app.use(useUmoEditor, {
});

app.component('DictTag', DictTag)
app.component("SidebarTree",SidebarTree)
app.component("TagList",TagInput)
app.mount('#app');
