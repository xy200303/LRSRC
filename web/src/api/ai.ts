import request from "@/utils/requests"

//总结摘要方法
export function summaryContent(data:any){
    return request({
      url: '/ai/summaryContent',
      headers: {
        isToken: true,
         'Content-Type': 'application/json',
      },
      method: 'post',
      data: data
    })
}

//AI功能
