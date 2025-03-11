import request from "@/utils/requests";

//获取基础设置方法
export function getBaseSysConfigMap(data:any){
    return request({
      url: '/getBaseSysConfigMap',
      headers: {
        isToken: false,
        'Content-Type': 'application/json'
      },
      method: 'get',
      data: data
    })
}
