// scripts/redis/clean.go
package main

import (
	"context"
	"fmt"
	"log"
	"seckill-system/internal/dao"

	"github.com/spf13/viper"
)

func init() {
	// 加载配置文件
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../../configs")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	// 初始化Redis连接
	dao.InitRedis()
}

func main() {
	ctx := context.Background()

	// 清理库存缓存
	stockKeys, err := dao.RDB.Keys(ctx, "seckill:stock:*").Result()
	if err != nil {
		log.Fatalf("获取库存键失败: %v", err)
	}

	// 清理用户购买记录
	userKeys, err := dao.RDB.Keys(ctx, "seckill:user:*").Result()
	if err != nil {
		log.Fatalf("获取用户记录键失败: %v", err)
	}

	// 合并所有需要删除的键
	allKeys := append(stockKeys, userKeys...)

	if len(allKeys) > 0 {
		// 批量删除
		if err := dao.RDB.Del(ctx, allKeys...).Err(); err != nil {
			log.Fatalf("清理缓存失败: %v", err)
		}
		fmt.Printf("成功清理 %d 个缓存键\n", len(allKeys))
	} else {
		fmt.Println("没有需要清理的缓存")
	}
}
