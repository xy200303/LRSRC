<template>
    <div class="not-found">
      <h1>404 - 页面未找到</h1>
      <p>抱歉，您访问的页面不存在。</p>
      <router-link to="/">返回首页</router-link>
    </div>
  </template>
  
  <script>
  export default {
    name: 'NotFound',
  };
  </script>
  
  <style scoped>
  .not-found {
    text-align: center;
    margin-top: 50px;
    font-family: Arial, sans-serif;
  }
  
  h1 {
    font-size: 48px;
    color: #333;
  }
  
  p {
    font-size: 24px;
    color: #666;
    margin-bottom: 20px;
  }
  
  a {
    text-decoration: none;
    color: #42b983;
    font-weight: bold;
  }
  
  a:hover {
    text-decoration: underline;
  }
  </style>