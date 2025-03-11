// router/index.ts 文件
import { createRouter,createWebHistory, createWebHashHistory, RouterOptions, Router, RouteRecordRaw } from 'vue-router'
//由于router的API默认使用了类型进行初始化，内部包含类型定义，所以本文内部代码中的所有数据类型是可以省略的
//RouterRecordRaw是路由组件对象
import { getToken } from '@/utils/auth';
import { useUserStore } from '@/stores/userStores';
import usePermissionStore from '../stores/modules/permission';
//默认公共路由地址
export const constantRoutes: RouteRecordRaw[] = [
     //登录页面
     { 
      path: '/login', 
      name: 'Login',
      component: () => import('@/views/Login.vue'),
      meta: { title: '用户登录', activeMenu: '/login',hidden:true }
    },
    //忘记密码页面
    { 
      path: '/forgotpwd',
      name: 'Forgotpwd',
      component: () => import('@/views/Forgotpwd.vue'),
      meta: { title: '忘记密码', activeMenu: '/forgotpwd',hidden:true }
    },
    //注册页面
    { 
      path: '/register',
      name: 'Register',
      component: () => import('@/views/Register.vue'),
      meta: { title: '用户注册', activeMenu: '/register',hidden:true }
    },
    //测试页面
    { 
      path: '/test',
      name: 'Test',
      component: () => import('@/views/admin/test.vue'),
      meta: { title: '测试', activeMenu: '/test',hidden:true }
    },
    //404页面
    { 
      path: '/404',
      name: '404',
      component: () => import('@/views/404.vue'),
      meta: { title: '404', activeMenu: '/404',hidden:true }
    },
];


// Router是路由对象类型
const router: Router = createRouter({
  history: createWebHistory(),
  routes:constantRoutes
})

router.beforeEach(async (to, from, next) => {
  const token = getToken(); // 获取用户 Token
  const isAuthRoute = ['/login', '/register', '/forgotpwd', '/404', '/503'].includes(to.path); // 是否为无需验证的路由
  // 验证 Token
  if (!token && !isAuthRoute) {
    return next('/login'); // 没有 Token 且不是登录相关页面时，跳转到登录页
  }
  // 路由刷新器逻辑：动态路由挂载
  if (to.matched.length === 0) {
    try {
      const userStore = useUserStore();
      const permissionStore = usePermissionStore();
      // 确保 userInfo 已加载（如果有异步操作）
      if (!userStore.userInfo) {
        await userStore.fetchUserInfo(); // 假设有一个异步方法获取用户信息
      }
      const { is_admin } = userStore.userInfo;

      // 根据权限动态挂载路由
      if (is_admin) {
        mountRoutes(permissionStore.getAdminRoutes());
      }
      mountRoutes(permissionStore.getHomeRoutes());
      mountRoutes(permissionStore.getUserCenterRoutes());

      // 再次检查目标路由是否匹配
      const matchedRoute = router.getRoutes().find((route) => route.path === to.path);
      if (!matchedRoute) {
        return next('/404'); // 如果目标路由仍然不存在，跳转到 404 页面
      }
      // 再次尝试导航到目标路径
      return next(to.fullPath); // 替换当前导航记录，避免重复触发
    } catch (error) {
      console.error('动态路由挂载失败:', error);
      return next('/404'); // 如果发生错误，跳转到 404 页面
    }
  }

  // 默认情况下继续导航
  next();
});


//挂在路由
function mountRoutes(routes: RouteRecordRaw[]) {
  routes.forEach((route) => {
    // 检查路由是否已存在
    console.log(route)
    const routeExists = router.getRoutes().some((r) => r.path === route.path);
    if (!routeExists) {
      try {
        router.addRoute(route); // 动态添加路由
        console.log(`Route added: ${route.path}`);
        // 使用 nextTick 确保路由已注册
        router.replace(router.currentRoute.value.fullPath);
      } catch (error) {
        console.error(`Failed to add route: ${route.path}`, error);
      }
    } else {
      console.warn(`Route already exists: ${route.path}`);
    }
  });
}
export default router