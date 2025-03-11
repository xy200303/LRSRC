import { defineStore } from 'pinia';
import { LoginData } from "@/user_types";
import { login,  logout,getMyProfile } from "@/api/login";
import { getToken, setToken, removeToken } from '@/utils/auth';

export const useUserStore = defineStore(
  'user',
  {
    state: () => ({
      userInfo: JSON.parse(localStorage.getItem("userInfo")||"{}"),
    }),
    actions: {
      // 登录
      login(userInfo: LoginData) {
        return new Promise((resolve, reject) => {
          login(userInfo).then((res:any) => {
            setToken(res.data.token);
            this.userInfo = res.data;
            resolve(res);
          }).catch((error: any) => {
            reject(error);
          });
        });
      },
      // 获取用户信息
      getMyProfile() {
        return new Promise((resolve, reject) => {
          getMyProfile().then((res:any) => {
            this.userInfo = res.data;
            resolve(res.data);
          }).catch((error: any) => {
            reject(error);
          });
        });
      },
      //获取token
      getToken(){
        return getToken()
      },
      // 删除用户信息
      removeUserInfo() {
        this.userInfo={}
        //删除用户信息
        localStorage.removeItem("user")
        removeToken();
        console.log("信息已经删除")
      },
      // 退出系统
      logOut() {
        return new Promise((resolve, reject) => {
          logout().then((res: any) => {
            localStorage.removeItem("user")
            removeToken();
            resolve(res);
          }).catch((error: any) => {
            reject(error);
          });
        });
      },
      isAdmin(){
        return this.userInfo.is_admin
      }
    },
    persist: {
      enabled: true,
      strategies: [
        {
          key: 'userInfo',
          storage: localStorage,
          paths: ['userInfo'],
        }
      ],
    },
  }
);


