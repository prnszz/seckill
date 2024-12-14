package main

import (
	"log"
	"seckill-system/internal/dao"
	"seckill-system/internal/handler"

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

	// 初始化数据库连接
	dao.InitMySQL()

	// 初始化Redis连接
	dao.InitRedis()
}

func main() {
	r := gin.Default()
	initRouter(r)

	port := viper.GetString("server.port")
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Server startup failed: %s", err)
	}
}

func initRouter(r *gin.Engine) {
	// 健康检查
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// 数据库测试
	r.GET("/test/db", handler.TestDB)

	// 商品管理路由
	productHandler := handler.NewProductHandler()
	v1 := r.Group("/api/v1")
	{
		products := v1.Group("/products")
		{
			products.POST("", productHandler.Create)       // 创建商品
			products.GET("/:id", productHandler.Get)       // 获取商品详情
			products.PUT("", productHandler.Update)        // 更新商品
			products.DELETE("/:id", productHandler.Delete) // 删除商品
			products.GET("", productHandler.List)          // 获取商品列表
		}
	}

	// 秒杀活动路由
	seckillHandler := handler.NewSeckillActivityHandler()
	v1 = r.Group("/api/v1")
	{
		seckill := v1.Group("/seckill")
		{
			seckill.POST("/activities", seckillHandler.Create) // 创建秒杀活动
			seckill.GET("/activities/:id", seckillHandler.Get) // 获取秒杀活动详情
			seckill.PUT("/activities", seckillHandler.Update)  // 更新秒杀活动
			seckill.GET("/activities", seckillHandler.List)    // 获取秒杀活动列表
		}
	}

	// 秒杀核心接口
	seckillCoreHandler := handler.NewSeckillCoreHandler()
	v1 = r.Group("/api/v1")
	{
		seckill := v1.Group("/seckill")
		{
			// ... 之前的秒杀活动管理接口 ...

			// 秒杀核心接口
			seckill.POST("/do/:id", seckillCoreHandler.Seckill)           // 执行秒杀
			seckill.POST("/preload/:id", seckillCoreHandler.PreloadStock) // 库存预热
		}
	}
}
