import request from "@/utils/requests"

//获取字典类型
export function getDictType(data:any){
	return request({
	  url: '/getDictType',
	  headers: {
	    isToken: true,
		'Content-Type': 'application/json'
	  },
	  method: 'get',
	  data: data
	})
}



//列出字典类型
export function listDictType(data:any){
	return request({
		url: '/listDictType',
		headers: {
		  isToken: true,
		  'Content-Type': 'application/json'
		},
		method: 'get',
		data:data
	  })
}

//添加数据
export function createDictType(data:any){
	return request({
		url: '/createDictType',
		headers: {
		  isToken: true,
		  'Content-Type': 'application/json'
		},
		method: 'post',
		data:data
	  })
}

//删除数据
export function deleteDictType(data:any){
	return request({
		url: '/deleteDictType',
		headers: {
		  isToken: true,
		  'Content-Type': 'application/json'
		},
		method: 'post',
		data:data
	  })
}

//更新数据
export function updateDictType(data:any){
	return request({
		url: '/updateDictType',
		headers: {
		  isToken: true,
		  'Content-Type': 'application/json'
		},
		method: 'post',
		data:data
	  })
}


//获取字典数据
export function getDictData(data:any){
	return request({
	  url: '/getDictData',
	  headers: {
	    isToken: true,
		'Content-Type': 'application/json'
	  },
	  method: 'get',
	  data: data
	})
}

//列出字典数据
export function listDictData(data:any){
	return request({
		url: '/listDictData',
		headers: {
		  isToken: true,
		  'Content-Type': 'application/json'
		},
		method: 'get',
		data:data
	  })
}

//添加数据
export function createDictData(data:any){
	return request({
		url: '/createDictData',
		headers: {
		  isToken: true,
		  'Content-Type': 'application/json'
		},
		method: 'post',
		data:data
	  })
}

//删除数据
export function deleteDictData(data:any){
	return request({
		url: '/deleteDictData',
		headers: {
		  isToken: true,
		  'Content-Type': 'application/json'
		},
		method: 'post',
		data:data
	  })
}

//更新数据
export function updateDictData(data:any){
	return request({
		url: '/updateDictData',
		headers: {
		  isToken: true,
		  'Content-Type': 'application/json'
		},
		method: 'post',
		data:data
	  })
}