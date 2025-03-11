import { defineConfig } from 'vite'

import path from 'path'
import createVitePlugins from './vite/plugins'
export default defineConfig(({ mode, command }) =>{
  const env=import.meta.env;
  return {
  build: {
    manifest: true,
    // 配置静态资源的公共路径
    assetsDir: 'static',
  },
  plugins: [
    createVitePlugins(env, command === 'build')
  ],
  resolve: {
    alias: {
      // 设置路径
      '~': path.resolve(__dirname, './'),
      // 设置别名
      '@': path.resolve(__dirname, './src')
    },
    // https://cn.vitejs.dev/config/#resolve-extensions
    extensions: ['.mjs', '.js', '.ts', '.jsx', '.tsx', '.json', '.vue']
  },
  server: {
      hmr: true, // 确保 HMR 启用（默认就是 true）
    },
}
})

