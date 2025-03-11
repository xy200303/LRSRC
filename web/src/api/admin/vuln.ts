import request from "@/utils/requests"

//列出漏洞类型
export function listVulnType(data:any){
    return request({
        url: '/vuln/listVulnType',
        headers: {
          isToken: true,
          'Content-Type': 'application/json'
        },
        method: 'post',
        data:data
      })
}

//添加漏洞类型数据
export function createVulnType(data:any){
    return request({
        url: '/vuln/createVulnType',
        headers: {
          isToken: true,
          'Content-Type': 'application/json'
        },
        method: 'post',
        data:data
      })
}

//删除漏洞类型数据
export function deleteVulnType(data:any){
    return request({
        url: '/vuln/deleteVulnType',
        headers: {
          isToken: true,
          'Content-Type': 'application/json'
        },
        method: 'post',
        data:data
      })
}

//更新漏洞类型数据
export function updateVulnType(data:any){
    return request({
        url: '/vuln/updateVulnType',
        headers: {
          isToken: true,
          'Content-Type': 'application/json'
        },
        method: 'post',
        data:data
      })
}

//获取树形数据
export function buildVulnTypeTree(data:any){
  return request({
      url: '/vuln/buildVulnTypeTree',
      headers: {
        isToken: true,
        'Content-Type': 'application/json'
      },
      method: 'get',
      data:data
    })
}


//提交漏洞数据
export function submitVuln(data:any){
  return request({
      url: '/vuln/submitVuln',
      headers: {
        isToken: true,
        'Content-Type': 'application/json'
      },
      method: 'post',
      data:data
    })
}
//查看漏洞数据
export function listVuln(data:any){
  return request({
      url: '/vuln/listVuln',
      headers: {
        isToken: true,
        'Content-Type': 'application/json'
      },
      method: 'post',
      data:data
    })
}

//获取漏洞数据
export function getVuln(data:any){
  return request({
      url: '/vuln/getVuln',
      headers: {
        isToken: true,
        'Content-Type': 'application/json'
      },
      method: 'get',
      data:data
    })
}
//更新漏洞数据
export function updateVuln(data:any){
  return request({
      url: '/vuln/updateVuln',
      headers: {
        isToken: true,
        'Content-Type': 'application/json'
      },
      method: 'post',
      data:data
    })
}
//删除漏洞数据
export function deleteVuln(data:any){
  return request({
      url: '/vuln/deleteVuln',
      headers: {
        isToken: true,
        'Content-Type': 'application/json'
      },
      method: 'post',
      data:data
    })
}

//删除漏洞数据
export function auditVuln(data:any){
  return request({
      url: '/vuln/auditVuln',
      headers: {
        isToken: true,
        'Content-Type': 'application/json'
      },
      method: 'post',
      data:data
    })
}