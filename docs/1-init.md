# 从零开始创建秒杀系统

```
mkdir -p cmd/api          # 主程序入口
mkdir -p internal/handler # HTTP处理器
mkdir -p internal/service # 业务逻辑层
mkdir -p internal/model   # 数据模型
mkdir -p internal/dao     # 数据访问层
mkdir -p configs         # 配置文件
mkdir -p scripts         # 脚本文件
mkdir -p docs           # 文档
mkdir -p test          # 测试文件

# 创建基础文件
touch go.mod
touch main.go
touch configs/config.yaml
touch README.md
```

## 初始化和依赖
```
go mod init seckill-system # 初始化go.mod

# Web框架
go get github.com/gin-gonic/gin

# 配置管理
go get github.com/spf13/viper

# 数据库驱动
go get gorm.io/gorm
go get gorm.io/driver/mysql

# Redis客户端
go get github.com/go-redis/redis/v8

# 日志
go get go.uber.org/zap
```


## 基础的配置文件`configs/config.yaml`
```
# configs/config.yaml
server:
  port: 8080
  mode: debug

mysql:
  host: localhost
  port: 3306
  username: root
  password: root
  database: seckill
  charset: utf8mb4

redis:
  host: localhost
  port: 6379
  password: ""
  db: 0
```
## 主程序入口`main.go`
```
package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	// 加载配置文件
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}
}

func main() {
	// 创建gin实例
	r := gin.Default()

	// 基础路由
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 启动服务器
	port := viper.GetString("server.port")
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Server startup failed: %s", err)
	}
}

```
## 确保MySQL和Redis已经安装

## 初步执行
```
go mod tidy  # 整理并下载依赖
go run main.go  # 运行项目
```