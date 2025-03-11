import request from "@/utils/requests"

//获取文章分组
export function getArticleGroup(data:any){
	return request({
	  url: '/article/getArticleGroup',
	  headers: {
	    isToken: true,
		'Content-Type': 'application/json'
	  },
	  method: 'get',
	  data: data
	})
}



//列出字典类型
export function listArticleGroup(data:any){
	return request({
		url: '/article/listArticleGroup',
		headers: {
		  isToken: true,
		  'Content-Type': 'application/json'
		},
		method: 'post',
		data:data
	  })
}

//添加数据
export function createArticleGroup(data:any){
	return request({
		url: '/article/createArticleGroup',
		headers: {
		  isToken: true,
		  'Content-Type': 'application/json'
		},
		method: 'post',
		data:data
	  })
}

//删除数据
export function deleteArticleGroup(data:any){
	return request({
		url: '/article/deleteArticleGroup',
		headers: {
		  isToken: true,
		  'Content-Type': 'application/json'
		},
		method: 'post',
		data:data
	  })
}

//更新数据
export function updateArticleGroup(data:any){
	return request({
		url: '/article/updateArticleGroup',
		headers: {
		  isToken: true,
		  'Content-Type': 'application/json'
		},
		method: 'post',
		data:data
	  })
}

//获取全部数据
export function listAllArticleGroup(data:any){
	return request({
		url: '/article/ListAllArticleGroup',
		headers: {
		  isToken: true,
		  'Content-Type': 'application/json'
		},
		method: 'post',
		data:data
	  })
}

//发表文章
export function createArticle(data:any){
	return request({
		url: '/article/createArticle',
		headers: {
		  isToken: true,
		  'Content-Type': 'application/json'
		},
		method: 'post',
		data:data
	  })
}
//获取文章信息
export function getArticle(data:any){
	return request({
		url: '/article/getArticle',
		headers: {
		  isToken: true,
		  'Content-Type': 'application/json'
		},
		method: 'get',
		data:data
	  })
}

//获取文章信息
export function updateArticle(data:any){
	return request({
		url: '/article/updateArticle',
		headers: {
		  isToken: true,
		  'Content-Type': 'application/json'
		},
		method: 'post',
		data:data
	  })
}
//列出文章
export function listArticle(data:any){
	return request({
		url: '/article/listArticle',
		headers: {
		  isToken: true,
		  'Content-Type': 'application/json'
		},
		method: 'post',
		data:data
	  })
}
//删除文章
export function deleteArticle(data:any){
	return request({
		url: '/article/deleteArticle',
		headers: {
		  isToken: true,
		  'Content-Type': 'application/json'
		},
		method: 'post',
		data:data
	  })
}

//创建文章评论
export function sendArticleComment(data:any){
	return request({
		url: '/article/sendArticleComment',
		headers: {
		  isToken: true,
		  'Content-Type': 'application/json'
		},
		method: 'post',
		data:data
	  })
}

//列出文章数据
export function listArticleComment(data:any){
	return request({
		url: '/article/listArticleComment',
		headers: {
		  isToken: true,
		  'Content-Type': 'application/json'
		},
		method: 'post',
		data:data
	  })
}

//删除文章数据
export function deleteMyArticleComment(data:any){
	return request({
		url: '/article/deleteMyArticleComment',
		headers: {
		  isToken: true,
		  'Content-Type': 'application/json'
		},
		method: 'post',
		data:data
	  })
}

//文章点赞
export function likeArticle(data:any){
	return request({
		url: '/article/likeArticle',
		headers: {
		  isToken: true,
		  'Content-Type': 'application/json'
		},
		method: 'post',
		data:data
	  })
}

//列出首页文章
export function listHomeArticle(data:any){
	return request({
		url: '/article/listHomeArticle',
		headers: {
		  isToken: true,
		  'Content-Type': 'application/json'
		},
		method: 'post',
		data:data
	  })
}

//列出首页文章接口
export function listHomeAllArticleGroup(data:any){
	return request({
		url: '/article/listHomeAllArticleGroup',
		headers: {
		  isToken: true,
		  'Content-Type': 'application/json'
		},
		method: 'post',
		data:data
	  })
}