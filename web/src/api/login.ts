import request from "@/utils/requests"
import { LoginData } from '@/user_types';

//登录方法
export function login(data:LoginData){
	return request({
	  url: '/login',
	  headers: {
	    isToken: false,
		'Content-Type': 'application/json'
	  },
	  method: 'post',
	  data: data
	})
}

//注册方法
export function register(data:LoginData){
	return request({
	  url: '/register',
	  headers: {
	    isToken: false,
	    'Content-Type': 'application/json'
	  },
	  method: 'post',
	  data: data
	})
}

//注销登录
export function logout(data:LoginData){
	return request({
	  url: '/logout',
	  headers: {
	    isToken: true,
	    'Content-Type': 'application/json'
	  },
	  method: 'get',
	  data: data
	})
}

//查询个人信息
export function getMyProfile(){
    return request({
        url: '/getMyProfile',
        headers: {
        isToken: true,
        'Content-Type': 'application/json'
        },
        method: 'get'
    })
}

//获取验证码
export function getCaptcha(data:any){
	return request({
        url: '/getCaptcha',
        headers: {
        isToken: true,
        'Content-Type': 'application/json'
        },
        method: 'get',
		data:data
    })
}

//重置密码
export function forgetPassword(data:any){
	return request({
        url: '/forgetPassword',
        headers: {
        isToken: true,
        'Content-Type': 'application/json'
        },
        method: 'post',
		data:data
    })
}

//更新密码
export function changePassword(data:any){
	return request({
        url: '/changePassword',
        headers: {
        isToken: true,
        'Content-Type': 'application/json'
        },
        method: 'post',
		data:data
    })
}

//更新个人信息
export function updateProfile(data:any){
	return request({
        url: '/updateProfile',
        headers: {
        isToken: true,
        'Content-Type': 'application/json'
        },
        method: 'post',
		data:data
    })
}