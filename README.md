# 零日信安-一企一案网络安全解决方案

## 项目简介
基于Golang+Vue3构建的Web应用系统，包含：
- 系统配置管理（SMTP/OSS/AI等）
- 多云存储方案支持（阿里云/腾讯云/华为云）
- OpenAI接口集成
- 用户权限管理系统
- 漏洞管理模块

## 功能特性
✅ 系统配置中心化管理  
✅ 文件分块上传与云存储  
✅ GPT-4o人工智能集成  
✅ JWT鉴权与RBAC权限控制  
✅ 多环境配置支持

## 环境要求
### 后端环境
- Go 1.20+ (推荐使用goenv管理版本)
- MySQL 8.0+ 
- Redis 7.0+
- Git 2.30+

### 前端环境
- Node.js 18+ (推荐使用nvm管理版本)
- npm 9+ 或 yarn 1.22+
- Vite 5.0+

## 安装步骤
```bash
# 克隆仓库
git clone https://github.com/xy200303/xiaoyun.git

# 系统准备
sudo apt-get install -y git build-essential

# 后端依赖安装
cd backend
export GOPROXY=https://goproxy.cn  # 国内镜像加速
go mod download -x  # 显示详细下载过程

# 前端依赖安装
cd ../web
npm config set registry https://registry.npmmirror.com  # 国内镜像源
npm install --verbose
```

## 配置说明
1. 复制环境模板文件：
```
cp config.yaml.example config.yaml
cp web/.env.example web/.env
```
2. 配置数据库连接（config.yaml）：
```yaml
database:
  host: 127.0.0.1
  port: 3306
  user: root
  password: yourpassword
```
3. 安全配置：
```bash
# 生成JWT加密密钥
openssl rand -base64 32 > config/jwt_secret.key

# 删除示例密钥（重要!）
sed -i '/yourpassword/d' config.yaml
```

## 开发指南
### 启动服务
```bash
# 启动后端
cd backend
go run main.go

# 启动前端
cd ../web
npm run dev
```

### 项目结构
```
├── backend/         # Go后端服务
│   ├── config/      # 配置加载
│   ├── controllers/ # 控制器层
│   ├── models/      # 数据库模型
│   └── service/     # 业务逻辑
├── web/             # Vue3前端
│   ├── src/        
│   │   ├── api/     # 接口定义
│   │   └── views/   # 页面组件
└── test/            # 单元测试
```


## 许可证
MIT License
