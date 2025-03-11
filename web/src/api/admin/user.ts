import request from "@/utils/requests";

//列出用户
export function listUser(data:any){
    return request({
        url: '/user/list',
        headers: {
          isToken: true,
          'Content-Type': 'application/json'
        },
        method: 'post',
        data:data
      })
}

//创建用户
export function createUser(data:any){
    return request({
        url: '/user/create',
        headers: {
          isToken: true,
          'Content-Type': 'application/json'
        },
        method: 'post',
        data:data
      })
}

//删除用户
export function deleteUser(data:any){
    return request({
        url: '/user/delete',
        headers: {
          isToken: true,
          'Content-Type': 'application/json'
        },
        method: 'post',
        data:data
      })
}

//更新用户
export function updateUser(data:any){
    return request({
        url: '/user/update',
        headers: {
          isToken: true,
          'Content-Type': 'application/json'
        },
        method: 'post',
        data:data
      })
}

//设置管理员权限
export function setUserAdmin(data:any){
  return request({
    url: '/user/setUserAdmin',
    headers: {
      isToken: true,
      'Content-Type': 'application/json'
    },
    method: 'post',
    data:data
  })
}
