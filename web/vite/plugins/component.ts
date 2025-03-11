import Components from 'unplugin-vue-components/vite';
import { ElementPlusResolver } from 'unplugin-vue-components/resolvers';
import IconsResolver from 'unplugin-icons/resolver';
export default function createComponents() {
    return Components({
        resolvers: [
          // 自动注册 Element Plus 组件
          ElementPlusResolver(),
          // 自动注册图标组件
          IconsResolver({
            enabledCollections: ['ep'], // 启用 Element Plus 图标集
          }),
        ],
    })
}