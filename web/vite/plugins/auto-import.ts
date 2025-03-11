
import AutoImport from 'unplugin-auto-import/vite'
import { ElementPlusResolver } from 'unplugin-vue-components/resolvers'
import IconsResolver from 'unplugin-icons/resolver'
export default function createAutoImport() {
    return AutoImport({
        resolvers: [
          // 自动导入函数
          ElementPlusResolver(),
          // 自动导入图标组件
          IconsResolver({
            prefix: 'Icon',
          }),
        ],
        imports: [
          'vue',
          'vue-router',
        ],
      })
}

