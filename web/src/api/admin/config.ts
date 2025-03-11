import request from "@/utils/requests"

//获取系统信息
export function getSysConfigMap(data:any){
	return request({
	  url: '/getSysConfigMap',
	  headers: {
	    isToken: true,
		'Content-Type': 'application/json'
	  },
	  method: 'get',
	  data: data
	})
}

//保存系统配置
export function updateSysConfigMap(data:any){
    return request({
        url: '/updateSysConfigMap',
        headers: {
          isToken: true,
          'Content-Type': 'application/json'
        },
        method: 'post',
        data: data
      })
}