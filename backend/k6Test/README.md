# 短链接数据生成器与性能测试

这个工具用于生成测试用的短链接数据，并提供高强度的性能和压力测试。

## 功能特点

- 生成50,000条随机短链接数据
- 批量插入到MongoDB数据库
- 提供进度显示和统计信息
- 支持高并发压力测试

## 使用方法

### 1. 数据生成

确保MongoDB服务已启动，默认连接到`mongodb://localhost:27019`，然后运行数据生成器:

```bash
cd backend/k6Test
go run data_generator.go
```

### 2. 性能测试

数据生成后，使用K6进行性能测试:

```bash
k6 run test.js
```

## 高强度测试特性

测试脚本包含多种数据库压力测试场景:

1. **高并发场景**
   - 最高支持1000个并发用户
   - 动态调整并发负载模拟峰值流量

2. **多样化压力测试**
   - 批量重复请求 (50个/批)
   - 随机混合请求 (30个/批)
   - 超大批量请求 (100个/批)
   - 长事务请求 (模拟长连接)
   - 复杂查询请求 (模拟统计分析)

3. **测试指标收集**
   - 数据库操作延迟 (db_operation_latency)
   - 数据库操作失败率 (db_operation_failures)
   - 缓存命中率分析
   - 响应时间分布

## 数据库监控

在运行测试的同时，建议监控MongoDB:

```bash
# 监控MongoDB容器资源使用
docker stats mongoContainer

# 查看数据库状态
docker exec -it mongoContainer mongo --port 27017 --eval "db.serverStatus()"

# 查看慢查询
docker exec -it mongoContainer mongo --port 27017 --eval "db.getSiblingDB('admin').system.profile.find({millis:{$gt:100}}).pretty()"
```

## 性能优化建议

如果测试结果显示性能瓶颈，可考虑:

1. 增加MongoDB索引
2. 启用Redis缓存层
3. 优化查询逻辑
4. 调整MongoDB连接池配置

## 生成的数据结构

每条短链接数据包含以下字段:
- 原始URL (随机选择不同来源)
- 6位字符的短链接URI
- 访问统计数据 (PV/UV/IP)
- 创建和更新时间

## 性能测试

数据生成后，可以使用K6进行性能测试:

```bash
k6 run test.js
```

测试脚本会模拟用户访问短链接，并收集响应时间、成功率等指标。 