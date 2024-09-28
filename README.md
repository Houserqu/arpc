# aRPC
基于 grpc 封装的针对 k8s 环境部署的微服务框架，特点：轻量、简单

## 主要特性

### gRPC
默认配置
```
grpc:
  addr: 0.0.0.0:8000
```

### HTTP
内置基于 grpc-gateway 的 HTTP 协议代理，需要在 proto 文件中配置 http path 和 method

默认配置
```
http:
  addr: 0.0.0.0:8080
```

### 服务发现
默认使用 k8s 的 dns 服务发现，端口 8000
可以通过配置自定义服务的 IP 和 端口，一般于开发环境

```
// 创建客户端
ipClient, err := arpc.GetServerClient("ip", ip.NewIPClient)
```

自定义服务的地址
```
discovery:
  ip: localhost:8000
```

### 配置
使用 viper 库加载本地 config.yaml 文件，可以通过环境变量 `CONFIG_PATH` 指定配置文件路径
建议使用 K8S 的 ConfigMap 管理配置文件，注意，如果想实现配置热更新，不要使用 subpath 挂载配置（subpath 不支持配置变更立即更新容器中的文件）

### 参数校验
基于 protoc-gen-validate 实现了全局参数校验，需要在 proto 中配置参数规则

### Mysql
内置 gorm 库作为 Mysql 客户端，最多支持创建 3 个 Mysql 客户端

使用方式
```
aprc.Mysql  # 客户端1
arpc.Mysql1 # 客户端1
arpc.Mysql2 # 客户端2
arpc.Mysql3 # 客户端3
```

配置示例
```
mysql: 
  host: localhost
  port: 3306
  database: test
  password: pass
  user: user
mysql2: 
  host: localhost
  port: 3306
  database: test2
  password: pass
  user: user
mysql2: 
  host: localhost
  port: 3306
  database: test3
  password: pass
  user: user
```

### Redis
内置 Redis 客户端，支持连接多个 DB

使用方式
```
arpc.Redis[1] # 1 是 DB 号
```

配置示例
```
redis:
  addr: localhost:6379
  username: user
  password: pass
  dbs: 
   - 0
   - 1
   - 2
```