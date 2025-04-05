# 项目名称: shortlink-go

## 项目简介

shortlink-go 是一个用于生成和管理短链接的应用程序。该项目包含前端和后端部分，前端使用现代Web技术构建，后端提供RESTful API服务。

## 目录结构

- `backend/`: 后端代码目录，包含RESTful API的实现。
- `frontend/`: 前端代码目录，包含用户界面的实现。
- `package.json`: 项目的依赖管理文件。

## 安装与运行

### 前端

1. 进入 `frontend/shortlink-frontend` 目录。
2. 安装依赖：
   ```bash
   npm install
   ```
3. 启动前端开发服务器：
   ```bash
   npm start
   ```

### 后端

1. 进入 `backend/restful` 目录。
2. 安装后端所需的依赖。
3. 启动后端服务器。

## 后端服务

后端涉及到两个RPC服务：

1. **UserRpcService**: 用户服务，使用 `user.User` 客户端进行通信。配置在 `shortlink-api.yaml` 中的 `UserRpcConf` 部分。

2. **GroupRpcService**: 群组服务，使用 `group.Group` 客户端进行通信。配置在 `shortlink-api.yaml` 中的 `GroupRpcConf` 部分。

这两个服务通过 `zrpc` 客户端进行连接，
