import request from "@/utils/requests"

//获取系统信息
export function getSysStatus(data:any){
	return request({
	  url: '/getSysStatus',
	  headers: {
	    isToken: true,
		'Content-Type': 'application/json'
	  },
	  method: 'get',
	  data: data
	})
}

//获取基本信息
export function getSysBaseInfo(data:any){
	return request({
	  url: '/getSysBaseInfo',
	  headers: {
	    isToken: true,
		'Content-Type': 'application/json'
	  },
	  method: 'get',
	})
}