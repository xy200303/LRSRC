import request from "@/utils/requests"

//上传方法
export function uploadFile(data:FormData){
	return request({
	  url: '/uploadFile',
	  headers: {
	    isToken: true,
		'Content-Type': 'multipart/form-data',
	  },
	  method: 'post',
	  data: data
	})
}

//下载方法
export function downloadFile(data){
	return request({
	  url: '/downloadFile',
	  headers: {
	    isToken: true,
		'Content-Type': 'application/json'
	  },
	  method: 'get',
	  data: data
	})
}
