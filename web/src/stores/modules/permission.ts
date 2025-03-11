import { defineStore } from "pinia";
import { RouteRecordRaw } from "vue-router"; // 使用 RouteRecordRaw 替代 RouteRecord
import { constantRoutes } from "@/router";
import { AdminRoutes, HomeRoutes } from "@/router/routes";
import { MunaCenterRoutes, UserCenterRoutes } from "@/router/routes";
const usePermissionStore = defineStore(
  'permission',
  {
    state: () => ({
      routes: [] as RouteRecordRaw[], // 明确指定类型
      defaultRoutes: constantRoutes,
      adminRouters: AdminRoutes,
      homeRouters: HomeRoutes,
      userCenterRoutes: UserCenterRoutes,
      munaCenterRoutes: MunaCenterRoutes,
    }),
    actions: {
      getAdminRoutes() {
        return this.adminRouters;
      },
      getHomeRoutes() {
        return this.homeRouters;
      },
      getUserCenterRoutes() {
        return this.userCenterRoutes;
      },
      getMunaCenterRoutes() {
        return this.munaCenterRoutes; // 修正拼写错误
      },
    },
  }
);

export default usePermissionStore;