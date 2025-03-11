import vue from '@vitejs/plugin-vue'

import createSvgIcon from './svg-icon'
import createAutoImport from "./auto-import"
import createComponent from "./component"
export default function createVitePlugins(viteEnv: any, isBuild = false) {
    const vitePlugins = [vue()]; // 默认添加 Vue 插件
    vitePlugins.push(createSvgIcon(isBuild)); // 添加 SVG 图标插件
    vitePlugins.push(createAutoImport()) //自动导入
    vitePlugins.push(createComponent()) //自动导入组件
    return vitePlugins;
  }
