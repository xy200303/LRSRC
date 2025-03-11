const TokenKey = 'Admin-Token';

// 获取 token
export function getToken() {
  return localStorage.getItem(TokenKey);
}

// 设置 token
export function setToken(token:string) {
  localStorage.setItem(TokenKey, token);
}

// 删除 token
export function removeToken() {
  localStorage.removeItem(TokenKey);
}