import axios from 'axios';
import { getToken, removeToken } from './auth';
import qs from 'qs';
const request = axios.create({
  baseURL: import.meta.env.VITE_APP_API_URL,
  timeout: 20000,
});


// 添加请求拦截器
request.interceptors.request.use(
  (config) => {
	  // 是否需要设置 token
	  const isToken = (config.headers || {}).isToken
    // 检查是否有 hasAuth 参数，并且其值为 true
    if (isToken && getToken()) {
        config.headers['Authorization'] = "Bearer "+getToken();
    }
    // 如果是 GET 请求，并且有 data 参数，则将其转换为查询字符串并附加到 URL 上
    if (config.method === 'get' && config.data) {
      const params = qs.stringify(config.data, { arrayFormat: 'brackets' });
      config.url = `${config.url}?${params}`;
      delete config.data; // 清除 data 属性，因为 GET 请求不应包含 body 数据
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// 添加响应拦截器
request.interceptors.response.use(
  (response) => {
    // 对响应数据进行处理
    const res = response.data;
    // 假设响应数据中包含 code 字段，0 表示成功
    if (res.code !== 200) {
      // 处理业务逻辑错误
	  ElMessage.error(res.msg)
      return Promise.reject(new Error(res.msg || 'Error'));
    }
    // 返回响应数据
    return res;
  },
  (error) => {
    // 对响应错误进行处理
    if (error.response) {
      // 请求已发出，但服务器响应状态码不在 2xx 范围内
      switch (error.response.status) {
        case 401:
          ElMessage.error(error.response.data.msg);
          // 可以在这里跳转到登录页面
          removeToken()
          // location.reload()
          break;
        case 404:
          ElMessage.error("不存在的资源接口")
          break;
        default:
          ElMessage.error( error.response.data.msg);
      }
    } else if (error.request) {
      // 请求已发出，但没有收到响应
      console.error('无响应:', error.request);
    } else {
      // 其他错误
      console.error('请求配置错误:', error.message);
    }
    return Promise.reject(error);
  }
);
export default request;